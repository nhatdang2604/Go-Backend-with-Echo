package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/constant"
	handler "github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/handlers"
	mw "github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/middlewares"
)

const ()

func main() {
	server := echo.New()

	//Middlewares must be registered before adding root path handler
	server.Use(middleware.Logger())

	//Define functor for middlewares
	isLoggedIn := middleware.JWT([]byte(constant.SECRET_KEY)) //building a logging checker middleware
	isAdmin := mw.AdminValidateMiddleware

	//Add handlers
	server.GET(constant.ROOT_PATH, handler.Hello, isLoggedIn) //root path handler, using isLoggedIn middleware to authorize for only logged user to use
	server.GET(constant.ADMIN_PATH, handler.Hello, isLoggedIn, isAdmin)
	server.POST(constant.LOGIN_PATH, handler.Login, middleware.BasicAuth(mw.BasicAuth))

	//Run the server
	server.Logger.Fatal(server.Start(":" + constant.PORT))
}
