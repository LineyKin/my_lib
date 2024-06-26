package db

import (
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

	// создаём таблицу authors
	createTableAuthors()
}

func createDbFile(dbPath string) {
	_, err := os.Create(dbPath)
	if err != nil {
		fmt.Println(err)
	}
}

func createTableAuthors() {
	db, err := GetConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	sql := `CREATE TABLE IF NOT EXISTS authors (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(256) NOT NULL DEFAULT "",
		father_name VARCHAR(256) NOT NULL DEFAULT "",
		last_name VARCHAR(256) NOT NULL DEFAULT ""
		
	);`

	_, err = db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
}
