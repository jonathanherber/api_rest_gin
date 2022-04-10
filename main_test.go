package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jonathanherber/golang_gin/controllers"
	"github.com/jonathanherber/golang_gin/database"
	"github.com/jonathanherber/golang_gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func RouteTestSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}
func CreateStudent() {
	student := models.Student{
		Name: "Test Student",
		CPF:  "12345678901",
		RG:   "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}
func DeleteStudent() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestVerifyGreetings(t *testing.T) {
	r := RouteTestSetup() //new instance of GIN
	r.GET("/:name", controllers.Greetings)
	request, _ := http.NewRequest("GET", "/Jonathan", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "Should be the same")
	mockResponse := `{"API:":"Paramater value Jonathan"}`
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, mockResponse, string(responseBody))
}
func TestGetStudents(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudent()
	defer DeleteStudent()
	r := RouteTestSetup()
	r.GET("/students", controllers.ShowStudents)
	request, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
	//fmt.Println(response.Body)
}
func TestGetStudentByCpf(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudent()
	defer DeleteStudent()
	r := RouteTestSetup()
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)
	request, _ := http.NewRequest("GET", "/students/cpf/03359110012", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}
func TestGetStudentById(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudent()
	defer DeleteStudent()
	r := RouteTestSetup()
	r.GET("/students/:id", controllers.GetStudentById)
	searchPath := "/students/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("GET", searchPath, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)
	assert.Equal(t, "Test Student", studentMock.Name, "names must be equal")
	assert.Equal(t, "12345678901", studentMock.CPF, "cpf must be equal")
	assert.Equal(t, "123456789", studentMock.RG, "rg must be equal")
	assert.Equal(t, http.StatusOK, response.Code)
}
func TestDeleteStudent(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudent()
	r := RouteTestSetup()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	searchPath := "/students/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("DELETE", searchPath, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}
func TestEditStudent(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudent()
	defer DeleteStudent()
	r := RouteTestSetup()
	r.PATCH("/students/:id", controllers.EditStudentById)
	student := models.Student{Name: "Test Student", CPF: "10987654321", RG: "123456700"}
	jsonValue, _ := json.Marshal(student)
	searchPath := "/students/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("PATCH", searchPath, bytes.NewBuffer(jsonValue))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	var studentMockUpdated models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMockUpdated)
	assert.Equal(t, "10987654321", studentMockUpdated.CPF)
	assert.Equal(t, "123456700", studentMockUpdated.RG)
	assert.Equal(t, "Test Student", studentMockUpdated.Name)
}
