package main

import (
	"fmt"
	hand_add "my_lib/handlers/add_book"
	hand_fs "my_lib/handlers/fileserver"
	"my_lib/helpers/env"
	db "my_lib/models/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Create()

	r := gin.Default()

	// ручка для главной страницы
	r.GET("/", hand_fs.FileServerHandler)
	r.Static("/style", "./web/style")
	r.Static("/js", "./web/js")

	// ручка добавления авторов
	r.POST("api/author/add", hand_add.AddAuthor)

	// ручка добавления книги
	r.POST("api/book/add", hand_add.AddBook)

	// ручка выгрузки авторов для подсказки в форме добавления книги
	r.GET("api/author/hint", hand_add.GetHint)

	port := env.GetPort()
	fmt.Printf("http://localhost:%s/\n", port)
	err := r.Run(":" + port)
	if err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}

}
