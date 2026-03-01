package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   uint
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type SerialNumber struct {
	ID        uint `gorm:"primaryKey"`
	Number    string
	ProductID uint
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	deletarTabelas(db)
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

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

	//create serial number
	serialNumber := SerialNumber{
		Number:    "123456789",
		ProductID: product.ID,
	}
	db.Create(&serialNumber)

	//has many
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		print(category.Name)
		for _, product := range category.Products {
			println(" -", product.Name)
		}
	}

	var products = []Product{}
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		println(product.Name, product.Price, product.Category.Name, product.SerialNumber.Number)
	}
}

func deletarTabelas(db *gorm.DB) {
	var tables []string
	db.Raw(`SHOW TABLES`).Scan(&tables)

	for _, table := range tables {
		db.Migrator().DropTable(table)
	}
}
