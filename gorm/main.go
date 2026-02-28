package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	db.Create(&Product{
		Name:  "Notebook",
		Price: 1999.90,
	})

	products := []Product{
		{Name: "Mouse", Price: 100.00},
		{Name: "Teclado", Price: 200.00},
		{Name: "Monitor", Price: 1500.00},
	}
	db.Create(&products)

	var product Product
	db.First(&product, "name = ?", "Notebook")
	fmt.Println(product)

	var productsList []Product
	db.Limit(2).Offset(2).Find(&productsList)

	for _, product := range productsList {
		fmt.Println(product)
	}

	db.Where("name = ?", "Teclado").Find(&products)
	for _, product := range productsList {
		fmt.Println(product)
	}
}
