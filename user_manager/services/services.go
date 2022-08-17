package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/beego/beego/orm"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/constant"
	"github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/models"
)

type UserService struct{}

func (service *UserService) Add(ctx echo.Context) error {

	//Get the user via json from the request
	user := &models.User{}
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

func (service *UserService) Get(ctx echo.Context) error {

	//Get the id from request
	raw, err := strconv.Atoi(ctx.QueryParam(constant.PARAM_USER_ID))
	if nil != err {
		glog.Errorf("Invalid id value: %v\r\n", err)
		return err
	}

	id := int32(raw)

	o := orm.NewOrm()
	user := &models.User{Id: id}

	err = o.Read(user)

	if nil != err {
		glog.Errorf("Error on get user from database: %v\r\n", err)
		return err
	}

	glog.Infof("Get user with id=%v from database: %v", id, user)
	return ctx.JSON(http.StatusOK, user)
}

func (service *UserService) Update(ctx echo.Context) error {

	user := &models.User{}
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

func (service *UserService) Delete(ctx echo.Context) error {

	raw, err := strconv.Atoi(ctx.QueryParam(constant.PARAM_USER_ID))
	if nil != err {
		glog.Errorf("Invalid id value: %v\r\n", err)
		return err
	}

	id := int32(raw)
	user := &models.User{Id: id}

	o := orm.NewOrm()
	_, err = o.Delete(user)

	if nil != err {
		glog.Errorf("Error on deleting user with id=%v: %v \r\n", id, err)
		return err
	}

	return ctx.String(http.StatusOK, "Delete successfully")
}

func (service *UserService) GetAll(ctx echo.Context) error {

	//Preparing for querying
	var users []*models.User
	o := orm.NewOrm()
	querySetter := o.QueryTable(constant.DB_TABLE_NAME_USER)

	size, err := querySetter.All(&users)

	if nil != err {
		glog.Errorf("Querying with error: %v", err)
		return err
	}

	//Header editing
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(ctx.Response())
	for _, user := range users {
		if err := encoder.Encode(user); nil != err {
			return err
		}

		ctx.Response().Flush()
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("There are %v users", size))
}
