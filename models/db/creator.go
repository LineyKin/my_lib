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
}

func createDbFile(dbPath string) {
	_, err := os.Create(dbPath)
	if err != nil {
		fmt.Println(err)
	}
}

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
