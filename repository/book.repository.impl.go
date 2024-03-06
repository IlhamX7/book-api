package repository

import (
	"book-api/data/request"
	"book-api/helper"
	"errors"
	"fmt"

	"book-api/model"

	"gorm.io/gorm"
)

type BooksRepositoryImpl struct {
	Db *gorm.DB
}

func NewBooksRepositoryImpl(Db *gorm.DB) BooksRepository {
	return &BooksRepositoryImpl{Db: Db}
}

// Delete implements BooksRepository
func (b *BooksRepositoryImpl) Delete(booksId int) {
	var books model.Books
	result := b.Db.Where("id = ?", booksId).Delete(&books)
	helper.ErrorPanic(result.Error)
}

// FindAll implements BooksRepository
func (b *BooksRepositoryImpl) FindAll() []model.Books {
	var books []model.Books
	result := b.Db.Find(&books)
	helper.ErrorPanic(result.Error)
	return books
}

// FindById implements BooksRepository
func (b *BooksRepositoryImpl) FindById(booksId int) (books model.Books, err error) {
	var book model.Books
	result := b.Db.Find(&book, booksId)
	fmt.Println("hasil ", &result)
	if result != nil {
		return book, nil
	} else {
		return book, errors.New("books is not found")
	}
}

// Save implements BooksRepository
func (b *BooksRepositoryImpl) Save(books model.Books) {
	result := b.Db.Create(&books)
	if result.Error != nil {
		panic(result.Error)
	}
}

// Update implements BooksRepository
func (b *BooksRepositoryImpl) Update(books model.Books) {
	var updateBook = request.UpdateBooksRequest{
		Id:       books.Id,
		Judul:    books.Judul,
		Penerbit: books.Penerbit,
		Rating:   books.Rating,
	}
	result := b.Db.Model(&books).Updates(updateBook)
	if result.Error != nil {
		panic(result.Error)
	}
}
