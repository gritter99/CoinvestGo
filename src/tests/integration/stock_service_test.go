package integration

import (
	repository "coinvest/src/repository"
	services "coinvest/src/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationUpdateStockDetails(t *testing.T) {
	stockRepo, err := repository.NewPostgresRepository()
	assert.Nil(t, err)

	service := services.NewStockService(stockRepo)

	err = service.UpdateStockDetails("AAPL")
	assert.Nil(t, err)
}
