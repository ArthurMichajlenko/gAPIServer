package main

import (
	"log"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.GET("/", hello)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}

// Handler
func hello(c echo.Context) error {
	var orders Orders
	err := db.Select(&orders, "SELECT * FROM orders")
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, &orders)
}
