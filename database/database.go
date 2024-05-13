package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func FilterByCategory(category uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("category_id = ?", category)
	}
}

func FilterByName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", name)
	}
}

func FilterByPriceBelow(price float32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price < ?", price)
	}
}

func FilterByPriceAbove(price float32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price > ?", price)
	}
}

// Create by hand because referential integrity doesn't work as it should
func createTables() {
	DB.Exec(`
		
		CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT
		);
	
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			price REAL,
			category_id INTEGER,
			FOREIGN KEY (category_id) REFERENCES categories(id)
		);
	
		CREATE TABLE IF NOT EXISTS basket (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			amount INTEGER,
			product_id INTEGER,
			FOREIGN KEY (product_id) REFERENCES products(id)
		);
	
	`)
}

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	
	if err != nil {
		panic(err)
	}
	db.Exec("PRAGMA foreign_keys = ON")
	DB = db
	createTables()
}
