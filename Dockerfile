FROM golang:1.15-alpine AS builder

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -mod=vendor -o singatureServer .

ENTRYPOINT [ "./singatureServer"]

EXPOSE 3000