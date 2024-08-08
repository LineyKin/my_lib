package add_book

import (
	"fmt"
	author "my_lib/models/author"
	"net/http"

	"github.com/gin-gonic/gin"
)

// выгрузка списка авторов в подсказу datalist
func GetHint(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	var author author.Author
	list, err := author.GetHintList()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hint_list": list})
}

// добавление нового автора
func AddAuthor(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	var author author.Author
	if err := c.BindJSON(&author); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка десериализации JSON"})
		return
	}

	authorId, err := author.Add()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": authorId})
}
