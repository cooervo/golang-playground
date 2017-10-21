package router

import (
	"net/http"
	"fmt"
	"html/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	generateHtml(w, nil, "layout", "nav", "index")
}


// Uses layout.html to generate a shared html layout for all views
func generateHtml(w http.ResponseWriter, data interface{}, templatesPaths ...string) {
	var files []string
	for _, t := range templatesPaths {
		files = append(files, fmt.Sprintf("templates/%s.html", t))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}
