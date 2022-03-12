package main

import (
	"echo_jtw_sample/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", handler.Login())

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", handler.Restricted())
	r.GET("/hello", handler.Hello())

	e.Logger.Fatal(e.Start(":3000"))
}
