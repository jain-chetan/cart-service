package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/jain-chetan/cart-service/interfaces"
	"github.com/jain-chetan/cart-service/lib/database"
	"github.com/jain-chetan/cart-service/model"
)

func main() {
	err := initDBClient()
	if err != nil {
		log.Fatal("DB Driver error", err)
	}
	//router initialization
	mux := http.NewServeMux()
	//Simple ping API
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		pingResponse := model.Response{
			Code:    200,
			Message: "OK",
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		json.NewEncoder(w).Encode(pingResponse)
	})
	http.ListenAndServe(":8000", mux)
}

func initDBClient() error {
	var config model.DBConfig
	//Read DB credentials from environment variables
	config.User = os.Getenv("DBUSER")
	config.Port = os.Getenv("PORT")
	config.Host = os.Getenv("HOST")
	interfaces.DBClient = new(database.DBRepo)
	err := interfaces.DBClient.DBConnect(config)
	return err
}
