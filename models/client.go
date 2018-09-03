package models

import (
	u "daas/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Client struct
type Client struct {
	gorm.Model

	Business       string
	TeamSize       string
	Location       string
	Specialization string
}

// Validate func
func (client *Client) Validate() (map[string]interface{}, bool) {

	if client.Business == "" {
		return u.Message(false, "Please provide your business details"), false
	}

	if client.TeamSize == "" {
		return u.Message(false, "Please provide current company's team size"), false
	}

	if client.Location == "" {
		return u.Message(false, "Please provide the company's location"), false
	}

	if client.Specialization == "" {
		return u.Message(false, "Please provide company's area of specialization"), false
	}

	return u.Message(true, "Details saved successfully"), true
}

// Create func
func (client *Client) Create() map[string]interface{} {

	if resp, ok := client.Validate(); !ok {
		return resp
	}

	GetDB().Create(client)

	resp := u.Message(true, "Success")
	resp["client"] = client
	return resp
}

// GetClients func
func GetClients() []*Client {

	clients := make([]*Client, 0)
	err := GetDB().Table("clients").Find(&clients).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return clients
}
