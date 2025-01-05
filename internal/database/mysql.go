package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Initialize() error {
	var err error
	DB, err = sql.Open("mysql", "root:key-1WIFI@tcp(127.0.0.1:3306)/portfolio_db")
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}
	log.Println("Database connected successfully!")

	return createTables()
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
	}

	for _, query := range queries {
		if _, err := DB.Exec(query); err != nil {
			return err
		}
	}

	return nil
}
