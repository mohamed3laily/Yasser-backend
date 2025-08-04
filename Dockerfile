FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go

RUN go build -o seeder ./cmd/seeder/main.go

FROM debian:bullseye-slim

WORKDIR /root/

COPY --from=builder /app/server .
COPY --from=builder /app/seeder .

EXPOSE 3000

CMD ["./server"]
