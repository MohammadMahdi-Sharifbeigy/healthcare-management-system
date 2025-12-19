.PHONY: help build run test clean docker-up docker-down migrate lint fmt install-deps

# Variables
BINARY_NAME=healthcare-api
GO=go
DOCKER_COMPOSE=docker-compose
PORT=8080

help:
	@echo "Healthcare Management System - Available Commands"
	@echo ""
	@echo "Development:"
	@echo "  make install-deps    Install Go dependencies"
	@echo "  make build           Build the application"
	@echo "  make run             Run the application locally"
	@echo "  make dev             Run with hot reload (requires air)"
	@echo ""
	@echo "Testing:"
	@echo "  make test            Run all tests"
	@echo "  make test-unit       Run unit tests only"
	@echo "  make test-integration Run integration tests only"
	@echo "  make coverage        Generate test coverage report"
	@echo ""
	@echo "Code Quality:"
	@echo "  make lint            Run linter"
	@echo "  make fmt             Format code"
	@echo "  make vet             Run go vet"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-up       Start Docker containers"
	@echo "  make docker-down     Stop Docker containers"
	@echo "  make docker-build    Build Docker image"
	@echo "  make docker-logs     View container logs"
	@echo ""
	@echo "Database:"
	@echo "  make migrate-up      Run migrations"
	@echo "  make migrate-down    Rollback migrations"
	@echo "  make db-seed         Seed sample data"
	@echo ""
	@echo "Cleanup:"
	@echo "  make clean           Remove build artifacts"
	@echo "  make clean-all       Remove build artifacts and containers"

install-deps:
	@echo "Installing Go dependencies..."
	$(GO) mod download
	$(GO) mod tidy
	@echo "Dependencies installed successfully"

build:
	@echo "Building $(BINARY_NAME)..."
	$(GO) build -o bin/$(BINARY_NAME) ./cmd/server/main.go
	@echo "Build complete: bin/$(BINARY_NAME)"

run: build
	@echo "Running $(BINARY_NAME)..."
	./bin/$(BINARY_NAME)

dev:
	@echo "Running with hot reload..."
	@command -v air >/dev/null 2>&1 || (echo "Installing air..." && go install github.com/cosmtrek/air@latest)
	air

test:
	@echo "Running all tests..."
	$(GO) test -v -race -cover ./...

test-unit:
	@echo "Running unit tests..."
	$(GO) test -v -race -short ./tests/unit/...

test-integration:
	@echo "Running integration tests..."
	$(GO) test -v -race ./tests/integration/...

coverage:
	@echo "Generating coverage report..."
	$(GO) test -v -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint:
	@echo "Running linter..."
	@command -v golangci-lint >/dev/null 2>&1 || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...

fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...
	goimports -w .

vet:
	@echo "Running go vet..."
	$(GO) vet ./...

docker-build:
	@echo "Building Docker image..."
	$(DOCKER_COMPOSE) build

docker-up:
	@echo "Starting Docker containers..."
	$(DOCKER_COMPOSE) up -d
	@echo "Containers started. API available at http://localhost:$(PORT)"

docker-down:
	@echo "Stopping Docker containers..."
	$(DOCKER_COMPOSE) down

docker-logs:
	@echo "Viewing Docker logs..."
	$(DOCKER_COMPOSE) logs -f api

migrate-up:
	@echo "Running database migrations..."
	@command -v migrate >/dev/null 2>&1 || (echo "Installing migrate CLI..." && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest)
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/healthcare_db?sslmode=disable" up

migrate-down:
	@echo "Rolling back database migrations..."
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/healthcare_db?sslmode=disable" down

db-seed:
	@echo "Seeding database with sample data..."
	psql -h localhost -U postgres -d healthcare_db -f migrations/sample_data.sql
	@echo "Database seeded successfully"

clean:
	@echo "Cleaning up..."
	rm -f bin/$(BINARY_NAME)
	$(GO) clean
	rm -f coverage.out coverage.html

clean-all: clean docker-down
	@echo "Full cleanup complete"

.DEFAULT_GOAL := help
