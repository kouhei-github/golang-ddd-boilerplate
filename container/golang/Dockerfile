FROM golang:1.20 as build

WORKDIR /app/public

#  Golang migrateのインストール
RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz \
    && tar -xzf migrate.linux-amd64.tar.gz \
    && mv migrate /usr/local/bin/migrate \
    && rm -rf migrate.linux-amd64.tar.gz LICENSE README.md


RUN go install github.com/cosmtrek/air@v1.29.0

COPY ./src/public/go.mod /app/public/
COPY ./src/public//go.sum /app/public/
RUN go mod download
COPY ./src /app
# RUN CGO_ENABLED=0 go build -o main main.go

