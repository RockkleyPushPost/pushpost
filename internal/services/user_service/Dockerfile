FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o user_service ./internal/services/user_service/cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/user_service .
COPY --from=builder /app/internal/services/user_service/config ./config/

EXPOSE ${PORT}

CMD ["./user_service"]