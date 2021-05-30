package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "./public/dist")
	e.GET("/ws", handleConn)
	go forwardBroadcast()
	e.Logger.Fatal(e.Start(":" + port))
}
