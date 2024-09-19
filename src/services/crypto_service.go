package services

import (
	"coinvest/src/repository"
	"fmt"
)

type CryptoService struct {
	cryptoRepo *repository.PostgresRepository
}

func NewCryptoService(cryptoRepo *repository.PostgresRepository) *CryptoService {
	return &CryptoService{cryptoRepo: cryptoRepo}
}

// Adicionar uma nova criptomoeda
func (s *CryptoService) AddCrypto(symbol, name string, price float64, volume int64, marketCap float64) error {
	err := s.cryptoRepo.AddCryptoDetail(symbol, name, price, marketCap)
	if err != nil {
		return fmt.Errorf("erro ao adicionar criptomoeda: %v", err)
	}
	return nil
}

// Buscar detalhes de uma criptomoeda por símbolo
func (s *CryptoService) GetCryptoBySymbol(symbol string) (repository.Crypto, error) {
	crypto, err := s.cryptoRepo.GetCryptoBySymbol(symbol)
	if err != nil {
		return crypto, fmt.Errorf("erro ao buscar criptomoeda: %v", err)
	}
	return crypto, nil
}
