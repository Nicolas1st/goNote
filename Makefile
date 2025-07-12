build: backend frontend

backend:
	go build -o main ./cmd/main.go

frontend:
	cd ./web
	npm run build

run: build
	./main
