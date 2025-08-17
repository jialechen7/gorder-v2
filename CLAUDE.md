# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is gorder-v2, a microservices-based order management system written in Go. The system implements Domain-Driven Design (DDD) principles with a hexagonal architecture pattern. It handles order processing, payment processing, stock management, and includes distributed tracing and event-driven communication.

## Architecture

### Microservices Structure
- **Order Service** (`internal/order/`): Core order management with HTTP and gRPC endpoints
- **Payment Service** (`internal/payment/`): Payment processing with Stripe integration, HTTP-only
- **Stock Service** (`internal/stock/`): Inventory management, gRPC-only
- **Kitchen Service** (`internal/kitchen/`): Minimal service (placeholder implementation)

### Shared Components (`internal/common/`)
- **Anti-corruption layer**: Abstracts external service integrations
- **Discovery**: Consul-based service discovery
- **Tracing**: Jaeger distributed tracing
- **Broker**: RabbitMQ event messaging
- **Config**: Viper-based configuration management
- **Server**: Common HTTP/gRPC server setup

### Communication Patterns
- **Synchronous**: gRPC for service-to-service communication
- **Asynchronous**: RabbitMQ for event-driven messaging
- **Service Discovery**: Consul for dynamic service registration

## Development Commands

### Code Generation
```bash
make gen          # Generate both protobuf and OpenAPI
make genproto     # Generate protobuf files only
make genopenapi   # Generate OpenAPI client code only
```

### Code Quality
```bash
make lint         # Run comprehensive linting (golangci-lint + go-cleanarch)
make fmt          # Format code with goimports
```

### Infrastructure
```bash
docker-compose up # Start Consul, RabbitMQ, and Jaeger
make listen       # Start Stripe webhook listener for payments
```

### Building Services
Each service has its own `go.mod` and can be built independently:
```bash
cd internal/order && go build -o tmp/main.exe .
cd internal/payment && go build -o tmp/main.exe .
cd internal/stock && go build -o tmp/main.exe .
```

## Configuration

All services use a shared configuration file at `internal/common/config/global.yaml`. Key settings:
- Service endpoints and ports
- External service connections (Consul, RabbitMQ, Jaeger)
- Environment-specific overrides via environment variables

Services can run in different modes:
- Order: HTTP + gRPC servers
- Payment: HTTP only (configured via `payment.server-to-run`)
- Stock: gRPC only (configured via `stock.server-to-run`)

## Code Organization

### Domain Layer
Each service follows DDD patterns:
- `domain/`: Core business entities and repository interfaces
- `app/`: Application services (commands/queries)
- `adapters/`: External integrations and repository implementations
- `ports/`: API handlers (HTTP/gRPC)

### Anti-corruption Layer
The `common/` module provides abstraction for:
- External service clients (`client/`)
- Cross-cutting concerns (`decorator/`, `middleware/`)
- Infrastructure components (`broker/`, `discovery/`, `tracing/`)

## Testing

No test framework is currently configured. When adding tests, check each service's `go.mod` for dependencies and follow Go testing conventions.

## Linting Configuration

The project uses golangci-lint with strict configuration (`.golangci.yaml`):
- Enforces Go naming conventions and best practices
- Includes staticcheck, revive, and stylecheck
- Custom rules for context handling and error management

Always run `make lint` before committing changes.