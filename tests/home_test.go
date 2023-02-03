package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apriliantocecep/ordent-restapi-artikel/config"
	"github.com/apriliantocecep/ordent-restapi-artikel/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func NewServer() *gin.Engine {
	configOptions := config.ConfigOptions{
		Path: "./..",
		Name: ".env.test",
	}
	handler := server.InitializedServer(configOptions)

	return handler
}

func TestHome(t *testing.T) {
	handler := NewServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	handler.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
