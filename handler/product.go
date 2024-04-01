package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zhikariz/go-commerce/data"
)

func GetProductByID(c echo.Context) error {
	req := struct {
		ID int `param:"id"`
	}{} // parsing parameter

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request body",
		})
	} // pengecekan / validasi

	product := data.GetProduct(req.ID) // mengambil data dari database

	return c.JSON(http.StatusOK, product) // return ke client berupa json
}

func GetProducts(c echo.Context) error {
	// parsing parameter
	// pagination filtering sorting ordering
	products := data.GetProducts() // ambil data dari database

	return c.JSON(http.StatusOK, products) // return
}

func CreateProduct(c echo.Context) error {
	type CreateProductRequest struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Price    int    `json:"price"`
		Quantity int    `json:"quantity"`
		Color    string `json:"color"`
		Size     string `json:"size"`
		Brand    string `json:"brand"`
		Model    string `json:"model"`
	}

	req := CreateProductRequest{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request body",
		})
	}

	data.CreateProduct(data.Product{
		ID:       req.ID,
		Name:     req.Name,
		Price:    req.Price,
		Quantity: req.Quantity,
		Color:    req.Color,
		Size:     req.Size,
		Brand:    req.Brand,
		Model:    req.Model,
	})

	res := struct {
		Message string       `json:"message"`
		Data    data.Product `json:"data"`
	}{
		Message: "Successfully connected to the server",
		Data:    *data.GetProduct(req.ID),
	}
	return c.JSON(http.StatusOK, res)
}

func UpdateProduct(c echo.Context) error {
	type UpdateProductRequest struct {
		ID       int    `param:"id"`
		Name     string `json:"name"`
		Price    int    `json:"price"`
		Quantity int    `json:"quantity"`
		Color    string `json:"color"`
		Size     string `json:"size"`
		Brand    string `json:"brand"`
		Model    string `json:"model"`
	}

	req := UpdateProductRequest{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request body",
		})
	}

	data.UpdateProduct(data.Product{
		ID:       req.ID,
		Name:     req.Name,
		Price:    req.Price,
		Quantity: req.Quantity,
		Color:    req.Color,
		Size:     req.Size,
		Brand:    req.Brand,
		Model:    req.Model,
	})

	res := struct {
		Message string       `json:"message"`
		Data    data.Product `json:"data"`
	}{
		Message: "Successfully connected to the server",
		Data:    *data.GetProduct(req.ID),
	}
	return c.JSON(http.StatusOK, res)
}

func DeleteProduct(c echo.Context) error {
	req := struct {
		ID int `param:"id"`
	}{} // parsing parameter

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request body",
		})
	} // pengecekan / validasi

	data.DeleteProduct(req.ID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully deleted product",
	})
}
