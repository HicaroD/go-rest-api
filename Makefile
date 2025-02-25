APP_NAME := server
BUILD_DIR := build
SRC := ./cmd/$(APP_NAME)
GO := go

# Default target
.PHONY: all
all: build

# Run the application
.PHONY: air
air:
	air

# Run the application
.PHONY: run
run:
	$(GO) run $(SRC)/*.go

# Build the application
.PHONY: build
build:
	$(GO) build -o $(BUILD_DIR)/$(APP_NAME) $(SRC)

# Clean build artifacts
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# Run tests
.PHONY: test
test:
	$(GO) test ./... -v

# Run tests with coverage
.PHONY: coverage
coverage:
	$(GO) test ./... -coverprofile=coverage.out
	$(GO) tool cover -html=coverage.out

# Format the code
.PHONY: fmt
fmt:
	$(GO) fmt ./...
	prettier --write .

# Lint the code
.PHONY: lint
lint:
	golangci-lint run ./...

# Run a static analysis
.PHONY: vet
vet:
	$(GO) vet ./...

# Install dependencies
.PHONY: deps
deps:
	$(GO) mod tidy
	$(GO) mod vendor

# Generate mocks or other code (e.g., using go generate)
.PHONY: generate
generate:
	$(GO) generate ./...

.PHONY: env
env:
	direnv allow .

# Help menu
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  air         Run the application with Air for hot reloading"
	@echo "  run         Run the application"
	@echo "  build       Build the application binary"
	@echo "  clean       Remove build artifacts"
	@echo "  test        Run unit tests"
	@echo "  coverage    Run tests and generate coverage report"
	@echo "  fmt         Format the code"
	@echo "  lint        Lint the code (requires golangci-lint)"
	@echo "  vet         Run static analysis"
	@echo "  deps        Install dependencies"
	@echo "  generate    Run code generation (if applicable)"
	@echo "  help        Show this help menu"


