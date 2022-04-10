package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jonathanherber/golang_gin/controllers"
	"github.com/stretchr/testify/assert"
)

func RouteTestSetup() *gin.Engine {
	routes := gin.Default()
	return routes
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
