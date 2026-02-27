# ============================================
# Stage 1: Build
# ============================================
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/event-backend .

# ============================================
# Stage 2: Runtime
# ============================================
FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy binary and migration files
COPY --from=builder /app/event-backend .
COPY --from=builder /app/migration/files ./migration/files
COPY --from=builder /app/seeder/files ./seeder/files

EXPOSE 8080

ENTRYPOINT ["./event-backend"]
