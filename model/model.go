package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//Response defines response code and message
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//Cart structure for cart document
type Cart struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID     string/*primitive.ObjectID*/ `json:"userID" bson:"userID"`
	GrandTotal float64    `json:"grandTotal" bson:"grandTotal"`
	Products   []Products `json:"products" bson:"products"`
}

//Products structure for sub-document inside cart document
type Products struct {
	ProductID   string/*primitive.ObjectID*/ `json:"productID" bson:"productID"`
	ProductName string  `json:"productName" bson:"productName"`
	Quantity    int     `json:"quantity" bson:"quantity"`
	SubTotal    float64 `json:"subTotal" bson:"subTotal"`
}

//DBConfig has information required to connect to DB
type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type GrandTotal struct {
	Total float64 `bson:"grandTotal"`
}

type Quantity struct {
	Quantity int `json:"quantity" bson:"quantity"`
}
