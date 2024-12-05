package service

import (
	"errors"
	"fmt"
	"my_lib/internal/storage/db/sqlite"
	"my_lib/models/author"
	"my_lib/models/book"
)

type myLibService struct {
	storage sqlite.StorageInterface
}

func NewService(s sqlite.StorageInterface) *myLibService {
	return &myLibService{
		storage: s,
	}
}

func (s *myLibService) GetBookList(limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error) {
	return s.storage.GetBookList(limit, offset, sortedBy, sortType)
}

func (s *myLibService) AddBook(b book.BookAdd) (book.BookAdd, error) {
	// 1. проверка заполнено ли издательство
	if b.PublishingHouse.IsEmpty() {
		return b, errors.New("поле 'Издательство' не заполнено")
	}

	// 2. для нового издательства получаем id (publishing_house_id)
	// после добавления в БД
	// в ином случае id прилетает с фронта сразу
	if b.PublishingHouse.IsNew() {
		err := s.addPublishingHouse(&b)
		if err != nil {
			return b, err
		}
	}

	// 3. заполним таблицу book
	// получим id книги
	err := s.addPhysicalBook(&b)
	if err != nil {
		return b, err
	}

	//4. работа с литературным произведением
	// проверяем, чтобы было заполнено хотя бы одно проле с названием
	if b.IsEmptyNameList() {
		return b, errors.New("поле 'Название' не заполнено")
	}

	// булева переменная, указывающая на наличие автора у книги
	hasAuthors := b.HasAuthors()

	// перебираем литературные произведение (названия) lw - literary work
	for _, lw := range b.Name {
		if lw.IsEmpty() {
			continue // одно название может быть пустым(на совести пользователя, может исправлю на фронте ещё)
		}

		// для нового литературного произведения получаем id (literary_work_id)
		// после добавления в БД
		// в ином случае id прилетает с фронта сразу
		if lw.IsNew() {
			err = s.addLiteraryWork(&lw)
			if err != nil {
				return b, err
			}
		}

		// далее заполняем связующте таблицы в БД

		// 4.1 заполним таблицу, связывающую физическую книгу и произведение, которые в ней (таблица book_and_literary_work)
		err := s.storage.LinkBookAndLiteraryWork(lw.Id, b.Id)
		if err != nil {
			return b, err
		}

		// случай, когда авторов у книги нет : Библия, например
		if !hasAuthors {
			continue
		}

		// 4.2 заполним таблицу, связывающую произведение и автора(ов)
		// перебираем авторов
		for _, authorId := range b.Author {
			err := s.storage.LinkAuthorAndLiteraryWork(authorId, lw.Id)
			if err != nil {
				return b, err
			}
		}
	}

	return b, nil
}

func (s *myLibService) addLiteraryWork(lw *book.LiteraryWork) error {
	id, err := s.storage.AddLiteraryWork(lw.Name)
	if err != nil {
		return err
	}

	lw.Id = id
	return nil
}

func (s *myLibService) addPhysicalBook(b *book.BookAdd) error {
	id, err := s.storage.AddPhysicalBook(b)
	if err != nil {
		return err
	}

	b.Id = id

	return nil
}

func (s *myLibService) addPublishingHouse(b *book.BookAdd) error {
	id, err := s.storage.AddPublishingHouse(b.PublishingHouse.Name)
	if err != nil {
		return err
	}
	b.PublishingHouse.Id = id
	return nil
}

func (s *myLibService) AddAuthor(a author.Author) (int, error) {

	isExists, err := s.IsAuthorExists(a)
	if err != nil {
		return 0, err
	}

	if isExists {
		return 0, fmt.Errorf("автор %s уже добавлен", a.GetName())
	}

	// TODO обработать ошибку
	id, err := s.storage.AddAuthor(a)
	fmt.Println(err)
	return id, nil
}

// проверяем, не заносили ли мы уже этого автора в БД
// исходим из того, что полные тёски среди авторов редкость
func (s *myLibService) IsAuthorExists(a author.Author) (bool, error) {
	idList, err := s.storage.GetAuthorByName(a)
	if err != nil {
		return false, err
	}

	len := len(idList)

	if len == 0 {
		return false, nil
	}

	if len == 1 {
		return true, nil
	}

	return true, fmt.Errorf("обнаружены дубликаты автора %s", a.GetName())
}

func (s *myLibService) GetAuthorList() ([]author.Author, error) {
	return s.storage.GetAuthorList()
}

func (s *myLibService) GetPublishingHouseList() ([]book.PublishingHouse, error) {
	return s.storage.GetPublishingHouseList()
}

func (s *myLibService) GetBookCount() (int, error) {
	return s.storage.GetBookCount()
}
