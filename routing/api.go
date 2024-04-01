package routing

import (
	"github.com/labstack/echo/v4"
	"github.com/zhikariz/go-commerce/handler"
)

func GenerateAPIV1(e *echo.Echo) {
	product := e.Group("/product")
	product.GET("", handler.GetProducts)
	product.GET("/:id", handler.GetProductByID)
	product.POST("", handler.CreateProduct)
	product.PUT("/:id", handler.UpdateProduct)
	product.DELETE("/:id", handler.DeleteProduct)
}
