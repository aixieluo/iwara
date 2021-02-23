FROM golang:latest

MAINTAINER xxy

ENV GOPROXY=https://goproxy.io,direct

WORKDIR $GOPATH/src/iwara

COPY . $GOPATH/src/iwara

RUN go get -u

EXPOSE 3939

ENTRYPOINT ["go", "run", "main.go"]

