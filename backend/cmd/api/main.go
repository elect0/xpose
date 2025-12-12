package main

import (
	"fmt"
	"log"

	"aidanwoods.dev/go-paseto"
	"github.com/elect0/xpose/backend/db"
	"github.com/elect0/xpose/backend/internal/api"
	"github.com/elect0/xpose/backend/internal/api/handlers"
	"github.com/elect0/xpose/backend/internal/config"
	"github.com/elect0/xpose/backend/internal/platform/auth"
	"github.com/elect0/xpose/backend/internal/platform/database"
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

	queries := database.New(db)

	tokenMaker, err := auth.NewTokenMaker(config.Paseto.Asymmetrical)
	if err != nil {
		logger.Fatal("Coudln't initialize token maker", zap.Error(err))
	}

	handlers := handlers.New(logger, config, db, queries, tokenMaker)

	api.SetupRoutes(e, handlers, db)

	logger.Info("server starting on: ", zap.String("Port", config.App.Port))
	e.Logger.Fatal(e.Start(config.App.Port))
}
