FROM golang:1.18-rc-alpine

RUN mkdir -p /home/app
COPY . /home/app
WORKDIR /home/app
RUN apk add build-base
RUN go mod download
RUN go build main.go
EXPOSE 8880
CMD ["./main"]
