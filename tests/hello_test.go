package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lakhinsu/gin-boilerplate/app"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
	zerolog.SetGlobalLevel(zerolog.Disabled)
}

func TestHello(t *testing.T) {
	router := app.SetupApp()
	w := httptest.NewRecorder()

	firstname := "foo"
	lastname := "bar"

	b := []byte(fmt.Sprintf(`{"firstname": "%s", "lastname":"%s"}`, firstname, lastname))

	reader := bytes.NewReader(b)

	req, _ := http.NewRequest("POST", "/v1/hello", reader)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	body, _ := io.ReadAll(w.Result().Body)
	var f interface{}
	json.Unmarshal(body, &f)

	myMap := f.(map[string]interface{})
	// validate response code
	assert.Equal(t, http.StatusOK, w.Code)
	// Validate response message
	assert.Equal(t, fmt.Sprintf("Hello %s %s", firstname, lastname), myMap["message"])

}

func TestInvalidPayload(t *testing.T) {
	router := app.SetupApp()
	w := httptest.NewRecorder()

	firstname := "foo"
	lastname := "bar"

	// Invalid payload
	b := []byte(fmt.Sprintf(`{"first": "%s", "lastname":"%s"}`, firstname, lastname))

	reader := bytes.NewReader(b)

	req, _ := http.NewRequest("POST", "/v1/hello", reader)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	// Validate response code
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
