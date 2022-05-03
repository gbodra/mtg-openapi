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

type AlertMessage struct {
	ID        string           `json:"_id" bson:"_id"`
	Threshold float64          `json:"threshold" bson:"threshold"`
	CreatedAt time.Time        `json:"created_at" bson:"created_at"`
	Cards     []CardPriceAlert `json:"cards" bson:"cards"`
}

type CardPriceAlert struct {
	ID                       string  `json:"id" bson:"id"`
	Name                     string  `json:"name" bson:"name"`
	LastPrice                float64 `json:"last_price" bson:"last_price"`
	NormalMovementMoney      float64 `json:"normal_movement_money" bson:"normal_movement_money"`
	NormalMovementPercentage float64 `json:"normal_movement_percentage" bson:"normal_movement_percentage"`
}
