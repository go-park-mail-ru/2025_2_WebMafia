FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/auth ./cmd/auth/main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /bin/auth /auth

EXPOSE 8080 9001

CMD ["/auth", "-f=/config"]