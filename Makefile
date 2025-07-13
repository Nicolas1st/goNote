build: backend

backend:
	go build -o main ./cmd/main.go

run: build
	./main

