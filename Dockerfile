FROM golang:1.18-rc-alpine

RUN apk add build-base
RUN mkdir -p /home/app

COPY . /home/app
WORKDIR /home/app

RUN go mod download \
	go build ./main/main.go \
	npm run build
WORKDIR /home/app/cmd

EXPOSE 8880
CMD ["./cmd/main"]
