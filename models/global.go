package models

import "time"

type Global struct {
	Date          time.Time `bson:"date" json:"date"`
	Motor         int       `bson:"motor" json:"motor"`
	Mobil         int       `bson:"mobil" json:"mobil"`
	Truk          int       `bson:"truk" json:"truk"`
	Bus           int       `bson:"bus" json:"bus"`
	Billable      int       `bson:"billable" json:"billable"`
	Transactions  int       `bson:"transactions" json:"transactions"`
	TotalDuration int       `bson:"t_duration" json:"t_duration"`
	MinDuration   int       `bson:"min_duration" json:"min_duration"`
	MaxDuration   int       `bson:"max_duration" json:"max_duration"`
}
