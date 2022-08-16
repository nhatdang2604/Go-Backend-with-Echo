package mw

import (
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/constant"
)

func AdminValidateMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := ctx.Get(constant.CONTEXT_KEY_DEFAULT_VALUE).(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		admin := claims[constant.ADMIN_KEY].(bool)

		log.Printf("Is Admin: %v\r\n", admin)

		if admin {
			next(ctx)
		} else {
			return echo.ErrUnauthorized
		}

		return nil
	}
}
