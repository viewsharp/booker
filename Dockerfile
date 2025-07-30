# Базовый образ для сборки
FROM golang:1.23-alpine AS builder

# Создание рабочей директории
WORKDIR /app

# Копирование файлов зависимостей
COPY go.mod go.sum ./

# Загрузка зависимостей
RUN go mod download

# Копирование исходного кода
COPY main.go next_date.go seater.go ./

# Сборка приложения
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o unspot-bot

# Финальный образ
FROM alpine:3.18

# Копирование бинарника из builder
COPY --from=builder /app/unspot-bot /usr/local/bin/unspot-bot

# Установка рабочей директории
WORKDIR /app

# Указание точки входа
ENTRYPOINT ["unspot-bot"]
