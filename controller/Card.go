package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gbodra/mtg-openapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoClient *mongo.Client

func FindCardById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cardsCollection := MongoClient.Database("mtg").Collection("cards")
	var result model.Card
	_ = cardsCollection.FindOne(context.TODO(), bson.D{{"id", vars["cardId"]}}).Decode(&result)

	pricesCollection := MongoClient.Database("mtg").Collection("prices")
	var resultPrice model.Price
	_ = pricesCollection.FindOne(context.TODO(), bson.D{{"id", vars["cardId"]}}).Decode(&resultPrice)

	result.Prices = getPrice(result.ID)
	resultJson, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resultJson))
}

func FindCardByName(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	log.Println(query)
	var result model.Card

	cardsCollection := MongoClient.Database("mtg").Collection("cards")
	_ = cardsCollection.FindOne(context.TODO(), bson.D{{"name", query}}).Decode(&result)

	result.Prices = getPrice(result.ID)
	resultsJson, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resultsJson))
}

func getPrice(cardId string) model.Price {
	pricesCollection := MongoClient.Database("mtg").Collection("prices")
	var resultPrice model.Price
	_ = pricesCollection.FindOne(context.TODO(), bson.D{{"id", cardId}}).Decode(&resultPrice)

	return resultPrice
}
