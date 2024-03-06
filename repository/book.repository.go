package repository

import "book-api/model"

type BooksRepository interface {
	Save(books model.Books)
	Update(books model.Books)
	Delete(booksId int)
	FindById(booksId int) (books model.Books, err error)
	FindAll() []model.Books
}
