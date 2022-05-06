package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gbodra/mtg-openapi/model"
	"github.com/gbodra/mtg-openapi/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoClient *mongo.Client

func FindCardById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cardsCollection := MongoClient.Database("mtg").Collection("cards")
	var result model.Card
	filter := bson.D{primitive.E{Key: "id", Value: vars["cardId"]}}
	_ = cardsCollection.FindOne(context.TODO(), filter).Decode(&result)

	pricesCollection := MongoClient.Database("mtg").Collection("prices")
	var resultPrice model.Price
	_ = pricesCollection.FindOne(context.TODO(), filter).Decode(&resultPrice)

	result.Prices = getPrice(result.ID)
	resultJson, err := json.Marshal(result)
	utils.HandleError(err, "Error transforming object into json on FindCardById")

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resultJson))
}

func FindCardByName(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	cardsCollection := MongoClient.Database("mtg").Collection("cards")
	matchStage := bson.D{{"$match", bson.D{{"$text", bson.D{{"$search", "\"" + query + "\""}}}}}}

	cursor, err := cardsCollection.Aggregate(context.TODO(), mongo.Pipeline{matchStage})

	utils.HandleError(err, "Error searching documents on FindCardByName")

	var results []model.Card
	cursor.All(context.TODO(), &results)

	if len(results) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	results[0].Prices = getPrice(results[0].ID)
	resultsJson, err := json.Marshal(results[0])
	utils.HandleError(err, "Error transforming object into json on FindCardByName")

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resultsJson))
}

func getPrice(cardId string) model.Price {
	pricesAggregatedCollection := MongoClient.Database("mtg").Collection("prices_aggregated")
	var resultPrice model.Price
	filter := bson.D{primitive.E{Key: "id", Value: cardId}}
	_ = pricesAggregatedCollection.FindOne(context.TODO(), filter).Decode(&resultPrice)

	return resultPrice
}
