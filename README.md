# Order System (Clean Architecture)

Aplicação de gerenciamento de pedidos em Go seguindo **Clean Architecture**, com três formas de acesso ao mesmo domínio:

- API **REST/HTTP**
- API **gRPC**
- API **GraphQL**

Além disso, ao criar um pedido a aplicação dispara um evento de domínio (`OrderCreated`) via RabbitMQ.

---

## Mapa de portas

| Componente | Porta |
|---|---:|
| REST (HTTP) | `8000` |
| gRPC | `50051` |
| GraphQL | `8080` |
| RabbitMQ UI | `15672` |
| MySQL | `3306` |

---

## Requisitos

- [Go 1.25+](https://go.dev/)
- [Docker](https://www.docker.com/) + Docker Compose

---

## Quickstart

### 1) Infra

```bash
docker compose up -d
```

### 2) Env

Faça a gestão das configurações do servidor no arquivo `cmd/ordersystem/.env`:

```env
DB_DRIVER=mysql
DB_HOST=mysql
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=orders
WEB_SERVER_PORT=:8000
GRPC_SERVER_PORT=50051
GRAPHQL_SERVER_PORT=8080
AMQP_URL=amqp://guest:guest@rabbitmq:5672/
```

### 3) Start

```bash
go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go
```

---

## Como usar

### REST (HTTP) — porta 8000

- Create Order

```http
POST http://localhost:8000/order
Content-Type: application/json

{
  "id": "order-999",
  "price": 999.9,
  "tax": 0.5
}
```

- List Order

```http
GET http://localhost:8000/order
```
---

### gRPC 

- Create Order:

```bash
grpcurl -plaintext \
  -proto internal/infra/grpc/protofiles/order.proto \
  -d '{"id":"order-300","price":300.5,"tax":0.5}' \
  localhost:50051 pb.OrderService/CreateOrder
```

- List Order:

```bash
grpcurl -plaintext \
  -proto internal/infra/grpc/protofiles/order.proto \
  localhost:50051 pb.OrderService/ListOrders
```

---

### GraphQ

- Create Order:

```graphql
mutation {
  createOrder(input: {id: "order-1", Price: 100.5, Tax: 0.5}) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

- List Order

```graphql
query {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```
