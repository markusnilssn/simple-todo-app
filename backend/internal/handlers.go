package internal

import "net/http"

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

func HandleByID(writer http.ResponseWriter, response *http.Request) {}

func getItems(w http.ResponseWriter)                         {}
func createItem(writer http.ResponseWriter, r *http.Request) {}
func deleteTodo(w http.ResponseWriter, id int)               {}
