package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/beego/beego/orm"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/constant"
)

type User struct {
	Id    int32  `orm:"auto" json:"id" form:"id"`
	Name  string `orm:"size(30)" json:"name" form:"name"`
	Age   int32  `json:"age" form:"age"`
	Phone string `orm:"size(11)" json:"phone" form:"phone"`
}

func init() {

	//init the model from database for Beego
	orm.RegisterModel(new(User))
}

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

func AddUser(ctx echo.Context) error {

	//Get the user via json from the request
	user := &User{}
	if err := ctx.Bind(user); nil != err {
		glog.Errorf("Binding user with error: %v\r\n", err)
		return err
	}

	//Insert the bind user
	o := orm.NewOrm()
	id, err := o.Insert(user)

	if nil != err {
		glog.Errorf("Insert user with error: %v\r\n", err)
		return err
	}

	glog.Infof("Insert user at row %d\r\n", id)

	return ctx.JSON(http.StatusOK, user)
}

func GetUser(ctx echo.Context) error {

	//Get the id from request
	raw, err := strconv.Atoi(ctx.QueryParam(constant.PARAM_USER_ID))
	if nil != err {
		glog.Errorf("Invalid id value: %v\r\n", err)
		return err
	}

	id := int32(raw)

	o := orm.NewOrm()
	user := &User{Id: id}

	err = o.Read(user)

	if nil != err {
		glog.Errorf("Error on get user from database: %v\r\n", err)
		return err
	}

	glog.Infof("Get user with id=%v from database: %v", id, user)
	return ctx.JSON(http.StatusOK, user)
}

func UpdateUser(ctx echo.Context) error {

	user := &User{}
	if err := ctx.Bind(user); nil != err {
		glog.Errorf("Binding user error: %v", err)
		return err
	}
	id := user.Id
	glog.Infof("Request update user with id=%v", id)
	o := orm.NewOrm()
	_, err := o.Update(user)

	if nil != err {
		glog.Errorf("Update user with id=%v failed: %v", id, err)
		return err
	}

	//Try to re-read the updated user
	err = o.Read(user)

	if nil != err {
		glog.Errorf("Error on re-reading the updated user: %v", err)
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func DeleteUser(ctx echo.Context) error {

	raw, err := strconv.Atoi(ctx.QueryParam(constant.PARAM_USER_ID))
	if nil != err {
		glog.Errorf("Invalid id value: %v\r\n", err)
		return err
	}

	id := int32(raw)
	user := &User{Id: id}

	o := orm.NewOrm()
	_, err = o.Delete(user)

	if nil != err {
		glog.Errorf("Error on deleting user with id=%v: %v \r\n", id, err)
		return err
	}

	return ctx.String(http.StatusOK, "Delete successfully")
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
