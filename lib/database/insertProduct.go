package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/jain-chetan/cart-service/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (dc *DBRepo) InsertProduct(userID string, product model.Products) error {
	var err error
	collection := dc.DBClient.Database("local").Collection("cart")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"userID": userID}

	log.Println(bson.M{"$push": bson.M{"products": product}})
	_, err = collection.UpdateOne(ctx, filter, bson.M{"$push": bson.M{"products": product}})
	if err != nil {
		return err
	}

	err = dc.UpdateGrandTotal(userID)
	return err
}

func (dc *DBRepo) UpdateGrandTotal(userID string) error {
	var total model.GrandTotal
	collection := dc.DBClient.Database("local").Collection("cart")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var totalArr []bson.M
	//match and group for aggregate method
	match := bson.D{{"$match", bson.D{{"userID", userID}}}}
	group := bson.D{{"$project", bson.D{{"grandTotal", bson.D{{"$sum", "$products.subTotal"}}}}}}
	totalResult, err := collection.Aggregate(ctx, mongo.Pipeline{match, group})
	if err = totalResult.All(ctx, &totalArr); err != nil {
		return err
	}
	grandTotal := totalArr[0]
	totalBytes, _ := bson.Marshal(grandTotal)
	bson.Unmarshal(totalBytes, &total)
	log.Println("grand total", total.Total)
	collection.UpdateOne(ctx, bson.M{"userID": userID}, bson.M{"$set": bson.M{"grandTotal": total.Total}})
	return err

}
