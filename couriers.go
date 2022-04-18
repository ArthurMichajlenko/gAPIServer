package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
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
	ID         string `json:"id" db:"id"`
	MacAddress string `json:"mac_address" db:"mac_address"`
	Tel        string `json:"tel" db:"tel"`
	Name       string `json:"name" db:"name"`
	CarNumber  string `json:"car_number" db:"car_number"`
	TimeStamp  string `json:"timestamp" db:"timestamp"`
}

// Geodatas is array of geodata
type Geodatas []Geodata

//UnmarshalGeodatas decode geodatas from JSON
func UnmarshalGeodatas(data []byte) (Geodatas, error) {
	var r Geodatas
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal encode geodatas to JSON
func (r *Geodatas) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Geodata geodata about courier
type Geodata struct {
	ID         int     `json:"id" db:"id"`
	MacAddress string  `json:"mac_address" db:"mac_address"`
	CourierID  string  `json:"courier_id" db:"courier_id"`
	Latitude   float64 `json:"latitude" db:"latitude"`
	Longitude  float64 `json:"longitude" db:"longitude"`
	TimeStamp  string  `json:"timestamp" db:"timestamp"`
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
	var id int
	if err := c.Bind(&geodata); err != nil {
		log.Println(err)
	}
	err := db.Get(&id, `SELECT id FROM geodata WHERE
	mac_address = ? AND courier_id = ? AND timestamp = ?`, geodata.MacAddress, geodata.CourierID, geodata.TimeStamp)
	switch err {
	case sql.ErrNoRows:
		_, err = db.NamedExec(`INSERT INTO geodata (mac_address, courier_id, latitude, longitude, timestamp) 
								VALUES (:mac_address, :courier_id, :latitude, :longitude, :timestamp)`, &geodata)
		if err != nil {
			log.Println(err)
		}
	default:
		if err != nil {
			log.Println(err)
		}
	}
	return c.NoContent(http.StatusNoContent)
}

func getGeodatas(c echo.Context) error {
	var geodata Geodatas
	err := db.Select(&geodata, "SELECT * FROM geodata WHERE mac_address = ? AND timestamp >= ? AND timestamp <= ? ORDER BY timestamp", c.QueryParam("mac_address"), c.QueryParam("data_start"), c.QueryParam("data_finish"))
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, geodata)
}
