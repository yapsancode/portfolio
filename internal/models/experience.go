// internal/models/experience.go
package models

import (
	"encoding/json"
	"portfolio/internal/database"
)

type Experience struct {
	ID           int
	Role         string
	Company      string
	Duration     string
	Description  string
	Skills       []string
	Achievements []string
}

func GetAllExperiences() ([]Experience, error) {
	rows, err := database.DB.Query("SELECT id, role, company, duration, description, skills, achievements FROM experiences ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var experiences []Experience
	for rows.Next() {
		var exp Experience
		var skillsJSON, achievementsJSON []byte

		if err := rows.Scan(&exp.ID, &exp.Role, &exp.Company, &exp.Duration, &exp.Description, &skillsJSON, &achievementsJSON); err != nil {
			return nil, err
		}

		// Unmarshal JSON arrays
		if err := json.Unmarshal(skillsJSON, &exp.Skills); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(achievementsJSON, &exp.Achievements); err != nil {
			return nil, err
		}

		experiences = append(experiences, exp)
	}

	return experiences, nil
}

func GetExperienceByID(id int) (Experience, error) {
	var exp Experience
	var skillsJSON, achievementsJSON []byte

	err := database.DB.QueryRow("SELECT id, role, company, duration, description, skills, achievements FROM experiences WHERE id = ?", id).
		Scan(&exp.ID, &exp.Role, &exp.Company, &exp.Duration, &exp.Description, &skillsJSON, &achievementsJSON)

	if err != nil {
		return Experience{}, err
	}

	// Unmarshal JSON arrays
	if err := json.Unmarshal(skillsJSON, &exp.Skills); err != nil {
		return Experience{}, err
	}
	if err := json.Unmarshal(achievementsJSON, &exp.Achievements); err != nil {
		return Experience{}, err
	}

	return exp, nil
}
