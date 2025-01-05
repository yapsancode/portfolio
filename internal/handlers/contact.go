// internal/handlers/contact.go
package handlers

import (
	"net/http"
	"portfolio/internal/database"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(
		"INSERT INTO contacts (name, email, message) VALUES (?, ?, ?)",
		r.FormValue("name"),
		r.FormValue("email"),
		r.FormValue("message"),
	)
	if err != nil {
		http.Error(w, "Unable to save contact message", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Thank you for your message!"))
}

func AMAHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(
		"INSERT INTO ama_questions (question) VALUES (?)",
		r.FormValue("question"),
	)
	if err != nil {
		http.Error(w, "Unable to save question", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Your question has been submitted!"))
}
