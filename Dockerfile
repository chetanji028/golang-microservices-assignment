# Stage 1: Build
FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o service cmd/main.go

# Stage 2: Run
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/service .
COPY ports.json .
CMD ["./service"]
