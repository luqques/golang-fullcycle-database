package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID uint
	Category   Category
	gorm.Model
}

type Category struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	//create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	//create product
	product := Product{
		Name:       "Notebook",
		Price:      1999.90,
		CategoryID: category.ID,
	}
	db.Create(&product)
}
