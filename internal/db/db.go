package db

import (
	"database/sql"
	"fmt"
	"log"
	"my_lib/lib/env"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", env.GetDbName())
	if err != nil {
		log.Fatal("can't open database:", err)
		return nil, fmt.Errorf("can't open database: %s", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("can't connect to database:", err)
		return nil, fmt.Errorf("can't connect to database: %s", err)
	}

	err = createTableAuthors(db)
	if err != nil {
		return nil, err
	}

	err = createTableBook(db)
	if err != nil {
		return nil, err
	}

	err = createTableLiteraryWork(db)
	if err != nil {
		return nil, err
	}

	err = createTableLiteraryWorkAndAuthors(db)
	if err != nil {
		return nil, err
	}

	err = createTableLiteraryWorkAndBook(db)
	if err != nil {
		return nil, err
	}

	err = createTablePublishingHouse(db)
	if err != nil {
		return nil, err
	}

	return db, nil

}

// связная таблица литературных произведений и физических книг
func createTableLiteraryWorkAndBook(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS book_and_literary_work (
		literary_work_id INTEGER,
		book_id INTEGER
	);`

	return createTable(s, q, "book_and_literary_work")
}

// связная таблица литературных произведений и авторов
func createTableLiteraryWorkAndAuthors(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS author_and_literary_work (
		author_id INTEGER,
		literary_work_id INTEGER
	);`

	return createTable(s, q, "author_and_literary_work")
}

// таблица (физических) книг
func createTableBook(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS book (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		year_of_publication INTEGER,
		publishing_house_id INTEGER
	);`

	return createTable(s, q, "book")
}

// таблица литературных произведений
func createTablePublishingHouse(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS publishing_house (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(256) NOT NULL DEFAULT ""
	);`

	return createTable(s, q, "publishing_house")
}

// таблица литературных произведений
func createTableLiteraryWork(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS literary_work (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(256) NOT NULL DEFAULT ""
	);`

	return createTable(s, q, "literary_work")
}

// таблица авторов
func createTableAuthors(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS authors (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(256) NOT NULL DEFAULT "",
		father_name VARCHAR(256) NOT NULL DEFAULT "",
		last_name VARCHAR(256) NOT NULL DEFAULT ""
	);`

	return createTable(s, q, "authors")
}

func createTable(s *sql.DB, query, tableName string) error {
	_, err := s.Exec(query)
	if err != nil {
		return fmt.Errorf("can't create table `%s`: %w", tableName, err)
	}

	return nil
}
