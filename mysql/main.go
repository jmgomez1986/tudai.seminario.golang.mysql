package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Car ...
	type Car struct {
		ID    int
		Brand string
		Model string
	}

	db, err := sql.Open("mysql", "root:root@/tudai")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	queryCreate := `CREATE TABLE IF NOT EXISTS cars (
										id    integer NOT NULL auto_increment,
										brand varchar(255) NOT NULL,
										model varchar(255) NOT NULL,
										PRIMARY KEY (id)
									);`

	stmtCreate, _ := db.Prepare(queryCreate)
	if _, err = stmtCreate.Exec(); err != nil {
		panic(err)
	}

	defer stmtCreate.Close()

	// newCar := Car{
	// 	ID:    2,
	// 	Brand: "Test3",
	// 	Model: "1987",
	// }

	// queryInsert := `INSERT INTO cars (brand, model) VALUES (?, ?);`

	// stmtInsert, _ := db.Prepare(queryInsert)
	// stmtInsert.Exec(newCar.Brand, newCar.Model)

	// defer stmtInsert.Close()

	querySelect1 := `SELECT * FROM cars WHERE id = ?`

	stmtSelect1, _ := db.Prepare(querySelect1)
	var aux Car
	if err = stmtSelect1.QueryRow(1).Scan(&aux.ID, &aux.Brand, &aux.Model); err != nil {
		panic(err)
	}
	fmt.Printf("\nResultado de SELECT de un registro: \n\t%+v\n", aux)
	defer stmtSelect1.Close()

	querySelect2 := `SELECT * FROM cars`
	stmtSelect2, _ := db.Prepare(querySelect2)
	rows, _ := stmtSelect2.Query()

	carsAgency := make(map[int]*Car)

	for rows.Next() {
		var aux1 Car
		if err = rows.Scan(&aux1.ID, &aux1.Brand, &aux1.Model); err != nil {
			panic(err)
		}
		carsAgency[aux1.ID] = &aux1
	}

	fmt.Println("\nResultado de SELECT de varios registros:")
	for k, v := range carsAgency {
		fmt.Printf("\t%v\t%+v\n", k, *v)
	}
	fmt.Print("\n")

}
