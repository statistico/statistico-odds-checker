# Step 1
FROM golang:1.13.9-alpine AS builder
RUN apk update && apk add git gcc g++
ENV GO111MODULE=on
WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/console
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/csv

# Step 2
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /opt
COPY ./bin ./bin
COPY ./.docker/crontab /etc/crontabs/root
COPY ./.docker/wait-for-it.sh .
COPY --from=builder /app/console .
COPY --from=builder /app/csv .

CMD ["/bin/sh"]