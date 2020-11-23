FROM golang:1.15.5-alpine3.12

RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD . /go/src/app

RUN apk add --no-cache \
        alpine-sdk \
        git \