// Banco de dados com GORM
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3307)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Name: "Dell XPS 13", Price: 1000})

	// Create many
	products := []Product{
		{Name: "Macbook Pro", Price: 2000},
		{Name: "Thinkpad X1", Price: 1500},
	}
	db.Create(&products)
}
