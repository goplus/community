FROM golang:1.19

RUN apt-get update

WORKDIR /community

# install goplus
RUN cd .. && git clone -b v1.2.0-pre.1 https://github.com/goplus/gop.git && cd gop && ./all.bash

# run goplus-community
COPY . .

# download account repository
RUN cd .. && git clone -b feat/init-account https://github.com/IRONICBo/account.git
CMD bash scripts/start.sh
