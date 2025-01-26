// internal/database/mysql.go
package database

import (
	"database/sql"
	"encoding/json"
	"portfolio/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Initialize() error {
	dbConfig := config.GetDBConfig()
	var err error
	DB, err = sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	if err = createTables(); err != nil {
		return err
	}
	if err = addSampleQuizQuestions(); err != nil {
		return err
	}
	return nil
}

func createTables() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS quiz_questions (
            id INT AUTO_INCREMENT PRIMARY KEY,
            question TEXT NOT NULL,
            options JSON NOT NULL,
            answer INT NOT NULL
        )`,
		`CREATE TABLE IF NOT EXISTS contacts (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            email VARCHAR(255) NOT NULL,
            message TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )`,
		`CREATE TABLE IF NOT EXISTS ama_questions (
            id INT AUTO_INCREMENT PRIMARY KEY,
            question TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )`,
		`CREATE TABLE IF NOT EXISTS random_advice (
			id INT AUTO_INCREMENT PRIMARY KEY,
			advice TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS quiz_users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL UNIQUE,
            total_score INT DEFAULT 0,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )`,
	}

	for _, query := range queries {
		if _, err := DB.Exec(query); err != nil {
			return err
		}
	}

	return nil
}

func addSampleQuizQuestions() error {
	// Check if we already have questions
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM quiz_questions").Scan(&count)
	if err != nil {
		return err
	}

	// If we already have questions, skip
	if count > 0 {
		return nil
	}

	// Sample questions
	questions := []struct {
		question string
		options  []string
		answer   int
	}{
		{
			question: "What is the correct syntax for declaring a variable in Go?",
			options:  []string{"var x = 5", "let x = 5", "x := 5", "All of the above"},
			answer:   2, // Index of "x := 5"
		},
		{
			question: "Which of these is NOT a valid Go data type?",
			options:  []string{"int", "float", "string", "boolean"},
			answer:   1, // Index of "float" (Go uses float32/float64)
		},
	}

	// Insert questions
	for _, q := range questions {
		optionsJSON, err := json.Marshal(q.options)
		if err != nil {
			return err
		}

		_, err = DB.Exec(`
            INSERT INTO quiz_questions (question, options, answer)
            VALUES (?, ?, ?)
        `, q.question, optionsJSON, q.answer)

		if err != nil {
			return err
		}
	}

	return nil
}
