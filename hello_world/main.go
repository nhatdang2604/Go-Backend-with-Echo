package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (

	//Address
	PORT = "8088"

	//Paths
	ROOT_PATH = "/"
)

type Message struct {
	Text string
}

//Handler for the root path
func hello(ctx echo.Context) error {
	data := Message{
		Text: "Hello World",
	}
	return ctx.JSON(http.StatusOK, data)
}

func main() {
	server := echo.New()

	server.GET(ROOT_PATH, hello)

	server.Logger.Fatal(server.Start(":" + PORT))
}
