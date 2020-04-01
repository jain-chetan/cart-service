package putHandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jain-chetan/cart-service/helpers"
	"github.com/jain-chetan/cart-service/interfaces"
	"github.com/jain-chetan/cart-service/model"
)

type PutHandler struct{}

func (g *PutHandler) InsertPdtHandler(w http.ResponseWriter, r *http.Request) {
	var product model.Products
	userID := r.Header.Get("userID")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &product)

	if err != nil {
		response := helpers.ResponseMapper(400, "error in getting response")
		json.NewEncoder(w).Encode(response)
	}
	//In case of API call to get product name and product price, use below code
	/*productResponse, err := http.Get("http://localhost:8000/cart/product" + product.ProductID)
	if err != nil {
		response := helpers.ResponseMapper(400, "error in getting response")
		json.NewEncoder(w).Encode(response)
	}
	byteData, _ := ioutil.ReadAll(productResponse.Body)
	json.Unmarshal(byteData, &productResponse)
	err = interfaces.DBClient.InsertProduct(userID, productResponse)
	*/
	//Call db to insert Product into cart for the user
	err = interfaces.DBClient.InsertProduct(userID, product)
	if err != nil {
		response := helpers.ResponseMapper(400, "Request error")
		json.NewEncoder(w).Encode(response)
	}
	response := helpers.ResponseMapper(200, "OK")
	json.NewEncoder(w).Encode(response)
}
