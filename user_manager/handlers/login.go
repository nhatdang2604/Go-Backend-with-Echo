package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/models"
)

func Login(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, &models.LoginResponse{
		Token: "123456",
	})
}
