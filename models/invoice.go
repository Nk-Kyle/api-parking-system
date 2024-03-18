package models

import "time"

type Invoice struct {
	VehicleID   string    `bson:"vehicle_id,omitempty" json:"vehicle_id,omitempty"`
	PlateNumber string    `bson:"plate_number,omitempty" json:"plate_number,omitempty"`
	Amount      int       `bson:"amount,omitempty" json:"amount,omitempty"`
	IsPaid      bool      `bson:"is_paid,omitempty" json:"is_paid,omitempty"`
	PaidAt      time.Time `bson:"paid_at,omitempty" json:"paid_at,omitempty"`
	Timestamps
}
