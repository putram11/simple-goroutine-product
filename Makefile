# Build stage
build:
	go build -o bin/server cmd/server/main.go

# Run the application
run:
	go run cmd/server/main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -cover ./...

# Run database migration
migrate:
	go run cmd/migrate/main.go

# Install dependencies
deps:
	go mod tidy

# Generate swagger docs
swagger:
	swag init -g cmd/server/main.go

# Docker commands
docker-build:
	docker build -t simple-product-api .

docker-run:
	docker run -p 8080:8080 simple-product-api

# Docker compose commands
compose-up:
	docker-compose up -d

compose-down:
	docker-compose down

compose-logs:
	docker-compose logs -f

# Clean up
clean:
	rm -rf bin/
	docker-compose down -v

# Lint code
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...

# Development setup
setup: deps swagger
	@echo "Development environment setup complete"

.PHONY: build run test test-coverage migrate deps swagger docker-build docker-run compose-up compose-down compose-logs clean lint fmt setup
