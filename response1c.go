package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
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
func (r *Response1C) FillFrom1C(src io.Reader) error {
	data, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}
	*r, err = UnmarshalResponse1C(data)
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
	ClientID      string `json:"client_id"`
	ClientName    string `json:"client_name"`
	ClientTel     string `json:"client_tel"`
	OrderID       string `json:"order_id"`
	PaymentMethod string `json:"payment_method"`
	OrderCost     int64  `json:"order_cost"`
	Delivered     string `json:"delivered"`
	DeliveryDelay string `json:"delivery_delay"`
	DateStart     string `json:"date_start"`
	DateFinish    string `json:"date_finish"`
	TimeStamp     string `json:"time_stamp"`
	Address       string `json:"address"`
}

// Consist1C ...
type Consist1C struct {
	ID        string `json:"id"`
	Product   string `json:"product"`
	Quantity  int64  `json:"quantity"`
	Price     int64  `json:"price"`
	EXTInfo   string `json:"ext_info"`
	Direction string `json:"direction"`
}
