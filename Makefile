lint: 
	golangci-lint run -v

test:
	go get -d
	go test -v ./...
run:
	go run main.go

build:
	go build main.go