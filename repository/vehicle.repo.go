package repository

import (
	"api-parking-system/models"
	"api-parking-system/mongodb"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUserByPlateNumber(plateNumber string) (*models.User, error) {
	// Structure is user.vehicles.plate_number

	var user *models.User

	filter := bson.M{"vehicles.plate_number": plateNumber}
	options := options.FindOne()
	err := mongodb.UserCol.FindOne(mongodb.Context, filter, options).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("Vehicle not found")
	}

	if err != nil {
		return nil, errors.Wrap(err, "Error querying MongoDB")
	}

	return user, nil
}

func GetVehicleByPlateNumber(user models.User, plateNumber string) (*models.Vehicle, error) {
	var vehicle *models.Vehicle

	for _, v := range user.Vehicles {
		if v.PlateNumber == plateNumber {
			vehicle = &v
			break
		}
	}

	return vehicle, nil
}

func PushParkingCollection(vehicle models.Vehicle, parkingLog models.ParkingLog) (*mongo.UpdateResult, error) {
	filter := bson.M{"vehicles.plate_number": vehicle.PlateNumber}
	update := bson.M{"$push": bson.M{"vehicles.$.parking_log": parkingLog}}

	result, err := mongodb.UserCol.UpdateOne(mongodb.Context, filter, update)

	if err != nil {
		return nil, errors.Wrap(err, "Error updating parking log")
	}

	return result, nil
}
