package models

import "time"

type Invoice struct {
	PlateNumber string    `bson:"plate_number,omitempty" json:"plate_number,omitempty"`
	Type        string    `bson:"type" json:"type"`
	PPM         int       `bson:"ppm" json:"price_per_min"`
	Duration    int       `bson:"duration" json:"duration"`
	Amount      int       `bson:"amount" json:"amount"`
	IsPaid      bool      `bson:"is_paid" json:"is_paid"`
	PaidAt      time.Time `bson:"paid_at,omitempty" json:"paid_at,omitempty"`
	Timestamps
}
