package main

import (
	"github.com/labstack/echo/v4"
	"ex4/database"
	"ex4/products"
	"ex4/categories"
	"ex4/basket"
)

func main() {
	database.ConnectDB()
	e := echo.New()
	

	e.GET("/products/:id", products.GetProduct, products.CheckIdMiddleware)
	e.PUT("/products/:id", products.UpdateProduct, products.CheckIdMiddleware)
	e.DELETE("/products/:id", products.DeleteProduct, products.CheckIdMiddleware)
	e.GET("/products", products.GetProducts)
	e.POST("/products", products.CreateProduct)

	e.GET("/categories/:id", categories.GetCategory, categories.CheckIdMiddleware)
	e.PUT("/categories/:id", categories.UpdateCategory, categories.CheckIdMiddleware)
	e.DELETE("/categories/:id", categories.DeleteCategory, categories.CheckIdMiddleware)
	e.GET("/categories", categories.GetCategories)
	e.POST("/categories", categories.CreateCategory)

	e.GET("/basket/:id", basket.GetBasketItem, basket.CheckIdMiddleware)
	e.PUT("/basket/:id", basket.UpdateBasketItem, basket.CheckIdMiddleware)
	e.DELETE("/basket/:id", basket.DeleteBasketItem, basket.CheckIdMiddleware)
	e.GET("/basket", basket.GetBasket)
	e.POST("/basket", basket.AddToBasket)

	e.Logger.Fatal(e.Start(":9000"))
}