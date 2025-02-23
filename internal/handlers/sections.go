package handlers

import (
	"html/template"
	"net/http"
)

// ProjectsHandler dynamically parses and renders the "projects.html" template
func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/pages/projects.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title: "Projects",
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
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

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
