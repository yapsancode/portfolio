// internal/handlers/quiz.go
package handlers

import (
	"fmt"
	"net/http"
	"portfolio/internal/database"
	"portfolio/internal/models"
)

func QuizQuestionHandler(w http.ResponseWriter, r *http.Request) {
	var question models.QuizQuestion
	err := database.DB.QueryRow(`
        SELECT id, question, options, answer 
        FROM quiz_questions 
        ORDER BY RAND() 
        LIMIT 1
    `).Scan(&question.ID, &question.Question, &question.Options, &question.Answer)

	if err != nil {
		http.Error(w, "Unable to fetch questions", http.StatusInternalServerError)
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
	html := `
    <section class="mt-8">
        <h2 class="text-2xl font-bold mb-4">Programming Quiz</h2>
        <div id="quizContent" class="bg-white p-6 rounded-lg shadow-md">
            <div class="mb-4">
                <span class="text-gray-600">Current Score: </span>
                <span id="currentScore" class="font-bold">0</span>
            </div>
            
            <div id="questionContainer" class="mb-6">
                <button 
                    hx-get="/quiz/question" 
                    hx-target="#questionContainer"
                    class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded">
                    Start Quiz
                </button>
            </div>

            <div id="quizStats" class="text-sm text-gray-600">
                <p>Questions Attempted: <span id="questionsAttempted">0</span></p>
                <p>Correct Answers: <span id="correctAnswers">0</span></p>
            </div>
        </div>
    </section>`

	fmt.Fprint(w, html)
}
