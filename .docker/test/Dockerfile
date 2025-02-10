FROM golang:1.23.3-alpine

RUN apk update && apk add git gcc g++

WORKDIR /opt

RUN go install gotest.tools/gotestsum@latest

CMD ["/bin/sh"]
