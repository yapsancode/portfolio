// internal/handlers/experience.go
package handlers

import (
	"html/template"
	"net/http"
	"portfolio/internal/models"
	"strconv"
	"strings"
)

type PageData struct {
	Title       string
	Experiences []models.Experience
}

func ExperienceHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/pages/experience.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Fetch experiences from database
	experiences, err := models.GetAllExperiences()
	if err != nil {
		http.Error(w, "Database Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:       "Where I Used to Work",
		Experiences: experiences,
	}

	tmpl.Execute(w, data)
}

// Handler for experience details
func ExperienceDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse experience ID from URL
	idStr := strings.TrimPrefix(r.URL.Path, "/api/experience/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Fetch from database
	experience, err := models.GetExperienceByID(id)
	if err != nil {
		http.Error(w, "Experience not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/partials/experience-details.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, experience)
}
