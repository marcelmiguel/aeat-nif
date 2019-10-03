FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/aeat-nif/
COPY . .

# Using go get.
RUN go get -d -v

# RUN go mod download
# RUN go mod verify

# Tests
RUN go test

