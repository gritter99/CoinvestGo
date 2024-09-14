package unit

import (
	"testing"

	config "coinvest/src/configs"
	"context"
)

func TestMongoDBConnection(t *testing.T) {
	client, err := config.ConnectMongoDB()
	if err != nil {
		t.Fatalf("Falha ao conectar ao MongoDB: %v", err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			t.Fatalf("Erro ao desconectar o MongoDB: %v", err)
		}
	}()

	t.Log("Conexão com o MongoDB estabelecida com sucesso")
}
