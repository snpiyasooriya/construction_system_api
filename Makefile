# Construction System API - Makefile

.PHONY: help build test test-unit test-integration test-api test-coverage clean dev logs

# Default target
help: ## Show this help message
	@echo "Construction System API - Available Commands:"
	@echo "============================================="
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Development
dev: ## Start development environment
	docker compose up --build

dev-detached: ## Start development environment in background
	docker compose up -d --build

logs: ## Show application logs
	docker compose logs -f api

# Building
build: ## Build the application
	docker compose build

# Testing
test: ## Run all tests
	./scripts/test.sh all

test-unit: ## Run unit tests only
	./scripts/test.sh unit

test-integration: ## Run integration tests only
	./scripts/test.sh integration

test-api: ## Run API tests only (Postman-like)
	./scripts/test.sh api

test-coverage: ## Run tests with coverage report
	./scripts/test.sh coverage

# Schedule API specific tests
test-schedule: ## Test schedule create endpoint specifically
	@echo "🧪 Testing Schedule Create API Endpoint..."
	@echo "=========================================="
	./scripts/test.sh api

# Database
db-reset: ## Reset database (remove volumes and restart)
	docker compose down -v
	docker compose up -d postgres

db-logs: ## Show database logs
	docker compose logs -f postgres

# Cleanup
clean: ## Clean up Docker containers and volumes
	docker compose down -v
	docker compose -f docker-compose.test.yml down -v
	docker system prune -f

clean-all: ## Clean up everything including images
	docker compose down -v --rmi all
	docker compose -f docker-compose.test.yml down -v --rmi all
	docker system prune -af

# Quick commands
up: dev-detached ## Alias for dev-detached
down: ## Stop all services
	docker compose down

restart: ## Restart all services
	docker compose restart

# Status
status: ## Show status of all containers
	docker compose ps

# API Testing (Postman-like commands)
api-health: ## Test API health endpoint
	@echo "🏥 Testing API Health..."
	curl -s http://localhost:8080/api/health | jq .

api-ping: ## Test API ping endpoint
	@echo "🏓 Testing API Ping..."
	curl -s http://localhost:8080/api/ping | jq .

# Schedule API manual tests
schedule-create-test: ## Manual test for schedule creation
	@echo "📅 Testing Schedule Creation..."
	curl -X POST http://localhost:8080/api/schedule/ \
		-H "Content-Type: application/json" \
		-d '{"name":"Manual Test Schedule","description":"Created via Makefile","project_id":1}' | jq .

schedule-get-test: ## Manual test for getting schedules by project
	@echo "📋 Testing Get Schedules by Project..."
	curl -s "http://localhost:8080/api/schedule/ByProject/?project_id=1" | jq .

# Development helpers
install-deps: ## Install development dependencies
	go mod tidy
	go mod download

format: ## Format Go code
	go fmt ./...

lint: ## Run linter
	golangci-lint run

# Documentation
docs: ## Generate API documentation
	@echo "📚 API Documentation will be available at http://localhost:8080/docs"
	@echo "Current endpoints:"
	@echo "  POST /api/schedule/           - Create schedule"
	@echo "  GET  /api/schedule/ByProject/ - Get schedules by project"
	@echo "  GET  /api/health              - Health check"
	@echo "  GET  /api/ping                - Ping test"
