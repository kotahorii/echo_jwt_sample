FROM golang:1.17.8-buster

RUN apt-get update && \
  apt-get install git && \
  apt install sudo && \
  apt-get install lsof

WORKDIR /app

COPY . /app

RUN go mod tidy
RUN go get github.com/cosmtrek/air
CMD ["air"]