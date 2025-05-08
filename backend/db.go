package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(dataSource string) {
	var err error
	db, err = sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func GetTasks() ([]Task, error) {
	rows, err := db.Query("SELECT id, text FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Text); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func AddTask(text string) error {
	_, err := db.Exec("INSERT INTO tasks (text) VALUES ($1)", text)
	return err
}
