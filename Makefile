all: test build

build:
	go build -o bin/gwennen cmd/gwennen/main.go
test:
	go test ./...
