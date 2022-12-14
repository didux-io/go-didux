# Build Geth in a stock Go builder container
FROM golang:1.19-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go-didux
RUN cd /go-didux && make didux

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-didux/build/bin/go-didux /usr/local/bin/

EXPOSE 21000 22000 23000 21000/udp
ENTRYPOINT ["go-didux", "--sport"]
