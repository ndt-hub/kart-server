FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o kart-server .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/kart-server .

COPY --from=builder /app/public ./public

EXPOSE 8080

CMD ["./kart-server"]
