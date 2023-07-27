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

runHostServer:
	cd hostServer && go run *.go
	cd ..

buildHostServer:
	cd hostServer && go build -o ../bin/hostServer *.go
	cd ..

prodHostServer:
	make buildHostServer
	./bin/hostServer

runWalletServer:
	cd walletServer && go run *.go
	cd ..

buildWalletServer:
	cd walletServer && go build -o ../bin/walletServer *.go
	cd ..

prodWalletServer:
	make buildWalletServer
	./bin/walletServer
