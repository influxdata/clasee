FROM golang:1.13.1-alpine as base

WORKDIR /influxdata/clasee

ADD go.mod .
ADD go.sum .

RUN go mod download

ADD main.go .

RUN go build -o clasee .

FROM alpine:3

COPY --from=base /influxdata/clasee/clasee /usr/local/bin/clasee

ENTRYPOINT [ "clasee" ]
