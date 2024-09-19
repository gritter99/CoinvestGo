package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	MongoHost     string
	MongoPort     string
	MongoUser     string
	MongoPassword string
	RedisHost     string
	RedisPort     string
	RedisPassword string
	InfluxURL     string
	InfluxToken   string
	FinnhubAPIKey string
)

// init é automaticamente chamado quando o pacote é importado
func init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	MongoHost = os.Getenv("MONGO_HOST")
	MongoPort = os.Getenv("MONGO_PORT")
	MongoUser = os.Getenv("MONGO_USER")
	MongoPassword = os.Getenv("MONGO_PASSWORD")
	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	InfluxURL = os.Getenv("INFLUX_URL")
	InfluxToken = os.Getenv("INFLUX_TOKEN")
	FinnhubAPIKey = os.Getenv("FINNHUB_API_KEY")
}
