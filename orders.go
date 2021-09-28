package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// Orders slice of Order
type Orders []Order

// UnmarshalOrders write Orders from JSON
func UnmarshalOrders(data []byte) (Orders, error) {
	var r Orders
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal write JSON from Orders
func (r *Orders) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Order ...
type Order struct {
	ID            string    `json:"id" db:"id"`
	OrderRoutlist string    `json:"order_routlist" db:"order_routlist"`
	OrderDate     string    `json:"order_date" db:"order_date"`
	CourierID     string    `json:"courier_id" db:"courier_id"`
	ClientID      string    `json:"client_id" db:"client_id"`
	PaymentMethod string    `json:"payment_method" db:"payment_method"`
	Consists      []Consist `json:"consists" db:"consists"`
	OrderCost     float64   `json:"order_cost" db:"order_cost"`
	Delivered     int       `json:"delivered" db:"delivered"`
	DeliveryDelay int       `json:"delivery_delay" db:"delivery_delay"`
	DateStart     string    `json:"date_start" db:"date_start"`
	DateFinish    string    `json:"date_finish" db:"date_finish"`
	TimeStamp     string    `json:"timestamp" db:"timestamp"`
	Address       string    `json:"address" db:"address"`
}

//Consist products of Order
type Consist struct {
	ID        int     `json:"id" db:"id"`
	Product   string  `json:"product" db:"product"`
	Quantity  float64 `json:"quantity" db:"quantity"`
	Price     float64 `json:"price" db:"price"`
	ExtInfo   string  `json:"ext_info" db:"ext_info"`
	Direction int     `json:"direction" db:"direction"`
	OrdersID  string  `json:"orders_id" db:"orders_id"`
}

func getOrders(c echo.Context) error {
	emptyConsists := []Consist{{ID:0, Product:"Empty", Quantity: 0.00, Price: 0.00, ExtInfo: "Empty", Direction: 0, OrdersID: "Empty"}}
	var orders Orders
	var couriers Couriers
	date := time.Now().Format("2006-01-02")
	// date := time.Now().AddDate(0,0,1).Format("2006-01-02")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	macAddress := claims["macAddress"].(string)

	var res1C Response1C
	// From real 1C
	err := res1C.FillFrom1C(FetchDataFromHTTP(url1CResp, macAddress), db)
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
		if len(order.Consists) != 0 {
			orders[i].Consists = order.Consists
		} else {
			log.Println("Consists is empty")
			emptyConsists[0].OrdersID=order.ID
			orders[i].Consists=emptyConsists
		}
	}
	if orders == nil {
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusOK, orders)
}

func postOrders(c echo.Context) error {
	var orders Orders

	if err := c.Bind(&orders); err != nil {
		log.Println(err)
	}
	for _, order := range orders {
		var post1C Request1C
		var client Client
		err := db.Get(&client, "SELECT * FROM clients WHERE id = ?", order.ClientID)
		if err != nil {
			log.Println(err)
		}
		post1C.ClientID = client.ID
		post1C.ClientName = client.Name
		post1C.ClientTel = client.Tel
		post1C.OrderID = order.ID
		post1C.OrderRoutlist = order.OrderRoutlist
		post1C.OrderDate = order.OrderDate
		post1C.PaymentMethod = order.PaymentMethod
		post1C.OrderCost = order.OrderCost
		post1C.Delivered = order.Delivered
		post1C.DeliveryDelay = order.DeliveryDelay
		post1C.DeliveryDelay = order.DeliveryDelay
		post1C.DateStart = order.DateStart
		post1C.DateFinish = order.DateFinish
		post1C.TimeStamp = order.TimeStamp
		post1C.Address = order.Address
		code := post1C.PostTo1C(url1CReq)
		if code != 200 {
			log.Printf("Error to connect 1C server. Code: %v\n", code)
		}
		_, err = db.NamedExec(`REPLACE INTO orders 
		(id, order_routlist, order_date, courier_id, client_id, payment_method, order_cost, delivered, delivery_delay, date_start, date_finish, address) VALUES 
		(:id, :order_routlist, :order_date, :courier_id, :client_id, :payment_method, :order_cost, :delivered, :delivery_delay, :date_start, :date_finish, :address)`, &order)
		if err != nil {
			log.Println(err)
		}
		for _, consist := range order.Consists {
			_, err = db.NamedExec(`REPLACE INTO consists
			(id, product, quantity, price, ext_info, direction, orders_id) VALUES
			(:id, :product, :quantity, :price, :ext_info, :direction, :orders_id)`, &consist)
			if err != nil {
				log.Println(err)
			}
		}
	}
	return c.NoContent(http.StatusOK)
}
