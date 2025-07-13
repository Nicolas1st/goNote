build: backend css

backend:
	go build -o main ./cmd/main.go

run: build
	./main

css: ./web/css/*
	cat ./web/css/* > ./static/style.css
