# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
SERVER_BINARY_NAME=gorpo-api
CLI_BINARY_NAME=gorpo-cli

all: test build

build-server: 
	$(GOBUILD) -o $(SERVER_BINARY_NAME) -v ./cmd/httpserver

build-cli: 
	$(GOBUILD) -o $(CLI_BINARY_NAME) -v ./cmd/cli

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

deps:
	$(GOGET) github.com/talento90/gorpo