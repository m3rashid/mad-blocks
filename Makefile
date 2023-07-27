run:
	clear
	go run main.go

build:
	go build -o bin/madBlocks main.go

prod:
	make build
	./bin/madBlocks

test:
	go test -v ./...
