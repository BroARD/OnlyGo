package service


type QuoteService interface {
	CreateQuote(quote Quote) error
	GetAllQuotes(authorParam string) ([]Quote, error)
	GetRandomQuote() (Quote, error)
	DeleteQuoteByID(quote_id string) error
}

type quoteService struct {
	repo QuoteRepository
}

func NewQuoteService(r QuoteRepository) QuoteService {
	return &quoteService{repo: r}
}

func (s *quoteService) DeleteQuoteByID(quote_id string) error {
	return s.repo.DeleteQuoteByID(quote_id)
}

func (s *quoteService) CreateQuote(quote Quote) error {
	return s.repo.CreateQuote(quote)
}

func (s *quoteService) GetAllQuotes(authorParam string) ([]Quote, error) {
	return s.repo.GetAllQuotes(authorParam)
}

func (s *quoteService) GetRandomQuote() (Quote, error) {
	return s.repo.GetRandomQuote()
}
