package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonathanherber/golang_gin/database"
	"github.com/jonathanherber/golang_gin/models"
)

func ShowStudents(c *gin.Context) { //c *gin.Context is a convention
	var students []models.Student
	database.DB.Find(&students) //find all students in the database
	c.JSON(200, students)
}

func Greetings(c *gin.Context) {
	name := c.Params.ByName("name") //search for the param from the url with gin.Context c by its name
	c.JSON(200, gin.H{
		"API: ": "Paramater value " + name,
	})
}
func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil { //ShouldBindJSON get the body of the request format JSON
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}
func GetStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")     //get the id from the url
	database.DB.First(&student, id) //search for the first match for the id in the student
	if student.ID == 0 {            //validate the id
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student does not exist"})
		return
	}

	c.JSON(http.StatusOK, student)
}
func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id) //deleted the student by the id
	if student.ID == 0 {             //validate
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Student deleted"})
}
func EditStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")                        //get the id from the url
	database.DB.First(&student, id)                    //try to find a row with this id
	if err := c.ShouldBindJSON(&student); err != nil { //if the id does not exist
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //print error
		return
	}
	database.DB.Model(&student).UpdateColumns(student) //update in the database
	c.JSON(http.StatusOK, student)                     //print the student
}
