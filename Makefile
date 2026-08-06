.PHONY: all build run test lint clean e2e e2e-go e2e-shell e2e-install help install-tools \
         build-docker-image docker-build docker-run docker-test docker-lint docker-e2e docker-e2e-go docker-e2e-shell docker-shell

BINARY_NAME=tui-e2e-demo
BINARY_DIR=bin
CMD_DIR=cmd/app

# ─── Docker settings ───────────────────────────────────────────────────────────
# Override DOCKER_IMAGE to push/pull from a registry, e.g.:
#   make build-docker-image DOCKER_IMAGE=registry.example.com/myteam/tui-e2e-demo:latest
DOCKER_IMAGE ?= tui-e2e-demo:dev

# Common flags for all non-interactive docker run invocations:
#   --rm            remove the container when it exits
#   -v $(PWD):/workspace  bind-mount the project source into /workspace
#   -w /workspace   set the working directory inside the container
# No -t or -i flags here → fully compatible with CI pipelines (no TTY needed).
DOCKER_RUN_ARGS = --rm -v $(PWD):/workspace -w /workspace

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
	@echo "  all              - Run lint, test, and build (default)"
	@echo "  build            - Build the application"
	@echo "  run              - Build and run the application"
	@echo "  test             - Run unit tests"
	@echo "  lint             - Run linter"
	@echo "  e2e              - Run all end-to-end tests (Go + Shell-Use)"
	@echo "  e2e-go           - Run Go-level E2E tests (teatest)"
	@echo "  e2e-shell        - Run Shell-Use E2E tests (Node.js)"
	@echo "  e2e-install      - Install Shell-Use test dependencies"
	@echo "  clean            - Remove build artifacts"
	@echo "  install-tools    - Install development tools"
	@echo "  deps             - Download and tidy dependencies"
	@echo "  fmt              - Format code"
	@echo "  coverage         - Generate test coverage report"
	@echo ""
	@echo "Docker targets (only Docker required — no Go/Node needed locally):"
	@echo "  build-docker-image - Build the Docker dev image ($(DOCKER_IMAGE))"
	@echo "  docker-build       - Compile the app binary inside Docker → writes to bin/"
	@echo "  docker-run         - Run the TUI app interactively inside Docker (needs a TTY)"
	@echo "  docker-test        - Run unit tests inside Docker"
	@echo "  docker-lint        - Run golangci-lint inside Docker"
	@echo "  docker-e2e         - Run all E2E tests inside Docker (CI-compatible)"
	@echo "  docker-e2e-go      - Run Go/teatest E2E tests inside Docker"
	@echo "  docker-e2e-shell   - Run Shell-Use/Vitest E2E tests inside Docker"
	@echo "  docker-shell       - Open an interactive bash shell inside the container"
	@echo "  help               - Show this help message"

# =============================================================================
# Docker targets
# =============================================================================

# Build (or rebuild) the Docker dev image.
# Re-run whenever Dockerfile, go.mod, go.sum, or package-lock.json change.
build-docker-image:
	@echo "Building Docker image $(DOCKER_IMAGE)..."
	docker build -t $(DOCKER_IMAGE) .
	@echo "Docker image built: $(DOCKER_IMAGE)"

# Compile the Go binary inside Docker and write it to bin/ on the host.
# Equivalent to 'make build' but requires no local Go installation.
# The CWD is bind-mounted into the container, so bin/$(BINARY_NAME) lands
# in your working directory exactly as if you had run 'make build' locally.
docker-build:
	@echo "Building $(BINARY_NAME) inside Docker..."
	@mkdir -p $(BINARY_DIR)
	docker run $(DOCKER_RUN_ARGS) $(DOCKER_IMAGE) \
		go build -o $(BINARY_DIR)/$(BINARY_NAME) ./$(CMD_DIR)
	@echo "Build complete: $(BINARY_DIR)/$(BINARY_NAME)"

# Run the compiled TUI app interactively.
# Requires a real TTY (-it), so this target is intended for local development
# only — not for CI pipelines (use docker-e2e-go / docker-e2e-shell there).
docker-run: docker-build
	@echo "Running TUI inside Docker (interactive — requires a TTY)..."
	docker run --rm -it -v $(PWD):/workspace -w /workspace $(DOCKER_IMAGE) \
		./$(BINARY_DIR)/$(BINARY_NAME)

# Run Go unit tests. Fully headless; no TTY required.
docker-test:
	@echo "Running unit tests inside Docker..."
	docker run -e CGO_ENABLED=1 $(DOCKER_RUN_ARGS) $(DOCKER_IMAGE) go test -v -race ./...

# Run golangci-lint. Fully headless; no TTY required.
docker-lint:
	@echo "Running linter inside Docker..."
	docker run $(DOCKER_RUN_ARGS) $(DOCKER_IMAGE) golangci-lint run ./...

# Run Go-level (teatest) E2E tests.
# teatest drives BubbleTea programmatically via an in-process pseudo-terminal;
# it does NOT require a host TTY and is fully CI-compatible.
docker-e2e-go: docker-build
	@echo "Running Go E2E tests (teatest) inside Docker..."
	docker run $(DOCKER_RUN_ARGS) $(DOCKER_IMAGE) go test -v -tags=e2e ./tests/e2e/...

# Run Shell-Use (Node/Vitest) E2E tests.
# shell-use spawns its own pseudo-terminal internally (/dev/pts is available
# in Docker containers by default), so no host TTY is needed — CI-compatible.
# node_modules were baked into the image at docker-build time; no npm install
# step is required here.
docker-e2e-shell: docker-build
	@echo "Running Shell-Use E2E tests inside Docker..."
	docker run $(DOCKER_RUN_ARGS) $(DOCKER_IMAGE) \
		sh -c 'cd tests/e2e-shell-use && npm test'

# Run the full E2E suite (Go + Shell-Use) inside Docker. CI-compatible.
docker-e2e: docker-e2e-go docker-e2e-shell

# Drop into an interactive bash shell inside the container for debugging.
docker-shell:
	@echo "Opening bash shell inside Docker container..."
	docker run --rm -it -v $(PWD):/workspace -w /workspace $(DOCKER_IMAGE) bash
