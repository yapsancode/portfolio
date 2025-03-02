// internal/handlers/routes.go
package handlers

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", HomeHandler)

	// Main section handlers (for HTMX requests)
	mux.HandleFunc("/experience", ExperienceHandler)
	mux.HandleFunc("/projects", ProjectsHandler)
	mux.HandleFunc("/api/project/", ProjectDetailsHandler)
	mux.HandleFunc("/contact", ContactSectionHandler)
	mux.HandleFunc("/ama", AMASectionHandler)
	mux.HandleFunc("/about", AboutHandler)

	// Form submission handlers
	mux.HandleFunc("/submit-contact", ContactHandler) // For contact form POST
	mux.HandleFunc("/submit-ama", AMAHandler)         // For AMA form POST

	// Other functionality
	mux.HandleFunc("/random-advice", RandomAdviceHandler)
	mux.HandleFunc("/working-on", WorkingOnHandler)
}
