# Build
FROM golang:1.19 AS builder

RUN apt-get update

WORKDIR /community

# install goplus
RUN cd .. && git clone -b v1.2.0-pre.1 https://github.com/goplus/gop.git && cd gop && ./all.bash

# build goplus-community
COPY . .

WORKDIR /community/cmd/gopcomm
RUN gop mod tidy
RUN gop build -o ./community 

# Run
FROM alpine:latest

RUN apk update && apk upgrade
RUN apk --no-cache add libc6-compat

WORKDIR /community/cmd/gopcomm

COPY --from=builder /community /community

CMD ["./community"]
