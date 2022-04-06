package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonathanherber/golang_gin/database"
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
func CreateStudent(c *gin.Context) {
	var aluno models.Student
	if err := c.ShouldBindJSON(&aluno); err != nil { //ShouldBindJSON get the body of the request format JSON
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}
