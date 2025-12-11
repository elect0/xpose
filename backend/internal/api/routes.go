package api

import (
	"database/sql"

	"github.com/elect0/xpose/backend/internal/api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title           Xpose API
// @version         1.0
// @description     This is the API for Xpose, a who's most likely to game.
// @host            localhost:8080
// @BasePath        /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func SetupRoutes(e *echo.Echo, h *handlers.Handlers, db *sql.DB) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	authGroup := e.Group("/auth")
	authGroup.POST("/register", h.RegisterUser)
	authGroup.POST("/verify", h.ValidateCode)
}
