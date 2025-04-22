package models

import "time"

type Face struct {
	ID         string    `json:"id" bson:"_id"`
	UserID     string    `json:"user_id" bson:"user_id" validate:"required"`
	Embedding  string    `json:"embedding" bson:"embedding" validate:"required"` // AES-256 encrypted
	Confidence float64   `json:"confidence" bson:"confidence" validate:"gte=0,lte=1"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
}
