package mw

import "github.com/labstack/echo/v4"

const (
	USERNAME_KEY = "username"
	ADMIN_KEY    = "admin"
)

func BasicAuth(username string, password string, ctx echo.Context) (bool, error) {

	//With admin role
	if "admin" == username && "admin" == password {
		ctx.Set(USERNAME_KEY, username)
		ctx.Set(ADMIN_KEY, true)
		return true, nil
	}

	//Without admin role
	if "test" == username && "test" == password {
		ctx.Set(USERNAME_KEY, username)
		ctx.Set(ADMIN_KEY, false)
		return true, nil
	}

	return false, nil
}
