package mongodb

import "go.mongodb.org/mongo-driver/mongo"

var UserCol *mongo.Collection

func InitCollections() {
	UserCol = DB.Collection("users")
}
