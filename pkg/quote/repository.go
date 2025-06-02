package quote

import (
	"gorm.io/gorm"
)

type QuoteRepository interface {
	CreateQuote(quote Quote) error
	GetAllQuotes(authorParam string) ([]Quote, error)
	GetRandomQuote() (Quote, error)
	DeleteQuoteByID(quote_id string) error
}

type quoteRepository struct {
	db *gorm.DB
}


func NewRepository(db *gorm.DB) QuoteRepository {
	return &quoteRepository{db: db}
}

func (r *quoteRepository) DeleteQuoteByID(quote_id string) error {
	return r.db.Delete(&Quote{}, "id = ?", quote_id).Error
}

func (r *quoteRepository) CreateQuote(quote Quote) error {
	return r.db.Create(&quote).Error
}

func (r *quoteRepository) GetAllQuotes(authorParam string) ([]Quote, error) {
	var quotes []Quote
	err := r.db.Where(&Quote{Author: authorParam}).Find(&quotes).Error
	return quotes, err
}

func (r *quoteRepository) GetRandomQuote() (Quote, error) {
	var quote Quote
	err := r.db.Order("RANDOM()").Limit(1).Find(&quote).Error
	return quote, err
}
