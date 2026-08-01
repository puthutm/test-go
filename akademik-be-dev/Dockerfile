# Step 1: Build the Go application
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# Private repo config
ENV GOPRIVATE=gitlab.unsia.ac.id
RUN git config --global url."https://gitlab.unsia.ac.id/".insteadOf "https://gitlab.unsia.ac.id/"

ARG NETRC_FILE=.netrc
COPY ${NETRC_FILE} /root/.netrc
RUN chmod 600 /root/.netrc

# Copy go mod dulu (biar caching jalan)
COPY go.mod go.sum ./

# Download dependency dulu
RUN go mod download

# Baru copy source
COPY . .

# Build binary
RUN go build -o main ./cmd/akademic

# Cleanup credential
RUN rm -f /root/.netrc

# Step 2: Runtime
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY .env .
COPY *.pem .

RUN apk add --no-cache bash

EXPOSE 9000

CMD ["./main"]