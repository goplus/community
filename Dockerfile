FROM golang:1.21.6

RUN apt-get update

WORKDIR /community

# install goplus
RUN cd .. && git clone https://github.com/goplus/gop.git && cd gop && ./all.bash

# install npm
RUN apt-get install npm -y

# run goplus-community
COPY . .
CMD cd cmd/gopcomm && gop run .
