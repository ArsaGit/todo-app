package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGETTodos(t *testing.T) {
	t.Run("возвращает todo c id=1", func(t *testing.T) {
		wantedTodo := Todo{"1", "Название задачи", "Описание задачи в деталях"}

		request, _ := http.NewRequest(http.MethodGet, "/todos/1", nil)
		response := httptest.NewRecorder()

		TodoServer(response, request)

		var got Todo
		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Ошибка при парсинге ответа в Todo, %s", err)
		}

		if response.Code != http.StatusOK {
			t.Fatalf("got %d, want %d", response.Code, http.StatusOK)
		}

		if !reflect.DeepEqual(got, wantedTodo) {
			t.Errorf("got %v, want %v", got, wantedTodo)
		}
	})

	t.Run("возвращает todo c id=2", func(t *testing.T) {
		wantedTodo := Todo{"2", "Название задачи с id2", "Описание задачи в деталях todo2"}

		request, _ := http.NewRequest(http.MethodGet, "/todos/2", nil)
		response := httptest.NewRecorder()

		TodoServer(response, request)

		var got Todo
		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Ошибка при парсинге ответа в Todo, %s", err)
		}

		if response.Code != http.StatusOK {
			t.Fatalf("got %d, want %d", response.Code, http.StatusOK)
		}

		if !reflect.DeepEqual(got, wantedTodo) {
			t.Errorf("got %v, want %v", got, wantedTodo)
		}
	})
}
