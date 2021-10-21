
default: build

build:
	go build -o sleuth ./cmd

test:
	go test -v -race -timeout 30m ./...

