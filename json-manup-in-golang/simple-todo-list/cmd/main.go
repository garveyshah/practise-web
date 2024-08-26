package main

import (
	"log"
	"net/http"

	"todolist/api"
)

func main() {
	http.HandlerFunc("/", api.PageHandler)

	// Routtes for the individual sections
	// http.HandlerFunc()

	log.Println("Server running on http://localhost:9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
