# Simple Goroutine Product API

A simple CRUD API for products built with Go, Echo framework, GORM, and PostgreSQL. This project demonstrates MVP (Model-View-Presenter) architecture pattern with Goroutines for asynchronous create and update operations.

## Features

- ✅ CRUD operations for products
- ✅ MVP architecture pattern
- ✅ Goroutines for create and update operations
- ✅ PostgreSQL database with GORM
- ✅ Echo web framework
- ✅ Swagger documentation
- ✅ Docker support
- ✅ Unit tests
- ✅ Environment configuration

## Tech Stack

- **Language**: Go 1.21
- **Web Framework**: Echo v4
- **ORM**: GORM
- **Database**: PostgreSQL
- **Documentation**: Swagger
- **Testing**: Testify
- **Containerization**: Docker & Docker Compose

## Project Structure

```
simple-goroutine-product/
├── cmd/
│   ├── server/           # Main application
│   └── migrate/          # Database migration
├── internal/
│   ├── models/           # Data models
│   ├── repositories/     # Data access layer
│   ├── presenters/       # Business logic layer (MVP)
│   ├── handlers/         # HTTP handlers (Views in MVP)
│   ├── routes/           # Route definitions
│   ├── database/         # Database connection
│   └── validators/       # Request validation
├── docs/                 # Swagger documentation
├── docker-compose.yml    # Docker services
├── Dockerfile           # Application container
├── .env                 # Environment variables
└── go.mod              # Go modules
```

## Quick Start

### Using Docker Compose (Recommended)

1. Clone the repository:
```bash
git clone <repository-url>
cd simple-goroutine-product
```

2. Start the services:
```bash
docker-compose up -d
```

3. The API will be available at `http://localhost:8080`
4. Swagger documentation at `http://localhost:8080/swagger/index.html`

### Manual Setup

1. Install dependencies:
```bash
go mod tidy
```

2. Start PostgreSQL database:
```bash
docker run --name postgres-db -e POSTGRES_DB=product_db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres:15-alpine
```

3. Run the application:
```bash
go run cmd/server/main.go
```

## API Endpoints

### Products

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST   | `/api/v1/products` | Create a new product |
| GET    | `/api/v1/products` | Get all products (with pagination) |
| GET    | `/api/v1/products/:id` | Get a product by ID |
| PUT    | `/api/v1/products/:id` | Update a product |
| DELETE | `/api/v1/products/:id` | Delete a product |

### Health Check

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/health` | Health check endpoint |

## Request/Response Examples

### Create Product
```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPhone 15",
    "description": "Latest iPhone model",
    "price": 999.99,
    "stock": 50
  }'
```

### Get Products
```bash
curl "http://localhost:8080/api/v1/products?page=1&limit=10"
```

### Update Product
```bash
curl -X PUT http://localhost:8080/api/v1/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPhone 15 Pro",
    "description": "Premium iPhone model",
    "price": 1199.99,
    "stock": 30
  }'
```

## Environment Variables

Create a `.env` file in the project root:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=product_db
APP_PORT=8080
```

## Testing

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

Run specific package tests:
```bash
go test ./internal/presenters/
go test ./internal/handlers/
```

## MVP Architecture

This project follows the MVP (Model-View-Presenter) pattern:

- **Models** (`internal/models/`): Data structures and business entities
- **Views** (`internal/handlers/`): HTTP handlers that handle requests and responses
- **Presenters** (`internal/presenters/`): Business logic layer that coordinates between models and views

### Goroutines Implementation

Create and Update operations are implemented using Goroutines for asynchronous processing:

- **Create Product**: Executes database insertion in a separate goroutine with timeout handling
- **Update Product**: Executes database update in a separate goroutine with timeout handling

Benefits:
- Non-blocking operations
- Better performance for concurrent requests
- Timeout handling for long-running operations

## Swagger Documentation

Access the interactive API documentation at:
```
http://localhost:8080/swagger/index.html
```

To regenerate Swagger docs:
```bash
swag init -g cmd/server/main.go
```

## Database Migration

The application automatically runs migrations on startup. To run migrations manually:
```bash
go run cmd/migrate/main.go
```

## Development

### Adding New Features

1. Add models in `internal/models/`
2. Create repository interface and implementation in `internal/repositories/`
3. Implement business logic in `internal/presenters/`
4. Create HTTP handlers in `internal/handlers/`
5. Add routes in `internal/routes/`
6. Write tests for all layers

### Code Quality

- Use `gofmt` for code formatting
- Use `golint` for linting
- Write unit tests for all business logic
- Follow Go best practices and conventions

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Run tests and ensure they pass
6. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
