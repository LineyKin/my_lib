package main

import (
	"fmt"
	hand_add "my_lib/handlers/addbook"
	hand_fs "my_lib/handlers/fileserver"
	"my_lib/helpers/env"

	"github.com/gin-gonic/gin"
)

func main() {
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
