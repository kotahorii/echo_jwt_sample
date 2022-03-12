package handler

import (
	"echo_jtw_sample/db"
	"echo_jtw_sample/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	}
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("name")
		password := c.FormValue("password")

		db := db.ConnectGORM()

		// it's necessary when table name is singular
		// db.SingularTable(true)
		user := []models.User{}
		db.Find(&user, "name=? and password=?", username, password)

		if len(user) > 0 && username == user[0].Name {
			token := jwt.New(jwt.SigningMethodHS256)

			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = username
			claims["admin"] = true
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, map[string]string{
				"token": t,
			})
		}
		return echo.ErrUnauthorized
	}
}

func Restricted() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		return c.String(http.StatusOK, "Welcome"+name+"!")
	}
}
