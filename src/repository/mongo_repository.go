package repository

import (
	"coinvest/src/configs"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	client *mongo.Client
}

func NewMongoRepository() (*MongoRepository, error) {
	client, err := configs.ConnectMongoDB()
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao MongoDB: %v", err)
	}

	return &MongoRepository{client: client}, nil
}

// função para inserir uma notícia de ativo
func (r *MongoRepository) AddAssetNews(symbol string, title string, content string) error {
	collection := r.client.Database("coinvestgo").Collection("news")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	news := bson.D{
		{Key: "symbol", Value: symbol},
		{Key: "title", Value: title},
		{Key: "content", Value: content},
		{Key: "date", Value: time.Now()},
	}

	_, err := collection.InsertOne(ctx, news)
	if err != nil {
		return fmt.Errorf("erro ao adicionar notícia: %v", err)
	}
	return nil
}
