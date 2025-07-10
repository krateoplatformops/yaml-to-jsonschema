FROM golang:1.24.4-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app

FROM golang:1.24.4-alpine3.21

COPY --from=builder /app/app /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/app"]