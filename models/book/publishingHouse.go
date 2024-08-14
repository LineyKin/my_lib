package book

import (
	"database/sql"
	"fmt"
	db "my_lib/models/db"
)

type PublishingHouse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

const phTableName string = "publishing_house"

// список издательств
func (ph PublishingHouse) GetList() ([]PublishingHouse, error) {
	db, err := db.GetConnection()
	if err != nil {
		return []PublishingHouse{}, err
	}
	defer db.Close()

	sql := fmt.Sprintf("SELECT * FROM %s ORDER BY name", phTableName)
	rows, err := db.Query(sql)
	if err != nil {
		return []PublishingHouse{}, err
	}

	list := []PublishingHouse{}
	for rows.Next() {
		phRow := PublishingHouse{}
		err := rows.Scan(&phRow.Id, &phRow.Name)

		if err != nil {
			return []PublishingHouse{}, err
		}

		list = append(list, phRow)
	}

	return list, nil
}

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
	sqlPattern = fmt.Sprintf(sqlPattern, phTableName)
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
