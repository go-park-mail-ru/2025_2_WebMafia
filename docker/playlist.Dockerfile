FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/playlist ./cmd/playlist/main.go

# ------------------------------

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY certs/sber-ca.pem /usr/local/share/ca-certificates/sber-ca.pem
COPY certs/russian-root.crt /usr/local/share/ca-certificates/
COPY certs/russian-sub.crt /usr/local/share/ca-certificates/
RUN update-ca-certificates

WORKDIR /

COPY --from=builder /bin/playlist /playlist

EXPOSE 8082 9002

CMD ["/playlist", "-f=/config"]
