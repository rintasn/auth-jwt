package main

import (
	"FinalTestLogin/controllers/authcontroller"
	"FinalTestLogin/middlewares"
	"FinalTestLogin/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()
	//log.Fatal(r)
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	//api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":3000", r))
}
