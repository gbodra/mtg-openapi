package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gbodra/mtg-openapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: refatorar gestao de erros

func OptinAlert(w http.ResponseWriter, r *http.Request) {
	chatId, _ := strconv.ParseInt(r.URL.Query().Get("chat_id"), 10, 64)

	alertObject := model.Alert{
		ID:        primitive.NewObjectID(),
		ChatID:    chatId,
		CreatedAt: time.Now(),
	}

	alertOptinCollection := MongoClient.Database("mtg").Collection("alert-optin")

	_, err := alertOptinCollection.InsertOne(context.TODO(), alertObject)
	if err != nil {
		panic(err)
	}
}

func GetAlerts(w http.ResponseWriter, r *http.Request) {
	alertOptinCollection := MongoClient.Database("mtg").Collection("alert-optin")

	cursor, err := alertOptinCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []model.Alert
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		panic(err)
	}

	resultsJson, _ := json.Marshal(&results)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resultsJson))
}

func AlertOptout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alertOptinCollection := MongoClient.Database("mtg").Collection("alert-optin")
	idPrimitive, _ := primitive.ObjectIDFromHex(vars["alertId"])

	filter := bson.D{primitive.E{Key: "_id", Value: idPrimitive}}

	_, err := alertOptinCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
}

func GetAlert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alertOptinCollection := MongoClient.Database("mtg").Collection("alert-optin")

	chatId, _ := strconv.ParseInt(vars["chatId"], 10, 0)
	filter := bson.D{primitive.E{Key: "chat_id", Value: chatId}}
	var alertObject model.Alert
	_ = alertOptinCollection.FindOne(context.TODO(), filter).Decode(&alertObject)

	alertJson, _ := json.Marshal(&alertObject)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(alertJson))
}

func GetTopPriceMovements(w http.ResponseWriter, r *http.Request) {
	alertsCollection := MongoClient.Database("mtg").Collection("alerts")
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "created_at", Value: -1}})

	var priceAlertObject model.AlertMessage
	_ = alertsCollection.FindOne(context.TODO(), bson.D{}, opts).Decode(&priceAlertObject)
	priceAlertJson, _ := json.Marshal(&priceAlertObject)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(priceAlertJson))
}
