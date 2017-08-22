package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tzmfreedom/go-todo/handlers"
)

func main() {
	db, err := dbInit()
	if err != nil {
		panic("DB error")
	}
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	product := handlers.Product{DB: db}
	r.GET("/", product.GetIndex)
	r.GET("/products", product.GetProducts)
	r.GET("/products/new", product.GetProductNew)
	r.POST("/products", product.PostProduct)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func dbInit() (*gorm.DB, error) {
	return gorm.Open("sqlite3", "test.db")
}
