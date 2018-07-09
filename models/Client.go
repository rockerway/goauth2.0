package models

import (
	"database/sql"
	"oauth/packages/database"
)

// Client struct
type Client struct {
	*database.Model
}

// GetClients is the method that get all users
func (client *Client) GetClients() *sql.Rows {
	defer client.Model.Init().Close()
	rows := client.Model.Query("select * from clients")
	return rows
}

// InitUser to init user
func InitUser() Client {
	return Client{Model: database.InitModel()}
}
