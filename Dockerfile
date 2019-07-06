# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.11 as builder

LABEL maintainer="Mahamadou Sylla <mahamadou.sylla3@gmail.com>"

# Copy local code to the container image.
WORKDIR /go/src/pyri
COPY . .

RUN go get github.com/golang/dep/cmd/dep
RUN dep init
RUN dep ensure -vendor-only

EXPOSE 8080:80
