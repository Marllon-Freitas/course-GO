// Banco de dados com GORM
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
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

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// Create category
	category := Category{Name: "Casa e Jardim"}
	db.Create(&category)

	// Create product
	product := Product{Name: "Mesa", Price: 500, CategoryID: category.ID}
	db.Create(&product)

	// Create serial number
	serialNumber := SerialNumber{ProductID: product.ID, Number: "123456"}
	db.Create(&serialNumber)

	var categories []Category
	// err = db.Model(&Category{}).Preload("Products").Preload("SerialNumber").Find(&categories).Error // isso não funciona
	// por que não existe um relacionamento entre Category e SerialNumber
	// o jeito correto é fazer o preload de Products e depois de SerialNumber
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Println(c.Name, ":")
		for _, p := range c.Products {
			fmt.Println(" - ", p.Name, p.Price, "Serial number: ", p.SerialNumber.Number)
		}
	}
}
