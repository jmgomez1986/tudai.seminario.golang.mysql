package main

import (
	// "fmt"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //You could import dialect
)

func main() {

	// Car ...
	type Car struct {
		gorm.Model
		Brand      string `gorm:"column:brand;type:varchar(255)"`
		BrandModel string `gorm:"column:model;type:varchar(255)"`
	}

	db, err := gorm.Open("mysql", "root:root@/tudai?parseTime=true")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&Car{})
	
	// Insertar una fila
	// db.Create(&Car{
	// 	Brand:      "lambogginni",
	// 	BrandModel: "diablo",
	// })

	// Buscar por ID
	var myCar Car
	db.First(&myCar, 1)

	// fmt.Printf("\nResultado de SELECT de un registro: \n\t%+v\n", myCar)

	// Buscar todos
	var myCars []Car
	db.Find(&myCars)

	// fmt.Println(myCars)

	// fmt.Println("\nResultado de SELECT de varios registros:")
	// for k, v := range myCars {
	// 	fmt.Printf("\t%v %+v\n", k, v)
	// }
	// fmt.Print("\n")

	// Buscar por columna
	var aux Car
	db.First(&aux, "model = ?", "golf")

	// fmt.Printf("%+v\n", aux)

	var rows []Car
	db.Find(&rows, "model = ?", "diablo")

	fmt.Println("\nResultado de SELECT de varios registros:")
	for k, v := range rows {
		fmt.Printf("\t%v %+v\n", k, v)
	}
	fmt.Print("\n")

}
