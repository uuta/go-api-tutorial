package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"example/Go-Api-Tutorial/app/books/domain/model"
	"example/Go-Api-Tutorial/app/books/domain/repository"
	"example/Go-Api-Tutorial/app/books/interface/handler/checkout"
)

func getBooks(c *gin.Context) {
	data := "example"
	err := validation.Validate(data,
		validation.Required,       // not empty
		validation.Length(5, 100), // length between 5 and 100
		is.URL,                    // is a valid URL
	)
	fmt.Println(err)

	c.IndentedJSON(http.StatusOK, model.Books)
}

func bookById(c *gin.Context) {
	id := c.Param("id") // "/books/2"
	book, err := repository.GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
	}

	book, err := repository.GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook model.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	model.Books = append(model.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkout.CheckoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}
