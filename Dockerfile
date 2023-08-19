FROM golang:1.21-alpine3.18 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o telegram ./cmd/websocket_to_telegram/main.go

FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
COPY --from=builder ["/build/telegram", "/"]

ENTRYPOINT ["/telegram"]
