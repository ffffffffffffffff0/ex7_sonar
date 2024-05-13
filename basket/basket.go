package basket

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
		basketItem := new(models.BasketItem)
		result := database.DB.First(&basketItem, id); if result.Error != nil {
			return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Basket item not found"}`))
		}
		c.Set("basketItem", basketItem)
		return next(c)
	}
}

func GetBasket(c echo.Context) error {
	var basketItems []models.BasketItem
	database.DB.Find(&basketItems)
	return c.JSON(http.StatusOK, basketItems)
}

func GetBasketItem(c echo.Context) error {
	basketItem := c.Get("basketItem").(*models.BasketItem)
	return c.JSON(http.StatusOK, basketItem)
}

func UpdateBasketItem(c echo.Context) error {
	basketItem := c.Get("basketItem").(*models.BasketItem)

	err := c.Bind(&basketItem); if err != nil {
		return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Basket item is not valid"}`))
	}
	
	result := database.DB.Save(&basketItem); if result.Error != nil {
		return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Cannot save basket item"}`))
	}
	return c.JSON(http.StatusOK, basketItem)
}

func AddToBasket(c echo.Context) error {
	basketItem := new(models.BasketItem)
	var existingItem models.BasketItem

	err := c.Bind(&basketItem); if err != nil {
		return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Error while adding item to basket"}`))
	}

	
	if err := database.DB.Where("product_id = ?", basketItem.ProductID).First(&existingItem).Error; err != nil {
		result := database.DB.Create(&basketItem); if result.Error != nil {
			return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Cannot save basket item"}`))
		}
		return c.JSON(http.StatusOK, basketItem)
	} else {
		existingItem.Amount += basketItem.Amount
		database.DB.Save(&existingItem)
		return c.JSON(http.StatusOK, existingItem)
	}

	// return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Error while adding item to basket"}`))

}

func DeleteBasketItem(c echo.Context) error {
	basketItem := c.Get("basketItem").(*models.BasketItem)
	result := database.DB.Delete(models.BasketItem{}, basketItem.ID); if result.Error != nil {
		return c.JSONBlob(http.StatusNotFound, []byte(`{"error":"Error"}`))
	}

	return c.JSONBlob(http.StatusNotFound, []byte(`{"Info":"Category deleted"}`))
}