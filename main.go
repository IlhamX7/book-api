package main

import (
	"book-api/config"
	"book-api/controller"
	"book-api/model"
	"book-api/repository"
	"book-api/router"
	"book-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server!")

	// Database
	db := config.DatabaseConnection()
	validate := validator.New()
	db.Table("books").AutoMigrate(&model.Books{})

	// Repository
	booksRepository := repository.NewBooksRepositoryImpl(db)

	// Service
	booksService := service.NewBooksServiceImpl(booksRepository, validate)

	// Controller
	booksController := controller.NewBooksController(booksService)

	// Router
	routes := router.NewRouter(booksController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err.Error != nil {
		panic(err.Error)
	}
}
