package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dc *DBRepo) UpdateQuantity(userID, productID string, quantity int, price float64) error {
	var err error
	var primitiveUserID primitive.ObjectID
	log.Println("update quantity", userID, quantity, price)
	primitivePdtID, errConversion := primitive.ObjectIDFromHex(productID)
	if errConversion != nil {
		return errConversion
	}
	primitiveUserID, errConversion = primitive.ObjectIDFromHex(userID)
	if errConversion != nil {
		return errConversion
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{"userID", primitiveUserID}, {"products.productID", primitivePdtID}}
	collection := dc.DBClient.Database("local").Collection("cart")
	//update := bson.D{{"$push", bson.D{{"$set", bson.D{{"products.$.quantity", quantity}, {"products.$.subTotal", float64(quantity) * price}}}}}}

	update := bson.D{primitive.E{Key: "$push", Value: bson.D{primitive.E{Key: "$set", Value: bson.D{{"products.$.quantity", quantity}, {"products.$.subTotal", float64(quantity) * price}}}}}}
	log.Println(filter)
	log.Println(update)
	//update := bson.D{{"$push", primitive.E{Key: "$set", Value: primitive.E{Key: "products.$.quantity", Value: quantity}, {Key: "products.$.subTotal", Value: float64(quantity) * price}}}}
	result, err := collection.UpdateOne(ctx, filter, update) //filter, update)
	log.Println(result, err)
	if err != nil {
		return err
	}
	dc.UpdateGrandTotal(userID)
	return nil
}
