package todo

import (
	"log"
	"todo_app/db"
)

type Todo struct {
	ID   int
	Task string
}

func AddTask(task string) {
	query := `INSERT INTO todos (task) VALUES (?)`
	_, err := db.DB.Exec(query, task)
	if err != nil {
		log.Fatalf("Ошибка добавления задания: %v", err)
	}
}

func GetTasks() []Todo {
	rows, err := db.DB.Query(`SELECT id, task FROM todos`)
	if err != nil {
		log.Fatalf("Ошибка получения заданий: %v", err)
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Task); err != nil {
			log.Fatalf("Ошибка сканирования строки: %v", err)
		}
		todos = append(todos, todo)
	}

	return todos
}

func DeleteTask(id int) {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatalf("Ошибка начала транзакции: %v", err)
	}

	query := `DELETE FROM todos WHERE id = ?`
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		log.Fatalf("Ошибка удаления задания: %v", err)
	}

	updateQuery := `
		UPDATE todos
		SET id = id - 1
		WHERE id > ?
	`
	_, err = tx.Exec(updateQuery, id)
	if err != nil {
		tx.Rollback()
		log.Fatalf("Ошибка обновления ID: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalf("Ошибка подтверждения транзакции: %v", err)
	}
}

func DeleteAllTasks() {
	query := `DELETE FROM todos`
	_, err := db.DB.Exec(query)
	if err != nil {
		log.Fatalf("Ошибка удаления всех заданий: %v", err)
	}

	resetQuery := `DELETE FROM sqlite_sequence WHERE name='todos'`
	_, err = db.DB.Exec(resetQuery)
	if err != nil {
		log.Fatalf("Ошибка сброса автоинкремента: %v", err)
	}
}
