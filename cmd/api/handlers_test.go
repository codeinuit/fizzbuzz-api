package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codeinuit/fizzbuzz-api/pkg/database/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetHealth(t *testing.T) {
	fb, err := setupRouter()
	assert.Nil(t, err)
	initRoutes(fb, mock.NewDatabaseMock())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	fb.engine.ServeHTTP(w, req)

	var want gin.H = gin.H{"message": "OK"}
	var got gin.H
	err = json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, want, got)
}

func TestGetFizzBuzzOK(t *testing.T) {
	fb, err := setupRouter()
	assert.Nil(t, err)
	initRoutes(fb, mock.NewDatabaseMock())

	data := getFizzBuzzBody{
		Int1:    3,
		Int2:    5,
		Int3:    3,
		String1: "fizz",
		String2: "buzz",
	}
	var b bytes.Buffer

	err = json.NewEncoder(&b).Encode(data)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/fizzbuzz", &b)
	fb.engine.ServeHTTP(w, req)

	var want string = "1,2,fizz"

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, want, w.Body.String())
}

func TestGetFizzBuzzMissingParameter(t *testing.T) {
	fb, err := setupRouter()
	assert.Nil(t, err)
	initRoutes(fb, mock.NewDatabaseMock())

	data := getFizzBuzzBody{
		Int1:    3,
		Int2:    5,
		String1: "fizz",
		String2: "buzz",
	}
	var b bytes.Buffer

	err = json.NewEncoder(&b).Encode(data)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/fizzbuzz", &b)
	fb.engine.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestGetFizzBuzzWrongParameter(t *testing.T) {
	fb, err := setupRouter()
	assert.Nil(t, err)
	initRoutes(fb, mock.NewDatabaseMock())

	data := getFizzBuzzBody{
		Int1:    3,
		Int2:    5,
		Int3:    -23,
		String1: "fizz",
		String2: "buzz",
	}
	var b bytes.Buffer

	err = json.NewEncoder(&b).Encode(data)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/fizzbuzz", &b)
	fb.engine.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
