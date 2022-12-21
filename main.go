package main

import (
	"log"
	"net/http"

	controller "online-transportation/controllers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// for load godotenv
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	router.HandleFunc("/users", controller.GetAllUsers).Methods("GET")

	log.Println("Starting on Port")

	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)

}
