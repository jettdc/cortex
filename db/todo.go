package db

import (
	"log"
)

type Todo struct {
	id       int
	message  string
	priority int8
}

func InsertTodo(message string, priority int8) {
	insertSQL := `INSERT INTO todos (message, priority) VALUES (?, ?)`
	_, err := database.Exec(insertSQL, message, priority)
	if err != nil {
		log.Fatal(err)
	}
}

func GetTodos() []Todo {
	getSql := `SELECT id, message, priority FROM todos ORDER BY priority`
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

		// Scan the current row into variables
		if err := rows.Scan(&id, &message, &priority); err != nil {
			log.Fatal(err)
		}

		results = append(results, Todo{
			id:       id,
			message:  message,
			priority: priority,
		})
	}

	return results
}
