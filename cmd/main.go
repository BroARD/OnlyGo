package main

import (
	"OnlyGo/logging"
	"OnlyGo/pkg/db"
	"OnlyGo/pkg/quote"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	logger := logging.GetLogger()

	logger.Info("Start db")
	database, err := db.InitDB()
	if err != nil {
		logger.Info("Could not start a DB")
	}

	logger.Info("registring a quote repository")
	quoteRepository := quote.NewRepository(database)
	logger.Info("registring a quote service")
	quoteService := quote.NewService(quoteRepository)
	logger.Info("registring a quote handlers")
	quoteHandler := quote.NewHandler(quoteService)

	r := mux.NewRouter()
	quoteHandler.Register(r)

	start(r)
}

func start(router *mux.Router) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}
