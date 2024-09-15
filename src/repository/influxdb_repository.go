package repository

import (
	"context"
	"fmt"
	"time"

	"coinvest/src/configs"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// InfluxRepository representa o repositório de dados do InfluxDB.
type InfluxRepository struct {
	client influxdb2.Client
}

// NewInfluxRepository cria uma nova instância do repositório InfluxDB.
func NewInfluxRepository() (*InfluxRepository, error) {
	client, err := configs.ConnectInfluxDB()
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao InfluxDB: %v", err)
	}

	return &InfluxRepository{
		client: client,
	}, nil
}

// AddAssetPrice insere o preço de um ativo em um bucket específico.
func (r *InfluxRepository) AddAssetPrice(bucket, symbol string, price float64) error {
	p := influxdb2.NewPointWithMeasurement("prices").
		AddTag("symbol", symbol).
		AddField("price", price).
		SetTime(time.Now())

	writeAPI := r.client.WriteAPIBlocking("", bucket)

	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return fmt.Errorf("erro ao adicionar preço do ativo: %v", err)
	}
	return nil
}
