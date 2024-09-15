package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Todo struct {
	Id          string
	Title       string
	Description string
}

func TodoServer(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")

	todo := GetTodo(id)
	json.NewEncoder(w).Encode(todo)
}

func GetTodo(id string) Todo {
	if id == "1" {
		return Todo{"1", "Название задачи", "Описание задачи в деталях"}
	}

	if id == "2" {
		return Todo{"2", "Название задачи с id2", "Описание задачи в деталях todo2"}
	}

	return Todo{}
}
