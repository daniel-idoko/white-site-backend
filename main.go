package main

import (
	"fmt"
	"log"
	controller "puppyspot-backend/pkg/controllers"

	"net/http"
	"os"

	"github.com/gorilla/mux"
	// "github.com/joho/godotenv"
)

func main() {
	// Only for local
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// 	panic(err)
	// }

	var router *mux.Router = mux.NewRouter()

	router.HandleFunc("/paypal-payment", controller.HandlePaypalSumbit).Methods("POST", "OPTIONS")
	router.HandleFunc("/crypto-currency-payment", controller.HandleCryptoSumbit).Methods("POST", "OPTIONS")
	router.HandleFunc("/bank-transfer-payment", controller.HandleBankTransferSumbit).Methods("POST", "OPTIONS")
	router.HandleFunc("/ask-about-mail", controller.HandleAskAboutMail).Methods("POST", "OPTIONS")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server on port.............. %d \n", 8080)
	http.ListenAndServe(":"+port, router)
}