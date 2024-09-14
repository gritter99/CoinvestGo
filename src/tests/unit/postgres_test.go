package unit

import (
	config "coinvest/src/configs"
	"context"
	"testing"
)

func TestPostgresConnection(t *testing.T) {
	conn, err := config.GetPostgresConnection()
	if err != nil {
		t.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	t.Log("Conexão com o PostgresDB estabelecida com sucesso")
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "SELECT 1")
	if err != nil {
		t.Fatalf("PostgreSQL connection test failed: %v", err)
	}

}
