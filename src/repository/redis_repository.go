package repository

import (
	"coinvest/src/configs"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository() (*RedisRepository, error) {
	client, err := configs.ConnectRedis()
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao Redis: %v", err)
	}

	return &RedisRepository{client: client}, nil
}

// função para armazenar o preço mais recente no Redis
func (r *RedisRepository) CacheAssetPrice(symbol string, price float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.client.Set(ctx, symbol, price, 10*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("erro ao armazenar preço no cache: %v", err)
	}
	return nil
}
