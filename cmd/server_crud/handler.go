package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"learning/cmd/server_crud/logic"
	"net/http"
	"strconv"
)

func (s *server) getAll(w http.ResponseWriter, r *http.Request) {
	todos, err := s.todoService.GetAll()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	//todosBytes, err := json.Marshal(todos)
	//if err != nil {
	//	w.WriteHeader(500)
	//	return
	//}

	//_, err = w.Write(todosBytes)
	//if err != nil {
	//	w.WriteHeader(500)
	//	return
	//}

}

func (s *server) getById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	todos, err := s.todoService.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func (s *server) create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	time, err := logic.StringToTodoTime(r.FormValue("time"))
	if err != nil {
		http.Error(w, "Unable to parse time", http.StatusBadRequest)
	}

	newTodo := logic.Todo{
		Id:        -1, //todo find first free id
		Task:      r.FormValue("task"),
		DueDate:   time,
		Priority:  r.FormValue("priority"),
		Completed: r.FormValue("completed") == "true",
	}

	todo, err := s.todoService.Create(&newTodo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if todo == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Wrong data to create todo")
		return
	}

	fmt.Fprintf(w, "Create todo successfully")
	fmt.Fprintf(w, "Id: %d\nTask: %s\nDate: %q\nPriority: %s\nCompletled: %t", todo.Id, todo.Task, newTodo.DueDate, newTodo.Priority, newTodo.Completed)
}

func (s *server) editById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "Wrong id to create todo", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Wrong id to create todo", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	time, err := logic.StringToTodoTime(r.FormValue("time"))
	if err != nil {
		http.Error(w, "Unable to parse time", http.StatusBadRequest)
		return
	}

	newTodo := logic.Todo{
		Id:        id,
		Task:      r.FormValue("task"),
		DueDate:   time,
		Priority:  r.FormValue("priority"),
		Completed: r.FormValue("completed") == "true",
	}

	editedTodo, err := s.todoService.EditById(idStr, &newTodo)
	if err != nil {
		http.Error(w, "Unable to update todo", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Edit todo successfully")
	fmt.Fprintf(w, "Id: %d\nTask: %s\nDate: %q\nPriority: %s\nCompletled: %t", editedTodo.Id, editedTodo.Task, editedTodo.DueDate, editedTodo.Priority, editedTodo.Completed)
}

func (s *server) deleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.todoService.DeleteById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
