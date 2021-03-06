package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gbodra/mtg-openapi/controller"
	"github.com/gbodra/mtg-openapi/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	tb "gopkg.in/tucnak/telebot.v2"
)

type App struct {
	Router      *mux.Router
	Port        string
	Mongo       *mongo.Client
	Bot         *tb.Bot
	NewRelicApp *newrelic.Application
}

func (a *App) Initialize() {
	err := godotenv.Load()
	utils.HandleError(err, "Error loading .env file")

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	a.Mongo, err = mongo.Connect(context.TODO(), clientOptions)
	utils.HandleError(err, "Error connecting to Mongo")

	a.Router = mux.NewRouter()
	a.initializeNewRelic()
	a.initializeRoutes()
	a.injectClients()
}

func (a *App) initializeRoutes() {
	// Management routes
	a.Router.HandleFunc(newrelic.WrapHandleFunc(a.NewRelicApp, "/health", controller.HealthCheck)).Methods("GET")

	// Cards
	a.Router.HandleFunc(newrelic.WrapHandleFunc(a.NewRelicApp, "/cards", controller.FindCardByName)).Methods("GET")
	a.Router.HandleFunc(newrelic.WrapHandleFunc(a.NewRelicApp, "/cards/{cardId}", controller.FindCardById)).Methods("GET")

	// Alerts
	a.Router.HandleFunc(newrelic.WrapHandleFunc(a.NewRelicApp, "/listAlerts", controller.GetAlerts)).Methods("GET")
	a.Router.HandleFunc(newrelic.WrapHandleFunc(a.NewRelicApp, "/alert/{chatId}", controller.GetAlert)).Methods("GET")
	a.Router.HandleFunc(newrelic.WrapHandleFunc(a.NewRelicApp, "/alert", controller.OptinAlert)).Methods("POST")
	a.Router.HandleFunc(newrelic.WrapHandleFunc(a.NewRelicApp, "/alert/{alertId}", controller.AlertOptout)).Methods("DELETE")

	// Price
	a.Router.HandleFunc(newrelic.WrapHandleFunc(a.NewRelicApp, "/price/top", controller.GetTopPriceMovements)).Methods("GET")
}

func (a *App) initializeNewRelic() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("mtg-open-api-prd"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC")),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	utils.HandleError(err, "Error initializing NewRelic")

	a.NewRelicApp = app
}

func (a *App) Run() {
	port := getPort()
	log.Println("App running on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, a.Router))
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
		log.Println("$PORT not set. Falling back to default " + port)
	}

	return port
}

func (a *App) injectClients() {
	controller.MongoClient = a.Mongo
}
