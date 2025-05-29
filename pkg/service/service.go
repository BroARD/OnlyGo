package service

type QuoteService interface {
	CreateQuote(quote Quote) error
	GetAllQuotes() ([]Quote, error)
	GetRandomQuote() (Quote, error)
}

type quoteService struct {
	repo QuoteRepository
}

func NewQuoteService(r QuoteRepository) QuoteService {
	return &quoteService{repo: r}
}

func (s *quoteService) CreateQuote(quote Quote) error {
	return s.repo.CreateQuote(quote)
}

func (s *quoteService) GetAllQuotes() ([]Quote, error) {
	return s.repo.GetAllQuotes()
}

func (s *quoteService) GetRandomQuote() (Quote, error) {
	return s.repo.GetRandomQuote()
}
