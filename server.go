package main

import (
	"fmt"
	"net/http"
)

func TodoServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"title": "Название задачи", "description": "Описание задачи в деталях"}`)
}
