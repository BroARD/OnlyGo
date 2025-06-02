package main

import (
	"OnlyGo/pkg/db"
	"OnlyGo/pkg/quote"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Could not connect to DB")
	}

	quoteRepository := quote.NewRepository(database)
	quoteService := quote.NewService(quoteRepository)
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
