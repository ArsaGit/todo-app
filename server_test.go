package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETTodos(t *testing.T) {
	t.Run("возвращает todo c id=1", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/todos", nil)
		response := httptest.NewRecorder()

		TodoServer(response, request)

		got := response.Body.String()
		want := `{"title": "Some title", "description": "Description of todo in details"}`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
