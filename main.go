package main

import (
	"github.com/labstack/echo"
	"github.com/mikamikuh/goannotator/controller"
)

func main() {
	e := echo.New()

	e.GET("/search", controller.SearchHandler)
	e.Logger.Fatal(e.Start(":80"))
}
