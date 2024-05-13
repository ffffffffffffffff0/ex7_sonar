package models

import "gorm.io/gorm"

type Category struct {
	ID 			uint 		`gorm:"primaryKey"`	
	Name 		string 		`json:"name"`
}

func (Category) TableName() string {
	return "categories"
}

type BasketItem struct {
	ID 			uint 		`gorm:"primaryKey"`
	Amount		uint		`json:"amount"`
	ProductID	uint		`json:"product"`	
}

func (BasketItem) TableName() string {
	return "basket"
}

type Product struct {
	ID 			uint 		`gorm:"primaryKey"`
	Name 		string 		`json:"name"`
	Price 		float32 	`json:"price"` 
	CategoryID 	uint 		`json:"category"`
}

func (Product) TableName() string {
	return "products"
}

func Exists(ID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", ID)
	}
}

func ProductInBasketExists(db *gorm.DB, ID uint) *gorm.DB {
	return db.Where("product_id = ?", ID)
}