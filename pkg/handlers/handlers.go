package handlers

import (
	"OnlyGo/pkg/service"
	"encoding/json"
	"net/http"
	"strings"
)

type quoteHandler struct {
	service service.QuoteService
}

func NewQuoteHandler(serv service.QuoteService) *quoteHandler {
	return &quoteHandler{service: serv}
}

func (h *quoteHandler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var quote service.Quote
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&quote)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = h.service.CreateQuote(quote)
	if err != nil{
		http.Error(w, "Bad request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *quoteHandler) GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
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

func (h *quoteHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
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

func (h *quoteHandler) DeleteQuoteByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")		//Считывание ID цитаты путём парсинга, не использую сторонние библиотеки(не совсем корректно, но по заданию)
    if len(parts) < 3 || parts[1] != "quotes" {
        http.NotFound(w, r)
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
