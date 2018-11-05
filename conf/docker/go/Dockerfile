FROM golang:latest

LABEL maintainer "yoshikawataiki@gmail.com"

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin

# build directories
RUN mkdir /go/src/github.com/Pluslab/fieldsensing-api
ADD . /go/src/github.com/Pluslab/fieldsensing-api
WORKDIR /go/src/github.com/Pluslab/fieldsensing-api

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure

# Build my app
RUN go build -o /go/src/github.com/Pluslab/fieldsensing-api/main .
CMD ["/go/src/github.com/Pluslab/fieldsensing-api/main"]