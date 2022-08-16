package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name string
	Age  int32
}

var users = []User{
	{Name: "test0", Age: 18},
	{Name: "test1", Age: 19},
	{Name: "test2", Age: 20},
	{Name: "test3", Age: 21},
	{Name: "test4", Age: 22},
	{Name: "test5", Age: 23},
	{Name: "test6", Age: 24},
}

func GetUser(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "API Get User")
}

func UpdateUser(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "API Update User")
}

func DeleteUser(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "API Delete User")
}
