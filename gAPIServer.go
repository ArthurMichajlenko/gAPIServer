package main

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("sqlite3", "gelibert.db")
	// db, err = sqlx.Connect("mysql", "root:Nfnmzyf@tcp(localhost:3306)/gelibert")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	log.Println("Hello World")
	var orders Orders
	err:=db.Select(&orders, "SELECT * FROM orders")
	if err != nil {
		log.Println(err)
	}
	log.Println(orders)
}

