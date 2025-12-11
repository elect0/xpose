package main

import (
	"log"
	"net/http"

	"github.com/elect0/xpose/backend/db"
	"github.com/elect0/xpose/backend/internal/config"
	"github.com/elect0/xpose/backend/internal/platform/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error while reading the config: %v", err)
	}

	logger, err := logger.New(config.App.Environment)
	if err != nil {
		log.Fatalf("Error while initializing logger: %v", err)
	}
	defer logger.Sync()

	db, err := db.New(config.DB.Source, config.App.Environment, logger)
	if err != nil {
		logger.Fatal("Couldn't initialize database", zap.Error(err))
	}

	defer db.Close()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.Start("localhost:2020")
}
