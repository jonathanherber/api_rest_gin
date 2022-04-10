package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jonathanherber/golang_gin/controllers"
)

func HandleRequests() { //function that creates the server, specify the http verb, the route and the controllers to that route
	server := gin.Default() //default configurations
	server.LoadHTMLGlob("templates/*")
	server.Static("/assets", "./assets")
	server.NoRoute(controllers.NotFoundRoute)
	server.GET("/students", controllers.ShowStudents)
	server.GET("/:name", controllers.Greetings)                   //getting the parameter with : and saving into the variable name
	server.POST("/students", controllers.CreateStudent)           //create a student
	server.GET("/students/:id", controllers.GetStudentById)       //get the student by id
	server.DELETE("/students/:id", controllers.DeleteStudent)     //delete student by id
	server.PATCH("/students/:id", controllers.EditStudentById)    //edit student by id
	server.GET("/students/cpf/:cpf", controllers.GetStudentByCPF) //get the student by cpf
	server.GET("/index", controllers.ShowIndexPage)
	server.Run(":8080") //default port
}
