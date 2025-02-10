# Step 1
FROM golang:1.23.3-alpine AS builder
RUN apk update && apk add git gcc g++
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/console

# Step 2
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /opt
COPY --from=builder /app/console .

CMD ["/bin/sh"]
