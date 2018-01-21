BINARY_NAME=gorpoapi

default: test build

test: 
	go test -v ./...

lint:

build: 
	go build -o $(BINARY_NAME) -v ./cmd/gorpoapi

docker:
	docker-compose up

docker-debug:
	docker-compose -f docker-compose.debug.yml up