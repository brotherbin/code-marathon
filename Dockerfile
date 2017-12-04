FROM golang:1.8

MAINTAINER songdabin@sf-express.com

ADD . /go/src/sfai/sf-code-marathon

RUN go install sfai/sf-code-marathon

WORKDIR /go/src/sfai/sf-code-marathon