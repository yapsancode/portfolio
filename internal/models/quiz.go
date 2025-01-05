package models

type QuizQuestion struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   int      `json:"answer"`
}

type QuizResponse struct {
	Question   string   `json:"question"`
	Options    []string `json:"options"`
	QuestionID int      `json:"questionId"`
}
