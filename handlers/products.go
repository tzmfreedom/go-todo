package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tzmfreedom/go-todo/models"
	"net/http"
	"strconv"
)

type Product struct {
	DB *gorm.DB
}

func (p *Product) GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}

func (p *Product) GetProducts(c *gin.Context) {
	products := model.GetProductsAll(p.DB)
	c.HTML(http.StatusOK, "products.tmpl", gin.H{
		"products": products,
	})
}

func (p *Product) GetProductNew(c *gin.Context) {
	c.HTML(http.StatusOK, "products.new.tmpl", gin.H{})
}

func (p *Product) PostProduct(c *gin.Context) {
	name := c.PostForm("name")
	code := c.PostForm("code")
	price, _ := strconv.ParseUint(c.PostForm("price"), 10, 32)

	product := model.Product{Name: name, Code: code, Price: uint(price)}
	p.DB.Create(&product)
	c.Redirect(http.StatusFound, "/products")
}
