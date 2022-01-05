package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/vicmanbrile/test/database"
)

type Stock struct {
	Id    int    `json:"id"`
	Clave string `json:"clave"`
	Valor string `json:"valor"`
}

func FilterStock(ID string) []byte {
	var s Stock

	conn := database.Connection()
	defer conn.Close()

	query := "SELECT id, clave, valor FROM logic.almacenes WHERE id = ?"
	if err := conn.QueryRow(query, ID).Scan(&s.Id, &s.Clave, &s.Valor); err != nil {
		log.Print(err)
	}

	result, _ := json.Marshal(s)

	return result
}

func InsertStock(clave, valor string) {
	conn := database.Connection()
	defer conn.Close()

	result, err := conn.Exec(`INSERT INTO logic.almacenes (clave, valor) VALUES (?, ?);`, clave, valor)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("%v\n", id)
}

func SelectAllStock() []byte {
	conn := database.Connection()
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM logic.almacenes`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var stock []Stock
	for rows.Next() {
		var u Stock

		err := rows.Scan(&u.Id, &u.Clave, &u.Valor)
		if err != nil {
			log.Fatal(err)
		}
		stock = append(stock, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	r, _ := json.Marshal(stock)

	return r
}

func DeleteStock(Num int) {
	conn := database.Connection()
	defer conn.Close()

	_, err := conn.Exec(`DELETE FROM logic.almacenes WHERE id = ?`, Num)
	if err != nil {
		log.Fatal(err)
	}
}
