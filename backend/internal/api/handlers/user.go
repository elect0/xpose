package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/elect0/xpose/backend/internal/platform/auth"
	"github.com/elect0/xpose/backend/internal/platform/database"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type AuthUserRequest struct {
	Email string `json:"email" example:"johndoe@example.com"`
}

type UpdateUserRequest struct {
	Username   string `json:"username" example:"john_doe"`
	Email      string `json:"email" example:"johndoe@example.com"`
	ProfilePic string `json:"profile_pic_url" example:"https:/"`
	Verified   bool   `json:"verified" example:"TRUE"`
}

type UserCodeRequest struct {
	UserID uuid.UUID `json:"id" example:"1a2388c5-9db0-42c5-9f62-915adcbec0b0"`
	Code   string    `json:"code" example:"jgfl12412"`
}

type UserResponse struct {
	ID         uuid.UUID `json:"id" example:"1a2388c5-9db0-42c5-9f62-915adcbec0b0"`
	Username   string    `json:"username" example:"john_doe"`
	Email      string    `json:"email" example:"johndoe@example.com"`
	ProfilePic string    `json:"profile_pic_url" example:"https:/"`
	Verified   bool      `json:"verified" example:"TRUE"`
}

type CodeResponse struct {
	Code string `json:"code" example:"FXMuYFtVxYtD"`
}

// RegisterUser godoc
// @Summary      Register User
// @Description  Registers a new user.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        customer body AuthUserRequest true "User details"
// @Success      201  {object} CodeResponse
// @Router       /auth/register [post]
func (h *Handlers) RegisterUser(c echo.Context) error {
	var req AuthUserRequest
	if err := c.Bind(&req); err != nil {
		h.Logger.Warn("Failed to bind request for user authentification", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide a valid message payload")
	}

	user, err := h.Queries.GetUserByEmail(c.Request().Context(), req.Email)
	if err == sql.ErrNoRows {
		user, err = h.Queries.CreateUser(c.Request().Context(), req.Email)
		if err != nil {
			h.Logger.Error("Failed to create user", zap.Error(err))
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
		}
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user")
	}

	code, err := auth.GenerateCode()

	userCode, err := h.Queries.CreateCode(c.Request().Context(), database.CreateCodeParams{
		Code:      code,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	h.Logger.Info("Code generated", zap.String("code", userCode.Code))
	// email service later

	return c.JSON(200, CodeResponse{
		Code: userCode,
	})
}

// ValidateCode godoc
// @Summary      Validate Code
// @Description  Validates an user's code.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body UserCodeRequest true "User details"
// @Success      201  {object} UserResponse
// @Router       /auth/verify [post]
func (h *Handlers) ValidateCode(c echo.Context) error {
	var req UserCodeRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide a valid message payload")
	}

	code, err := h.Queries.GetCode(c.Request().Context(), database.GetCodeParams{
		UserID: req.UserID,
		Code:   req.Code,
	})

	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusBadRequest, "The provided code is invalid")
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	// transaction
	tx, err := h.DB.BeginTx(c.Request().Context(), nil)

	qtx := h.Queries.WithTx(tx)

	defer tx.Rollback()

	err = qtx.MarkCodeUsed(c.Request().Context(), code.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	user, err := qtx.VerifyUserById(c.Request().Context(), code.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")

	}

	if err = tx.Commit(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Commit failed")
	}

	token, err := h.TokenMaker.CreateToken(user.ID, time.Duration(h.Config.Paseto.DurationMinutes))
	if err != nil {
		h.Logger.Warn("Failed to create user token", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	cookie := new(http.Cookie)
	cookie.Name = "access_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Duration(h.Config.Paseto.DurationMinutes))
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.Path = "/"
	cookie.SameSite = http.SameSiteStrictMode

	c.SetCookie(cookie)

	return c.JSON(200, UserResponse{
		ID:         user.ID,
		Username:   user.Username.String,
		Email:      user.Email,
		ProfilePic: user.ProfilePicUrl.String,
	})
}
