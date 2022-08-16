package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name string
	Age  int32
}

// func init() {

// 	//init the model from database for Beego
// 	orm.RegisterModel(new(User))
// }

//Mock data
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

func GetAllUser(ctx echo.Context) error {

	//Header editing
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(ctx.Response())
	for _, user := range users {
		if err := encoder.Encode(user); nil != err {
			return err
		}

		ctx.Response().Flush()

		//Sleep to simulate heavyweight process
		time.Sleep(1 * time.Second)
	}

	return nil
}
