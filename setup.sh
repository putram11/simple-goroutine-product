#!/bin/bash

# Simple Product API Setup Script

echo "🚀 Setting up Simple Product API..."

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker is not running. Please start Docker and try again."
    exit 1
fi

# Check if docker-compose is available
if ! command -v docker-compose &> /dev/null; then
    echo "❌ docker-compose is not installed. Please install docker-compose and try again."
    exit 1
fi

echo "✅ Docker is running"

# Build and start services
echo "🔧 Building and starting services..."
docker-compose up -d

# Wait for database to be ready
echo "⏳ Waiting for database to be ready..."
sleep 10

# Check if services are running
if docker-compose ps | grep -q "Up"; then
    echo "✅ Services are running successfully!"
    echo ""
    echo "🌐 API is available at: http://localhost:8080"
    echo "📚 Swagger documentation: http://localhost:8080/swagger/index.html"
    echo "❤️  Health check: http://localhost:8080/health"
    echo ""
    echo "📝 Example API calls:"
    echo "# Create a product"
    echo "curl -X POST http://localhost:8080/api/v1/products \\"
    echo "  -H 'Content-Type: application/json' \\"
    echo "  -d '{\"name\":\"iPhone 15\",\"description\":\"Latest iPhone\",\"price\":999.99,\"stock\":50}'"
    echo ""
    echo "# Get all products"
    echo "curl http://localhost:8080/api/v1/products"
    echo ""
    echo "🛑 To stop services: docker-compose down"
else
    echo "❌ Failed to start services. Check docker-compose logs for details."
    docker-compose logs
fi
