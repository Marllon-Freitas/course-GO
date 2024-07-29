// lock otimista vs pessimista
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:product_categories;"`
	gorm.Model
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:product_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3307)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}, &Category{})

	tx := db.Begin() // inicia uma transação
	if tx.Error != nil {
		panic(tx.Error)
	}

	var category Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&category, 1).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	category.Name = "Eletronicos"
	err = tx.Debug().Save(&category).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	tx.Commit()
}
