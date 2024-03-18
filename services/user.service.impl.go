package services

import (
	"api-parking-system/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.userCollection.InsertOne(u.ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) GetUser(id string) (*models.User, error) {
	var user models.User
	err := u.userCollection.FindOne(u.ctx, id).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserServiceImpl) UpdateUser(id string, user *models.User) error {
	_, err := u.userCollection.UpdateOne(u.ctx, id, user)
	if err != nil {
		return err
	}
	return nil
}
