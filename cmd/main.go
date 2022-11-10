package main

import (
	logger "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"todos-api-go/config"
	"todos-api-go/server"
)

func init() {
	if GO_ENV := os.Getenv("GO_ENV"); GO_ENV == "local" {
		logger.SetLevel(logger.DebugLevel)
		logger.SetFormatter(&logger.TextFormatter{
			FullTimestamp: true,
			ForceColors:   true,
			ForceQuote:    true,
		})
		logger.SetOutput(os.Stdout)
	} else {
		logger.SetLevel(logger.DebugLevel)
		logger.SetFormatter(&logger.JSONFormatter{
			PrettyPrint:      true,
			DisableTimestamp: false,
		})
	}
}

func main() {
	// Initialize Environment Variable
	settings := config.NewSetting()
	config := config.NewConfig(settings)

	server.Init(config)

	quit := make(chan os.Signal, 1)
	// k8s termination call - sigint, sigterm
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
