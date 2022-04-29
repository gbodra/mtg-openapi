package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gbodra/mtg-openapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()
var MongoClient *mongo.Client

func FindCards(w http.ResponseWriter, r *http.Request) {
	coll := MongoClient.Database("mtg").Collection("cards")
	var result model.Card
	_ = coll.FindOne(context.TODO(), bson.D{{"name", "Assembled Ensemble"}}).Decode(&result)
	resultJson, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resultJson))
}

func FindCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cardsCollection := MongoClient.Database("mtg").Collection("cards")
	var result model.Card
	_ = cardsCollection.FindOne(context.TODO(), bson.D{{"id", vars["cardId"]}}).Decode(&result)

	pricesCollection := MongoClient.Database("mtg").Collection("prices")
	var resultPrice model.Price
	_ = pricesCollection.FindOne(context.TODO(), bson.D{{"id", vars["cardId"]}}).Decode(&resultPrice)

	result.Prices = resultPrice
	resultJson, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resultJson))
}
