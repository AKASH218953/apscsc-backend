package models

import "time"

type GunnyBag struct {
	ID        string    `json:"id" bson:"id"`
	Count     int       `json:"count" bson:"count" validate:"required,gte=0"`
	Warehouse string    `json:"warehouse" bson:"warehouse" validate:"required"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
