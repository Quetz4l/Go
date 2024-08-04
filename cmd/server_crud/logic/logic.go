package logic

type TodoService struct {
	Repository ITodoRepository
}

func (s *TodoService) GetAll() ([]*Todo, error) {
	return s.Repository.GetAll()
}

func (s *TodoService) GetById(id string) (*Todo, error) {
	return s.Repository.Get(id)
}

func (s *TodoService) Create(newTodo *Todo) (*Todo, error) {
	return s.Repository.Create(newTodo)
}

func (s *TodoService) EditById(id string, todo *Todo) (*Todo, error) {
	return s.Repository.Edit(id, todo)
}

func (s *TodoService) DeleteById(id string) error {
	return s.Repository.Delete(id)
}
