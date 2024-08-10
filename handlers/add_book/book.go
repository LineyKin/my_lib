package add_book

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddBook(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}
}
