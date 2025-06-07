# 👷 Stage 1: Build your Go app
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the app, output binary as /bin/Bank inside builder stage
RUN go build -o /bin/Bank .

# 🧼 Stage 2: Use small image to run the app
FROM alpine:latest

WORKDIR /app

# Copy binary from builder's /bin/Bank to final image's /bin/Bank
COPY --from=builder /bin/Bank /bin/Bank

EXPOSE 8001

CMD ["/bin/Bank"]
