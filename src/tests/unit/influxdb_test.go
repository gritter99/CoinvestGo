package unit

import (
	"coinvest/src/configs"
	"testing"
)

func TestInfluxDBConnection(t *testing.T) {
	client, err := configs.ConnectInfluxDB()
	if err != nil {
		t.Fatalf("Falha ao conectar ao InfluxDB: %v", err)
	}

	defer client.Close()

	t.Log("Conexão com o InfluxDB estabelecida com sucesso")
}
