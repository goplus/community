# Build
FROM golang:1.19 AS builder

RUN apt-get update

WORKDIR /community

# install goplus
RUN cd .. && git clone -b v1.2.0-pre.1 https://github.com/goplus/gop.git && cd gop && ./all.bash

# run goplus-community
COPY . .

# download account repository
# RUN cd .. && git clone https://github.com/IRONICBo/account.git && git checkout feat/init-account
# CMD bash scripts/start.sh

WORKDIR /community/cmd/gopcomm
RUN go mod tidy
RUN go build -o /community/community

# Run
FROM alpine:lastest

WORKDIR /community

COPY --from=builder /community /community

CMD ["./community"]