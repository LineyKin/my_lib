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

	//s := []rune("w")
	////s := "1"
	//
	//fmt.Println(unicode.IsDigit(s[0]))
	//
	//return
	// проверяем БД и в случае отсутствия создаём её с таблицей
	db.Create()

	r := gin.Default()

	// ручка для главной страницы
	r.GET("/", hand_fs.FileServerHandler)
	r.Static("/style", "./web/style")
	r.Static("/js", "./web/js")

	// ручка добавления авторов
	r.POST("api/author/add", hand_add.AddAuthor)

	port := env.GetPort()
	fmt.Printf("http://localhost:%s/\n", port)
	err := r.Run(":" + port)
	if err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}

}
