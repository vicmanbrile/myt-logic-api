package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connection() (db *sql.DB) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file")
	}

	usr := os.Getenv("USER_MYSQL")
	psswd := os.Getenv("PASSWORD_MYSQL")
	schm := os.Getenv("SCHEMA_MYSQL")

	db, err = sql.Open(Credential(usr, psswd, schm))

	if err != nil {
		log.Print(err)
	}

	if err := db.Ping(); err != nil {
		log.Print(err)
	}

	return
}

func Credential(user, password, dbname string) (db, credential string) {
	credential = fmt.Sprintf("%s:%s@tcp(192.168.1.31:3306)/%s?parseTime=true", user, password, dbname)
	db = "mysql"

	return
}
