FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Соберите приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o order_service ./main.go

# Создайте финальный образ
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/order_service .

CMD ["./order_service"]

