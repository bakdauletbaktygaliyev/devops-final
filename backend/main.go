package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dataSource := "postgres://" + dbUser + ":" + dbPass + "@" + dbHost + "/" + dbName + "?sslmode=disable"
	InitDB(dataSource)

	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			tasks, err := GetTasks()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(tasks)

		case http.MethodPost:
			var t Task
			err := json.NewDecoder(r.Body).Decode(&t)
			if err != nil || t.Text == "" {
				http.Error(w, "invalid input", http.StatusBadRequest)
				return
			}
			if err := AddTask(t.Text); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
