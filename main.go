package main

import (
	"github.com/labstack/echo/v4"
	"github.com/zhikariz/go-commerce/routing"
)

func main() {
	e := echo.New()
	routing.GenerateAPIV1(e)
	e.Logger.Fatal(e.Start(":8080"))
}
