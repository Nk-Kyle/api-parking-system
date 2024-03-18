package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database
var Context context.Context

func ConnectDB() {
	mongoURI := os.Getenv("MONGO_URI")
	Context = context.Background()
	if mongoURI == "" {
		log.Println("MONGO_URI environment variable not set")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	Client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Println(err)
	}

	// Send a ping to confirm a successful connection
	if err := Client.Database("Admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		log.Println(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	DB = Client.Database(os.Getenv("MONGO_DBNAME"))
}
