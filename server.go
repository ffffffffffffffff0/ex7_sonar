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
	p := "/products/:id"
	c := "/categories/:id"
	b := "/basket/:id"

	e.GET(p, products.GetProduct, products.CheckIdMiddleware)
	e.PUT(p, products.UpdateProduct, products.CheckIdMiddleware)
	e.DELETE(p, products.DeleteProduct, products.CheckIdMiddleware)
	e.GET("/products", products.GetProducts)
	e.POST("/products", products.CreateProduct)

	e.GET(c, categories.GetCategory, categories.CheckIdMiddleware)
	e.PUT(c, categories.UpdateCategory, categories.CheckIdMiddleware)
	e.DELETE(c, categories.DeleteCategory, categories.CheckIdMiddleware)
	e.GET("/categories", categories.GetCategories)
	e.POST("/categories", categories.CreateCategory)

	e.GET(b, basket.GetBasketItem, basket.CheckIdMiddleware)
	e.PUT(b, basket.UpdateBasketItem, basket.CheckIdMiddleware)
	e.DELETE(b, basket.DeleteBasketItem, basket.CheckIdMiddleware)
	e.GET("/basket", basket.GetBasket)
	e.POST("/basket", basket.AddToBasket)

	e.Logger.Fatal(e.Start(":9000"))
}