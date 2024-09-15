package repository

import (
	"coinvest/src/configs"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

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

// função para obter informações de um ativo
func (r *PostgresRepository) GetAssetBySymbol(symbol string) (string, error) {
	var name string
	err := r.conn.QueryRow(context.Background(), "SELECT name FROM assets WHERE symbol=$1", symbol).Scan(&name)
	if err != nil {
		return "", fmt.Errorf("erro ao buscar ativo: %v", err)
	}
	return name, nil
}
