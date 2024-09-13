package main

import (
	"fmt"
	"net/http"
)

func TodoServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"title": "Some title", "description": "Description of todo in details"}`)
}
