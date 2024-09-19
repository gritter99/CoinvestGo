package configs

import (
	"context"
	"fmt"

	config "coinvest/src/helpers"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// ConnectInfluxDB faz a conexão com o InfluxDB e retorna um cliente.
//
// Esta função lê as credenciais do InfluxDB das variáveis de ambiente INFLUXDB_URL e INFLUXDB_TOKEN,
// verifica se há uma conexão válida com o servidor do InfluxDB e retorna um cliente conectado.
//
// Retorna:
//   - influxdb2.Client: Cliente conectado ao InfluxDB
//   - error: Se ocorrer algum erro durante a conexão
func ConnectInfluxDB() (influxdb2.Client, error) {
	influxURL := config.InfluxURL
	influxToken := config.InfluxToken

	client := influxdb2.NewClient(influxURL, influxToken)

	status, err := client.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("não foi possível verificar a conexão com o InfluxDB: %v", err)
	}

	if !status {
		return nil, fmt.Errorf("conexão com o InfluxDB falhou")
	}

	return client, nil
}
