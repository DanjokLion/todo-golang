package service

import (
	"github.com/DanjokLion/todo-go"
	"github.com/DanjokLion/todo-go/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetByID(userId, listId)
	if err != nil {
		return 0, nil
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

// func (s *TodoItemService) GetByID(userId int, listId int) (todo.TodoList, error) {
// 	return s.repo.GetByID(userId, listId)
// } 

// func (s *TodoItemService) Delete(userId int, listId int) error {
// 	return s.repo.Delete(userId, listId)
// }

// func (s *TodoItemService) Update(userId int, listId int, input todo.UpdateListInput) error {
// 	if err := input.Validate(); err != nil {
// 		return err
// 	}

// 	return s.repo.Update(userId, listId, input)
// }