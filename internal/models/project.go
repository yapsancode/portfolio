// internal/models/project.go
package models

import (
	"encoding/json"
	"portfolio/internal/database"
	"time"
)

type Project struct {
	ID           int
	Title        string
	Description  string
	ImageURL     string
	LiveURL      string
	GithubURL    string
	Technologies []string
	Highlights   []string
	CreatedAt    time.Time
}

func GetAllProjects() ([]Project, error) {
	rows, err := database.DB.Query("SELECT id, title, description, image_url, live_url, github_url, technologies, highlights, created_at FROM projects ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var proj Project
		var techJSON, highlightsJSON []byte

		if err := rows.Scan(
			&proj.ID,
			&proj.Title,
			&proj.Description,
			&proj.ImageURL,
			&proj.LiveURL,
			&proj.GithubURL,
			&techJSON,
			&highlightsJSON,
			&proj.CreatedAt,
		); err != nil {
			return nil, err
		}

		// Unmarshal JSON arrays
		if err := json.Unmarshal(techJSON, &proj.Technologies); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(highlightsJSON, &proj.Highlights); err != nil {
			return nil, err
		}

		projects = append(projects, proj)
	}

	return projects, nil
}

func GetProjectByID(id int) (Project, error) {
	var proj Project
	var techJSON, highlightsJSON []byte

	err := database.DB.QueryRow(
		"SELECT id, title, description, image_url, live_url, github_url, technologies, highlights, created_at FROM projects WHERE id = ?",
		id,
	).Scan(
		&proj.ID,
		&proj.Title,
		&proj.Description,
		&proj.ImageURL,
		&proj.LiveURL,
		&proj.GithubURL,
		&techJSON,
		&highlightsJSON,
		&proj.CreatedAt,
	)

	if err != nil {
		return Project{}, err
	}

	// Unmarshal JSON arrays
	if err := json.Unmarshal(techJSON, &proj.Technologies); err != nil {
		return Project{}, err
	}
	if err := json.Unmarshal(highlightsJSON, &proj.Highlights); err != nil {
		return Project{}, err
	}

	return proj, nil
}
