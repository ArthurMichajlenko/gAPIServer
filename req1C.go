package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// UnmarshalRequest1C ...
func UnmarshalRequest1C(data []byte) (Request1C, error) {
	var r Request1C
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal ...
func (r *Request1C) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//PostTo1C post JSON data to 1C server
func (r *Request1C) PostTo1C(url string) int {
	body, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic TG9naXN0aWM6MTIzNA==")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	log.Println(res.Status)
	return res.StatusCode
}

// Request1C ...
type Request1C struct {
	OrderRoutlist string  `json:"order_routlist"`
	DateRoutlist  string  `json:"date_routlist"`
	ClientID      string  `json:"client_id"`
	ClientName    string  `json:"client_name"`
	ClientTel     string  `json:"client_tel"`
	OrderID       string  `json:"order_id"`
	OrderDate     string  `json:"order_date"`
	PaymentMethod string  `json:"payment_method"`
	OrderCost     float64 `json:"order_cost"`
	Delivered     int     `json:"delivered"`
	DeliveryDelay int     `json:"delivery_delay"`
	DateStart     string  `json:"date_start"`
	DateFinish    string  `json:"date_finish"`
	TimeStamp     string  `json:"time_stamp"`
	Address       string  `json:"address"`
}
