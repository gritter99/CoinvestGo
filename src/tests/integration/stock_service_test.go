package integration

import (
	repository "coinvest/src/repository"
	services "coinvest/src/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationUpdateStockDetails(t *testing.T) {
	var mainStockTickers = []string{
		"AAPL", "MSFT", "GOOGL", "AMZN",
		"TSLA", "NVDA", "ADBE", "PYPL", "NFLX",
		"UNH", "JPM", "V", "MA", "DIS",
		"PG", "HD", "KO", "JNJ", "MRK",
		"INTC", "AMD", "CRM", "SQ", "BABA",
		"TSM", "TXN", "AVGO", "MU", "AMD",
		"CSCO", "ORCL", "IBM", "ACN", "HRL",
		"COST", "WMT", "PEP", "CLX", "EMR",
	}
	stockRepo, err := repository.NewPostgresRepository()
	assert.Nil(t, err)

	service := services.NewStockService(stockRepo)

	for _, ticker := range mainStockTickers {
		err := service.UpdateStockDetails(ticker)
		if err != nil {
			t.Errorf("Failed to update stock details for %s: %v", ticker, err)
		}
	}
}
