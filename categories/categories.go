package categories

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"ex4/database"
	"ex4/models"
	// "fmt"
)

func CheckIdMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		category := new(models.Category)
		result := database.DB.First(&category, id); if result.Error != nil {
			return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Category not found"}`))
		}
		c.Set("category", category)
		return next(c)
	}
}

func GetCategories(c echo.Context) error {
	var categories []models.Category
	database.DB.Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

func GetCategory(c echo.Context) error {
	category := c.Get("category").(*models.Category)
	return c.JSON(http.StatusOK, category)
}

func UpdateCategory(c echo.Context) error {
	category := c.Get("category").(*models.Category)

	err := c.Bind(&category); if err != nil {
		return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Category is not valid"}`))
	}
	
	result := database.DB.Save(&category); if result.Error != nil {
		return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Cannot save category"}`))
	}
	return c.JSON(http.StatusOK, category)
}

func CreateCategory(c echo.Context) error {
	category := new(models.Category)

	err := c.Bind(&category); if err != nil {
		return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Error while creating category"}`))
	}

	result := database.DB.Save(&category); if result.Error != nil {
		return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Cannot save category"}`))
	}
	return c.JSON(http.StatusOK, category)
}

func DeleteCategory(c echo.Context) error {
	category := c.Get("category").(*models.Category)
	result := database.DB.Delete(models.Category{}, category.ID); if result.Error != nil {
		return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Cannot delete category if product of that category exists"}`))
	}

	return c.JSONBlob(http.StatusNotFound, []byte(`{"Info":"Category deleted"}`))
}