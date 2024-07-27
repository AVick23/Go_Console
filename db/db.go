package db

import (
	"database/sql"
	"log"
	"todo_app/config"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase(cfg config.Config) {
	var err error
	DB, err = sql.Open("sqlite3", cfg.Database.Path)
	if err != nil {
		log.Fatalf("Ошибка открытия базы данных: %v", err)
	}

	createTable()
}

func createTable() {
	query := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        task TEXT NOT NULL
    );`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы: %v", err)
	}
}
