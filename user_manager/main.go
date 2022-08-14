package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handler "github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/handlers"
	mw "github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/middlewares"
)

const (

	//Address
	PORT = "8088"

	//Paths
	ROOT_PATH  = "/"
	LOGIN_PATH = "/login"
)

func main() {
	server := echo.New()

	//Middlewares must be registered before adding root path handler
	server.Use(middleware.Logger())

	//Add handlers
	server.GET(ROOT_PATH, handler.Hello) //root path handler
	server.POST(LOGIN_PATH, handler.Login, middleware.BasicAuth(mw.BasicAuth))

	//Run the server
	server.Logger.Fatal(server.Start(":" + PORT))
}
