# Build
FROM golang:1.19 AS builder

RUN apt-get update

WORKDIR /community

# install goplus
RUN cd .. && git clone -b v1.2.5 https://github.com/goplus/gop.git && cd gop && ./all.bash

# build goplus-community
COPY . .

WORKDIR /community/cmd/gopcomm
RUN gop mod tidy
RUN gop build -o ./community 

# Run
FROM ubuntu:22.04

RUN apt-get update

WORKDIR /community/cmd/gopcomm

RUN ln -s /config/.env .env

COPY --from=builder /community /community

CMD ["./community"]
