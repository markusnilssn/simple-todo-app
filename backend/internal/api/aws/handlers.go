package aws

import (
	"backend/internal/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var database = "todos"

func Handle(writer http.ResponseWriter, response *http.Request) {
	switch response.Method {
	case http.MethodGet:
		getItems(writer)
	case http.MethodPost:
		createItem(writer, response)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleByID(writer http.ResponseWriter, response *http.Request) {
	if response.Method != http.MethodDelete {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(response.URL.Path, "/todos/")
	id, _ := strconv.Atoi(idStr)

	deleteItem(writer, id)
}

func getItems(writer http.ResponseWriter) {
	query := "SELECT id, title, description, priority, completed FROM todos"

	rows, err := SQLDatabase.QueryContext(context.TODO(), query)
	if err != nil {
		log.Printf("failed to execute sql %s, error: %s", query, err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		var priority int

		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&priority,
			&todo.Completed,
		)
		if err != nil {
			log.Printf("failed to scan row: %s", err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		todo.Priority = models.Priority(priority)
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("row iteration error: %s", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(todos)
}

func createItem(writer http.ResponseWriter, request *http.Request) {
	var newTodo models.Todo

	err := json.NewDecoder(request.Body).Decode(&newTodo)
	if err != nil {
		log.Printf("failed to decode request body: %s", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO todos (title, description, priority, completed)
		VALUES ($1, $2, $3, $4)
	`

	_, err = SQLDatabase.ExecContext(
		context.TODO(),
		query,
		newTodo.Title,
		newTodo.Description,
		int(newTodo.Priority),
		newTodo.Completed,
	)

	if err != nil {
		log.Printf("failed to execute sql %s, error: %s", query, err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func deleteItem(writer http.ResponseWriter, id int) {
	query := "DELETE FROM todos WHERE id = $1"

	_, err := SQLDatabase.ExecContext(
		context.TODO(),
		query,
		id,
	)

	if err != nil {
		log.Printf("failed to execute sql %s, error: %s", query, err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("deleted todo %d\n", id)
	writer.WriteHeader(http.StatusOK)
}

func stringPtr(s string) *string {
	return &s
}
