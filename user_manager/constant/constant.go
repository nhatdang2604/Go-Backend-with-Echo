package constant

const (

	//Address
	PORT = "8088"

	//Paths
	ROOT_PATH      = "/"
	LOGIN_PATH     = "/login/"
	ADMIN_PATH     = "/admin/"
	HELLO_PATH     = "/hello/"
	GROUP_API_PATH = "/v2/"

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
