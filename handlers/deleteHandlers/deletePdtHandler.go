package deleteHandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jain-chetan/cart-service/helpers"
)

type DeleteHandler struct{}

func (g *DeleteHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	log.Println(userID)

	//Get path parameter
	params := mux.Vars(r)
	productID := params["productID"]
	log.Println(productID)

	//Call db to delete item from Cart details for the user

	response := helpers.ResponseMapper(200, "OK")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(response)
}
