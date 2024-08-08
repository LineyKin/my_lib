package author

import (
	"database/sql"
	"fmt"
	db "my_lib/models/db"
)

const tableName string = "authors"

type Author struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	FatherName string `json:"fatherName,omitempty"`
	LastName   string `json:"lastName"`
}

// список авторов без отчества
func (a Author) GetHintList() ([]Author, error) {
	db, err := db.GetConnection()
	if err != nil {
		return []Author{}, err
	}
	defer db.Close()

	sql := fmt.Sprintf("SELECT id, name, last_name  FROM %s ORDER BY last_name", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		return []Author{}, err
	}

	list := []Author{}
	for rows.Next() {
		authorRow := Author{}
		err := rows.Scan(&authorRow.Id, &authorRow.Name, &authorRow.LastName)

		if err != nil {
			return []Author{}, err
		}

		list = append(list, authorRow)
	}

	return list, nil
}

// добавление автора
func (a Author) Add() (int, error) {
	db, err := db.GetConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	sqlPattern := `INSERT INTO %s (name, father_name, last_name) VALUES(:name, :father_name, :last_name)`
	sqlPattern = fmt.Sprintf(sqlPattern, tableName)
	res, err := db.Exec(
		sqlPattern,
		sql.Named("name", a.Name),
		sql.Named("father_name", a.FatherName),
		sql.Named("last_name", a.LastName))

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
