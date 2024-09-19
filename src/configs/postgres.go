package configs

import (
	config "coinvest/src/helpers"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetPostgresConnection() (*pgx.Conn, error) {
	dbHost := config.DBHost
	dbPort := config.DBPort
	dbUser := config.DBUser
	dbPassword := config.DBPassword
	dbName := config.DBName

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
