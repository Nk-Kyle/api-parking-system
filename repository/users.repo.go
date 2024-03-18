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

func GetUserByEmail(email string) (*models.User, error) {
	var user *models.User

	filter := bson.M{"email": email}
	options := options.FindOne().SetSort(bson.M{"created_at": -1})
	err := mongodb.UserCol.FindOne(mongodb.Context, filter, options).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("User not found")
	} else if err != nil {
		return nil, errors.Wrap(err, "Error querying MongoDB")
	}

	return user, nil
}
