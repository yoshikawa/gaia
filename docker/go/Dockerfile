FROM golang:latest

# Set go bin which doesn't appear to be set already.
ENV GOPATH=/go \
        GOBIN=/go/bin \
        GO111MODULE=on

# build directories
ADD ./server $GOPATH/src/github.com/Pluslab/gaia/server
WORKDIR /go/src/github.com/Pluslab/gaia/server

RUN go get bitbucket.org/liamstask/goose/cmd/goose