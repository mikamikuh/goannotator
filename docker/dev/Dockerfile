FROM golang:1.8.3

RUN curl https://glide.sh/get | sh
RUN go get github.com/derekparker/delve/cmd/dlv

ENV APP_PATH /go/src/github.com/mikamikuh/goannotator
RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH

COPY . .
RUN glide up

CMD ["bash", "docker/dev/run.sh"]
