package models

import "time"

type Plate struct {
	ID          string    `json:"id" bson:"id"`
	PlateNumber string    `json:"plate_number" bson:"plate_number" validate:"required"`
	VehicleType string    `json:"vehicle_type" bson:"vehicle_type"`
	Confidence  float64   `json:"confidence" bson:"confidence" validate:"gte=0,lte=1"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
