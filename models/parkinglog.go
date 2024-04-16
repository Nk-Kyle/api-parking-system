package models

type ParkingLog struct {
	State    ParkingState `bson:"state,omitempty" json:"state,omitempty"`
	ImageURL string       `bson:"image_url,omitempty" json:"image_url,omitempty"`
	Timestamps
}
