package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"todos-api-go/api"
	"todos-api-go/config"
	"todos-api-go/services"
	"todos-api-go/utils"
)

type TodoServer struct {
	*echo.Echo
	config.Config
}

func NewServer(cfg config.Config, todoService *services.TodoService) (*TodoServer, error) {
	server := echo.New()
	server.Use(
		middleware.Recover(),
		// TODO: logger middleware
	)
	server.Validator = &utils.Validator{Validator: validator.New()}

	health := server.Group("/status")
	{
		health.GET("", api.HealthCheck())
	}

	//v1 := server.Group("api/v1/todos/v1")

	return &TodoServer{
		Echo:   server,
		Config: cfg,
	}, nil
}
