package main

import (
	"fmt"
	"log"
	"net/http"

	"portfolio/internal/database"
	"portfolio/internal/handlers"
)

func main() {
	// Initialize database
	if err := database.Initialize(); err != nil {
		log.Fatal(err)
	}

	// Set up routes
	mux := http.NewServeMux()

	// Serve static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Register routes
	handlers.RegisterRoutes(mux)

	// Start server
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
