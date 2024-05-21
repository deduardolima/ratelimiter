FROM golang:1.20-buster as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

FROM debian:buster-slim as app

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env ./

EXPOSE 8080

CMD ["./main"]

FROM golang:1.20-buster as test

WORKDIR /app

COPY --from=builder /app ./
COPY --from=builder /app/.env ./

CMD ["go", "test", "./test"]
