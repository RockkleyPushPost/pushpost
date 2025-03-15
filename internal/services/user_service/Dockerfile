FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download



RUN go build -o user_service ./internal/services/user_service/cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/user_service .

COPY --from=builder /app/configs ./configs

EXPOSE 8080

CMD ["./user_service"]