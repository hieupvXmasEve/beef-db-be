.PHONY: all build run clean test coverage lint sqlc help migrate-create migrate-up migrate-down migrate-force migrate-version

# Include .env file
include .env
export

# Go parameters
BINARY_NAME=beef-db-be
MAIN_PACKAGE=./cmd/api
GO_FILES=$(shell find . -name '*.go' -not -path "./vendor/*")

# Build flags
LDFLAGS=-ldflags "-w -s"

all: lint test build

help:
	@echo "Available commands:"
	@echo "  make build          - Build the application"
	@echo "  make run           - Run the application"
	@echo "  make clean         - Clean build files"
	@echo "  make test          - Run tests"
	@echo "  make coverage      - Run tests with coverage"
	@echo "  make lint          - Run linter"
	@echo "  make sqlc          - Generate SQLC code"
	@echo "  make migrate-create NAME=migration_name - Create a new migration"
	@echo "  make migrate-up    - Run database migrations up"
	@echo "  make migrate-down  - Run database migrations down"
	@echo "  make migrate-force VERSION=x - Force migration version"
	@echo "  make migrate-version - Show current migration version"
	@echo "  make docker-build  - Build Docker image"
	@echo "  make docker-run    - Run Docker container"
	@echo "  make help          - Show this help message"

build:
	@echo "Building..."
	go build $(LDFLAGS) -o bin/$(BINARY_NAME) $(MAIN_PACKAGE)

run:
	@echo "Running..."
	go run $(MAIN_PACKAGE)

clean:
	@echo "Cleaning..."
	go clean
	rm -f bin/$(BINARY_NAME)
	rm -f coverage.out

test:
	@echo "Running tests..."
	go test -v ./...

coverage:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

lint:
	@echo "Running linter..."
	golangci-lint run

sqlc:
	@echo "Generating SQLC code..."
	sqlc generate

migrate-create:
	@echo "Creating migration files..."
	@if [ -z "$(NAME)" ]; then \
		echo "Please provide a migration name. Example: make migrate-create NAME=create_users_table"; \
		exit 1; \
	fi
	cd scripts && ./migrate.sh create $(NAME)

migrate-up:
	@echo "Running migrations up..."
	cd scripts && ./migrate.sh up

migrate-down:
	@echo "Running migrations down..."
	cd scripts && ./migrate.sh down

migrate-force:
	@echo "Forcing migration version..."
	@if [ -z "$(VERSION)" ]; then \
		echo "Please provide a version number. Example: make migrate-force VERSION=1"; \
		exit 1; \
	fi
	cd scripts && ./migrate.sh force $(VERSION)

migrate-version:
	@echo "Checking migration version..."
	cd scripts && ./migrate.sh version

# Development tools installation
install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Docker commands
docker-build:
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME) .

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 --env-file .env $(BINARY_NAME)

# Watch for changes and restart server (requires air)
dev:
	@echo "Starting development server..."
	air

# Install air for live reloading
install-air:
	@echo "Installing air..."
	go install github.com/cosmtrek/air@latest
