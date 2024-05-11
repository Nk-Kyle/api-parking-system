package repository

import (
	"api-parking-system/models"
	"api-parking-system/mongodb"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetVehicleByPlateNumber(plateNumber string) (*models.Vehicle, error) {
	// Structure is user.vehicles.plate_number

	var vehicle *models.Vehicle

	filter := bson.M{"vehicles.plate_number": plateNumber}
	options := options.FindOne()
	err := mongodb.UserCol.FindOne(mongodb.Context, filter, options).Decode(&vehicle)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("Vehicle not found")
	}

	if err != nil {
		return nil, errors.Wrap(err, "Error querying MongoDB")
	}

	return vehicle, nil
}
