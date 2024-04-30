package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Phone    string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Nik      string             `bson:"nik,omitempty" json:"nik,omitempty"`
	Invoices []Invoice          `bson:"invoices,omitempty" json:"invoices,omitempty"`
	Timestamps
}
