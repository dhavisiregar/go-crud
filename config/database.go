package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:Redline-2@/go_products?parseTime=true")
	if err != nil {
	panic(err)
	}

	log.Println("Database connected")
	DB = db
}