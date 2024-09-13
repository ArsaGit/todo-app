package main

import (
	"fmt"
	"net/http"
	"strings"
)

func TodoServer(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")

	todo := GetTodo(id)

	fmt.Fprint(w, todo)
}

func GetTodo(id string) string {
	if id == "1" {
		return `{"title": "Название задачи", "description": "Описание задачи в деталях"}`
	}

	if id == "2" {
		return `{"title": "Название задачи с id2", "description": "Описание задачи в деталях todo2"}`
	}

	return ""
}
