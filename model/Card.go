package model

type Card struct {
	ID          string   `json:"id" bson:"id"`
	Name        string   `json:"name" bson:"name"`
	ScryfallURL string   `json:"scryfall_uri" bson:"scryfall_uri"`
	TypeLine    string   `json:"type_line" bson:"type_line"`
	Colors      []string `json:"colors" bson:"colors"`
	Rarity      string   `json:"rarity" bson:"rarity"`
	Set         string   `json:"set" bson:"set"`
	SetName     string   `json:"set_name" bson:"set_name"`
	Prices      Price    `json:"prices" bson:"prices"`
}
