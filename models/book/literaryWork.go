package book

import (
	"database/sql"
	"fmt"
	db "my_lib/models/db"
)

type LiteraryWork struct {
	Id   int
	Name string
}

const lwTableName string = "literary_work"

func (lw LiteraryWork) isEmpty() bool {
	if lw.Id == 0 && lw.Name == "" {
		return true
	}

	return false
}

func (lw LiteraryWork) isNew() bool {
	if lw.Id == 0 && lw.Name != "" {
		return true
	}

	return false
}

func (lw *LiteraryWork) Add() error {
	db, err := db.GetConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlPattern := `INSERT INTO %s (name) VALUES(:name)`
	sqlPattern = fmt.Sprintf(sqlPattern, lwTableName)
	res, err := db.Exec(
		sqlPattern,
		sql.Named("name", lw.Name))

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	lw.Id = int(id)

	return nil
}
