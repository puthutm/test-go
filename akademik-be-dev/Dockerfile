# Step 1: Build the Go application
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Ensure .env and public_key.pem exist so stage 2 COPY never fails
RUN touch .env public_key.pem

RUN go build -o main ./cmd/akademic

# Step 2: Runtime
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/public_key.pem .

RUN apk add --no-cache bash

EXPOSE 8080

CMD ["./main"]