package logic

import (
	"strconv"
	"time"
)

type ITodoRepository interface {
	Get(Id string) (todo *Todo, err error)
	GetAll() (todos []*Todo, err error)
	Create(newTodo *Todo) (todo *Todo, err error)
	Delete(Id string) (err error)
	Edit(Id string, updateTodo *Todo) (todo *Todo, err error)
}
type InMemoryTodoRepository struct {
	todos []*Todo
}

func (i *InMemoryTodoRepository) Get(idStr string) (todo *Todo, err error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	return i.todos[id-1], nil
}

func (i *InMemoryTodoRepository) GetAll() (todos []*Todo, err error) {
	return i.todos, nil
}

func (i *InMemoryTodoRepository) Create(newTodo *Todo) (todo *Todo, err error) {
	//todo create todo and find this it in bd
	return newTodo, nil
}

func (i *InMemoryTodoRepository) Edit(Id string, updateTodo *Todo) (todo *Todo, err error) {
	//todo edit todo in json/bd
	return updateTodo, nil
}

func (i *InMemoryTodoRepository) Delete(Id string) (err error) {
	//todo delete from json/db
	return nil
}

func (i *InMemoryTodoRepository) InitData() {
	i.todos = append(i.todos, &Todo{
		Id:        1,
		Task:      "Do something, touch a grass!",
		DueDate:   time.Now(),
		Priority:  "High",
		Completed: false,
	})
}
