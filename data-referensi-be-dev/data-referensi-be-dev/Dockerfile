# Step 1: Build the Go application
FROM golang:1.23.3-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

RUN apk add --no-cache git

ARG NETRC_FILE=.netrc

# Salin file netrc (sementara untuk build)
COPY ${NETRC_FILE} /root/.netrc

# Amankan file
RUN chmod 600 /root/.netrc


# Copy the Go Modules and Sum files
COPY go.mod go.sum ./

# Download all the dependencies. Dependencies will be cached if the go.mod and go.sum are not changed
RUN go mod tidy

# Copy the entire source code
COPY . .
COPY main.go ./

# Build the Go app
RUN go build -o main .

RUN rm -f /root/.netrc

# Step 2: Create a minimal runtime container
FROM alpine:latest  

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .
COPY .env .
COPY *.pem .
# Expose port 8080 (or whatever port your app runs on)
EXPOSE 3000

# Command to run the executable
CMD ["/bin/sh", "-c", "apk update && apk add bash"]
CMD ["./main"]
