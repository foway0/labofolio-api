FROM golang:1.15.6

ENV HOME=/usr/src/app
WORKDIR $HOME

ADD main.go $HOME/main.go
ADD go.mod $HOME/go.mod
ADD go.sum $HOME/go.sum
ADD grpc_specs $HOME/grpc_spec

RUN go mod download

EXPOSE 4000