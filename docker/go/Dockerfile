FROM golang:latest

# Set go bin which doesn't appear to be set already.
ENV GOPATH /go
ENV GOBIN /go/bin

# build directories
ADD . $GOPATH/src/github.com/Pluslab/fieldsensing
WORKDIR /go/src/github.com/Pluslab/fieldsensing/server

# Go dep
RUN go get -u github.com/golang/dep/... && dep ensure -vendor-only