package controllers

import (
	"daas/models"
	u "daas/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	engineer := &models.Engineer{}

	err := json.NewDecoder(r.Body).Decode(engineer)
	if err != nil {
		u.Respond(w, u.Message(false, "An error occured while decoding your request"))
		fmt.Println(err)
		return
	}
	resp := engineer.Create()
	u.Respond(w, resp)
}

// // GetEngineers function
// func GetEngineers(w http.ResponseWriter, r *http.Request) {
// 	var engineers []Engineer
// 	db.Find(&engineers)
// 	json.NewEncoder(w).Encode(&engineers)
// }

// // GetEngineer function
// func GetEngineer(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var engineer Engineer
// 	db.First(&engineer, params["id"])
// 	json.NewEncoder(w).Encode(&engineer)
// }

// // CreateAccount function
// func CreateAccount(w http.ResponseWriter, r *http.Request) {
// 	var engineer Engineer
// 	json.NewDecoder(r.Body).Decode(&engineer)
// 	db.Create(&engineer)
// 	json.NewEncoder(w).Encode(&engineer)
// }

// // DeleteAccount function
// func DeleteAccount(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var engineer Engineer
// 	db.First(&engineer, params["id"])
// 	db.Delete(&engineer)

// 	var engineers []Engineer
// 	db.Find(&engineers)
// 	json.NewEncoder(w).Encode(&engineers)
// }

// // CreatePortifolio function
// func CreatePortifolio(w http.ResponseWriter, r *http.Request) {
// 	var details EngineerPortfolio
// 	json.NewDecoder(r.Body).Decode(&details)
// 	db.Create(&details)
// 	json.NewEncoder(w).Encode(&details)
// }

// // ClientAccountInfo function
// // Creates Client's info
// func ClientAccountInfo(w http.ResponseWriter, r *http.Request) {
// 	var clientaccount ClientAccount
// 	json.NewDecoder(r.Body).Decode(&clientaccount)
// 	db.Create(&clientaccount)
// 	json.NewEncoder(w).Encode(&clientaccount)
// }

// // ClientDescriptionInfo function
// func ClientDescriptionInfo(w http.ResponseWriter, r *http.Request) {
// 	var clientdescription ClientDescription
// 	json.NewDecoder(r.Body).Decode(&clientdescription)
// 	db.Create(&clientdescription)
// 	json.NewEncoder(w).Encode(&clientdescription)
// }
