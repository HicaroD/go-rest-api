APP_NAME := server
BUILD_DIR := build
SRC := ./cmd/$(APP_NAME)
GO := go

# Default target
.PHONY: all
all: build

.PHONY: reqs
reqs:
	@echo "This command requires the Golang toolchain and npm to be installed and declared on your PATH"
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/echo-swagger
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/air-verse/air@latest
	npm install -g prettier

# Run the application with Hot reloading
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

# Update all environment variables
.PHONY: env
env:
	direnv allow .

# Start containers
.PHONY: db-up
db-up:
	docker compose up

# Stop containers
.PHONY: db-down
db-down:
	docker compose down

.PHONY: db-purge
db-purge:
	docker compose down -v

# Add a new migration
.PHONY: migrate-new
migrate-new:
	@read -p "Enter migration name: " name; \
	goose create $$name sql

# Apply migrations
.PHONY: migrate-up
migrate-up:
	goose up

# Unapply migrations
.PHONY: migrate-down
migrate-down:
	goose down

# Redo migrations
.PHONY: migrate-redo
migrate-redo:
	goose redo

# Retrieve migration status
.PHONY: migrate-status
migrate-status:
	goose status
	
# Generate Swagger docs
.PHONY: swag
swag:
	./scripts/setup_swagger.sh

# Help menu
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  reqs              Install all Golang requirements"
	@echo "  air               Run the application with Air for hot reloading"
	@echo "  run               Run the application"
	@echo "  build             Build the application binary"
	@echo "  clean             Remove build artifacts"
	@echo "  test              Run unit tests"
	@echo "  coverage          Run tests and generate coverage report"
	@echo "  fmt               Format the code"
	@echo "  lint              Lint the code (requires golangci-lint)"
	@echo "  vet               Run static analysis"
	@echo "  deps              Install dependencies"
	@echo "  generate          Run code generation (if applicable)"
	@echo "  help              Show this help menu"
	@echo "  help              Show this help menu"
	@echo "  env               Update environment variables (requires direnv)"
	@echo "  db-up             Start containers"
	@echo "  db-down           Stop containers"
	@echo "  db-purge          Stop containers and remove all volumes (careful!)"
	@echo "  migration-new     Create a new database migration"
	@echo "  migrate-up        Apply all pending migrations"
	@echo "  migrate-down      Roll back the last applied migration"
	@echo "  migrate-redo      Reapply the last migration (down then up)"
	@echo "  migrate-status    Show migration status"
	@echo "  swag              Generate Swagger docs"
	@echo "  help              Show this help menu"
