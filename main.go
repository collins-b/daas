package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"daas/controllers"
	"daas/internal"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

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

// EngineerPortfolio struct
type EngineerPortfolio struct {
	gorm.Model

	Skills       string
	Experience   string
	Availability string
}

// ClientAccount struct
type ClientAccount struct {
	gorm.Model

	Name         string
	EmailAddress string
	Username     string
	Password     string
}

// ClientDescription struct
type ClientDescription struct {
	gorm.Model

	Business       string
	TeamSize       string
	Location       string
	Specialization string
}

var db *gorm.DB
var err error

func main() {
	// os.Setenv("HOST", "localhost")
	// os.Setenv("DBNAME", "daas")
	// os.Setenv("PASSWORD", "")
	// os.Setenv("PORT", "3000")
	router := mux.NewRouter()
	// db, err = gorm.Open(
	// 	"postgres",
	// 	"host="+os.Getenv("HOST")+" user="+os.Getenv("USER")+
	// 		" dbname="+os.Getenv("DBNAME")+" sslmode=disable password="+os.Getenv("PASSWORD"))

	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// defer db.Close()

	// db.AutoMigrate(&Engineer{}, &EngineerPortfolio{})

	router.HandleFunc("/api/engineers", controllers.GetAllEngineers).Methods("GET")
	router.HandleFunc("/api/engineers/{id}", controllers.GetEngineer).Methods("GET")
	router.HandleFunc("/api/users/signup", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/engineers/details", controllers.EngineerDetails).Methods("POST")
	router.HandleFunc("/api/users/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/clients/details", controllers.ClientDetails).Methods("POST")
	router.Use(internal.JwtAuthentication)
	// router.HandleFunc("/api/engineers/{id}", controllers.DeleteAccount).Methods("DELETE")

	// router.HandleFunc("/api/engineers/details", CreatePortifolio).Methods("POST")

	// handler := cors.Default().Handler(router)

	// log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Server running at: http://127.0.0.1:3000")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
