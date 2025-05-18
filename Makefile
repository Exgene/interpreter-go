.PHONY: all test build run clean

all: test run

test:
	@echo "Running tests..."
	@go test -v ./tests/...

build:
	@echo "Building..."
	@go build -o bin/main ./cmd/forge/main.go

run: build
	@echo "Running..."
	@./bin/main
