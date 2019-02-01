BINARY_NAME=imgartapi
GO_PACKAGES=$(shell ls -d */ | grep -v vendor)

default: quality

.PHONY: quality
quality:
	go test -v -race ./...
	go vet ./...
	golint -set_exit_status $(go list ./...)
	gocyclo -over 12 $(GO_PACKAGES)
	# megacheck ./...
	
.PHONY: clean
clean:
	go clean
	rm $(BINARY_NAME)

.PHONY: deps
deps:
	go mod download
	go get github.com/golang/lint/golint
	go get honnef.co/go/tools/cmd/megacheck
	go get github.com/fzipp/gocyclo

.PHONY: build
build: deps
	go build -o $(BINARY_NAME) -v ./cmd/imgartapi

.PHONY: docker
docker:
	docker-compose up

.PHONY: docker-debug
docker-debug:
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml up