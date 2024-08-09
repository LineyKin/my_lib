package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

func Create() {
	dbPath := getDbPath()

	// проверяем, есть ли файл БД
	_, err := os.Stat(dbPath)
	if err != nil {
		createDbFile(dbPath) // создаём файл БД, если его нет
	}

	db, err := GetConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// создаём таблицу authors
	createTableAuthors(db)

	// создаём таблицу literary_work
	createTableLiteraryWork(db)

	// создаём таблицу publishing_house
	createTablePublishingHouse(db)

	// создаём таблицу book
	createTableBook(db)

	// создаём связную таблицу author_and_literary_work
	createTableLiteraryWorkAndAuthors(db)

	// создаём связную таблицу book_and_literary_work
	createTableLiteraryWorkAndBook(db)
}

func createDbFile(dbPath string) {
	_, err := os.Create(dbPath)
	if err != nil {
		fmt.Println(err)
	}
}

// связная таблица литературных произведений и физических книг
func createTableLiteraryWorkAndBook(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS author_and_literary_work (
		literary_work_id INTEGER,
		book_id INTEGER
	);`

	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
}

// связная таблица литературных произведений и авторов
func createTableLiteraryWorkAndAuthors(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS author_and_literary_work (
		author_id INTEGER,
		literary_work_id INTEGER
	);`

	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
}

// таблица (физических) книг
func createTableBook(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS book (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		year_of_publication INTEGER,
		publishing_house_id INTEGER
	);`

	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
}

// таблица литературных произведений
func createTablePublishingHouse(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS publishing_house (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(256) NOT NULL DEFAULT ""
	);`

	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
}

// таблица литературных произведений
func createTableLiteraryWork(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS literary_work (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(256) NOT NULL DEFAULT ""
	);`

	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
}

// таблица авторов
func createTableAuthors(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS authors (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(256) NOT NULL DEFAULT "",
		father_name VARCHAR(256) NOT NULL DEFAULT "",
		last_name VARCHAR(256) NOT NULL DEFAULT ""
	);`

	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
}
