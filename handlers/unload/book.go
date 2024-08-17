package unload

import (
	"my_lib/models/book"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBookCount(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	var b book.Book

	count, err := b.Count()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}

func GetBookList(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	var lp book.ListParam
	if err := c.BindJSON(&lp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка десериализации JSON " + err.Error()})
		return
	}

	var b book.BookUnload
	list, err := b.GetList(lp)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book_list": list})
}
