FROM golang:1.18-alpine

ENV GOPATH="/go"
ENV PATH="/go/bin:$PATH"

WORKDIR /api

RUN apk add build-base

RUN go install github.com/cosmtrek/air@v1.40
RUN go install github.com/go-delve/delve/cmd/dlv@v1.9
RUN go install github.com/swaggo/swag/cmd/swag@v1.8

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

EXPOSE 2345
EXPOSE 4000

CMD [ "air" ]
