package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"learning/cmd/server_crud/logic"
	"net/http"
	"time"
)

const todoPath string = "static/todo.json"

type server struct {
	todoService *logic.TodoService
}

func (s *server) run() {
	fmt.Println("Server is listening...", time.Now())
	http.ListenAndServe(":8181", nil)
}

func (s *server) registerHandler() {
	router := mux.NewRouter()
	todoRouter := router.PathPrefix("/todo").Subrouter()

	todoRouter.HandleFunc("", s.getAll).Methods("GET")
	todoRouter.HandleFunc("/{id}", s.getById).Methods("GET")
	todoRouter.HandleFunc("", s.create).Methods("POST")

	todoRouter.HandleFunc("/{id}", s.editById).Methods("PUT")
	todoRouter.HandleFunc("/{id}", s.deleteById).Methods("DELETE")

	http.Handle("/", router)

}

func main() {

	r := logic.InMemoryTodoRepository{}
	r.InitData()

	s := &server{todoService: &logic.TodoService{Repository: &r}}
	s.registerHandler()
	s.run()
}
