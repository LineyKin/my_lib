package main

import (
	"log"
	"my_lib/internal/db"
	"my_lib/internal/pkg/app"
)

func main() {
	dbsqlite, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbsqlite.Close()

	a, err := app.New(dbsqlite)
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
