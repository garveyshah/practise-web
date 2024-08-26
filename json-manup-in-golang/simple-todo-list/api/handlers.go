package api

import (
	"html/template"
	"net/http"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	var err error

	switch r.URL.Path {
	case "/":
		tmpl = template.Must(template.ParseFiles("../views/template/home.html"))
		err = tmpl.Execute(w, nil) // No data needed for the welcome page
	// case "/create":
	//	tmpl = template.Must(template.ParseFiles("../views/template/home.html"))
	//	err = tmpl.Execute(w, )
	default:
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
