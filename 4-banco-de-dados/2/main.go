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
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3307)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	// Create
	// db.Create(&Product{Name: "Dell XPS 13", Price: 1000})

	// Create many
	// products := []Product{
	// 	{Name: "Macbook Pro", Price: 2000},
	// 	{Name: "Thinkpad X1", Price: 1500},
	// }
	// db.Create(&products)

	// Select one
	// var product Product
	// db.First(&product, 1)
	// db.First(&product, "name = ?", "Dell XPS 13")
	// fmt.Println(product)

	// Select many
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// limit and offset
	// var products []Product
	// db.Limit(2).Offset(1).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// Where
	// var product []Product
	// db.Where("price > ?", 1000).Find(&product)
	// for _, p := range product {
	// 	fmt.Println(p)
	// }
	// db.Where("name LIKE ?", "%Dell%").Find(&product)
	// for _, p := range product {
	// 	fmt.Println(p)
	// }

	// Update
	// var product Product
	// db.First(&product, 1)
	// product.Price = 1100
	// db.Save(&product)

	// var products02 Product
	// db.First(&products02, 1)

	// Delete
	// db.Delete(&products02)
}
