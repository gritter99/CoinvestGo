package unit

import (
	config "coinvest/src/configs"
	"context"
	"testing"
)

func TestRedisConnection(t *testing.T) {
	client, err := config.ConnectRedis()
	if err != nil {
		t.Fatalf("Falha ao conectar ao Redis: %v", err)
	}

	defer func() {
		if err = client.Close(); err != nil {
			t.Fatalf("Erro ao desconectar do Redis: %v", err)
		}
	}()

	ctx := context.Background()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		t.Fatalf("Falha ao pingar o Redis: %v", err)
	}
}
