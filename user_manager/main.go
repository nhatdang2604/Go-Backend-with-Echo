package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/constant"
	handler "github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/handlers"
	mw "github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/middlewares"
)

const (

	//Address
	PORT = "8088"

	//Paths
	ROOT_PATH  = "/"
	LOGIN_PATH = "/login/"
	ADMIN_PATH = "/admin/"
)

func main() {
	server := echo.New()

	//Middlewares must be registered before adding root path handler
	server.Use(middleware.Logger())

	//Define functor for middlewares
	isLoggedIn := middleware.JWT([]byte(constant.SECRET_KEY)) //building a logging checker middleware
	isAdmin := mw.AdminValidateMiddleware

	//Add handlers
	server.GET(ROOT_PATH, handler.Hello, isLoggedIn) //root path handler, using isLoggedIn middleware to authorize for only logged user to use
	server.GET(ADMIN_PATH, handler.Hello, isLoggedIn, isAdmin)
	server.POST(LOGIN_PATH, handler.Login, middleware.BasicAuth(mw.BasicAuth))

	//Run the server
	server.Logger.Fatal(server.Start(":" + PORT))
}
