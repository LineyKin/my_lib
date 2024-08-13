package book

import (
	"errors"
	db "my_lib/models/db"
)

type Book struct {
	Id              string              `json:"id,omitempty"`
	Name            []map[string]string `json:"name"`
	Author          []int               `json:"author"`
	PublishingHouse PublishingHouse     `json:"publishingHouse"`
	PublishingYear  string              `json:"publishingYear"`
}

// добавление книги

// возвращает id издательства и произведения
func (b Book) Add() (map[string]int, error) {
	db, err := db.GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// работа с издательством
	if b.PublishingHouse.isEmpty() {
		return nil, errors.New("издательство не заполнено")
	}

	if b.PublishingHouse.isNew() {
		err = b.PublishingHouse.Add()
		if err != nil {
			return map[string]int{}, err
		}
	}

	// 2. Добавляем физическую книгу в БД, получаем book_id

	// 3. Если name.id == 0, добавляем новое произведение в БД, получаем literary_work_id

	// 4. заполняем связующую таблицу book_and_literary_work
	// book_id и literary_work_id у нас уже получены

	// 5. Заполняем связующую таблицу author_and_literary_work
	// literary_work_id и author_id у нас уже есть

	// 6. Возвращаем publishing_house_id и literary_work_id для даталистов Названия и Издательства на фронте

	return nil, nil
}
