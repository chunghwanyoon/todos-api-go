package dto

import (
	"github.com/volatiletech/null/v8"
	"todos-api-go/models"
)

type TodoResponse struct {
	Id        int         `json:"id"`
	Title     string      `json:"title"`
	Content   null.String `json:"content"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
}

type TodoListResponse struct {
	todos []*TodoResponse
}

func NewTodoResponse(todo *models.Todo) *TodoResponse {
	return &TodoResponse{
		Id:        todo.ID,
		Title:     todo.Title,
		Content:   todo.Content,
		CreatedAt: todo.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: todo.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func NewTodoListResponse(todos []*models.Todo) *TodoListResponse {
	var responses []*TodoResponse
	for _, todo := range todos {
		responses = append(responses, NewTodoResponse(todo))
	}
	return &TodoListResponse{
		todos: responses,
	}
}
