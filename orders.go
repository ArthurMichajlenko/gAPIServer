package main

import (
	"encoding/json"
	"time"
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
	DateStart     time.Time `json:"date_start" db:"date_start"`
	DateFinish    time.Time `json:"date_finish" db:"date_finish"`
	TimeStamp     time.Time `json:"timestamp" db:"timestamp"`
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
