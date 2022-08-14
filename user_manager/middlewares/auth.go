package mw

import "github.com/labstack/echo/v4"

func BasicAuth(username string, password string, ctx echo.Context) (bool, error) {
	if "admin" != username ||
		"admin" != password {
		return false, nil
	}

	return true, nil
}
