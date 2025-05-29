package main

import (
	"OnlyGo/pkg/db"
	"OnlyGo/pkg/handlers"
	"OnlyGo/pkg/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Could not connect to DB")
	}

	quoteRepository := service.NewQuoteRepository(database)
	quoteService := service.NewQuoteService(quoteRepository)
	quoteHandler := handlers.NewQuoteHandler(quoteService)

	r := mux.NewRouter()
	r.HandleFunc("/quotes/{id}", quoteHandler.DeleteQuoteByID).Methods("DELETE")
	r.HandleFunc("/quotes", quoteHandler.CreateQuote).Methods("POST")
	r.HandleFunc("/quotes", quoteHandler.GetAllQuotes).Methods("GET")
	r.HandleFunc("/quotes/random", quoteHandler.GetRandomQuote).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
