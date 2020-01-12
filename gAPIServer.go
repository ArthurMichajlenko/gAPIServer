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
	e.GET("/couriers", getCouriers)
	e.GET("/clients", getClients)
	e.GET("/orders", getOrders)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}

// Handlers
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func getCouriers(c echo.Context) error {
	var couriers Couriers
	err := db.Select(&couriers, "SELECT * FROM couriers")
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, &couriers)
}

func getClients(c echo.Context) error {
	var clients Clients
	err := db.Select(&clients, "SELECT * FROM clients")
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, &clients)
}

func getOrders(c echo.Context) error {
	var orders Orders
	err := db.Select(&orders, "SELECT * FROM orders")
	if err != nil {
		log.Println(err)
	}
	for i, order := range orders {
		err := db.Select(&order.ConsistsTo, "SELECT product, quantity, price FROM consists_to WHERE id = ?", order.ID)
		if err != nil {
			log.Println(err)
		}
		orders[i].ConsistsTo = order.ConsistsTo
	}
	for i, order := range orders {
		err := db.Select(&order.ConsistsFrom, "SELECT product, quantity, price FROM consists_from WHERE id = ?", order.ID)
		if err != nil {
			log.Println(err)
		}
		orders[i].ConsistsFrom = order.ConsistsFrom
	}
	return c.JSON(http.StatusOK, &orders)
}

