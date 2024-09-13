package main

import (
	"fmt"
	"net/http"
	"strings"
)

func TodoServer(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")

	if id == "1" {
		fmt.Fprintf(w, `{"title": "Название задачи", "description": "Описание задачи в деталях"}`)
	}

	if id == "2" {
		fmt.Fprintf(w, `{"title": "Название задачи с id2", "description": "Описание задачи в деталях todo2"}`)
	}
}
