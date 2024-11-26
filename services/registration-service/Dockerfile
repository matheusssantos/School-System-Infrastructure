FROM golang:latest

WORKDIR /go/app

RUN apt-get update

RUN go mod tidy

CMD ["tail", "-f", "/dev/null"]