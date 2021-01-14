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
	ID         string    `json:"id" db:"id"`
	MacAddress string    `json:"mac_address" db:"mac_address"`
	Tel        string    `json:"tel" db:"tel"`
	Name       string    `json:"name" db:"name"`
	CarNumber  string    `json:"car_number" db:"car_number"`
	TimeStamp  time.Time `json:"timestamp" db:"timestamp"`
}

// Geodata geodata about courier
type Geodata struct {
	MacAddress string    `json:"mac_address" db:"mac_address"`
	CourierID  string    `json:"courier_id" db:"courier_id"`
	Latitude   float64   `json:"latitude" db:"latitude"`
	Longitude  float64   `json:"longitude" db:"longitude"`
	TimeStamp  time.Time `json:"timestamp" db:"timestamp"`
}

// UnmarshalGeodata decode Geodata from JSON
func UnmarshalGeodata(data []byte) (Geodata, error) {
	var r Geodata
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal encode Geodata to JSON
func (r *Geodata) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
