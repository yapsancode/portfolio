// internal/handlers/projects.go
package handlers

import (
	"html/template"
	"net/http"
	"portfolio/internal/models"
	"strconv"
	"strings"
)

type ProjectsPageData struct {
	Title    string
	Projects []models.Project
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/pages/projects.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Fetch projects from database
	projects, err := models.GetAllProjects()
	if err != nil {
		http.Error(w, "Database Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := ProjectsPageData{
		Title:    "My Projects",
		Projects: projects,
	}

	tmpl.Execute(w, data)
}

// Handler for project details
func ProjectDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse project ID from URL
	idStr := strings.TrimPrefix(r.URL.Path, "/api/project/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Fetch from database
	project, err := models.GetProjectByID(id)
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/partials/project-details.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, project)
}
