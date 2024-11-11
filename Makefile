# Local
config:
	go version

install:
	go mod tidy
	
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0 run -v

start:
	go run main.go

# Docker
net:
	docker network create go_todo_network

up:
	docker-compose up -d --build

down:
	docker-compose down
