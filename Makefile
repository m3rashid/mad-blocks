run:
	go run main.go

build:
	go build -o bin/madBlocks main.go

test:
	go test -v ./...
