package checkout

import (
	"example/Go-Api-Tutorial/app/books/usecase/checkout"

	"github.com/gin-gonic/gin"
)

// Params
type Params struct {
	ID string
}

func CheckoutBook(c *gin.Context) {
	id, _ := c.GetQuery("id")

	err := validate(Params{ID: id})
	if err != nil {
		c.IndentedJSON(422, gin.H{"message": err})
		return
	}

	checkout.PatchCheckout(c, id)
}
