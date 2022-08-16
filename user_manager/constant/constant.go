package constant

const (

	//Address
	PORT = "8088"

	//Paths
	ROOT_PATH        = "/"
	LOGIN_PATH       = "/login/"
	ADMIN_PATH       = "/admin/"
	HELLO_PATH       = "/hello/"
	GROUP_API_PATH   = "/v2"       //removing foward slash at the end
	USER_GROUP_PATH  = "/api/user" //removing foward slash at the end
	USER_GET_PATH    = "/get-user/"
	USER_UPDATE_PATH = "/update-user/"
	USER_DELETE_PATH = "/delete-user/"

	//jwt keys
	USERNAME_KEY              = "username"
	ADMIN_KEY                 = "admin"
	EXPIRATION_KEY            = "exp"
	CONTEXT_KEY_DEFAULT_VALUE = "user"

	//secret key
	// Note: storing secret key hardcodely is a bad design.
	//	Only using for learning purpose
	SECRET_KEY = "secret_key"
)
