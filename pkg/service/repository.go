package service

import (
	"gorm.io/gorm"
)

type QuoteRepository interface {
	CreateQuote(quote Quote) error
	GetAllQuotes() ([]Quote, error)
	GetRandomQuote() (Quote, error)
}

type quoteRepository struct {
	db *gorm.DB
}

func NewQuoteRepository(db *gorm.DB) QuoteRepository {
	return &quoteRepository{db: db}
}


func (r *quoteRepository) CreateQuote(quote Quote) error {
	return r.db.Create(&quote).Error
}

func (r *quoteRepository) GetAllQuotes() ([]Quote, error) {
	var quotes []Quote
	err := r.db.Find(&quotes).Error
	return quotes, err
}

func (r *quoteRepository) GetRandomQuote() (Quote, error) {
	var quote Quote
	err := r.db.Order("RANDOM()").Limit(1).Find(&quote).Error
	return quote, err
}
