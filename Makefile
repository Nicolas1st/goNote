build: backend frontend css

backend:
	go build -o main ./cmd/main.go

frontend: ./web/js/*.js
	cd ./web
	npm run build

run: build
	./main

css: ./web/css/*
	cat ./web/css/* > ./static/style.css
