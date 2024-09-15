package configs

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

// ConnectRedis faz a conexão com o Redis e retorna um cliente
func ConnectRedis() (*redis.Client, error) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		return nil, fmt.Errorf("variáveis de ambiente INFLUX_URL ou INFLUX_TOKEN não foram definidas")
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	if redisHost == "" || redisPort == "" {
		return nil, fmt.Errorf("variáveis de ambiente REDIS_HOST ou REDIS_PORT não foram definidas")
	}

	// Cria o cliente Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword, // vazio se não houver senha
		DB:       0,             // usa o banco de dados padrão
	})

	// Testa a conexão com o Redis
	ctx := context.Background()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("não foi possível conectar ao Redis: %v", err)
	}

	return client, nil
}
