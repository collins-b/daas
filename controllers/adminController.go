package controllers

import (
	"daas/models"
	u "daas/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var GetAllEngineers = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetEngineers()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetEngineer = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	data := models.GetEngineer(id)
	if data == nil {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "An engineer with that email address doesn't exist.")
	}
}

var GetAllClients = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetClients()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
