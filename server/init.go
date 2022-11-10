package server

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"todos-api-go/config"
)

func Init(cfg config.Config) {
	settings := cfg.Setting()

	//Initialize DB instance
	db, err := config.NewDB(settings)
	if err != nil {
		logger.Fatalf("Creating new DB instance failed - %+v", err)
	}

	// Initialize Http Echo Server
	server, err := NewServer(cfg, db)
	if err != nil {
		logger.Fatalf("Creating new server error - %+v", err)
	}

	go func() {
		port := settings.PORT
		if err := server.Start(fmt.Sprintf(":%s", port)); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Starting server error - %+v", err)
		}
	}()
}
