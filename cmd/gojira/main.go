package main

import (
	"log"
	"net/http"

	"github.com/sodle/gojira/internal/db"
	"github.com/sodle/gojira/internal/views"
)

func main() {
	db := db.InitDb()
	defer db.Close()

	http.HandleFunc("/", views.ListProjects)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
