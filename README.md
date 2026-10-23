# Flash Sales

Microservice backend for flash sales: limited stock, many buyers at the same second. Six Go services talk over gRPC, events go through Kafka, everything starts with one `docker compose up`.

![Go](https://img.shields.io/badge/Go-1.23-00ADD8?logo=go&logoColor=white)
![gRPC](https://img.shields.io/badge/gRPC-protobuf-244c5a)
![Kafka](https://img.shields.io/badge/Kafka-events-231F20?logo=apachekafka&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-compose-2496ED?logo=docker&logoColor=white)
[![ci](https://github.com/HowkaCoder/flash_sales/actions/workflows/ci.yml/badge.svg)](https://github.com/HowkaCoder/flash_sales/actions/workflows/ci.yml)

## Architecture

```
                 HTTP :3000
client ─────────► gateway ──── JWT check ────► auth_service      :50051
                     │
                     ├──── gRPC ─────────────► warehouse_service :50052
                     ├──── gRPC ─────────────► order_service     :50053 ──► auth, warehouse
                     └──── gRPC ─────────────► sales_service     :50054 ──► warehouse

auth, order, sales ── events ──► Kafka "notifications" ──► notification_service
```

| Service | Responsibility |
|---|---|
| `gateway` | REST API, validates JWT through `auth_service`, routes to gRPC services |
| `auth_service` | users, bcrypt passwords, JWT issue and validation |
| `warehouse_service` | products and stock |
| `order_service` | orders: reserves stock first, then saves the order |
| `sales_service` | time-limited discounts for products |
| `notification_service` | consumes events from Kafka and stores them |

Each service owns its data (separate SQLite database) and has its own Go module.

## How overselling is prevented

The classic flash sale bug is "check stock, then decrease it": two buyers read `quantity = 1` at the same time and both get the item.

Here the warehouse decreases stock with one conditional statement:

```sql
UPDATE products SET quantity = quantity - ? WHERE id = ? AND quantity >= ?
```

If no row was changed, there is not enough stock and the order is rejected with `409`. The order service reserves stock **before** writing the order; if writing fails, the stock is returned. Changing the quantity of an order moves only the difference, deleting an order returns the items.

`warehouse_service` has a test where 50 buyers compete for 10 items in parallel: exactly 10 succeed and the stock ends at zero.

## Getting started

```bash
cp .env.example .env    # set JWT_SECRET
docker compose up --build
```

```bash
curl -X POST localhost:3000/register -d '{"name":"Ali","email":"ali@example.com","password":"secret123","address":"Tashkent"}' -H 'Content-Type: application/json'
TOKEN=$(curl -s -X POST localhost:3000/login -d '{"email":"ali@example.com","password":"secret123"}' -H 'Content-Type: application/json' | jq -r .token)

curl -X POST localhost:3000/warehouse/product -H "Authorization: Bearer $TOKEN" -d '{"name":"Phone","desc":"128GB","price":300,"quantity":10}' -H 'Content-Type: application/json'
curl -X POST localhost:3000/orders -H "Authorization: Bearer $TOKEN" -d '{"product_id":1,"quantity":2}' -H 'Content-Type: application/json'
```

Every service also runs locally with `go run .`; addresses default to `localhost` and the ports above.

## Configuration

| Variable | Used by | Default |
|---|---|---|
| `PORT` | gateway | `3000` |
| `GRPC_ADDR` | gRPC services | `:50051` … `:50054` |
| `AUTH_ADDR`, `WAREHOUSE_ADDR`, `ORDER_ADDR`, `SALES_ADDR` | clients | `localhost:<port>` |
| `JWT_SECRET` | auth_service | required |
| `KAFKA_BROKERS` | auth, order, sales, notification | events are off if empty |
| `DB_PATH` | services with storage | `<service>.db` |

## API

| Method | Path | Auth |
|---|---|---|
| `POST` | `/register`, `/login`, `/validate` | — |
| `GET` | `/warehouse/products`, `/warehouse/product/:id`, `/warehouse/product/:id/inventory` | — |
| `POST`, `PATCH`, `DELETE` | `/warehouse/product…` | JWT |
| `GET` | `/sales`, `/sales/:id` | — |
| `POST`, `PATCH`, `DELETE` | `/sales…` | JWT |
| `GET`, `POST`, `PATCH`, `DELETE` | `/orders…` | JWT, the buyer is taken from the token |
| `GET`, `PATCH`, `DELETE` | `/users…` | JWT |

## Tests

```bash
cd warehouse_service && go test -race ./...
cd order_service && go test -race ./...
```

Order service tests use in-memory fakes of the auth and warehouse gRPC clients.

## What's next

- Roles: right now any authenticated user can manage products and sales
- Apply active sale prices when an order is created
- Transactional outbox for events, so an event is never lost when Kafka is down
- PostgreSQL instead of SQLite for production
