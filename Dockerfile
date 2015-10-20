FROM golang:latest

ADD . /go/src/HelloWorld

WORKDIR /go/src/HelloWorld/
RUN go build
CMD ./HelloWorld

EXPOSE 3000
