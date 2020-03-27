package database

import (
	"context"
	"log"
	"time"

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

	//(ctx, bson.M{"userID": userID})
	return err
}
