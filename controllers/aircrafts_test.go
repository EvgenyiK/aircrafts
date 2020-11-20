package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetResourceById(t *testing.T) {
	router:= Find()

	w:= httptest.NewRecorder()
	req,_:= http.NewRequest("GET","/aircrafts", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}