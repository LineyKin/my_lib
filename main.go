package main

import (
	"fmt"
	hand_fs "my_lib/handlers/fileserver"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ручка для главной страницы
	r.GET("/", hand_fs.FileServerHandler)
	r.Static("/style", "./web/style")
	r.Static("/js", "./web/js")

	port := "1991"
	fmt.Printf("http://localhost:%s/\n", port)
	err := r.Run(":" + port)
	if err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}

}
