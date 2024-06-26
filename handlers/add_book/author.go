package add_book

import (
	"fmt"
	author "my_lib/models/author"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	fmt.Println(author.ID)
	fmt.Println(author.Name)
	fmt.Println(author.FatherName)
	fmt.Println(author.LastName)

	authorId, _ := author.Add()

	c.JSON(http.StatusOK, gin.H{"id": authorId})
}
