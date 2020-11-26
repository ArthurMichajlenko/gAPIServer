package main

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"

	"net/http"
	"net/url"

	"github.com/dimchansky/utfbom"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sqlx.DB

func init() {
	var err error
	// db, err = sqlx.Connect("mysql", "root:Nfnmzyf@tcp(localhost:3306)/gelibert?parseTime=true&loc=Local")
	db, err = sqlx.Connect("mysql", "gelibert:gelibert@tcp(localhost:3306)/gelibert?parseTime=true&loc=Local")
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
// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello World")
// }
func hello(c echo.Context) error {
	var res Response1C
	jsonFile, err := os.Open("Response1C.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	res.FillFrom1C(jsonFile, db)
	return c.JSON(http.StatusOK, res)
}

func getCouriers(c echo.Context) error {
	var couriers Couriers
	var courier CourierCl
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	macAddress := claims["macAddress"].(string)
	err := db.Select(&couriers, "SELECT id, mac_address, tel, name, car_number FROM couriers WHERE mac_address = ?", macAddress)
	courier.ID = couriers[0].ID
	courier.MacAddress = couriers[0].MacAddress
	courier.Tel = couriers[0].Tel
	courier.Name = couriers[0].Name
	courier.CarNumber = couriers[0].CarNumber
	if err != nil {
		log.Println(err)
	}
	// return c.JSON(http.StatusOK, couriers)
	return c.JSON(http.StatusOK, courier)
}

func putCouriers(c echo.Context) error {
	var courier CourierCl
	if err := c.Bind(&courier); err != nil {
		log.Println(err)
	}
	_, err := db.NamedExec(`INSERT INTO geodata (mac_address, courier_id, latitude, longitude, address) 
							VALUES (:mac_address, :id, :latitude, :longitude, :address)`, &courier)
	if err != nil {
		log.Println(err)
	}
	return c.NoContent(http.StatusNoContent)
}

func getClients(c echo.Context) error {
	var courierID string
	var clientIDs []string
	var client Client
	var clients Clients
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	macAddress := claims["macAddress"].(string)
	err := db.QueryRow("SELECT id FROM couriers WHERE mac_address = ?", macAddress).Scan(&courierID)
	if err != nil {
		log.Println(err)
	}
	err = db.Select(&clientIDs, "SELECT client_id FROM orders WHERE courier_id = ?", courierID)
	if err != nil {
		log.Println(err)
	}
	for _, clientID := range clientIDs {
		err := db.QueryRow("SELECT * FROM clients WHERE id = ?", clientID).Scan(&client.ID, &client.Name, &client.Tel)
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
	macAddress := claims["macAddress"]
	err := db.Select(&couriers, "SELECT * FROM couriers WHERE mac_address = ?", macAddress)
	if err != nil {
		log.Println(err)
	}
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
		err := db.Select(&order.Consists, "SELECT * FROM consists WHERE orders_id = ?", order.ID)
		if err != nil {
			log.Println(err)
		}
		orders[i].Consists = order.Consists
	}
	return c.JSON(http.StatusOK, orders)
}

func putOrders(c echo.Context) error {
	var (
		order  Order
		orders Orders
	)
	if err := c.Bind(&order); err != nil {
		log.Println(err)
	}
	err := db.Select(&orders, "SELECT * FROM orders WHERE id = ?", order.ID)
	if err != nil {
		log.Println(err)
	}
	// err = db.Select(&order.Consists, "SELECT product, quantity, price, ext_info FROM consists WHERE orders_id = ?", order.ID)
	err = db.Select(&order.Consists, "SELECT product, quantity, price, ext_info FROM consists WHERE orders_id = ?", order.ID)
	if err != nil {
		log.Println(err)
	}
	orders[0].Consists = order.Consists
	return c.JSON(http.StatusOK, orders)
}

func login(c echo.Context) error {
	var couriers Couriers
	var res1C Response1C
	// macAddress := c.QueryParam("macAddress")
	macAddress := c.FormValue("macAddress")

	url := "http://10.10.11.158/trade/hs/ObmenLogistica/V1/Document?IMEI=" + url.QueryEscape(macAddress)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", "Basic TG9naXN0aWM6MTIzNA==")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	res1C.FillFrom1C(utfbom.SkipOnly(res.Body), db)

	err = db.Select(&couriers, "SELECT * FROM couriers WHERE mac_address = ?", macAddress)
	if err != nil {
		log.Println(err)
	}
	if couriers == nil {
		return echo.ErrUnauthorized
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["macAddress"] = macAddress
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	t, err := token.SignedString([]byte("gelibert"))
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
