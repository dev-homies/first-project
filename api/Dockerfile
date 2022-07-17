FROM golang:1.18-alpine

ENV GOPATH="/go"
ENV PATH="/go/bin:$PATH"

WORKDIR /api

RUN go install github.com/mitranim/gow@latest

EXPOSE 4000

CMD [ "gow", "run", "." ]
