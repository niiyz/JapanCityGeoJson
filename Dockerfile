FROM osgeo/gdal:alpine-small-latest

RUN apk add --update --no-cache alpine-sdk git go

RUN mkdir -p /go/src/app

WORKDIR /go/src/app

COPY ./ /go/src/app

RUN go version
