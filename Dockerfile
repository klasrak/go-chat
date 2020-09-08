FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# New Stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /usr/src/app/main .
COPY --from=builder /usr/src/app/.env .

EXPOSE 8080

CMD [ "./main" ]
