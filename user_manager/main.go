package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
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
func hello(ctx echo.Context) error {
	data := Message{
		Text: "Hello World",
	}
	return ctx.JSON(http.StatusOK, data)
}

func login(ctx echo.Context) error {

	request := new(LoginRequest)
	ctx.Bind(request)

	log.Printf("Request data: %v\r\n", request)

	if "admin" != request.Username ||
		"admin" != request.Password {
		return ctx.String(http.StatusUnauthorized, "login failed")
	}

	return ctx.JSON(http.StatusOK, &LoginResponse{
		Token: "123456",
	})
}

func main() {
	server := echo.New()
	server.GET(ROOT_PATH, hello)
	server.POST(LOGIN_PATH, login)
	server.Logger.Fatal(server.Start(":" + PORT))
}
