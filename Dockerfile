FROM golang:1.20.1-alpine

WORKDIR /go/src

RUN apk update && apk add git

COPY . /go/src

RUN cd /go/src
CMD ["go","run","main.go"]

