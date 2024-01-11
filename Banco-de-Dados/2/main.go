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
	dsn := "root:root@tcp(localhost:3307)/cursogo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	//create
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000.00,
	// })

	//create batch
	// products := []Product{
	// 	{Name: "Notebook", Price: 1000.00},
	// 	{Name: "Mouse", Price: 50.00},
	// 	{Name: "Keyboard", Price: 100.00},
	// }
	// db.Create(products)

	//select one
	// var product Product
	// // db.First(&product, 3)
	// db.First(&product, "name = ?", "Keyboard")
	// fmt.Println(product)

	//select all
	// var products []Product
	// db.Find(&products)

	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//utilizando Limit para limitar a consulta e Offset para paginaçao
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// where
	// var products []Product
	// db.Where("price > ?", 90).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Where("name LIKE ?", "%book%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// println(p2.Name)
	// db.Delete(&p2)
}
