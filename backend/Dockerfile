FROM golang:1.24.0-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Install essential build dependencies
RUN apk add --no-cache gcc musl-dev git

# Copy only dependency files first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -ldflags="-s -w" -o main cmd/main.go cmd/wire_gen.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Install runtime dependencies
RUN apk --no-cache add ca-certificates

# Copy the binary
COPY --from=builder /app/main .

# Configure networking
ENV GIN_HOST=0.0.0.0

EXPOSE 8080

CMD ["./main"]