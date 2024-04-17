package main

import (
	"TodoApp/Router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	//"github.com/rs/cors"
)

func main() {
	
	e := echo.New()
	e.HideBanner=true
	e.Use(middleware.CORS())
	Router.RouterHander(e)
	e.Start(":8080")
	
}