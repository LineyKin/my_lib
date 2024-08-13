package book

import (
	"database/sql"
	"fmt"
	db "my_lib/models/db"
)

type PublishingHouse struct {
	Id   int
	Name string
}

const tableName string = "publishing_house"

func (ph PublishingHouse) isEmpty() bool {
	if ph.Id == 0 && ph.Name == "" {
		return true
	}

	return false
}

func (ph PublishingHouse) isNew() bool {
	if ph.Id == 0 && ph.Name != "" {
		return true
	}

	return false
}

func (ph *PublishingHouse) Add() error {
	db, err := db.GetConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlPattern := `INSERT INTO %s (name) VALUES(:name)`
	sqlPattern = fmt.Sprintf(sqlPattern, tableName)
	res, err := db.Exec(
		sqlPattern,
		sql.Named("name", ph.Name))

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	ph.Id = int(id)

	return nil
}
