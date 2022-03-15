FROM golang:1.18-rc-alpine

RUN apk add build-base
RUN mkdir -p /home/app

COPY . /home/app
WORKDIR /home/app

RUN go mod download
RUN cd cmd && go build main.go
WORKDIR /home/app/cmd

EXPOSE 8880
CMD ["./main"]
