package main

import (
	"encoding/json"
	"time"
)

// Couriers is array of courier
type Couriers []Courier

// UnmarshalCouriers decode couriers from JSON
func UnmarshalCouriers(data []byte) (Couriers, error) {
	var r Couriers
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal encode couriers to JSON
func (r *Couriers) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Courier is a single courier
type Courier struct {
	ID        int       `json:"id" db:"id"`
	Imei      int64     `json:"imei" db:"imei"`
	Tel       string    `json:"tel" db:"tel"`
	Name      string    `json:"name" db:"name"`
	CarNumber string    `json:"car_number" db:"car_number"`
	Latitude  float64   `json:"latitude" db:"latitude"`
	Longitude float64   `json:"longitude" db:"longitude"`
	Address   string    `json:"address" db:"address"`
	TimeStamp time.Time `json:"timestamp" db:"timestamp"`
}
