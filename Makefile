include .env.*

build:
	go build -o ${BINARY} main.go

run:
	./${BINARY}

restart: build run