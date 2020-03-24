package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jain-chetan/cart-service/interfaces"
	"github.com/jain-chetan/cart-service/lib/database"
	"github.com/jain-chetan/cart-service/model"
	api "github.com/jain-chetan/cart-service/reciever"
	"github.com/joho/godotenv"
)

func main() {
	err := initDBClient()
	if err != nil {
		log.Fatal("DB Driver error", err)
	}
	//router initialization
	router := mux.NewRouter()
	//Simple ping API
	router.HandleFunc("/cart/ping", api.Get.PingHandler).Methods("GET")
	router.HandleFunc("/cart/", api.Get.GetCartHandler).Methods("GET")
	router.HandleFunc("/cart/", api.Update.InsertPdtHandler).Methods("PUT")
	router.HandleFunc("/cart/updateQuantity/{productID}", api.Update.UpdateQuantityHandler).Methods("PUT")
	router.HandleFunc("/cart/{productID}", api.Delete.DeleteProductHandler).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}

func initDBClient() error {
	var config model.DBConfig
	err := godotenv.Load()
	if err != nil {
		return err
	}
	//Read DB credentials from environment variables
	config.User = os.Getenv("DBUSER")
	config.Port = os.Getenv("PORT")
	config.Host = os.Getenv("HOST")
	interfaces.DBClient = new(database.DBRepo)
	err = interfaces.DBClient.DBConnect(config)
	return err
}
