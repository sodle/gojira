package main

import "github.com/sodle/gojira/internal/db"

func main() {
	db := db.InitDb()
	defer db.Close()
}
