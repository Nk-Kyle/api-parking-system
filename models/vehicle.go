package models

type Vehicle struct {
	PlateNumber string       `bson:"plate_number" json:"plate_number"`
	Type        string       `bson:"type" json:"type"`
	ParkingLog  []ParkingLog `bson:"parking_log" json:"parking_log"`
	Timestamps
}
