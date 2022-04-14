package model

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var Books = []Book{
	{ID: "1", Title: "In search of Lost Time", Author: "Marcel Proust", Quantity: 1},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 2},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 3},
}
