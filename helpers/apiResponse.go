package helpers

import (
	"log"

	"github.com/jain-chetan/cart-service/model"
)

func ResponseMapper(code int, message string) model.Response {
	var response model.Response
	response = model.Response{
		Code:    code,
		Message: message,
	}
	log.Println(response)
	return response
}
