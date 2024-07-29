// Banco de dados com GORM
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	ProductID int
	Number    string
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category `gorm:"foreignKey:CategoryID"`
	SerialNumber SerialNumber
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3307)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Category{}, &Product{}, &SerialNumber{})

	// Create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// Create product
	product := Product{Name: "Mouse", Price: 100, CategoryID: 1}
	db.Create(&product)

	// Create serial number
	serialNumber := SerialNumber{ProductID: 1, Number: "123456"}
	db.Create(&serialNumber)

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, p := range products {
		fmt.Println(p.Name, p.Category.Name, p.SerialNumber.Number)
	}
}
