package repository

import (
	"errors"
	"example/Go-Api-Tutorial/app/books/domain/model"
)

func GetBookById(id string) (*model.Book, error) {
	for i, b := range model.Books {
		if b.ID == id {
			return &model.Books[i], nil
		}
	}

	return nil, errors.New("book not found")
}
