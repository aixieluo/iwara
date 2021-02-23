FROM golang:latest as go6

MAINTAINER xxy

ENV GOPROXY=https://goproxy.io,direct

WORKDIR $GOPATH/src/iwara

COPY . $GOPATH/src/iwara

VOLUME ["$GOPATH/src/iwara", "/go/src/iwara"]

#RUN go get -u
RUN go build main.go

EXPOSE 3939


ENTRYPOINT ["go", "run", "main.go"]
#ENTRYPOINT ["/go/src/iwara/main"]

