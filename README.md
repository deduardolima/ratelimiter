# 🚦 Rate Limiter em Go

Rate limiter de alta performance desenvolvido em **Go**, com armazenamento em **Redis**,
suporte a limitação por **IP** e por **Token de acesso**, e bloqueio configurável por janela de tempo.

---

## 🏗️ Arquitetura

```
ratelimiter/
├── cmd/
│   └── server/
│       └── main.go              # Entrypoint da aplicação
├── configs/
│   └── config.go                # Carregamento de variáveis de ambiente
├── infra/
│   └── database/
│       ├── interface.go         # Contrato da camada de persistência
│       └── redis_store.go       # Implementação com Redis
├── internal/
│   ├── limiter/
│   │   └── ratelimiter.go       # Lógica central do rate limiting
│   └── middleware/
│       └── ratelimiter_middleware.go  # Middleware HTTP
└── test/
    └── ratelimiter_test.go      # Testes de integração
```

---

## ⚙️ Stack Técnica

| Camada          | Tecnologia              |
|-----------------|-------------------------|
| Linguagem       | Go 1.21+                |
| Armazenamento   | Redis                   |
| Containerização | Docker + Docker Compose |

---

## ✨ Funcionalidades

- **Limitação por IP** — controla requisições por endereço IP do cliente
- **Limitação por Token** — tokens de acesso têm limites independentes e prioritários
- **Bloqueio configurável** — tempo de bloqueio definido via variável de ambiente
- **Persistência em Redis** — estado compartilhado, pronto para ambientes distribuídos
- **Middleware HTTP plugável** — integração transparente com qualquer handler Go

---

## 🔑 Decisão Técnica — Token tem prioridade sobre IP

Quando uma requisição possui um **token de acesso** no header, o limiter usa
a regra do token e **ignora** o limite por IP. Isso permite políticas diferenciadas
por cliente sem complexidade adicional:

```
Requisição com Token → aplica RATE_LIMIT_TOKEN
Requisição sem Token → aplica RATE_LIMIT_IP
```

**Por que Redis?**
- Estado compartilhado entre múltiplas instâncias da aplicação
- Operações atômicas via comandos nativos (`INCR`, `EXPIRE`)
- TTL nativo — bloqueio expira automaticamente sem job de limpeza

---

## 🚀 Como Rodar

### 1. Clone o repositório

```bash
git clone https://github.com/deduardolima/ratelimiter.git
cd ratelimiter
```

### 2. Configure as variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto:

```env
REDIS_ADDR=localhost:6379
RATE_LIMIT_IP=5          # máximo de requisições/segundo por IP
RATE_LIMIT_TOKEN=10      # máximo de requisições/segundo por Token
BLOCK_TIME=300           # tempo de bloqueio em segundos
```

### 3. Suba os containers

```bash
docker-compose up --build
```

| Serviço | URL |
|---------|-----|
| API     | http://localhost:8080 |
| Redis   | localhost:6379 |

---

## 📡 Como Usar

### Requisição por IP
```bash
curl http://localhost:8080/
```

### Requisição com Token
```bash
curl http://localhost:8080/ -H "API_KEY: meu-token-aqui"
```

### Resposta ao atingir o limite
```
HTTP 429 Too Many Requests

you have reached the maximum number of requests or actions allowed within a certain time frame
```

---

## 🧪 Testes

```bash
docker-compose run test
```

Sobe um container isolado e executa a suíte de testes da aplicação.

---

## 🧠 O que Este Projeto Demonstra

- Implementação de **rate limiting** do zero, sem biblioteca externa
- Uso de **Redis** como store distribuído com TTL nativo
- Separação entre **lógica de negócio** (`limiter/`) e **infraestrutura** (`infra/`)
- **Interface de repositório** desacoplada — fácil trocar Redis por outro store
- **Middleware HTTP** reutilizável e configurável via environment
