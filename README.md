# CoinvestGo

Projeto de obtenção de dados de criptomoedas e ações.

Funcionalidade de comparação e visualização de tendências.

## CoinvestGo - TODO List

Este é o **TODO list** para o desenvolvimento do projeto *CoinvestGo*. Abaixo estão os próximos passos organizados por tarefas:

### 1. Definir as Entidades do Projeto

- [ ] Definir entidades que representam ações e criptomoedas:
  - [ ] Criar o modelo `Stock` (Ações)
  - [ ] Criar o modelo `CryptoCoin` (Criptomoedas)
  - [ ] Definir campos relevantes: preço, volume, data de transação, etc.

### 2. Implementar Repositórios

- [ ] Criar repositórios para manipulação de dados:
  - [ ] Implementar CRUD para PostgreSQL (`postgres_repository.go`)
  - [ ] Implementar CRUD para MongoDB (`mongo_repository.go`)
  - [ ] Implementar operações de inserção, busca, atualização e remoção de registros.
  - [ ] Criar consultas específicas, como filtro por intervalo de tempo para ações e criptomoedas.

### 3. Desenvolver a Lógica dos Serviços

- [ ] Implementar a lógica de negócios nos serviços:
  - [ ] Criar `StockService` para manipular dados de ações.
  - [ ] Criar `CryptoService` para manipular dados de criptomoedas.
  - [ ] Desenvolver funções para comparar desempenhos entre ações e criptomoedas (taxa de crescimento, volatilidade, etc.).
  - [ ] Implementar lógica de previsão de tendências de mercado.
  - [ ] **Usar Redis como cache** para armazenar dados acessados frequentemente (ações e criptomoedas).
  - [ ] **Armazenar dados históricos de séries temporais no InfluxDB** (preço, volume, variação de ações e criptomoedas).

### 4. Desenvolver os Handlers da API

- [ ] Expor a lógica dos serviços via API:
  - [ ] Criar `stock_handler.go` para expor endpoints relacionados a ações.
  - [ ] Criar `crypto_handler.go` para expor endpoints relacionados a criptomoedas.
  - [ ] Criar rotas de API:
    - [ ] `/stocks` para buscar informações sobre ações.
    - [ ] `/cryptos` para buscar informações sobre criptomoedas.
    - [ ] `/compare` para comparar desempenho entre ações e criptomoedas.

### 5. Testes Unitários e de Integração

- [ ] Desenvolver testes para garantir o bom funcionamento da aplicação:
  - [ ] Escrever testes unitários para os repositórios.
  - [ ] Escrever testes unitários para os serviços.
  - [ ] Implementar testes de integração para verificar a interação entre serviços e bancos de dados.
  - [ ] **Testar integração com Redis** para cache.
  - [ ] **Testar integração com InfluxDB** para séries temporais.
  - [ ] Garantir que a API esteja funcionando corretamente com testes de integração.

### 6. Implementar Autenticação e Autorização (Opcional)

- [ ] Adicionar autenticação e autorização (caso necessário):
  - [ ] Implementar autenticação com JWT.
  - [ ] Utilizar Redis para armazenar tokens de sessão.
  - [ ] Criar controle de autorização para acesso aos dados.

### 7. Configurar Logs e Monitoramento

- [ ] Implementar logging e monitoramento da aplicação:
  - [ ] Adicionar logs nas camadas de serviço e repositório.
  - [ ] Usar InfluxDB para armazenar métricas de uso da API.
  - [ ] Monitorar a saúde da aplicação e gerar gráficos sobre o desempenho.
