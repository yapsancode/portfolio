// internal/handlers/quiz.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"portfolio/internal/database"
	"portfolio/internal/models"
)

func QuizQuestionHandler(w http.ResponseWriter, r *http.Request) {
	var question models.QuizQuestion
	var optionsJSON []byte // Temporary variable to hold JSON data

	err := database.DB.QueryRow(`
        SELECT id, question, options, answer
        FROM quiz_questions
        ORDER BY RAND()
        LIMIT 1
    `).Scan(&question.ID, &question.Question, &optionsJSON, &question.Answer)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No questions available", http.StatusNotFound)
		} else {
			http.Error(w, "Unable to fetch questions: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Parse the JSON options into the string slice
	err = json.Unmarshal(optionsJSON, &question.Options)
	if err != nil {
		http.Error(w, "Invalid options format: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := `
        <div class="space-y-4">
            <p class="text-lg font-semibold mb-4">%s</p>
            <div class="space-y-2">
                %s
            </div>
        </div>
    `

	var optionsHTML string
	for i, opt := range question.Options {
		optionsHTML += fmt.Sprintf(`
            <button
                class="w-full text-left p-3 rounded border hover:bg-gray-50"
                hx-post="/quiz/answer"
                hx-vals='{"questionId": %d, "answer": %d}'
                hx-target="#questionContainer"
            >%s</button>
        `, question.ID, i, opt)
	}

	fmt.Fprintf(w, tmpl, question.Question, optionsHTML)
}

func QuizAnswerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	questionID := r.FormValue("questionId")
	userAnswer := r.FormValue("answer")

	var correctAnswer int
	err := database.DB.QueryRow("SELECT answer FROM quiz_questions WHERE id = ?", questionID).Scan(&correctAnswer)
	if err != nil {
		http.Error(w, "Unable to verify answer", http.StatusInternalServerError)
		return
	}

	isCorrect := fmt.Sprintf("%d", correctAnswer) == userAnswer

	w.Header().Set("HX-Trigger", fmt.Sprintf(`{
        "updateScore": {"correct": %v},
        "showFeedback": {"message": "%s", "isCorrect": %v}
    }`, isCorrect,
		map[bool]string{true: "Correct!", false: "Incorrect. Try the next question."}[isCorrect],
		isCorrect))

	QuizQuestionHandler(w, r)
}

func QuizSectionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/pages/quiz.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}
