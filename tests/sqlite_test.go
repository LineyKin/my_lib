package tests

// go test -count=1 ./tests

import (
	"database/sql"
	"log"
	"my_lib/lib/env"
	"my_lib/models/author"
	"my_lib/storage/db/sqlite"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetAuthorByName struct {
	author author.Author
	want   []int
}

func TestGetAuthorByName(t *testing.T) {
	dbsqlite, err := sql.Open("sqlite", "../"+env.GetDbName())
	if err != nil {
		log.Fatal(err)
	}
	defer dbsqlite.Close()

	Storage := sqlite.New(dbsqlite)

	testData := []GetAuthorByName{
		{author: author.Author{Name: "Михаил", FatherName: "Афанасьевич", LastName: "Булгаков"}, want: []int{9}},
		{author: author.Author{Name: "Михаил", LastName: "Булгаков"}, want: []int{}},
		{author: author.Author{Name: "Карл", LastName: "Маркс"}, want: []int{4}},
		{author: author.Author{Name: "Святозар", LastName: "Стёркин"}, want: []int{}},
		{author: author.Author{Name: "Вера", FatherName: "Викторовна", LastName: "Камша"}, want: []int{2}},
		{author: author.Author{Name: "Вера", FatherName: "Викторовна", LastName: "Камша22"}, want: []int{}},
	}

	for _, v := range testData {
		idList, err := Storage.GetAuthorByName(v.author)
		assert.Equal(t, v.want, idList)
		assert.NoError(t, err)
	}
}
