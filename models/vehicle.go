package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vehicle struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	PlateNumber string             `bson:"plate_number,omitempty" json:"plate_number,omitempty"`
	ParkingLog  []ParkingLog       `bson:"parking_log,omitempty" json:"parking_log,omitempty"`
	Timestamps
}
