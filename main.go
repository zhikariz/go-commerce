package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zhikariz/go-commerce/handler"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		res := struct {
			Message string `json:"message"`
			Data    string `json:"data"`
		}{
			Message: "Successfully connected to the server",
			Data:    "Hello World",
		}
		return c.JSON(http.StatusOK, res)
	})
	e.GET("/product", handler.GetProducts)
	e.GET("/product/:id", handler.GetProductByID)
	// e.POST("/product", handler.CreateProduct)
	// e.PUT("/product/:id", handler.UpdateProduct)
	// e.DELETE("/product/:id", handler.DeleteProduct)

	e.Logger.Fatal(e.Start(":8080"))
}
