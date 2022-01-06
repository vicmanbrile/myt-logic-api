package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (db *sql.DB, err error) {
	usr := os.Getenv("USER_MYSQL")
	psswd := os.Getenv("PASSWORD_MYSQL")
	schm := os.Getenv("SCHEMA_MYSQL")

	db, err = sql.Open(Credential(usr, psswd, schm))

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return
}

func Credential(user, password, dbname string) (db, credential string) {
	credential = fmt.Sprintf("%s:%s@tcp(192.168.1.31:3306)/%s?parseTime=true", user, password, dbname)
	db = "mysql"

	return
}
