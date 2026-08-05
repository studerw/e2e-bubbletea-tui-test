.PHONY: all build run test lint clean e2e e2e-go e2e-shell e2e-install help install-tools

BINARY_NAME=tui-e2e-demo
BINARY_DIR=bin
CMD_DIR=cmd/app

all: lint test build

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BINARY_DIR)
	go build -o $(BINARY_DIR)/$(BINARY_NAME) ./$(CMD_DIR)
	@echo "Build complete: $(BINARY_DIR)/$(BINARY_NAME)"

run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_DIR)/$(BINARY_NAME)

test:
	@echo "Running unit tests..."
	go test -v -race ./...

lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Run 'make install-tools' to install."; \
		go vet ./...; \
	fi

e2e: e2e-go e2e-shell

e2e-go: build
	@echo "Running Go E2E tests (teatest)..."
	go test -v -tags=e2e ./tests/e2e/...

e2e-shell: build
	@echo "Running Shell-Use E2E tests..."
	cd tests/e2e-shell-use && npm test

e2e-install:
	@echo "Installing Shell-Use E2E test dependencies..."
	cd tests/e2e-shell-use && npm install

clean:
	@echo "Cleaning..."
	rm -rf $(BINARY_DIR)
	go clean -cache -testcache

install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Tools installed successfully"

deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

fmt:
	@echo "Formatting code..."
	go fmt ./...

coverage:
	@echo "Generating coverage report..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

help:
	@echo "Available targets:"
	@echo "  all           - Run lint, test, and build (default)"
	@echo "  build         - Build the application"
	@echo "  run           - Build and run the application"
	@echo "  test          - Run unit tests"
	@echo "  lint          - Run linter"
	@echo "  e2e           - Run all end-to-end tests (Go + Shell-Use)"
	@echo "  e2e-go        - Run Go-level E2E tests (teatest)"
	@echo "  e2e-shell     - Run Shell-Use E2E tests (Node.js)"
	@echo "  e2e-install   - Install Shell-Use test dependencies"
	@echo "  clean         - Remove build artifacts"
	@echo "  install-tools - Install development tools"
	@echo "  deps          - Download and tidy dependencies"
	@echo "  fmt           - Format code"
	@echo "  coverage      - Generate test coverage report"
	@echo "  help          - Show this help message"
