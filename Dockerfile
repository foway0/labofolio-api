FROM golang:1.14.7

ENV HOME=/usr/src/app
WORKDIR $HOME

ADD main.go $HOME/main.go

RUN go get -u github.com/oxequa/realize github.com/gin-gonic/gin

EXPOSE 3000

CMD ["realize", "start", "--no-config", "--run"]