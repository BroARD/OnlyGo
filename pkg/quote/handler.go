package quote

import (
	"OnlyGo/logging"
	"OnlyGo/pkg/handlers"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type quoteHandler struct {
	logger logging.Logger
	service QuoteService
}

func NewHandler(serv QuoteService, logger logging.Logger) handlers.Handler {
	return &quoteHandler{service: serv, logger: logger}
}

func (h *quoteHandler) Register(router *mux.Router) {
	router.HandleFunc("/quotes/{id}", h.DeleteQuoteByID).Methods("DELETE")
	router.HandleFunc("/quotes", h.CreateQuote).Methods("POST")
	router.HandleFunc("/quotes", h.GetAllQuotes).Methods("GET")
	router.HandleFunc("/quotes/random", h.GetRandomQuote).Methods("GET")
}

func (h *quoteHandler) CreateQuote(w http.ResponseWriter, router *http.Request) {
	var quote Quote
	decoder := json.NewDecoder(router.Body)
	err := decoder.Decode(&quote)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	
	h.logger.Info("Create new quote")
	err = h.service.CreateQuote(quote)
	if err != nil {
		http.Error(w, "Bad request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *quoteHandler) GetAllQuotes(w http.ResponseWriter, router *http.Request) {
	author := router.URL.Query().Get("author")
	
	h.logger.Info("Get quotes from DB")
	quotes, err := h.service.GetAllQuotes(author)
	if err != nil {
		http.Error(w, "Could not get quotes", http.StatusInternalServerError)
		return
	}
	jsonQuotes, err := json.Marshal(quotes)
	if err != nil {
		http.Error(w, "Could not transform from json to bytes", http.StatusInternalServerError)
		return
	}
	w.Write(jsonQuotes)
	w.WriteHeader(http.StatusOK)
}

func (h *quoteHandler) GetRandomQuote(w http.ResponseWriter, router *http.Request) {
	quote, err := h.service.GetRandomQuote()
	if err != nil {
		http.Error(w, "Could not get random quote from db", http.StatusInternalServerError)
		return
	}
	jsonQuote, err := json.Marshal(quote)
	if err != nil {
		http.Error(w, "Could not transform from json to bytes", http.StatusInternalServerError)
		return
	}
	w.Write(jsonQuote)
	w.WriteHeader(http.StatusOK)
}

func (h *quoteHandler) DeleteQuoteByID(w http.ResponseWriter, router *http.Request) {
	parts := strings.Split(router.URL.Path, "/") //Считывание ID цитаты путём парсинга, не использую сторонние библиотеки(не совсем корректно, но по заданию)
	if len(parts) < 3 || parts[1] != "quotes" {
		http.NotFound(w, router)
		return
	}
	id := parts[2]
	err := h.service.DeleteQuoteByID(id)
	if err != nil {
		http.Error(w, "Could not delete quotes by id", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
