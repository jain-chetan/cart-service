package getHandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jain-chetan/cart-service/helpers"
	"github.com/jain-chetan/cart-service/interfaces"
)

func (g *GetHandler) GetCartHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	log.Println(userID)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//Call db to get Cart details for the user
	cartItems, err := interfaces.DBClient.GetCartQuery(userID)
	if err != nil {
		response := helpers.ResponseMapper(400, "Bad Request")
		json.NewEncoder(w).Encode(response)
	}
	//response := helpers.ResponseMapper(200, "OK")

	json.NewEncoder(w).Encode(cartItems)
}
