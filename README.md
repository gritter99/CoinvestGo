# CoinvestGo

Projeto de obtenção de dados de criptomoedas e ações, com principal objetivo de manipular e aprender a linguagem Golang.

## CoinvestGo - TODO List

Este é o **TODO list** para o desenvolvimento do projeto *CoinvestGo*. Abaixo estão os próximos passos organizados por tarefas:

### 1. Definir as Entidades do Projeto

- [x] Definir entidades que representam ações e criptomoedas:
  - [x] Criar o modelo `Stock` (Ações)
  - [x] Criar o modelo `CryptoCoin` (Criptomoedas)
  - [x] Definir campos relevantes: preço, volume, data de transação, etc.

### 2. Implementar Repositórios

- [x] Criar repositórios para manipulação de dados:
  - [x] Implementar repositórios
  - [x] Implementar operações de inserção de registros.

### 3. Desenvolver a Lógica dos Serviços

- [ ] Implementar a lógica de negócios nos serviços:
  - [x] Criar `StockService` para manipular dados de ações.
  - [x] Criar `CryptoService` para manipular dados de criptomoedas.
  - [ ] **Usar Redis como cache** para armazenar dados acessados frequentemente (ações e criptomoedas).
  - [ ] **Armazenar dados históricos de séries temporais no InfluxDB** (preço, volume, variação de ações e criptomoedas).

### 5. Testes Unitários e de Integração

- [x] Desenvolver testes para garantir o bom funcionamento da aplicação:
  - [x] Escrever testes unitários para os repositórios.
  - [x] Implementar testes de integração para verificar a interação entre serviços e bancos de dados.

### 7. Configurar Logs e Monitoramento

- [x] Implementar logging e monitoramento da aplicação:
  - [x] Adicionar logs nas camadas de serviço e repositório.
