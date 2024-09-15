package repository

import (
	"coinvest/src/configs"
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type InfluxRepository struct {
	client api.WriteAPIBlocking
}

func NewInfluxRepository() (*InfluxRepository, error) {
	client, err := configs.ConnectInfluxDB()
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao InfluxDB: %v", err)
	}

	writeAPI := client.WriteAPIBlocking("", "coinvestgo/prices")
	return &InfluxRepository{client: writeAPI}, nil
}

// função para inserir preço de um ativo
func (r *InfluxRepository) AddAssetPrice(symbol string, price float64) error {
	p := influxdb2.NewPointWithMeasurement("prices").
		AddTag("symbol", symbol).
		AddField("price", price).
		SetTime(time.Now())

	err := r.client.WritePoint(context.Background(), p)
	if err != nil {
		return fmt.Errorf("erro ao adicionar preço do ativo: %v", err)
	}
	return nil
}
