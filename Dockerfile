# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o data-digger .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy executable
COPY --from=builder /app/data-digger /usr/local/bin/data-digger

# Copy mock database files
COPY --from=builder /app/scraping/*.json /app/scraping/

ENTRYPOINT ["data-digger"]