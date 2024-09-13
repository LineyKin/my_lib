package main

import (
	"fmt"
	"log"
	"my_lib/handlers"
	"my_lib/internal/db"
	"my_lib/lib/env"
	"my_lib/service"
	"my_lib/storage/db/sqlite"

	"github.com/gin-gonic/gin"
)

func main() {

	dbsqlite, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbsqlite.Close()

	Storage := sqlite.New(dbsqlite)
	Service := service.New(Storage)
	Handlers := handlers.New(Service)

	r := gin.Default()

	// ручка для главной страницы
	r.GET("/", Handlers.FileServer)
	r.Static("/style", "./web/style")
	r.Static("/js", "./web/js")

	// ручка добавления авторов
	r.POST("api/author/add", Handlers.AddAuthor)

	// ручка для выгрузки списка книг
	r.GET("api/book/list", Handlers.GetBookList)

	// ручка для выгрузки количества книг
	r.GET("api/book/count", Handlers.GetBookCount)

	// ручка добавления книги
	r.POST("api/book/add", Handlers.AddBook)

	// ручка выгрузки авторов для подсказки в форме добавления книги
	r.GET("api/author/hint", Handlers.GetAuthorList)

	// ручка выгрузки списка издательств
	r.GET("api/publishingHouse/list", Handlers.GetPublishingHouseList)

	port := env.GetPort()
	fmt.Printf("http://localhost:%s/\n", port)
	err = r.Run(":" + port)
	if err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}

}
