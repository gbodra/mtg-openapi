package model

import "time"

type Price struct {
	ID     string      `json:"id" bson:"id"`
	Prices []PriceItem `json:"price_history" bson:"price_history"`
}

type PriceItem struct {
	PrintingType       string    `json:"printing_type" bson:"printingType"`
	MarketPrice        float64   `json:"market_price" bson:"marketPrice"`
	BuylistMarketPrice float64   `json:"buylist_market_price" bson:"buylistMarketPrice"`
	ListedMedianPrice  float64   `json:"listed_median_price" bson:"listedMedianPrice"`
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
}
