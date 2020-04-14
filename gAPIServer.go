package main

import (
	"log"

	"github.com/dgrijalva/jwt-go"

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
	// db, err = sqlx.Connect("sqlite3", "gelibert.db")
	db, err = sqlx.Connect("mysql", "root:Nfnmzyf@tcp(localhost:3306)/gelibert?parseTime=true&loc=Local")
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
	e.POST("/login", login)
	g := e.Group("/data", middleware.JWT([]byte("gelibert")))
	g.GET("/couriers", getCouriers)
	g.PUT("/couriers", putCouriers)
	g.GET("/clients", getClients)
	g.GET("/orders", getOrders)
	g.PUT("/orders", putOrders)
	// Start server
	// e.HideBanner=true
	e.Logger.Fatal(e.Start(":1323"))

}

// Handlers
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func getCouriers(c echo.Context) error {
	var couriers Couriers
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	imei := claims["imei"].(string)
	err := db.Select(&couriers, "SELECT * FROM couriers WHERE imei = ?", imei)
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, couriers)
}

func putCouriers(c echo.Context) error {
	var courier Courier
	// var couriers Couriers
	if err := c.Bind(&courier); err != nil {
	// if err := c.Bind(&couriers); err != nil {
		log.Println(err)
	}
	log.Println(courier.ID)
	return c.NoContent(http.StatusNoContent)
	// return c.JSON(http.StatusOK, courier.ID)
	// return c.String(http.StatusOK, "Test PUT courier")
}

func getClients(c echo.Context) error {
	var courierID int
	var clientIDs []int
	var client Client
	var clients Clients
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	imei := claims["imei"].(string)
	row := db.QueryRowx("SELECT id FROM couriers WHERE imei = ?", imei)
	err := row.Scan(&courierID)
	if err != nil {
		log.Println(err)
	}
	err = db.Select(&clientIDs, "SELECT client_id FROM orders WHERE courier_id = ?", courierID)
	if err != nil {
		log.Println(err)
	}
	for _, clientID := range clientIDs {
		row := db.QueryRowx("SELECT * FROM clients WHERE id = ?", clientID)
		err := row.Scan(&client.ID, &client.Name, &client.Tel, &client.Address)
		if err != nil {
			log.Println(err)
		}
		clients = append(clients, client)
	}
	return c.JSON(http.StatusOK, clients)
}

func getOrders(c echo.Context) error {
	var orders Orders
	var couriers Couriers
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	imei := claims["imei"].(string)
	err := db.Select(&couriers, "SELECT * FROM couriers WHERE imei = ?", imei)
	if err != nil {
		log.Println(err)
	}
	// log.Println(imei, couriers[0].Name, couriers[0].ID)
	if c.QueryParam("client") == "" {
		err := db.Select(&orders, "SELECT * FROM orders WHERE courier_id = ?", couriers[0].ID)
		if err != nil {
			log.Println(err)
		}
	} else {
		err := db.Select(&orders, "SELECT * FROM orders WHERE client_id = ? AND courier_id = ?", c.QueryParam("client"), couriers[0].ID)
		if err != nil {
			log.Println(err)
		}
	}
	// err = db.Select(&orders, "SELECT * FROM orders WHERE courier_id = ? AND client_id = ?", c.QueryParam("courier"), c.QueryParam("client"))
	// if err != nil {
	// 	log.Println(err)
	// }
	for i, order := range orders {
		err := db.Select(&order.ConsistsTo, "SELECT product, quantity, price, ext_info FROM consists_to WHERE id = ?", order.ID)
		if err != nil {
			log.Println(err)
		}
		orders[i].ConsistsTo = order.ConsistsTo
	}
	for i, order := range orders {
		err := db.Select(&order.ConsistsFrom, "SELECT product, quantity, price, ext_info FROM consists_from WHERE id = ?", order.ID)
		if err != nil {
			log.Println(err)
		}
		orders[i].ConsistsFrom = order.ConsistsFrom
	}
	return c.JSON(http.StatusOK, orders)
}

func putOrders(c echo.Context) error {
	return c.String(http.StatusOK, "Test Put orders")
}

func login(c echo.Context) error {
	var couriers Couriers
	// imei := c.QueryParam("imei")
	imei := c.FormValue("imei")
	err := db.Select(&couriers, "SELECT * FROM couriers WHERE imei = ?", imei)
	if err != nil {
		log.Println(err)
	}
	if couriers == nil {
		return echo.ErrUnauthorized
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["imei"] = imei
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	t, err := token.SignedString([]byte("gelibert"))
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
