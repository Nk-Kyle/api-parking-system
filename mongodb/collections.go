package mongodb

import "go.mongodb.org/mongo-driver/mongo"

var UserCol *mongo.Collection
var GlobalCol *mongo.Collection

func InitCollections() {
	UserCol = DB.Collection("users")
	GlobalCol = DB.Collection("global")
}
