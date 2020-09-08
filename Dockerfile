FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080
