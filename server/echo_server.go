package server

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"todos-api-go/api"
	"todos-api-go/config"
	"todos-api-go/controllers"
	"todos-api-go/utils"
)

type TodoServer struct {
	*echo.Echo
	config.Config
}

func NewServer(cfg config.Config, db *sql.DB) (*TodoServer, error) {
	server := echo.New()
	server.Use(
		middleware.Recover(),
		// TODO: other middlewares
	)
	server.Validator = &utils.CustomValidator{Validator: validator.New()}

	health := server.Group("/status")
	{
		health.GET("", api.HealthCheck())
	}

	// Initialize Controllers
	todoController := controllers.NewTodoController(db, cfg.Setting())

	v1 := server.Group("api/v1/todos")
	{
		v1.GET("", api.HandleGetTodos(todoController))
		v1.GET("/:todoId", api.HandleGetTodoById(todoController))
		//v1.POST("", api.HandleCreateTodo(todoController))
		//v1.PUT("", api.HandleUpdateTodo(todoController))
		//v1.DELETE("", api.HandleDeleteTodo(todoController))
	}

	return &TodoServer{
		Echo:   server,
		Config: cfg,
	}, nil
}
