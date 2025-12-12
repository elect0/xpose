package handlers

import (
	"database/sql"

	"github.com/elect0/xpose/backend/internal/config"
	"github.com/elect0/xpose/backend/internal/platform/auth"
	"github.com/elect0/xpose/backend/internal/platform/database"
	"github.com/elect0/xpose/backend/internal/platform/logger"
)

type ErrorResponse struct {
	Message string `json:"messsage" example:"An error occurred"`
}

type Handlers struct {
	Logger     *logger.Logger
	Config     *config.Config
	DB         *sql.DB
	Queries    *database.Queries
	TokenMaker *auth.TokenMaker
}

func New(logger *logger.Logger, config *config.Config, db *sql.DB, queries *database.Queries, tokenMaker *auth.TokenMaker) *Handlers {
	return &Handlers{
		Logger:     logger,
		Config:     config,
		DB:         db,
		Queries:    queries,
		TokenMaker: tokenMaker,
	}
}
