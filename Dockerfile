FROM golang:1.22-bookworm AS builder

# Create and change to the app directory
WORKDIR /app

# Copy go.mog and go.sum
COPY go.* ./
RUN go mod download

# Copy code to the container image
COPY . ./

# Build the binary
RUN go build -o api cmd/api/main.go

FROM debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage
COPY --from=builder /app/api /app/api

EXPOSE 8080

# Run the api on container startup
CMD ["/app/api"]