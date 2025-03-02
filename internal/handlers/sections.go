package handlers

import (
	"html/template"
	"net/http"
)

// ContactSectionHandler dynamically parses and renders the "contact_me.html" template
func ContactSectionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/pages/contact_me.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// AMASectionHandler dynamically parses and renders the "ask_me_anything.html" template
func AMASectionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/pages/ask_me_anything.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
