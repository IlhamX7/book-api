package router

import (
	"book-api/controller"

	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(booksController *controller.BooksController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	booksRouter := baseRouter.Group("/books")
	booksRouter.GET("", booksController.FindAll)
	booksRouter.GET("/:bookId", booksController.FindById)
	booksRouter.POST("", booksController.Create)
	booksRouter.PATCH("/:bookId", booksController.Update)
	booksRouter.DELETE("/:bookId", booksController.Delete)

	return router
}
