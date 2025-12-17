FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/playlist ./cmd/playlist/main.go

# ------------------------------

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=builder /bin/playlist /playlist

EXPOSE 8082 9002

CMD ["/playlist", "-f=/config"]
