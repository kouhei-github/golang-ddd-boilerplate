FROM golang:1.20 as build

WORKDIR /app/public

RUN go install github.com/cosmtrek/air@v1.29.0

COPY ./src/public/go.mod /app/public/
COPY ./src/public//go.sum /app/public/
RUN go mod download
COPY ./src /app
RUN CGO_ENABLED=0 go build -o main main.go

