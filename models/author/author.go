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
	FatherName string `json:"fatherName"`
	LastName   string `json:"lastName"`
}

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
