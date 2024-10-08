package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETTodos(t *testing.T) {
	t.Run("возвращает todo c id=1", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/todos/1", nil)
		response := httptest.NewRecorder()

		TodoServer(response, request)

		got := response.Body.String()
		want := `{"title": "Название задачи", "description": "Описание задачи в деталях"}`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("возвращает todo c id=2", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/todos/2", nil)
		response := httptest.NewRecorder()

		TodoServer(response, request)

		got := response.Body.String()
		want := `{"title": "Название задачи с id2", "description": "Описание задачи в деталях todo2"}`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
