package constant

import (
	"time"
)

const (

	//Address
	PORT = "8088"

	//Paths
	ROOT_PATH         = "/"
	LOGIN_PATH        = "/login/"
	ADMIN_PATH        = "/admin/"
	HELLO_PATH        = "/hello/"
	USER_ADD_PATH     = "/add/"
	USER_GET_PATH     = "/get/"
	USER_UPDATE_PATH  = "/update/"
	USER_DELETE_PATH  = "/delete/"
	USER_GET_ALL_PATH = "/get-all/"

	//Grouping Paths
	//Removing foward slash at the end
	GROUP_API_PATH  = "/v2"
	USER_GROUP_PATH = "/api/user"

	//jwt keys
	USERNAME_KEY              = "username"
	ADMIN_KEY                 = "admin"
	EXPIRATION_KEY            = "exp"
	CONTEXT_KEY_DEFAULT_VALUE = "user"

	//Database config
	DB_NAME    = "echotest"
	DB_USER    = "test"
	DB_PWD     = "test"
	DB_CHARSET = "utf8"

	//Database schema's name
	DB_TABLE_NAME_USER = "user"

	//secret key
	// Note: storing secret key hardcodely is a bad design.
	//	Only using for learning purpose
	SECRET_KEY = "secret_key"

	//Parameter from user's request
	PARAM_USER_ID = "id"

	//Cache config
	CACHE_HOST       = "localhost"
	CACHE_PORT       = "6379"
	CACHE_ADDR       = CACHE_HOST + ":" + CACHE_PORT
	CACHE_DB         = 0
	CACHE_EXPIRATION = 5 * time.Minute
)
