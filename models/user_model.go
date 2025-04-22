package models

import "time"

type User struct {
	ID        string    `json:"id" bson:"_id"`
	Username  string    `json:"username" bson:"username" validate:"required"`
	Email     string    `json:"email" bson:"email" validate:"required,email"`
	Role      string    `json:"role" bson:"role" validate:"required,oneof=admin user"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
