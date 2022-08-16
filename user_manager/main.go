package main

import (
	"github.com/beego/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/constant"
	handler "github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/handlers"
	mw "github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/middlewares"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	//Database config
	user := constant.DB_USER
	password := constant.DB_PWD
	dbName := constant.DB_NAME
	charset := constant.DB_CHARSET

	//Connect the databse
	connectString := user + ":" + password + "@/" + dbName + "?charset=" + charset
	err := orm.RegisterDataBase("default", "mysql", connectString)

	if nil != err {
		glog.Fatal("Failed to register the database: %v", err)
	}

	//Synchronize config in connection
	name := "default" //database alias
	force := false    //drop the table and re-create after running code
	verbose := true   //logging on
	err = orm.RunSyncdb(name, force, verbose)

	if nil != err {
		glog.Fatal("Failed to sync the database: %v", err)
	}

}

func main() {

	server := echo.New()

	//Middlewares must be registered before adding root path handler
	server.Use(middleware.Logger())

	//Define functor for middlewares
	isLoggedIn := middleware.JWT([]byte(constant.SECRET_KEY)) //building a logging checker middleware
	isAdmin := mw.AdminValidateMiddleware

	//Add handlers
	//NOTE: It's very important to put the middlewares in the function parameters's order in correct position
	server.GET(constant.ROOT_PATH, handler.Hello, isLoggedIn) //root path handler, using isLoggedIn middleware to authorize for only logged user to use
	server.GET(constant.ADMIN_PATH, handler.Hello, isLoggedIn, isAdmin)
	server.POST(constant.LOGIN_PATH, handler.Login, middleware.BasicAuth(mw.BasicAuth))

	//Grouping API testing
	groupv2 := server.Group(constant.GROUP_API_PATH)
	groupv2.GET(constant.HELLO_PATH, handler.Hello2)

	//Keep grouping for User APIs
	groupUser := server.Group(constant.USER_GROUP_PATH, isLoggedIn)
	groupUser.GET(constant.USER_GET_PATH, handler.GetUser)
	groupUser.GET(constant.USER_UPDATE_PATH, handler.UpdateUser, isAdmin)
	groupUser.GET(constant.USER_DELETE_PATH, handler.DeleteUser, isAdmin)
	groupUser.GET(constant.USER_GET_ALL_PATH, handler.GetAllUser)

	//Run the server
	server.Logger.Fatal(server.Start(":" + constant.PORT))
}
