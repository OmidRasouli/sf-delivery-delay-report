package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"sf-delivery-delay-report/internal/app"
	"sf-delivery-delay-report/internal/db"
	"sf-delivery-delay-report/internal/db/models"
	"testing"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

type APIResponse struct {
	Data  models.DelayReport `json:"data"`
	Msg   string             `json:"msg"`
	Error string             `json:"error"`
}

type AgentResponse struct {
	Data  models.DelayReport `json:"data"`
	Error string             `json:"error"`
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	router = gin.Default()
	dbInstance, err := db.Initialize()
	if err != nil {
		panic("Failed to connect to the database")
	}

	app.Initialize(router, dbInstance)
	exitCode := m.Run()

	os.Exit(exitCode)
}

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	return writer
}
