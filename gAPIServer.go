package main

import (
	"log"

	"github.com/dgrijalva/jwt-go"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sqlx.DB
var url1CResp string
var url1CReq string

func init() {
	var err error
	url1CResp = "http://10.10.11.158/trade/hs/ObmenLogistica/V1/Document?IMEI="
	url1CReq = "http://10.10.11.158/trade/hs/ObmenLogistica/postjson"
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

func login(c echo.Context) error {
	// macAddress := c.QueryParam("macAddress")
	macAddress := c.FormValue("macAddress")
	var couriers Couriers

	var res1C Response1C
	// From real 1C
	err := res1C.FillFrom1C(FetchDataFromHTTP(url1CResp, macAddress), db)
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

