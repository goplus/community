FROM golang:1.19

RUN apt-get update

WORKDIR /community

# install goplus
RUN cd .. && git clone https://github.com/goplus/gop.git && cd gop && ./all.bash

# run goplus-community
COPY . .

# download account repository
RUN cd .. && git clone https://github.com/IRONICBo/account.git && git checkout feat/init-account
CMD bash scripts/start.sh
