package main

import (
	"fmt"
	hand_fs "my_lib/handlers/fileserver"
	"my_lib/helpers/env"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Author struct {
	Name       string `json:"name"`
	FatherName string `json:"fatherName"`
	LastName   string `json:"lastName"`
}

func addAuthor(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	var author Author
	if err := c.BindJSON(&author); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка десериализации JSON"})
		return
	}

	fmt.Println(author.Name)
	fmt.Println(author.FatherName)
	fmt.Println(author.LastName)

	c.JSON(http.StatusOK, gin.H{"id": "1917"})
}

func main() {
	r := gin.Default()

	// ручка для главной страницы
	r.GET("/", hand_fs.FileServerHandler)
	r.Static("/style", "./web/style")
	r.Static("/js", "./web/js")

	// ручка добавления авторов
	r.POST("api/author/add", addAuthor)

	port := env.GetPort()
	fmt.Printf("http://localhost:%s/\n", port)
	err := r.Run(":" + port)
	if err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}

}
