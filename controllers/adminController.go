package controllers

import (
	"daas/models"
	u "daas/utils"
	"net/http"
)

var GetAllEngineers = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetEngineers()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
