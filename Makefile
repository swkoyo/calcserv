BINARY_NAME=calcserv

build:
	go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}

dependencies:
	go mod install

dev:
	go run main.go
