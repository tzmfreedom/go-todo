package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name  string
	Code  string
	Price uint
}

func GetProductsAll(db *gorm.DB) []Product {
	products := []Product{}
	db.Find(&products)
	return products
}
