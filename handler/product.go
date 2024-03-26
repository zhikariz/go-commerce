package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zhikariz/go-commerce/data"
)

func GetProductByID(c echo.Context) error {
	req := struct {
		ID int `json:"id"`
	}{} // parsing parameter

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request body",
		})
	} // pengecekan / validasi

	product := data.GetProduct() // mengambil data dari database

	return c.JSON(http.StatusOK, product) // return ke client berupa json
}

func GetProducts(c echo.Context) error {
	// parsing parameter
	// pagination filtering sorting ordering
	products := data.GetProducts() // ambil data dari database

	return c.JSON(http.StatusOK, products) // return
}

func CreateProduct(c echo.Context) error {
	req := struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	}{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request body",
		})
	}

	type DataProduct struct {
		Name  string
		Price int `json:"price"`
	}
	res := struct {
		Message string      `json:"message"`
		Data    DataProduct `json:"data"`
	}{
		Message: "Successfully connected to the server",
		Data: DataProduct{
			Name:  req.Name,
			Price: req.Price,
		},
	}
	return c.JSON(http.StatusOK, res)
}
