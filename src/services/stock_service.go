package services

import (
	"coinvest/src/api/external"
	"coinvest/src/repository"
	"fmt"
)

type StockService struct {
	stockRepo *repository.PostgresRepository
}

func NewStockService(stockRepo *repository.PostgresRepository) *StockService {
	return &StockService{
		stockRepo: stockRepo,
	}
}

// Atualiza os detalhes de uma ação no banco de dados
func (s *StockService) UpdateStockDetails(symbol string) error {
	data, err := external.GetStockData(external.InitFinnhubClient(), symbol)
	if err != nil {
		return fmt.Errorf("failed to get stock data for symbol: %s, error: %v", symbol, err)
	}

	// Check if company name exists
	companyName := data.CompanyName
	if companyName == "" {
		return fmt.Errorf("company name not found for symbol: %s", symbol)
	}

	price := data.Price
	if price == 0 {
		return fmt.Errorf("invalid price for symbol: %s", symbol)
	}

	marketCap := data.MarketCap
	if marketCap == 0 {
		return fmt.Errorf("invalid market cap for symbol: %s", symbol)
	}

	err = s.stockRepo.AddStockDetail(
		symbol,
		companyName,
		float64(price),
		float64(marketCap),
	)
	if err != nil {
		return fmt.Errorf("error adding stock details to repository: %v", err)
	}

	return nil
}
