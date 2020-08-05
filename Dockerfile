FROM golang:1.14.6

ENV HOME=/usr/src/app
WORKDIR $HOME

ADD main.go $HOME/main.go

RUN go get -u github.com/gin-gonic/gin

EXPOSE 3000