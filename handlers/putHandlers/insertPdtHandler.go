package putHandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jain-chetan/cart-service/helpers"
)

type PutHandler struct{}

func (g *PutHandler) InsertPdtHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	log.Println(userID)

	//Call db to insert Product into cart for the user

	response := helpers.ResponseMapper(200, "OK")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(response)
}
