package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/models"
)

func Hello(ctx echo.Context) error {
	data := models.Message{
		Text: "Hello World",
	}
	return ctx.JSON(http.StatusOK, data)
}
