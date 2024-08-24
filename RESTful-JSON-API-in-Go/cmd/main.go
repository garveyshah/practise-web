package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../templates/home.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			log.Printf("Template execution failed: %v", err)
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	})

	log.Printf("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
