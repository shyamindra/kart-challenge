# backend/Dockerfile

# --- Build Stage ---
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Run the data preparation script to download coupon files
# We need bash for the script, which is available in golang:1.21-alpine
RUN chmod +x scripts/prepare_data.sh && ./scripts/prepare_data.sh

# Build the Go application
RUN go build -o /app/api ./cmd/api

# --- Final Stage ---
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/api /usr/local/bin/api

# Copy the prepared data files
COPY --from=builder /app/data /app/data

# Expose the port the Go server listens on (assuming 8080)
EXPOSE 8080

# Command to run the executable
CMD ["api"]