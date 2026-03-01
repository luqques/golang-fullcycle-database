package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:product_categories;"`
	gorm.Model
}

type Category struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:product_categories;"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	deletarTabelas(db)
	db.AutoMigrate(&Product{}, &Category{})

	//create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	category2 := Category{Name: "Informatica"}
	db.Create(&category2)

	//create product
	product := Product{
		Name:       "Notebook",
		Price:      1999.90,
		Categories: []Category{category, category2},
	}
	db.Create(&product)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.Name)
		for _, product := range category.Products {
			println("-", product.Name)
		}
	}

	//lock pessimista
	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	c.Name = "Eletronicos e Informatica"
	err = tx.Debug().Save(&c).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
}

func deletarTabelas(db *gorm.DB) {
	var tables []string
	db.Raw(`SHOW TABLES`).Scan(&tables)

	for _, table := range tables {
		db.Migrator().DropTable(table)
	}
}
