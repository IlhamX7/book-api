package service

import (
	"book-api/data/request"
	"book-api/data/response"
	"book-api/helper"
	"book-api/model"
	"book-api/repository"

	"github.com/go-playground/validator/v10"
)

type BooksServiceImpl struct {
	BooksRepository repository.BooksRepository
	Validate        *validator.Validate
}

func NewBooksServiceImpl(bookRepository repository.BooksRepository, validate *validator.Validate) BooksService {
	return &BooksServiceImpl{
		BooksRepository: bookRepository,
		Validate:        validate,
	}
}

// Create implements BooksService
func (b *BooksServiceImpl) Create(books request.CreateBooksRequest) {
	err := b.Validate.Struct(books)
	if err != nil {
		panic(err.Error)
	}
	bookModel := model.Books{
		Judul:    books.Judul,
		Penerbit: books.Penerbit,
		Rating:   books.Rating,
	}
	b.BooksRepository.Save(bookModel)
}

// Delete implements BooksService
func (b *BooksServiceImpl) Delete(booksId int) {
	b.BooksRepository.Delete(booksId)
}

// FindAll implements BooksService
func (b *BooksServiceImpl) FindAll() []response.BooksResponse {
	result := b.BooksRepository.FindAll()

	var books []response.BooksResponse
	for _, value := range result {
		book := response.BooksResponse{
			Id:       value.Id,
			Judul:    value.Judul,
			Penerbit: value.Penerbit,
			Rating:   value.Rating,
		}
		books = append(books, book)
	}

	return books
}

// FindById implements BooksService
func (b *BooksServiceImpl) FindById(booksId int) response.BooksResponse {
	bookData, err := b.BooksRepository.FindById(booksId)
	helper.ErrorPanic(err)

	bookResponse := response.BooksResponse{
		Id:       bookData.Id,
		Judul:    bookData.Judul,
		Penerbit: bookData.Penerbit,
		Rating:   bookData.Rating,
	}
	return bookResponse
}

// Update implements BooksService
func (b *BooksServiceImpl) Update(books request.UpdateBooksRequest) {
	bookData, err := b.BooksRepository.FindById(books.Id)
	if err != nil {
		panic(err.Error)
	}
	bookData.Judul = books.Judul
	bookData.Penerbit = books.Penerbit
	bookData.Rating = books.Rating
	b.BooksRepository.Update(bookData)
}
