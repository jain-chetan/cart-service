package putHandlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jain-chetan/cart-service/helpers"
	"github.com/jain-chetan/cart-service/interfaces"
	"github.com/jain-chetan/cart-service/model"
)

type PutHandler struct{}

func (g *PutHandler) InsertPdtHandler(w http.ResponseWriter, r *http.Request) {
	var pingResponse model.Response
	var product model.Products
	userID := r.Header.Get("userID")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &product)

	if err != nil {
		response := helpers.ResponseMapper(400, "error in getting response")
		json.NewEncoder(w).Encode(response)
	}

	apiResponse, err := http.Get("http://localhost:8000/cart/ping")
	if err != nil {
		response := helpers.ResponseMapper(400, "error in getting response")
		json.NewEncoder(w).Encode(response)
	}
	byteData, _ := ioutil.ReadAll(apiResponse.Body)
	json.Unmarshal(byteData, &pingResponse)
	log.Println(pingResponse)
	//Call db to insert Product into cart for the user
	err = interfaces.DBClient.InsertProduct(userID, product)
	if err != nil {
		response := helpers.ResponseMapper(400, "Request error")
		json.NewEncoder(w).Encode(response)
	}
	response := helpers.ResponseMapper(200, "OK")
	json.NewEncoder(w).Encode(response)
}
