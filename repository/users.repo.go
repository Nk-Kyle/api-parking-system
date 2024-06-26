package repository

import (
	"api-parking-system/models"
	"api-parking-system/mongodb"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUser(user *models.User) (*mongo.InsertOneResult, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result, err := mongodb.UserCol.InsertOne(mongodb.Context, user)

	if err != nil {
		return nil, errors.Wrap(err, "Error creating user")
	}

	return result, nil
}

func GetUserByEmailorNik(email string, nik string) (*models.User, error) {
	var user *models.User

	filter := bson.M{"$or": []bson.M{
		{"email": email},
		{"nik": nik},
	}}

	options := options.FindOne().SetSort(bson.M{"created_at": -1})
	err := mongodb.UserCol.FindOne(mongodb.Context, filter, options).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("User not found")
	} else if err != nil {
		return nil, errors.Wrap(err, "Error querying MongoDB")
	}

	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user *models.User

	filter := bson.M{"email": email}
	options := options.FindOne()
	err := mongodb.UserCol.FindOne(mongodb.Context, filter, options).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("User not found")
	} else if err != nil {
		return nil, errors.Wrap(err, "Error querying MongoDB")
	}

	return user, nil
}

func UpdateUser(user *models.User) (*mongo.UpdateResult, error) {
	user.UpdatedAt = time.Now()

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}

	result, err := mongodb.UserCol.UpdateOne(mongodb.Context, filter, update)

	if err != nil {
		return nil, errors.Wrap(err, "Error updating user")
	}

	return result, nil
}

func GetThisWeekUser(email string) (*mongo.Cursor, error) {
	startDate := time.Now().AddDate(0, 0, -7)

	// Define the aggregation pipeline to filter and project only the invoices
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"email": email}}},
		{{Key: "$project", Value: bson.M{
			"invoices": bson.M{
				"$filter": bson.M{
					"input": "$invoices",
					"as":    "invoice",
					"cond":  bson.M{"$gte": bson.A{"$$invoice.timestamps.created_at", startDate}},
				},
			},
			"vehicle": bson.M{"$first": "$vehicles"},
		}}},
	}

	// Perform the aggregation
	res, err := mongodb.UserCol.Aggregate(mongodb.Context, pipeline)
	return res, err
}

func GetUserImages(email string) (*mongo.Cursor, error) {

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"email": email}}},
		{{Key: "$project", Value: bson.M{
			"images": bson.M{
				"$arrayElemAt": bson.A{
					"$vehicles.parking_log.image_url", 0,
				},
			}}},
		},
	}

	res, err := mongodb.UserCol.Aggregate(mongodb.Context, pipeline)
	return res, err
}
