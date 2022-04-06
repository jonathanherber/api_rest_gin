package database

import (
	"log"

	"github.com/jonathanherber/golang_gin/models"
	"gorm.io/driver/postgres" //postgres driver
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB //DB pointing to value of gorm
	err error
)

func ConnectToDatabase() {
	ConnectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable" //db credentials from docker-compose
	DB, err = gorm.Open(postgres.Open(ConnectionString))                                               //opening connection
	if err != nil {
		log.Panic("Error opening connection")
	}
	DB.AutoMigrate(&models.Student{}) //create a table following the model in the struct, AutoMigrate has 1 parameter: memory address of the struct, instantiating
}
