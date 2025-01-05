// internal/handlers/routes.go
package handlers

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/contact", ContactHandler)
	mux.HandleFunc("/ama", AMAHandler)
	mux.HandleFunc("/random-advice", RandomAdviceHandler)
	mux.HandleFunc("/working-on", WorkingOnHandler)
	mux.HandleFunc("/quiz/question", QuizQuestionHandler)
	mux.HandleFunc("/quiz/answer", QuizAnswerHandler)
	mux.HandleFunc("/sections/experience", ExperienceHandler)
	mux.HandleFunc("/sections/quiz", QuizSectionHandler)
}
