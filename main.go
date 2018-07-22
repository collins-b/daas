package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Engineer struct
type Engineer struct {
	gorm.Model

	FirstName    string
	LastName     string
	EmailAddress string
	Username     string
	Password     string
}

var db *gorm.DB
var err error

func main() {
	os.Setenv("HOST", "localhost")
	os.Setenv("DBNAME", "daas")
	os.Setenv("PASSWORD", "")
	os.Setenv("PORT", "3000")
	router := mux.NewRouter()
	db, err = gorm.Open(
		"postgres",
		"host="+os.Getenv("HOST")+" user="+os.Getenv("USER")+
			" dbname="+os.Getenv("DBNAME")+" sslmode=disable password="+os.Getenv("PASSWORD"))

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Engineer{})

	router.HandleFunc("/api/engineers", GetEngineers).Methods("GET")
	router.HandleFunc("/api/engineers/{id}", GetEngineer).Methods("GET")
	router.HandleFunc("/api/engineers", CreateAccount).Methods("POST")
	router.HandleFunc("/api/engineers/{id}", DeleteAccount).Methods("DELETE")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}

// GetEngineers function
func GetEngineers(w http.ResponseWriter, r *http.Request) {
	var engineers []Engineer
	db.Find(&engineers)
	json.NewEncoder(w).Encode(&engineers)
}

// GetEngineer function
func GetEngineer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var engineer Engineer
	db.First(&engineer, params["id"])
	json.NewEncoder(w).Encode(&engineer)
}

// CreateAccount function
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var engineer Engineer
	json.NewDecoder(r.Body).Decode(&engineer)
	db.Create(&engineer)
	json.NewEncoder(w).Encode(&engineer)
}

// DeleteAccount function
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var engineer Engineer
	db.First(&engineer, params["id"])
	db.Delete(&engineer)

	var engineers []Engineer
	db.Find(&engineers)
	json.NewEncoder(w).Encode(&engineers)
}
