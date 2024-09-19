package external

import (
	"context"
	"fmt"

	config "coinvest/src/helpers"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

type StockData struct {
	Symbol      string
	CompanyName string
	Price       float32
	MarketCap   float32
}

// InitFinnhubClient inicializa o cliente Finnhub
func InitFinnhubClient() *finnhub.DefaultApiService {
	apiKey := config.FinnhubAPIKey
	if apiKey == "" {
		fmt.Println("API key not set")
	}
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", apiKey)
	client := finnhub.NewAPIClient(cfg).DefaultApi
	return client
}

// GetStockData consome a API da Finnhub e retorna os dados de uma ação
func GetStockData(client *finnhub.DefaultApiService, symbol string) (*StockData, error) {
	companyProfile, _, err := client.CompanyProfile2(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter perfil da empresa: %v", err)
	}

	quote, _, err := client.Quote(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter cotação da empresa: %v", err)
	}

	stockData := &StockData{
		Symbol:      symbol,
		CompanyName: companyProfile.GetName(),
		Price:       quote.GetC(),
		MarketCap:   companyProfile.GetMarketCapitalization(),
	}

	return stockData, nil
}
