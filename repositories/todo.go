package repositories

import (
	"context"
	"database/sql"
	"todos-api-go/models"
)

type TodoRepository struct {
	*sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		DB: db,
	}
}

func (r *TodoRepository) FindTodos(ctx context.Context) (todos []*models.Todo, err error) {
	todos, err = models.Todos().All(ctx, r)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) FindTodoById(ctx context.Context, todoId int) (todo *models.Todo, err error) {
	todo, err = models.Todos(models.TodoWhere.ID.EQ(todoId)).One(ctx, r)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
