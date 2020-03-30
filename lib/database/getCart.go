package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jain-chetan/cart-service/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (dc *DBRepo) GetCartQuery(userID string) (model.Cart, error) {
	var cart model.Cart
	collection := dc.DBClient.Database("local").Collection("cart")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor := collection.FindOne(ctx, bson.M{"userID": userID})
	if err := cursor.Err(); err != nil {
		return cart, errors.New("DB error")
	}
	cursor.Decode(&cart)
	log.Println(cart)
	return cart, nil
}

func (dc *DBRepo) GetProductQuery(userID string, productID string) bool {
	var cart []model.Cart
	collection := dc.DBClient.Database("local").Collection("cart")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{"userID": userID, "products.productID": productID})
	if err != nil {
		return false
	}
	cursor.Decode(&cart)
	log.Println(cart)
	if len(cart) > 0 {
		return true
	}
	return false
}
