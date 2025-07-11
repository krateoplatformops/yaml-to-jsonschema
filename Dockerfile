# Build stage
FROM golang:1.24.4-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app

# UPX stage
FROM alpine:3.21 AS shrinker

RUN apk add --no-cache upx

COPY --from=builder /app/app /app/app
RUN upx --best --lzma /app/app

# Final Image
FROM golang:1.24.4-alpine3.21

COPY --from=shrinker /app/app /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/app"]