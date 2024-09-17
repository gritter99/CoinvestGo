package repository

import (
	"coinvest/src/configs"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Estruturas para armazenar dados de ações e criptomoedas
type Stock struct {
	Symbol      string
	CompanyName string
	Price       float64
	Volume      int64
	MarketCap   float64
	CreatedAt   string
}

type Crypto struct {
	Symbol    string
	Name      string
	Price     float64
	Volume    int64
	MarketCap float64
	CreatedAt string
}
type PostgresRepository struct {
	conn *pgx.Conn
}

func NewPostgresRepository() (*PostgresRepository, error) {
	conn, err := configs.GetPostgresConnection()
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao PostgreSQL: %v", err)
	}

	return &PostgresRepository{conn: conn}, nil
}

// GetStockBySymbol obtém informações completas sobre uma ação.
func (r *PostgresRepository) GetStockBySymbol(symbol string) (Stock, error) {
	var stock Stock
	query := `SELECT symbol, company_name, price, volume, market_cap, created_at FROM stocks WHERE symbol=$1`
	err := r.conn.QueryRow(context.Background(), query, symbol).Scan(
		&stock.Symbol,
		&stock.CompanyName,
		&stock.Price,
		&stock.Volume,
		&stock.MarketCap,
		&stock.CreatedAt,
	)
	if err != nil {
		return Stock{}, fmt.Errorf("erro ao buscar ação: %v", err)
	}
	return stock, nil
}

// GetCryptoBySymbol obtém informações completas sobre uma criptomoeda.
func (r *PostgresRepository) GetCryptoBySymbol(symbol string) (Crypto, error) {
	var crypto Crypto
	query := `SELECT symbol, name, price, volume, market_cap, created_at FROM cryptos WHERE symbol=$1`
	err := r.conn.QueryRow(context.Background(), query, symbol).Scan(
		&crypto.Symbol,
		&crypto.Name,
		&crypto.Price,
		&crypto.Volume,
		&crypto.MarketCap,
		&crypto.CreatedAt,
	)
	if err != nil {
		return Crypto{}, fmt.Errorf("erro ao buscar criptomoeda: %v", err)
	}
	return crypto, nil
}

// AddStockDetail insere ou atualiza detalhes sobre uma ação.
func (r *PostgresRepository) AddStockDetail(symbol, companyName string, price float64, volume float64, marketCap float64) error {
	query := `
		INSERT INTO stocks (symbol, company_name, price, volume, market_cap)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (symbol) DO UPDATE
		SET company_name = EXCLUDED.company_name,
		    price = EXCLUDED.price,
		    volume = EXCLUDED.volume,
		    market_cap = EXCLUDED.market_cap,
		    created_at = EXCLUDED.created_at;
	`

	_, err := r.conn.Exec(context.Background(), query, symbol, companyName, price, volume, marketCap)
	if err != nil {
		return fmt.Errorf("erro ao inserir detalhes da ação: %v", err)
	}

	return nil
}

// AddCryptoDetail insere ou atualiza detalhes sobre uma criptomoeda.
func (r *PostgresRepository) AddCryptoDetail(symbol, name string, price float64, volume int64, marketCap float64) error {
	query := `
		INSERT INTO cryptos (symbol, name, price, volume, market_cap)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (symbol) DO UPDATE
		SET name = EXCLUDED.name,
		    price = EXCLUDED.price,
		    volume = EXCLUDED.volume,
		    market_cap = EXCLUDED.market_cap,
		    created_at = EXCLUDED.created_at;
	`

	_, err := r.conn.Exec(context.Background(), query, symbol, name, price, volume, marketCap)
	if err != nil {
		return fmt.Errorf("erro ao inserir detalhes da criptomoeda: %v", err)
	}

	return nil
}
