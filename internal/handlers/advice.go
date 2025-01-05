// internal/handlers/advice.go
package handlers

import (
	"net/http"
	"portfolio/internal/database"
)

func RandomAdviceHandler(w http.ResponseWriter, r *http.Request) {
	var advice string
	err := database.DB.QueryRow("SELECT advice FROM random_advice ORDER BY RAND() LIMIT 1").Scan(&advice)
	if err != nil {
		http.Error(w, "Unable to fetch advice", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(advice))
}
