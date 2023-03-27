FROM golang:1.20.1-alpine

WORKDIR /go/src

RUN apk update && apk add git

COPY go.mod ./go/src
RUN go mod download

COPY . /go/src

RUN go build -o main .
CMD ["./main"]

