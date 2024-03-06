package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbName   = "latihan_db"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println("result ", db)
	if err != nil {
		panic(err.Error)
	}

	return db
}
