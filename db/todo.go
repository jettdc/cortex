package db

import (
	"log"
)

type Todo struct {
	Id       int
	Message  string
	Priority int8
	Status   string
}

func InsertTodo(message string, priority int8) {
	insertSQL := `INSERT INTO todos (Message, Priority) VALUES (?, ?)`
	_, err := database.Exec(insertSQL, message, priority)
	if err != nil {
		log.Fatal(err)
	}
}

func GetTodos() []Todo {
	getSql := `SELECT id, message, priority, status FROM todos ORDER BY Priority`
	rows, err := database.Query(getSql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	results := []Todo{}
	for rows.Next() {
		var id int
		var message string
		var priority int8
		var status string

		// Scan the current row into variables
		if err := rows.Scan(&id, &message, &priority, &status); err != nil {
			log.Fatal(err)
		}

		results = append(results, Todo{
			Id:       id,
			Message:  message,
			Priority: priority,
			Status:   status,
		})
	}

	return results
}
