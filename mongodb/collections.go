package mongodb

import "go.mongodb.org/mongo-driver/mongo"

var UserCol *mongo.Collection
var VehicleCol *mongo.Collection

func InitCollections() {
	UserCol = DB.Collection("users")
	VehicleCol = DB.Collection("vehicles")
}
