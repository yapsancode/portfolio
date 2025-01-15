// internal/handlers/about.go
package handlers

import (
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/pages/about.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
	}{
		Title: "About Me 🚀",
	}

	tmpl.Execute(w, data)
}
