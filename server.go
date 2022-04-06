package main

import (
	"fmt"

	"github.com/jonathanherber/golang_gin/database"
	"github.com/jonathanherber/golang_gin/routes"
)

func main() {
	database.ConnectToDatabase() //func to connect to database and execute the AutoMigrate
	routes.HandleRequests()
	fmt.Println("Server starting...")
}
