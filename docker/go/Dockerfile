FROM golang:latest

# Set go bin which doesn't appear to be set already.
ENV GOPATH=/go \
        GOBIN=/go/bin

# build directories
ADD . $GOPATH/src/github.com/Pluslab/gaia
WORKDIR /go/src/github.com/Pluslab/gaia/server

RUN go get bitbucket.org/liamstask/goose/cmd/goose  && \
        go get -u github.com/golang/dep/cmd/dep
