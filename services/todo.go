package services

import (
	"context"
	"database/sql"
	"todos-api-go/config"
	"todos-api-go/models"
	"todos-api-go/repositories"
)

type TodoService struct {
	repository *repositories.TodoRepository
	config.Settings
}

func NewTodoService(db *sql.DB, settings config.Settings) *TodoService {
	return &TodoService{
		repository: repositories.NewTodoRepository(db),
		Settings:   settings,
	}
}

func (s *TodoService) GetTodos(ctx context.Context) (todos []*models.Todo, err error) {
	// TODO: handle sql errors
	return s.repository.FindTodos(ctx)
}

func (s *TodoService) GetTodoById(ctx context.Context, todoId int) (todo *models.Todo, err error) {
	// TODO: handle sql errors
	return s.repository.FindTodoById(ctx, todoId)
}
