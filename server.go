package main

import (
	"fmt"

	"github.com/jonathanherber/golang_gin/database"
	"github.com/jonathanherber/golang_gin/models"
	"github.com/jonathanherber/golang_gin/routes"
)

func main() {
	database.ConnectToDatabase() //func to connect to database and execute the AutoMigrate
	models.Students = []models.Student{
		{Name: "Jonathan", CPF: "123456789", RG: "123123123"},
		{Name: "Daniel", CPF: "987654321", RG: "321321321"},
	}
	routes.HandleRequests()
	fmt.Println("Server starting...")
}
