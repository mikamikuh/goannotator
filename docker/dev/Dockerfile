FROM golang:1.8.3

RUN curl https://glide.sh/get | sh
RUN go get github.com/derekparker/delve/cmd/dlv

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . .
RUN glide up

CMD ["bash", "docker/dev/run.sh"]
