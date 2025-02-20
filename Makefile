# Variables
BINARY_NAME = mygrep
SRC = ./...

# Build the binary
build:
	@go build -o $(BINARY_NAME) cmd/main.go

# Run the program (build first, suppress command logs)
run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BINARY_NAME) <body> test

# Run tests
test:
	@go test -v $(SRC)

# Format the code
fmt:
	@go fmt $(SRC)

# Lint the code (requires golangci-lint)
lint:
	@golangci-lint run

# Clean build files
clean:
	@rm -f $(BINARY_NAME)

# Install dependencies
deps:
	@go mod tidy

# Rebuild everything
rebuild: clean build

# Default target
.PHONY: all
all: build
