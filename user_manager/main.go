package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (

	//Address
	PORT = "8088"

	//Paths
	ROOT_PATH  = "/"
	LOGIN_PATH = "/login"
)

type Message struct {
	Text string
}

type LoginRequest struct {
	Username string `json:"username" form:"username" xml:"username" query: "username"`
	Password string `json:"password" form:"password" xml:"password" query: "password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

//Handler for the root path
func Hello(ctx echo.Context) error {
	data := Message{
		Text: "Hello World",
	}
	return ctx.JSON(http.StatusOK, data)
}

func Login(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, &LoginResponse{
		Token: "123456",
	})
}

func AuthValidator(username string, password string, ctx echo.Context) (bool, error) {
	if "admin" != username ||
		"admin" != password {
		return false, nil
	}

	return true, nil
}

func main() {
	server := echo.New()

	server.GET(ROOT_PATH, Hello)
	server.POST(LOGIN_PATH, Login, middleware.BasicAuth(AuthValidator))
	server.Logger.Fatal(server.Start(":" + PORT))
}
