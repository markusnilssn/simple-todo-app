package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

var todos []Todo = []Todo{}

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
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	deleteTodo(writer, id)
}

func getItems(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(todos) // write :))
}

func createItem(writer http.ResponseWriter, response *http.Request) {
	var newTodo Todo
	json.NewDecoder(response.Body).Decode(&newTodo)
	todos = append(todos, newTodo)
	writer.WriteHeader(201)
}

func deleteTodo(w http.ResponseWriter, id int) {
	removeIndex := -1
	for index, _ := range todos {
		if index == id {
			removeIndex = index
		}
	}

	if removeIndex > -1 {
		todos = slices.Delete(todos, removeIndex, 1)
		fmt.Printf("removed index %d\n", removeIndex)
	} else {
		fmt.Printf("failed to remove index\n")
	}
}
