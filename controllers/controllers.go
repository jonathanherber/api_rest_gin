package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jonathanherber/golang_gin/models"
)

func ShowStudents(c *gin.Context) { //c *gin.Context is a convention
	c.JSON(200, models.Students) //return the struct from the controller instantiated in the server.go file
}

func Greetings(c *gin.Context) {
	name := c.Params.ByName("name") //search for the param from the url with gin.Context c by its name
	c.JSON(200, gin.H{
		"API: ": "Paramater value " + name,
	})
}
