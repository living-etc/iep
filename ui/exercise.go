package ui

import (
	"context"
	"database/sql"

	"github.com/charmbracelet/log"
)

type Exercise struct {
	Id, title, description string
}

func NewExercise(attributes map[string]string) Exercise {
	return Exercise{
		Id:          attributes["Id"],
		title:       attributes["title"],
		description: attributes["description"],
	}
}

func (i Exercise) Title() string { return i.title }

func (i Exercise) FilterValue() string { return i.title }

func (i Exercise) Description() string { return i.description }

func (i Exercise) Tests(conn *sql.DB, logger *log.Logger) []Test {
	ctx := context.Background()

	rows, err := conn.QueryContext(
		ctx,
		"SELECT * FROM tests WHERE exercise_id = ?;",
		i.Id,
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer rows.Close()

	var tests []Test
	for rows.Next() {
		var e Test
		if err := rows.Scan(
			&e.Id,
			&e.Name,
			&e.ExerciseId,
			&e.ResourceType,
			&e.ResourceName,
			&e.ResourceAttribute,
			&e.ResourceAttributeValue,
			&e.Negation,
			&e.Result,
		); err != nil {
			logger.Fatal(err)
		}
		tests = append(tests, e)
	}

	return tests
}
