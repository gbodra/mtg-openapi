package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gbodra/mtg-openapi/controller"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	tb "gopkg.in/tucnak/telebot.v2"
)

type App struct {
	Router *mux.Router
	Port   string
	Mongo  *mongo.Client
	Bot    *tb.Bot
}

func (a *App) Initialize() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env")
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	a.Mongo, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Println(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
	a.injectClients()
}

func (a *App) initializeRoutes() {
	// Management routes
	a.Router.HandleFunc("/health", controller.HealthCheck).Methods("GET")

	// Cards
	a.Router.HandleFunc("/findCards", controller.FindCards).Methods("GET")
	a.Router.HandleFunc("/cards/{cardId}", controller.FindCard).Methods("GET")

	// Alerts
	a.Router.HandleFunc("/findAlerts", controller.HealthCheck).Methods("GET")
	a.Router.HandleFunc("/alert", controller.HealthCheck).Methods("POST")
	a.Router.HandleFunc("/alert", controller.HealthCheck).Methods("PUT")
	a.Router.HandleFunc("/alert/{alertId}", controller.HealthCheck).Methods("DELETE")
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