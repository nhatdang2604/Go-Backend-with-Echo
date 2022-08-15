package handler

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/models"
)

const (
	CONTEXT_KEY_DEFAULT_VALUE = "user"
)

func Hello(ctx echo.Context) error {

	user := ctx.Get(CONTEXT_KEY_DEFAULT_VALUE).(*jwt.Token)

	claims := user.Claims.(jwt.MapClaims)

	username := claims[USERNAME_KEY].(string)
	admin := claims[ADMIN_KEY].(bool)

	message := fmt.Sprintf("Hello %s is admin %v", username, admin)

	data := models.Message{
		Text: message,
	}
	return ctx.JSON(http.StatusOK, data)
}
