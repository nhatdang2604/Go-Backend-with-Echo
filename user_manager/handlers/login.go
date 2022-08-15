package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	mw "github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/middlewares"
	"github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/models"
)

const (

	//jwt keys
	USERNAME_KEY   = "username"
	ADMIN_KEY      = "admin"
	EXPIRATION_KEY = "exp"

	//secret key
	// Note: storing secret key hardcodely is a bad design
	//	only using for learning purpose
	SECRET_KEY = "secret_key"
)

func Login(ctx echo.Context) error {

	username := ctx.Get(mw.USERNAME_KEY).(string)
	log.Printf("Login with username: %v\r\n", username)

	admin := ctx.Get(mw.ADMIN_KEY).(bool)
	log.Printf("Login with admin: %v\r\n", admin)

	token := jwt.New(jwt.SigningMethodHS256)

	//set claim
	claims := token.Claims.(jwt.MapClaims)
	claims[USERNAME_KEY] = username
	claims[ADMIN_KEY] = admin
	claims[EXPIRATION_KEY] = time.Now().Add(2 * time.Minute).Unix() //after 10 minutes => token out of date

	sign, err := token.SignedString([]byte(SECRET_KEY))

	if nil != err {
		log.Printf("Sign token error: %v", err)
		return err
	}

	return ctx.JSON(http.StatusOK, &models.LoginResponse{
		Token: sign,
	})
}
