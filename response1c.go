package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Response1C ...
type Response1C []Response1CElement

// UnmarshalResponse1C write response from 1C to structure
func UnmarshalResponse1C(data []byte) (Response1C, error) {
	var r Response1C
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal method write structure to jSON
func (r *Response1C) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// FillFrom1C ...
func (r *Response1C) FillFrom1C(data []byte, db *sqlx.DB) error {
	var err error
	*r, err = UnmarshalResponse1C(data)
	for _, res := range *r {
		// for _, client := range (*r)[0].Clients {
		var courier Courier
		if res.CourierID == "" {
			continue
		}
		err := db.Get(&courier, "SELECT id FROM couriers WHERE id = ?", res.CourierID)
		switch err {
		case nil:
			// log.Println("Courier found")
			_, err1 := db.Exec(`UPDATE couriers 
								SET tel=?, name=?, car_number=?, mac_address=? 
								WHERE id=?`, res.CourierTel, res.CourierName, res.CourierCarNumber, res.CourierImei, res.CourierID)
			if err1 != nil {
				log.Println(err1)
			}
		case sql.ErrNoRows:
			// log.Println("Courier not found")
			_, err1 := db.Exec(`INSERT INTO 
								couriers (id, mac_address, tel, name, car_number) 
								VALUES (?, ?, ?, ?, ?)`, res.CourierID, res.CourierImei, res.CourierTel, res.CourierName, res.CourierCarNumber)
			if err1 != nil {
				log.Println(err1)
			}
		default:
			log.Println(err)
		}
		for _, client := range res.Clients {
			err := db.Get(&client, "SELECT id FROM clients WHERE id = ?", client.ClientID)
			switch err {
			case nil:
				// log.Println("Client found")
				_, err1 := db.Exec(`UPDATE clients
									SET name=?, tel=?
									WHERE id=?`, client.ClientName, client.ClientTel, client.ClientID)
				if err1 != nil {
					log.Println(err1)
				}
			case sql.ErrNoRows:
				// log.Println("Client not found")
				_, err1 := db.Exec(`INSERT INTO 
									clients (id, name, tel) 
									VALUES (?, ?, ?)`, client.ClientID, client.ClientName, client.ClientTel)
				if err1 != nil {
					log.Println(err1)
				}
			default:
				log.Println(err)
			}
			err = db.Get(&client, "SELECT id FROM orders WHERE id = ?", client.OrderID)
			switch err {
			case nil:
				// log.Println("Order found")
				_, err1 := db.Exec(`UPDATE orders 
									SET courier_id=?, payment_method=?, order_cost=?, address=?, order_routlist=?, order_date=? 
									WHERE id=?`, res.CourierID, client.PaymentMethod, client.OrderCost, client.Address, client.OrderRoutlist, client.OrderDate, client.OrderID)
				if err1 != nil {
					log.Println(err1)
				}
			case sql.ErrNoRows:
				// log.Println("Orders not found")
				_, err1 := db.Exec(`INSERT INTO 
									orders (id, courier_id, client_id, payment_method, order_cost, address, order_routlist, order_date, date_start) 
									VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
									`, client.OrderID, res.CourierID, client.ClientID, client.PaymentMethod, client.OrderCost, client.Address, client.OrderRoutlist, client.OrderDate, time.Now().Format("2006-01-02 15:04:05"))
				if err1 != nil {
					log.Println(err1)
				}
			default:
				log.Println(err)
			}
		}
		for _, consist := range res.Consists {
			var direction int
			// var tmpConsist Consist
			var tmpID int
			if consist.Direction == "" {
				direction = 0
			} else {
				direction = 1
			}
			err := db.Get(&tmpID, `SELECT id FROM consists 
			WHERE product = ? AND quantity = ? AND price = ? AND orders_id = ? AND direction = ?`, consist.Product, consist.Quantity, consist.Price, consist.ID, direction)
			if err == sql.ErrNoRows && tmpID == 0 {
				_, err := db.Exec(`INSERT INTO 
								consists (product, quantity, price, ext_info, orders_id, direction) 
								VALUES (?, ?, ?, ?, ?, ?)`, consist.Product, consist.Quantity, consist.Price, consist.ExtInfo, consist.ID, direction)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
	return err
}

// Response1CElement ...
type Response1CElement struct {
	CourierID        string      `json:"courier_id"`
	CourierImei      string      `json:"courier_imei"`
	CourierTel       string      `json:"courier_tel"`
	CourierName      string      `json:"courier_name"`
	CourierCarNumber string      `json:"courier_car_number"`
	CourierTimestamp string      `json:"courier_timestamp"`
	Clients          []Client1C  `json:"Clients"`
	Consists         []Consist1C `json:"Consists"`
}

// Client1C ...
type Client1C struct {
	OrderRoutlist string  `json:"order_routlist"`
	ClientID      string  `json:"client_id" db:"id"`
	ClientName    string  `json:"client_name"`
	ClientTel     string  `json:"client_tel"`
	OrderID       string  `json:"order_id"`
	OrderDate     string  `json:"order_date"`
	PaymentMethod string  `json:"payment_method"`
	OrderCost     float64 `json:"order_cost"`
	Delivered     string  `json:"delivered"`
	DeliveryDelay string  `json:"delivery_delay"`
	DateStart     string  `json:"date_start"`
	DateFinish    string  `json:"date_finish"`
	TimeStamp     string  `json:"time_stamp"`
	Address       string  `json:"address"`
}

// Consist1C ...
type Consist1C struct {
	ID        string  `json:"id" db:"orders_id"`
	Product   string  `json:"product"`
	Quantity  float64 `json:"quantity"`
	Price     float64 `json:"price"`
	ExtInfo   string  `json:"ext_info" db:"ext_info"`
	Direction string  `json:"direction"`
}
