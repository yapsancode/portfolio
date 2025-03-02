// internal/database/mysql.go
package database

import (
	"database/sql"
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
	return nil
}

func createTables() error {
	queries := []string{
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
		`CREATE TABLE IF NOT EXISTS experiences (
            id INT AUTO_INCREMENT PRIMARY KEY,
            role VARCHAR(255) NOT NULL,
            company VARCHAR(255) NOT NULL,
            duration VARCHAR(50) NOT NULL,
            description TEXT NOT NULL,
            skills JSON NOT NULL,
            achievements JSON NOT NULL
        )`,
	}

	for _, query := range queries {
		if _, err := DB.Exec(query); err != nil {
			return err
		}
	}

	return nil
}
