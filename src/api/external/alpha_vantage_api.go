package external

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	alphaVantageBaseURL = "https://www.alphavantage.co/query?"
)

// Struct para armazenar o Global Quote
type GlobalQuoteResponse struct {
	GlobalQuote struct {
		Symbol string `json:"01. symbol"`
		Price  string `json:"05. price"`
		Volume string `json:"06. volume"`
	} `json:"Global Quote"`
}

// Struct para armazenar o Overview
type OverviewResponse struct {
	Symbol    string `json:"Symbol"`
	Name      string `json:"Name"`
	MarketCap string `json:"MarketCapitalization"`
}

// Struct final para retornar os dados necessários
type StockData struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	Volume    string `json:"volume"`
	MarketCap string `json:"market_cap"`
}

// Função para obter os dados de GLOBAL_QUOTE
func GetGlobalQuote(symbol string) (*GlobalQuoteResponse, error) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
	}

	url := fmt.Sprintf("%sfunction=GLOBAL_QUOTE&symbol=%s&apikey=%s", alphaVantageBaseURL, symbol, os.Getenv("ALPHA_VANTAGE_API_KEY"))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var quoteResp GlobalQuoteResponse
	err = json.Unmarshal(body, &quoteResp)
	if err != nil {
		return nil, err
	}

	if quoteResp.GlobalQuote.Symbol == "" {
		return nil, errors.New("no data found for symbol")
	}

	return &quoteResp, nil
}

// Função para obter os dados de OVERVIEW
func GetOverview(symbol string) (*OverviewResponse, error) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
	}
	url := fmt.Sprintf("%sfunction=OVERVIEW&symbol=%s&apikey=%s", alphaVantageBaseURL, symbol, os.Getenv("ALPHA_VANTAGE_API_KEY"))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var overviewResp OverviewResponse
	err = json.Unmarshal(body, &overviewResp)
	if err != nil {
		return nil, err
	}

	if overviewResp.Symbol == "" {
		return nil, errors.New("no overview data found for symbol")
	}

	return &overviewResp, nil
}

// Função para consolidar e retornar os dados finais
func GetStockData(symbol string) (*StockData, error) {
	// Obter os dados de GLOBAL_QUOTE
	quote, err := GetGlobalQuote(symbol)
	if err != nil {
		return nil, err
	}

	// Obter os dados de OVERVIEW
	overview, err := GetOverview(symbol)
	if err != nil {
		return nil, err
	}

	// Montar o objeto final StockData
	stockData := &StockData{
		Symbol:    quote.GlobalQuote.Symbol,
		Name:      overview.Name,
		Price:     quote.GlobalQuote.Price,
		Volume:    quote.GlobalQuote.Volume,
		MarketCap: overview.MarketCap,
	}

	return stockData, nil
}
