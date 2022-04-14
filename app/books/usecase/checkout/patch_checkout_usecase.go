package checkout

import (
	"example/Go-Api-Tutorial/app/books/domain/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PatchCheckout(c *gin.Context, id string) {

	book, err := repository.GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}
