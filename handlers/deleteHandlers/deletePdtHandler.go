package deleteHandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jain-chetan/cart-service/helpers"
	"github.com/jain-chetan/cart-service/interfaces"
)

type DeleteHandler struct{}

func (g *DeleteHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	log.Println(userID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//Get path parameter
	params := mux.Vars(r)
	productID := params["productID"]
	log.Println(productID)

	//Call db to delete item from Cart details for the user
	err := interfaces.DBClient.DeleteProduct(userID, productID)
	if err != nil {
		response := helpers.ResponseMapper(400, "Database error")
		json.NewEncoder(w).Encode(response)
	} else {
		response := helpers.ResponseMapper(200, "OK")
		json.NewEncoder(w).Encode(response)
	}
}
