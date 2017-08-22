package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tzmfreedom/go-todo/models"
	"os"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Rollback!!\n")
		return
	}
	// Migrate the schema
	db.AutoMigrate(&model.Product{})
}
