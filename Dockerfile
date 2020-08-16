FROM golang:1.14.7

ENV HOME=/usr/src/app
WORKDIR $HOME

ADD main.go $HOME/main.go
ADD src $HOME/src

RUN go get -u github.com/oxequa/realize github.com/gin-gonic/gin github.com/kelseyhightower/envconfig

EXPOSE 3000

CMD ["realize", "start", "--no-config", "--run"]