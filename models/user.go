package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Phone    string             `bson:"phone" json:"phone"`
	Nik      string             `bson:"nik" json:"nik"`
	Vehicles []Vehicle          `bson:"vehicles" json:"vehicles"`
	Invoices []Invoice          `bson:"invoices" json:"invoices"`
	Timestamps
}
