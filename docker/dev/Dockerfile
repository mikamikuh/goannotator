FROM golang:1.8.3

RUN go get github.com/derekparker/delve/cmd/dlv

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

CMD ["dlv", "debug", "--headless", "--listen=:2345", "--log"]