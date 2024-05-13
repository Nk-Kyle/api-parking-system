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

func GetGlobal() (*models.Global, error) {
	var global *models.Global
	currentDate := time.Now()
	filter := bson.M{"date": time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.UTC)}
	options := options.FindOne()
	err := mongodb.GlobalCol.FindOne(mongodb.Context, filter, options).Decode(&global)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "Error querying MongoDB")
	}

	return global, nil
}

func UpdateOrCreateGlobal(global *models.Global) (*mongo.UpdateResult, error) {
	currentDate := time.Now()
	filter := bson.M{"date": time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.UTC)}
	update := bson.M{"$set": global}
	upsert := true

	result, err := mongodb.GlobalCol.UpdateOne(mongodb.Context, filter, update, &options.UpdateOptions{
		Upsert: &upsert,
	})

	if err != nil {
		return nil, errors.Wrap(err, "Error updating global")
	}

	return result, nil
}
