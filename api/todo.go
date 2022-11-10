package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"todos-api-go/common/response"
	"todos-api-go/controllers"
)

// API Handlers accept controllers then return echo.HandlerFunc
// parse query, path, body params and call controller methods
// handle http responses and errors
func HandleGetTodos(c *controllers.TodoController) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		todoListResponse, err := c.GetTodos(ctx.Request().Context())
		if err != nil {
			return response.ResponseError(ctx, err)
		}
		return response.Response(ctx, http.StatusOK, todoListResponse)
	}
}

func HandleGetTodoById(c *controllers.TodoController) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		todoId, err := strconv.Atoi(ctx.Param("todoId"))
		if err != nil {
			return response.ResponseError(ctx, err)
		}
		todoResponse, err := c.GetTodoById(ctx.Request().Context(), todoId)
		if err != nil {
			return response.ResponseError(ctx, err)
		}
		return response.Response(ctx, http.StatusOK, todoResponse)
	}
}

func HandleCreateTodo(c *controllers.TodoController) echo.HandlerFunc {
	return nil
}

func HandleUpdateTodo(c *controllers.TodoController) echo.HandlerFunc {
	return nil
}

func HandleDeleteTodo(c *controllers.TodoController) echo.HandlerFunc {
	return nil
}
