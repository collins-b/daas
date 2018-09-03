package controllers

import (
	"daas/models"
	u "daas/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// ClientDetails func
var ClientDetails = func(w http.ResponseWriter, r *http.Request) {
	client := &models.Client{}

	err := json.NewDecoder(r.Body).Decode(client)
	if err != nil {
		u.Respond(w, u.Message(false, "An error occured while decoding your request"))
		fmt.Println(err)
		return
	}
	resp := client.Create()
	u.Respond(w, resp)
}
