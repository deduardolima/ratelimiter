# Rate Limiter em Go

## Descrição

Este projeto implementa um rate limiter em Go que limita o número de requisições por segundo com base em um endereço IP específico ou em um token de acesso.


## Estrutura do Projeto

- ratelimiter/
  - cmd/
    - server/
      - main.go
  - configs
    - config.go 
  - infra/
    - database/
      - inteface.go
      - redis_store.go          
  - internal/
    - limiter/
      - ratelimiter.go
    - middleware/
      - ratelimiter_middleware.go
  - test/
    - ratelimiter_test.go
  - .env  
  - docker-compose.yml
  - Dockerfile
  - go.mod
  - go.sum
  - README.md


## Pré-requisitos
- Docker
- Docker Compose

## Configuração
```
git clone https://github.com/deduardolima/ratelimiter.git
cd ratelimiter
```

## Instalação e Execução com Docker
Construa e inicie os containers:
```
docker-compose up --build
```

isso irá construir a imagem do aplicativo e iniciar os serviços definidos no docker-compose.yml, incluindo o banco de dados e o aplicativo.

A aplicação estará acessível em http://localhost:8080

## Testes

```
docker-compose run test
```
Esse comando ira criar um container para rodar testes da aplicação


### Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis exemplo :

- REDIS_ADDR=localhost:6379
- RATE_LIMIT_IP=5
- RATE_LIMIT_TOKEN=10
- BLOCK_TIME=300

