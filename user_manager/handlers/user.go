package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "API Get User")
}

func UpdateUser(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "API Update User")
}

func DeleteUser(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "API Delete User")
}
