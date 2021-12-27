FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd/*.go ./

RUN go build -o main

EXPOSE 4000

ENTRYPOINT [ "./main" ]