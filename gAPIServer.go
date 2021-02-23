package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

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
var url1C string

func init() {
	var err error
	url1C = "http://10.10.11.158/trade/hs/ObmenLogistica/V1/Document?IMEI="
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
	g.POST("/geodata", postGeodata)
	g.GET("/clients", getClients)
	g.GET("/orders", getOrders)
	g.POST("/orders", postOrders)
	// Start server
	// e.HideBanner=true
	e.Logger.Fatal(e.Start(":1323"))

}

// Handlers
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func getCouriers(c echo.Context) error {
	var couriers Couriers
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	macAddress := claims["macAddress"].(string)
	err := db.Select(&couriers, "SELECT id, mac_address, tel, name, car_number FROM couriers WHERE mac_address = ?", macAddress)
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, couriers[0])
}

func postGeodata(c echo.Context) error {
	var geodata Geodata
	if err := c.Bind(&geodata); err != nil {
		log.Println(err)
	}
	_, err := db.NamedExec(`REPLACE INTO geodata (id, mac_address, courier_id, latitude, longitude, timestamp) 
							VALUES (:id, :mac_address, :courier_id, :latitude, :longitude, :timestamp)`, &geodata)
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
	date := time.Now().Format("2006-01-02")
	// date := time.Now().AddDate(0,0,1).Format("2006-01-02")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	macAddress := claims["macAddress"].(string)

	var res1C Response1C
	// From real 1C
	err := res1C.FillFrom1C(FetchDataFromHTTP(url1C, macAddress), db)
	if err != nil {
		log.Println(err)
	}
	// For debug JSON file
	// res1C.FillFrom1C(FetchDataFromFile("Response1C_oppo.json", macAddress), db)

	err = db.Select(&couriers, "SELECT * FROM couriers WHERE mac_address = ?", macAddress)
	if err != nil {
		log.Println(err)
	}
	if c.QueryParam("client") == "" {
		err := db.Select(&orders, `SELECT * FROM orders 
								  WHERE courier_id = ? AND date_start >= ? AND date_start < DATE_ADD(?, INTERVAL 1 DAY)`, couriers[0].ID, date, date)
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
	if orders == nil {
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusOK, orders)
}

func postOrders(c echo.Context) error {
	var (
		// order  Order
		orders Orders
	)
	if err := c.Bind(&orders); err != nil {
		log.Println(err)
	}
	log.Println(orders)
	return c.NoContent(http.StatusOK)
}

func login(c echo.Context) error {
	// macAddress := c.QueryParam("macAddress")
	macAddress := c.FormValue("macAddress")
	var couriers Couriers
	
	var res1C Response1C
	// From real 1C
	err := res1C.FillFrom1C(FetchDataFromHTTP(url1C, macAddress), db)
	if err != nil {
		log.Println(err)
	}
	// For debug JSON file
	// err := res1C.FillFrom1C(FetchDataFromFile("Response1C_oppo.json", macAddress), db)
	// if err != nil {
	// 	log.Print(err)
	// }

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

// FetchDataFromHTTP read data from 1C server
func FetchDataFromHTTP(srcURL string, macAddress string) []byte {
	url := srcURL + url.QueryEscape(macAddress)
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
	content, err := ioutil.ReadAll(utfbom.SkipOnly(res.Body))
	if err != nil {
		log.Println(err)
	}
	return content
}

// FetchDataFromFile read JSON file for debug purpose
func FetchDataFromFile(file string, macAddress string) []byte {
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()
	content, err := ioutil.ReadAll(utfbom.SkipOnly(jsonFile))
	if err != nil {
		log.Println(err)
	}
	return content
}
