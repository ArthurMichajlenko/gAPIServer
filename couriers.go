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
	ID        int    `json:"id" db:"id"`
	Imei      int64  `json:"imei" db:"imei"`
	Tel       string `json:"tel" db:"tel"`
	Name      string `json:"name" db:"name"`
	CarNumber string `json:"car_number" db:"car_number"`
	TimeStamp time.Time `json:"timestamp" db:"timestamp"`
}

// Geodata geodata about courier
type Geodata struct {
	ID        int     `json:"id" db:"id"`
	Imei      int64   `json:"imei" db:"imei"`
	CourierID int     `json:"courier_id" db:"courier_id"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
	Address   string  `json:"address" db:"address"`
	TimeStamp time.Time `json:"timestamp" db:"timestamp"`
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

// CourierCl courier data for REST from mobile client
type CourierCl struct {
	ID        int     `json:"id" db:"id"`
	Imei      int64   `json:"imei" db:"imei"`
	Tel       string  `json:"tel" db:"tel"`
	Name      string  `json:"name" db:"name"`
	CarNumber string  `json:"car_number" db:"car_number"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
	Address   string  `json:"address" db:"address"`
	TimeStamp time.Time `json:"timestamp" db:"timestamp"`
}

// UnmarshalCourierCl decode CourierCl from JSON
func UnmarshalCourierCl(data []byte) (CourierCl, error) {
	var r CourierCl
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal encode CourierCl to JSON
func (r *CourierCl) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
