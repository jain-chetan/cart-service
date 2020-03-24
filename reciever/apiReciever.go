package reciever

import (
	delete "github.com/jain-chetan/cart-service/handlers/deleteHandlers"
	get "github.com/jain-chetan/cart-service/handlers/getHandlers"
	update "github.com/jain-chetan/cart-service/handlers/putHandlers"
)

var Get get.GetHandler
var Delete delete.DeleteHandler
var Update update.PutHandler
