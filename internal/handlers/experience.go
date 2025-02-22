// internal/handlers/experience.go
package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Experience struct {
	ID           int
	Role         string
	Company      string
	Duration     string
	Description  string
	Skills       []string
	Achievements []string
}

type PageData struct {
	Title       string
	Experiences []Experience
}

func ExperienceHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/pages/experience.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title: "Where I Used to Work",
		Experiences: []Experience{
			{
				ID:          1,
				Role:        "Software Developer",
				Company:     "ABC Co.",
				Duration:    "2020-2023",
				Description: "Developed scalable web applications and collaborated with cross-functional teams.",
				Skills:      []string{"Go", "HTMX", "Tailwind"},
				Achievements: []string{
					"Improved application performance by 40%",
					"Led team of 5 developers",
				},
			},
			// Add more experiences here
		},
	}

	tmpl.Execute(w, data)
}

// New handler for experience details
func ExperienceDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse experience ID from URL
	idStr := strings.TrimPrefix(r.URL.Path, "/api/experience/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// In a real app, you'd fetch this from a database
	experience := Experience{
		ID:          id,
		Role:        "Software Developer",
		Company:     "ABC Co.",
		Duration:    "2020-2023",
		Description: "Developed scalable web applications and collaborated with cross-functional teams.",
		Skills:      []string{"Go", "HTMX", "Tailwind"},
		Achievements: []string{
			"Improved application performance by 40%",
			"Led team of 5 developers",
		},
	}

	tmpl, err := template.ParseFiles("templates/partials/experience-details.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, experience)
}
