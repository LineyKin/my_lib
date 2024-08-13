package book

import (
	"database/sql"
	"errors"
	"fmt"
	db "my_lib/models/db"
)

type Book struct {
	Id              int             `json:"id,omitempty"`
	Name            []LiteraryWork  `json:"name"`
	Author          []int           `json:"author"`
	PublishingHouse PublishingHouse `json:"publishingHouse"`
	PublishingYear  string          `json:"publishingYear"`
}

const tableName = "book"

// добавление книги

// возвращает id издательства и произведения
func (b Book) Add() (Book, error) {
	db, err := db.GetConnection()
	if err != nil {
		return Book{}, err
	}
	defer db.Close()

	// работа с издательством, проверка на пустоту
	if b.PublishingHouse.isEmpty() {
		return Book{}, errors.New("поле 'Издательство' не заполнено")
	}

	// для нового издательства получаем id (publishing_house_id)
	// после добавления в БД
	// в ином случае id прилетает с фронта сразу
	if b.PublishingHouse.isNew() {
		err = b.PublishingHouse.Add()
		if err != nil {
			return Book{}, err
		}
	}

	// заполним таблицу book
	// получим id книги
	err = b.addPhysicalBook()
	if err != nil {
		return Book{}, err
	}

	// 2. работа с литературным произведением
	if isEmptyNameList(b.Name) {
		return Book{}, errors.New("поле 'Название' не заполнено")
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
				return Book{}, err
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
			return Book{}, err
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
				return Book{}, err
			}
		}

	}

	return b, nil
}

// добавление физической книги
func (b *Book) addPhysicalBook() error {
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
