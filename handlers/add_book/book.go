package add_book

import (
	"fmt"
	book "my_lib/models/book"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddBook(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	var book book.Book
	if err := c.BindJSON(&book); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка десериализации JSON"})
		return
	}

	response, err := book.Add()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ids": response})
}
