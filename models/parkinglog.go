package models

type ParkingLog struct {
	State ParkingState `bson:"state,omitempty" json:"state,omitempty"`
	Timestamps
}
