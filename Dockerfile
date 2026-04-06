FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/ordersystem/

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server ./
COPY --from=builder /app/cmd/ordersystem/.env ./
EXPOSE 8000 8080 5672
CMD ["sh", "-c","sleep 60 && exec ./server"]