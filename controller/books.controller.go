package controller

import (
	"book-api/data/request"
	"book-api/data/response"
	"book-api/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type BooksController struct {
	booksService service.BooksService
}

func NewBooksController(service service.BooksService) *BooksController {
	return &BooksController{
		booksService: service,
	}
}

func (controller *BooksController) Create(ctx *gin.Context) {
	log.Info().Msg("create books")
	createBooksRequest := request.CreateBooksRequest{}
	err := ctx.ShouldBindJSON(&createBooksRequest)
	if err != nil {
		panic(err.Error)
	}

	controller.booksService.Create(createBooksRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BooksController) Update(ctx *gin.Context) {
	log.Info().Msg("update books")
	updateBooksRequest := request.UpdateBooksRequest{}
	err := ctx.ShouldBindJSON(&updateBooksRequest)
	if err != nil {
		panic(err.Error)
	}

	bookId := ctx.Param("bookId")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		panic(err.Error)
	}
	updateBooksRequest.Id = id

	controller.booksService.Update(updateBooksRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BooksController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete books")
	bookId := ctx.Param("bookId")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		panic(err.Error)
	}
	controller.booksService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BooksController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid books")
	bookId := ctx.Param("bookId")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("masuk panic")
		panic(err.Error)
	}

	bookResponse := controller.booksService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   bookResponse,
	}

	nilResponse := response.Response{
		Code:   http.StatusNotFound,
		Status: "Not Found",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	if bookResponse.Id == 0 {
		ctx.JSON(http.StatusNotFound, nilResponse)
	} else {
		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *BooksController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll books")
	bookResponse := controller.booksService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   bookResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
