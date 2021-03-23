FROM golang:1.15.6-alpine
RUN apk update && apk add git
RUN mkdir /go/src/work
WORKDIR /go/src/work
ADD . /go/src/work