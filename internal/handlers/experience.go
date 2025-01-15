// internal/handlers/experience.go
package handlers

import (
	"html/template"
	"net/http"
)

func ExperienceHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/pages/experience.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// You can pass data to template if needed
	data := struct {
		Title       string
		Experiences []struct {
			Role        string
			Description string
		}
	}{
		Title: "Where I used to ",
		Experiences: []struct {
			Role        string
			Description string
		}{
			{
				Role:        "Software Developer at ABC Co.",
				Description: "Developed scalable web applications and collaborated with cross-functional teams.",
			},
			{
				Role:        "Intern at XYZ Inc.",
				Description: "Worked on IoT projects and implemented innovative solutions for client needs.",
			},
		},
	}

	tmpl.Execute(w, data)
}
