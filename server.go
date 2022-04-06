package main

import (
	"github.com/gin-gonic/gin"
)

func ShowStudents(c *gin.Context) { //controller
	//c *gin.Context is a convention
	c.JSON(200, gin.H{ //H is a shortcut for map[string]interface{}
		"id":   "1",
		"name": "Jonathan",
	})
}

func main() { //function that creates the server, specify the http verb, the route and the controllers to that route
	server := gin.Default() //default configurations
	server.GET("/students", ShowStudents)
	server.Run(":8080") //default port
}
