package services

import (
	"coinvest/src/api/external"
	"coinvest/src/repository"
	"fmt"
	"strconv"
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
	// Consome a API Alpha Vantage para obter os dados da ação
	data, err := external.GetStockData(symbol)
	if err != nil {
		return fmt.Errorf("failed to get stock data for symbol: %s, error: %v", symbol, err)
	}

	// Validando e extraindo os dados necessários
	companyName := data.Name
	if companyName == "" {
		return fmt.Errorf("company name not found for symbol: %s", symbol)
	}

	price, err := parseFloat(data.Price)
	if err != nil {
		return fmt.Errorf("invalid price for symbol: %s, error: %v", symbol, err)
	}

	volume, err := parseFloat(data.Volume)
	if err != nil {
		return fmt.Errorf("invalid volume for symbol: %s, error: %v", symbol, err)
	}

	marketCap, err := parseFloat(data.MarketCap)
	if err != nil {
		return fmt.Errorf("invalid market cap for symbol: %s, error: %v", symbol, err)
	}

	// Adiciona os dados da ação no banco de dados
	err = s.stockRepo.AddStockDetail(
		symbol,
		companyName,
		price,
		volume,
		marketCap,
	)
	if err != nil {
		return fmt.Errorf("error adding stock details to repository: %v", err)
	}

	return nil
}

// Função auxiliar para converter string em float64
func parseFloat(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}
