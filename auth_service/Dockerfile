# Используйте официальный образ Go в качестве базового
FROM golang:1.23.1 AS builder

# Установите рабочую директорию
WORKDIR /app

# Копируйте go.mod и go.sum
COPY go.mod go.sum ./

# Загрузите зависимости
RUN go mod download

# Копируйте все файлы в рабочую директорию
COPY . .

# Скомпилируйте исполняемый файл
RUN  GOOS=linux go build -o auth_service ./main.go

# Используйте минимальный образ для выполнения приложения
FROM alpine:latest

WORKDIR /root/

# Копируйте скомпилированный исполняемый файл из образа builder
COPY --from=builder /app/auth_service .

# Установите команду по умолчанию
CMD ["./auth_service"]

