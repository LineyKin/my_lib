package book

import (
	"database/sql"
	"errors"
	"fmt"
	db "my_lib/models/db"
)

type Book struct {
	Id             int    `json:"id,omitempty"`
	PublishingYear string `json:"publishingYear"`
}

// формат добавления
type BookAdd struct {
	Book
	Name            []LiteraryWork  `json:"name"`
	Author          []int           `json:"author"`
	PublishingHouse PublishingHouse `json:"publishingHouse"`
}

// формат чтения из списка
type BookUnload struct {
	Book
	Name            string `json:"name"`
	Author          string `json:"author"`
	PublishingHouse string `json:"publishingHouse"`
}

type ListParam struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

const tableName = "book"

func (*Book) Count() (int, error) {
	db, err := db.GetConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	sql := `SELECT COUNT(*) AS count FROM book`
	row, err := db.Query(sql)
	if err != nil {
		return 0, err
	}

	var count int

	row.Next()
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (b BookUnload) GetList(lp ListParam) ([]BookUnload, error) {
	db, err := db.GetConnection()
	if err != nil {
		return []BookUnload{}, err
	}
	defer db.Close()
	sqlPattern := `
	SELECT
 		b.id AS id,
 		GROUP_CONCAT(IFNULL(a.name || ' ' || a.last_name, '-'), ', ') AS author,
 		lw.name AS name,
 		ph.name AS publishing_house,
 		b.year_of_publication
	FROM book AS b
	LEFT JOIN publishing_house AS ph ON ph.id = b.publishing_house_id
	LEFT JOIN book_and_literary_work AS blw ON blw.book_id = b.id
	LEFT JOIN literary_work AS lw ON lw.id = blw.literary_work_id
	LEFT JOIN author_and_literary_work AS alw ON alw.literary_work_id = lw.id
	LEFT JOIN authors AS a ON a.id = alw.author_id
	GROUP BY b.id
	ORDER BY a.last_name, lw.name
	LIMIT %d OFFSET %d;`
	sql := fmt.Sprintf(sqlPattern, lp.Limit, lp.Offset)
	rows, err := db.Query(sql)
	if err != nil {
		return []BookUnload{}, err
	}

	list := []BookUnload{}
	for rows.Next() {
		bRow := BookUnload{}
		err := rows.Scan(&bRow.Id, &bRow.Author, &bRow.Name, &bRow.PublishingHouse, &bRow.PublishingYear)

		if err != nil {
			return []BookUnload{}, err
		}

		list = append(list, bRow)
	}

	return list, nil

}

// возвращает BookAdd
func (b BookAdd) Add() (BookAdd, error) {
	db, err := db.GetConnection()
	if err != nil {
		return BookAdd{}, err
	}
	defer db.Close()

	// работа с издательством, проверка на пустоту
	if b.PublishingHouse.isEmpty() {
		return BookAdd{}, errors.New("поле 'Издательство' не заполнено")
	}

	// для нового издательства получаем id (publishing_house_id)
	// после добавления в БД
	// в ином случае id прилетает с фронта сразу
	if b.PublishingHouse.isNew() {
		err = b.PublishingHouse.Add()
		if err != nil {
			return BookAdd{}, err
		}
	}

	// заполним таблицу book
	// получим id книги
	err = b.addPhysicalBook()
	if err != nil {
		return BookAdd{}, err
	}

	// 2. работа с литературным произведением
	if isEmptyNameList(b.Name) {
		return BookAdd{}, errors.New("поле 'Название' не заполнено")
	}

	// перебираем литературные произведения
	for _, lw := range b.Name {
		if lw.isEmpty() {
			continue
		}

		// для нового литературного произведения получаем id (literary_work_id)
		// после добавления в БД
		// в ином случае id прилетает с фронта сразу
		if lw.isNew() {
			err = lw.Add()
			if err != nil {
				return BookAdd{}, err
			}
		}

		// заполним в цикле связные таблицы
		// book_and_literary_work
		sqlPattern := `INSERT INTO book_and_literary_work (literary_work_id, book_id) VALUES(:lw_id, :b_id)`
		_, err := db.Exec(
			sqlPattern,
			sql.Named("lw_id", lw.Id),
			sql.Named("b_id", b.Id),
		)

		if err != nil {
			return BookAdd{}, err
		}

		// author_and_literary_work
		for _, aId := range b.Author {
			sqlPattern := `INSERT INTO author_and_literary_work (author_id, literary_work_id) VALUES(:a_id, :lw_id)`
			_, err := db.Exec(
				sqlPattern,
				sql.Named("lw_id", lw.Id),
				sql.Named("a_id", aId),
			)

			if err != nil {
				return BookAdd{}, err
			}
		}

	}

	return b, nil
}

// добавление физической книги
func (b *BookAdd) addPhysicalBook() error {
	db, err := db.GetConnection()
	if err != nil {
		return err

	}
	defer db.Close()

	sqlPattern := `INSERT INTO %s (publishing_house_id, year_of_publication) VALUES(:ph_id, :yop)`
	sqlPattern = fmt.Sprintf(sqlPattern, tableName)
	res, err := db.Exec(
		sqlPattern,
		sql.Named("ph_id", b.PublishingHouse.Id),
		sql.Named("yop", b.PublishingYear),
	)

	if err != nil {
		return err
	}

	bookId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	b.Id = int(bookId)

	return nil
}

// проверка на пустоту списка произведений
func isEmptyNameList(lwList []LiteraryWork) bool {
	for _, lw := range lwList {
		if !lw.isEmpty() {
			return false
		}
	}

	return true
}
