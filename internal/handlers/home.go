package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Get all template files
	templates := []string{
		"templates/layouts/base.html",
		"templates/partials/sidebar.html",
		"templates/pages/index.html",
	}

	// Parse all templates
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		log.Printf("Template parsing error: %v", err)
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	// Check if the request is from HTMX
	if r.Header.Get("HX-Request") == "true" {
		// Render only the content block
		err = tmpl.ExecuteTemplate(w, "content", nil)
		if err != nil {
			log.Printf("Template execution error: %v", err)
			http.Error(w, "Unable to render template", http.StatusInternalServerError)
		}
		return
	}

	// Otherwise, render the full base template
	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func WorkingOnHandler(w http.ResponseWriter, r *http.Request) {
	currentWork := "I create this portfolio using Go, HTMX, Tailwind CSS, and MySQL!"
	w.Write([]byte(currentWork))
}
