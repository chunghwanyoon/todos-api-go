package controllers

import (
	"context"
	"database/sql"
	"todos-api-go/api/dto"
	"todos-api-go/config"
	"todos-api-go/services"
)

type TodoController struct {
	todoService *services.TodoService
}

func NewTodoController(db *sql.DB, settings config.Settings) *TodoController {
	return &TodoController{
		todoService: services.NewTodoService(db, settings),
	}
}

func (c *TodoController) GetTodos(ctx context.Context) (response *dto.TodoListResponse, err error) {
	todos, err := c.todoService.GetTodos(ctx)
	if err != nil {
		return nil, err
	}
	return dto.NewTodoListResponse(todos), nil
}

func (c *TodoController) GetTodoById(ctx context.Context, todoId int) (response *dto.TodoResponse, err error) {
	todo, err := c.todoService.GetTodoById(ctx, todoId)
	if err != nil {
		return nil, err
	}
	return dto.NewTodoResponse(todo), nil
}
