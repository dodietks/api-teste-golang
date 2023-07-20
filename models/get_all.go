package models

import (
	"teste.com/apitestgo/db"
)

func GetAll() ([]Todo, error) {
	var todos []Todo

	conn, err := db.OpenConnection()
	if err != nil {
		return todos, err
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return todos, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}
	if err = rows.Err(); err != nil {
		return todos, err
	}
	return todos, nil
}
