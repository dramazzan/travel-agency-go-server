FROM golang:1.23-alpine

LABEL authors="ramazandautbek"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main ./cmd

CMD ["./main"]
