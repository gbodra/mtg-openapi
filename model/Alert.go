package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Alert struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	ChatID    int64              `json:"chat_id,omitempty" bson:"chat_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
