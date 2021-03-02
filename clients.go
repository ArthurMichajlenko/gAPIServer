package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// Clients is array of client
type Clients []Client

// UnmarshalClients decode clients from JSON
func UnmarshalClients(data []byte) (Clients, error) {
	var r Clients
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal encode clients to JSON
func (r *Clients) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Client is a single clients
type Client struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Tel  string `json:"tel" db:"tel"`
}

func getClients(c echo.Context) error {
	var courierID string
	var clientIDs []string
	var client Client
	var clients Clients
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	macAddress := claims["macAddress"].(string)
	err := db.QueryRow("SELECT id FROM couriers WHERE mac_address = ?", macAddress).Scan(&courierID)
	if err != nil {
		log.Println(err)
	}
	err = db.Select(&clientIDs, "SELECT client_id FROM orders WHERE courier_id = ?", courierID)
	if err != nil {
		log.Println(err)
	}
	for _, clientID := range clientIDs {
		err := db.QueryRow("SELECT * FROM clients WHERE id = ?", clientID).Scan(&client.ID, &client.Name, &client.Tel)
		if err != nil {
			log.Println(err)
		}
		clients = append(clients, client)
	}
	return c.JSON(http.StatusOK, clients)
}
