package router

import (
	"net/http"
	"html/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	generateHtml(w, nil, "index")

}

func Documentary(w http.ResponseWriter, r *http.Request) {
	generateHtml(w, nil, "documentary")
}

// Uses layout.html to generate a shared html layout for all views
func generateHtml(w http.ResponseWriter, data interface{}, content string) {
	files := []string{
		"templates/layout.html",
		"templates/nav.html",
		"templates/" + content + ".html",
	}

	parsed := template.Must(template.ParseFiles(files...))
	parsed.ExecuteTemplate(w, "layout", data)
}
