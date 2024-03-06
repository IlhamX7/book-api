package service

import (
	"book-api/data/request"
	"book-api/data/response"
)

type BooksService interface {
	Create(books request.CreateBooksRequest)
	Update(books request.UpdateBooksRequest)
	Delete(booksId int)
	FindById(booksId int) response.BooksResponse
	FindAll() []response.BooksResponse
}
