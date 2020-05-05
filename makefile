export GOPATH := $(shell pwd)

build:
	@echo "--> Building..."
	@mkdir -p bin/
	go build -v -o bin/vdf vdf/main.go
	@chmod 755 bin/*

clean:
	@echo "--> Cleaning..."
	@go clean
	@rm -f bin/*

.PHONY: build clean install fmt test coverage
