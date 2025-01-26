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
	"strings"
)

func CheckUserHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("quiz_user")
	if err != nil {
		// No user found, trigger dialog
		w.Header().Set("HX-Trigger", "showUsernameDialog")
		return
	}
	// Use the cookie value to verify the user exists in database
	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM quiz_users WHERE name = ?)", cookie.Value).Scan(&exists)
	if err != nil || !exists {
		// User not found in database or error occurred
		w.Header().Set("HX-Trigger", "showUsernameDialog")
		return
	}
	// User exists, start quiz
	QuizQuestionHandler(w, r)
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Modified query - remove the LAST_INSERT_ID() part since we don't need it
	_, err := database.DB.Exec(`
        INSERT INTO quiz_users (name) 
        VALUES (?) 
        ON DUPLICATE KEY UPDATE name=name`,
		username)

	if err != nil {
		http.Error(w, "Unable to register user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "quiz_user",
		Value:    username,
		MaxAge:   86400, // 1 day
		Path:     "/",
		HttpOnly: true,
	})

	// Hide dialog and start quiz
	w.Header().Set("HX-Trigger", "hideUsernameDialog")
	QuizQuestionHandler(w, r)
}

func GetLeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
        SELECT name, total_score 
        FROM quiz_users 
        ORDER BY total_score DESC 
        LIMIT 5
    `)
	if err != nil {
		http.Error(w, "Unable to fetch leaderboard", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var leaderboardHTML strings.Builder
	rank := 1
	for rows.Next() {
		var name string
		var score int
		if err := rows.Scan(&name, &score); err != nil {
			continue
		}
		leaderboardHTML.WriteString(fmt.Sprintf(`
            <div class="flex justify-between items-center p-2 bg-gray-50 rounded">
                <span class="font-medium">%d. %s</span>
                <span class="text-gray-600">%d points</span>
            </div>
        `, rank, name, score))
		rank++
	}

	fmt.Fprint(w, leaderboardHTML.String())
}

func QuizQuestionHandler(w http.ResponseWriter, r *http.Request) {
	var question models.QuizQuestion
	var optionsJSON []byte

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
                class="w-full text-left p-3 rounded border hover:bg-gray-50 transition-colors"
                hx-post="/quiz/answer"
                hx-trigger="click"
                hx-vals='{"questionId": "%d", "answer": "%d"}'
                hx-target="#questionContainer"
                hx-swap="innerHTML"
            >%s</button>
        `, question.ID, i, opt)
	}

	fmt.Fprintf(w, tmpl, question.Question, optionsHTML)
}

func QuizAnswerHandler(w http.ResponseWriter, r *http.Request) {
	questionID := r.FormValue("questionId")
	userAnswer := r.FormValue("answer")

	var correctAnswer int
	err := database.DB.QueryRow("SELECT answer FROM quiz_questions WHERE id = ?", questionID).Scan(&correctAnswer)
	if err != nil {
		http.Error(w, "Unable to verify answer", http.StatusInternalServerError)
		return
	}

	isCorrect := fmt.Sprintf("%d", correctAnswer) == userAnswer

	// Update score if correct
	if isCorrect {
		if cookie, err := r.Cookie("quiz_user"); err == nil {
			_, err = database.DB.Exec(`
                UPDATE quiz_users 
                SET total_score = total_score + 1 
                WHERE name = ?`,
				cookie.Value)
			if err == nil {
				w.Header().Set("HX-Trigger", "refreshLeaderboard")
			}
		}
	}

	// Respond with feedback message and next question
	feedback := map[bool]string{true: "Correct!", false: "Incorrect. Try the next question."}[isCorrect]

	// Generate the next question
	var nextQuestionHTML strings.Builder
	nextQuestionHTML.WriteString(fmt.Sprintf(`
        <div class="p-4 rounded-lg %s mb-4">
            %s
        </div>
    `,
		map[bool]string{true: "bg-green-100 text-green-900", false: "bg-red-100 text-red-900"}[isCorrect],
		feedback,
	))

	// Update the score and attempts counters
	w.Header().Add("HX-Trigger-After-Swap", "updateScore")
	w.Header().Add("HX-Trigger-After-Swap", "updateAttempts")

	// Fetch and append the next question
	currentQuestion := 0
	if session, _ := r.Cookie("current_question"); session != nil {
		fmt.Sscanf(session.Value, "%d", &currentQuestion)
	}

	if currentQuestion < 3 {
		// Increment question number
		http.SetCookie(w, &http.Cookie{
			Name:     "current_question",
			Value:    fmt.Sprintf("%d", currentQuestion+1),
			MaxAge:   86400,
			Path:     "/",
			HttpOnly: true,
		})

		var question models.QuizQuestion
		var optionsJSON []byte

		err := database.DB.QueryRow(`
            SELECT id, question, options, answer
            FROM quiz_questions
            ORDER BY RAND()
            LIMIT 1
        `).Scan(&question.ID, &question.Question, &optionsJSON, &question.Answer)

		if err != nil {
			http.Error(w, "Unable to fetch next question: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(optionsJSON, &question.Options)
		if err != nil {
			http.Error(w, "Invalid options format: "+err.Error(), http.StatusInternalServerError)
			return
		}

		nextQuestionHTML.WriteString(`
            <div class="space-y-4">
                <p class="text-lg font-semibold mb-4">` + question.Question + `</p>
                <div class="space-y-2">
        `)

		for i, opt := range question.Options {
			nextQuestionHTML.WriteString(fmt.Sprintf(`
                <button
                    class="w-full text-left p-3 rounded border hover:bg-gray-50 transition-colors"
                    hx-post="/quiz/answer"
                    hx-trigger="click"
                    hx-vals='{"questionId": "%d", "answer": "%d"}'
                    hx-target="#questionContainer"
                    hx-swap="innerHTML"
                >%s</button>
            `, question.ID, i, opt))
		}

		nextQuestionHTML.WriteString(`
                </div>
            </div>
        `)
	} else {
		// Quiz is complete
		nextQuestionHTML.WriteString(`
            <div class="p-6 rounded-lg shadow-md bg-green-100 text-green-900">
                <h2 class="text-2xl font-bold mb-4">Congrats!</h2>
                <p>You have completed the quiz!</p>
            </div>
        `)
	}

	fmt.Fprint(w, nextQuestionHTML.String())
}

func QuizSectionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/layouts/base.html",
		"templates/partials/sidebar.html",
		"templates/pages/quiz.html",
		"templates/components/quiz_dialog.html",
	)
	if err != nil {
		http.Error(w, "Unable to load template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, "Unable to render template: "+err.Error(), http.StatusInternalServerError)
	}
}
