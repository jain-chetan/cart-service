package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (dc *DBRepo) UpdateQuantity(userID, productID string, quantity int, price float64) error {
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{"userID", userID}, {"products.productID", productID}}
	collection := dc.DBClient.Database("local").Collection("cart")
	update := bson.D{{"$push", bson.D{{"$set", bson.D{{"products.$.quantity", quantity}, {"products.$.subTotal", float64(quantity) * price}}}}}}
	_, err = collection.UpdateMany(ctx, filter, update)
	return err
}
