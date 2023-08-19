FROM golang:1.21-alpine3.18 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64
RUN go build -ldflags="-s -w" -o telegram .

FROM scratch
COPY --from=builder ["/build/telegram", "/"]
LABEL ARCH=arm64

ENTRYPOINT ["/telegram"]
