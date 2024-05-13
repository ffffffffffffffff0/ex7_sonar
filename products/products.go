package products

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"ex4/database"
	"ex4/models"
	//"fmt"
	"strconv"
)

func CheckIdMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		product := new(models.Product)
		result := database.DB.First(&product, id); if result.Error != nil {
			return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Product not found"}`))
		}
		c.Set("product", product)
		return next(c)
	}
}

func CheckPrice(price_str string) float32 {
	if floatV, err := strconv.ParseFloat(price_str, 64); err == nil {
		return float32(floatV)
	}
	return 0.0
}

func CheckCategory(cat string) uint {
	if uintV, err := strconv.Atoi(cat); err == nil {
		return uint(uintV)
	}
	return 1
}

func GetProducts(c echo.Context) error {
	var products []models.Product
	category := c.QueryParam("cat")
	name := c.QueryParam("name")
	above := c.QueryParam("above")
	below := c.QueryParam("below")

	query := database.DB.Model(&models.Product{})

	if category != "" {
		query.Scopes(database.FilterByCategory(CheckCategory(category)))
	}
	if name != "" {
		query.Scopes(database.FilterByName(name))
	}
	if above != "" {
		query.Scopes(database.FilterByPriceAbove(CheckPrice(above)))
	}
	if below != "" {
		query.Scopes(database.FilterByPriceBelow(CheckPrice(below)))
	}

	query.Find(&products)
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	product := c.Get("product").(*models.Product)
	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	product := c.Get("product").(*models.Product)

	err := c.Bind(&product); if err != nil {
		return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Product is not valid"}`))
	}
	
	result := database.DB.Save(&product); if result.Error != nil {
		return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Cannot save Product, check if category is correct"}`))
	}
	return c.JSON(http.StatusOK, product)
}

func CreateProduct(c echo.Context) error {
	product := new(models.Product)

	err := c.Bind(&product); if err != nil {
		return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Error while creating product"}`))
	}

	result := database.DB.Create(&product); if result.Error != nil {
		return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Cannot save Product, check if category is correct"}`))
	}
	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	product := c.Get("product").(*models.Product)
	database.DB.Delete(models.Product{}, product.ID)
	return c.JSONBlob(http.StatusNotFound, []byte(`{"Info":"Product deleted"}`))
}