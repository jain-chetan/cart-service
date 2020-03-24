package getHandlers

import (
	"encoding/json"
	"net/http"

	"github.com/jain-chetan/cart-service/helpers"
)

type GetHandler struct{}

func (g *GetHandler) PingHandler(w http.ResponseWriter, r *http.Request) {
	response := helpers.ResponseMapper(200, "OK")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(response)
}
