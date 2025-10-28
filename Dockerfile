# Dockerfile para API Go
FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o api_inmemory main.go

EXPOSE 8091

ENV API_PORT=8091

CMD ["./api_inmemory"]
