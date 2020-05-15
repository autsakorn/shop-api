# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.13-alpine

RUN apk add git

ENV CGO_ENABLED=0

# Set the Current Working Directory inside the container
WORKDIR /go/src/shop-api

# Copy go mod and sum files
COPY go.mod go.sum /go/src/shop-api/

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

RUN go get github.com/cweill/gotests/...

RUN go get github.com/golang/mock/mockgen

RUN GO111MODULE=off go get -u github.com/beego/bee

# Copy the source from the current directory to the Working Directory inside the container
COPY . ./shop-api