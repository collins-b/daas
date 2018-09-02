package controllers

import (
	"daas/models"
	u "daas/utils"
	"encoding/json"
	"net/http"
)

// var cookieHandler = securecookie.New(
// 	securecookie.GenerateRandomKey(64),
// 	securecookie.GenerateRandomKey(32))

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	// setSession(account.Email, w)
	u.Respond(w, resp)
}

// func setSession(userName string, response http.ResponseWriter) {
// 	value := map[string]string{
// 		"name": userName,
// 	}
// 	if encoded, err := cookieHandler.Encode("session", value); err == nil {
// 		cookie := &http.Cookie{
// 			Name:  "session",
// 			Value: encoded,
// 		}
// 		http.SetCookie(response, cookie)
// 		fmt.Println(cookie)
// 	}
// }

// func getUserName(request *http.Request) (userName string) {
// 	if cookie, err := request.Cookie("session"); err == nil {
// 		cookieValue := make(map[string]string)
// 		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
// 			userName = cookieValue["name"]
// 		}
// 	}
// 	return userName
// }

// func InternalPageHandler(response http.ResponseWriter, request *http.Request) {
// 	userName := getUserName(request)
// 	fmt.Println(userName)
// 	if userName != "" {
// 		fmt.Fprintf(response, userName)
// 	} else {
// 		http.Redirect(response, request, "/api/engineers", 302)
// 	}
// }
