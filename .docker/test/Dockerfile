FROM golang:1.19-rc-alpine

RUN apk update && apk add git gcc g++

WORKDIR /opt

RUN go install gotest.tools/gotestsum@latest

CMD ["/bin/sh"]
