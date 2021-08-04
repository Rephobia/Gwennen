all: test build

build:
	go build -o bin/gwennen cmd/gwennen/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/gwennen.exe cmd/gwennen/main.go
test:
	go test ./...
