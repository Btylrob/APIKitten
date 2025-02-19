# Go parameters
BINARY_NAME=myapp
GO_FILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Build the binary
build:
	@echo "Building the project..."
	go build -o $(BINARY_NAME) main.go cli.go

# Run the application
run:
	@echo "Running the application..."
	./$(BINARY_NAME)

# Run tests
.PHONY: test

test:
	@echo "Running tests..."
	go test ./...

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	@echo "Linting code..."
	golangci-lint run

# Clean up binaries
clean:
	@echo "Cleaning up..."
	@if [ "$(OS)" = "Windows_NT" ]; then del $(BINARY_NAME).exe; else rm -f $(BINARY_NAME); fi

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy

# Generate code (if applicable)
gen:
	@echo "Generating code..."
	go generate ./...

# Default target
.DEFAULT_GOAL := build