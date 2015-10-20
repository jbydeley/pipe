package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler(w, req)
	if !strings.Contains(w.Body.String(), "Hello, World!") {
		t.Error(w.Body.String())
	}
}
