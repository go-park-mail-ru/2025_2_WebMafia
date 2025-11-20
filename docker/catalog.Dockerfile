FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/catalog ./cmd/catalog/main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /bin/catalog /catalog

EXPOSE 8081 9000

CMD ["/catalog", "-f=/config"]