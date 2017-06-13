FROM golang:1.8.3

RUN curl https://glide.sh/get | sh
RUN go get github.com/derekparker/delve/cmd/dlv
COPY glide.yaml .
COPY glide.lock .
RUN glide install

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

CMD ["dlv", "debug", "--headless", "--listen=:2345", "--log"]