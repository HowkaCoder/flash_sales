FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o sales_service ./main.go


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/sales_service .

CMD ["./sales_service"]

