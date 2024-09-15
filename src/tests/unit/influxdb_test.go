package unit

import (
	"coinvest/src/configs"
	"os"
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

func TestConnectToInfluxDB(t *testing.T) {
	tests := []struct {
		name        string
		envVars     map[string]string
		wantErr     bool
		expectedMsg string
	}{
		{
			name:        "Failed connection due to missing URL",
			envVars:     map[string]string{"INFLUX_URL": "", "INFLUX_TOKEN": "token"},
			wantErr:     true,
			expectedMsg: "variáveis de ambiente INFLUXDB_URL ou INFLUXDB_TOKEN não foram definidas",
		},
		{
			name:        "Failed connection due to invalid token",
			envVars:     map[string]string{"INFLUX_URL": "http://invalid-url.com", "INFLUX_TOKEN": "invalid_token"},
			wantErr:     true,
			expectedMsg: "não foi possível conectar ao InfluxDB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.envVars {
				os.Setenv(k, v)
			}

			_, err := configs.ConnectInfluxDB()
			if (err == nil) == tt.wantErr {
				t.Errorf("ConnectInfluxDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer func() {
				for _, k := range []string{"INFLUX_URL", "INFLUX_TOKEN"} {
					if v := os.Getenv(k); v != "" {
						os.Unsetenv(k)
					}
				}
			}()
		})
	}
}
