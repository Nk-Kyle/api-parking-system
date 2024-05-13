package repository

import (
	"api-parking-system/models"
	"api-parking-system/mongodb"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddInvoice(invoice *models.Invoice, user models.User) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$push": bson.M{"invoices": invoice}}

	result, err := mongodb.UserCol.UpdateOne(mongodb.Context, filter, update)

	if err != nil {
		return nil, errors.Wrap(err, "Error updating invoice")
	}

	return result, nil
}
