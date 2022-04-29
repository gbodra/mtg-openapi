package model

type Alert struct {
	CardID string  `json:"card_id" bson:"card_id"`
	Price  float64 `json:"price" bson:"price"`
	Email  string  `json:"email" bson:"email"`
}
