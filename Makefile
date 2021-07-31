all: test
	go build -o bin/gwennen cmd/gwennen/main.go
test:
	go test internal/...
