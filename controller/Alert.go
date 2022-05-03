package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gbodra/mtg-openapi/model"
	"github.com/gbodra/mtg-openapi/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OptinAlert(w http.ResponseWriter, r *http.Request) {
	chatId, err := strconv.ParseInt(r.URL.Query().Get("chat_id"), 10, 64)
	utils.HandleError(err, "Error parsing chat_id on OptinAlert method")

	alertObject := model.Alert{
		ID:        primitive.NewObjectID(),
		ChatID:    chatId,
		CreatedAt: time.Now(),
	}

	alertOptinCollection := MongoClient.Database("mtg").Collection("alert-optin")

	_, err = alertOptinCollection.InsertOne(context.TODO(), alertObject)
	utils.HandleError(err, "Error inserting AlertOptin on Mongo")
}

func GetAlerts(w http.ResponseWriter, r *http.Request) {
	alertOptinCollection := MongoClient.Database("mtg").Collection("alert-optin")

	cursor, err := alertOptinCollection.Find(context.TODO(), bson.D{})
	utils.HandleError(err, "Error loading alerts on GetAlerts")

	var results []model.Alert
	err = cursor.All(context.TODO(), &results)
	utils.HandleError(err, "Error parsing alerts to object on GetAlerts")

	resultsJson, err := json.Marshal(&results)
	utils.HandleError(err, "Error transforming object into json on GetAlerts")

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resultsJson))
}

func AlertOptout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alertOptinCollection := MongoClient.Database("mtg").Collection("alert-optin")
	idPrimitive, err := primitive.ObjectIDFromHex(vars["alertId"])
	utils.HandleError(err, "Error parsing ObjectId on AlertOptout")

	filter := bson.D{primitive.E{Key: "_id", Value: idPrimitive}}

	_, err = alertOptinCollection.DeleteOne(context.TODO(), filter)
	utils.HandleError(err, "Error deleting alert from Mongo on AlertOptout")
}

func GetAlert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alertOptinCollection := MongoClient.Database("mtg").Collection("alert-optin")

	chatId, _ := strconv.ParseInt(vars["chatId"], 10, 0)
	filter := bson.D{primitive.E{Key: "chat_id", Value: chatId}}
	var alertObject model.Alert
	_ = alertOptinCollection.FindOne(context.TODO(), filter).Decode(&alertObject)

	alertJson, err := json.Marshal(&alertObject)
	utils.HandleError(err, "Error transforming object into json on GetAlert")

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(alertJson))
}

func GetTopPriceMovements(w http.ResponseWriter, r *http.Request) {
	alertsCollection := MongoClient.Database("mtg").Collection("alerts")
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "created_at", Value: -1}})

	var priceAlertObject model.AlertMessage
	_ = alertsCollection.FindOne(context.TODO(), bson.D{}, opts).Decode(&priceAlertObject)
	priceAlertJson, err := json.Marshal(&priceAlertObject)
	utils.HandleError(err, "Error transforming object into json on GetTopPriceMovements")

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(priceAlertJson))
}
