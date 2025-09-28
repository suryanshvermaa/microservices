# Go Microservices E-commerce Platform

A modern microservices-based e-commerce platform built with Go, gRPC, GraphQL, PostgreSQL, and Elasticsearch. This project demonstrates a scalable architecture pattern with separate services for account management, product catalog, order processing, and a unified GraphQL API gateway.

## 🏗️ Architecture Overview

The system consists of four main microservices:

- **Account Service**: User account management with PostgreSQL storage
- **Catalog Service**: Product catalog with Elasticsearch for search capabilities  
- **Order Service**: Order processing and management with PostgreSQL storage
- **GraphQL Gateway**: Unified API gateway that aggregates data from all services

```
                ┌─────────────────────┐
                │   GraphQL API       │
                │    Gateway :8000    │
                └─────────┬───────────┘
                          │
        ┌─────────────────┼─────────────────┐
        │                 │                 │
┌───────▼───────┐ ┌───────▼───────┐ ┌──────▼───────┐
│ Account       │ │ Catalog       │ │ Order        │
│ Service :8080 │ │ Service :8080 │ │ Service :8080│
└───────┬───────┘ └───────┬───────┘ └──────┬───────┘
        │                 │                │
┌───────▼───────┐ ┌───────▼───────┐ ┌──────▼───────┐
│ PostgreSQL DB │ │ Elasticsearch │ │ PostgreSQL DB│
│ (account_db)  │ │ (catalog_db)  │ │ (order_db)   │
└───────────────┘ └───────────────┘ └──────────────┘
```

## 🚀 Services Overview

### Account Service
- **Purpose**: Manages user accounts and authentication
- **Technology**: Go + gRPC + PostgreSQL
- **Features**:
  - Create new accounts
  - Retrieve account information
  - List accounts with pagination
- **Database**: PostgreSQL with `accounts` table

### Catalog Service  
- **Purpose**: Product catalog management and search
- **Technology**: Go + gRPC + Elasticsearch
- **Features**:
  - Add products to catalog
  - Search products by text query
  - Retrieve product details
  - List products with pagination
- **Database**: Elasticsearch for full-text search capabilities

### Order Service
- **Purpose**: Order processing and management
- **Technology**: Go + gRPC + PostgreSQL
- **Features**:
  - Create orders with multiple products
  - Calculate total order price
  - Retrieve orders by account
- **Database**: PostgreSQL with `orders` and `order_products` tables

### GraphQL Gateway
- **Purpose**: Unified API gateway
- **Technology**: Go + GraphQL (gqlgen)
- **Features**:
  - Single endpoint for all client interactions
  - Aggregates data from all microservices
  - Type-safe GraphQL schema
  - Real-time queries and mutations

## 🛠️ Technology Stack

- **Language**: Go 1.24.2
- **Communication**: gRPC for inter-service communication
- **API Gateway**: GraphQL with gqlgen
- **Databases**: 
  - PostgreSQL for transactional data
  - Elasticsearch for search functionality
- **Containerization**: Docker & Docker Compose
- **Key Libraries**:
  - `google.golang.org/grpc` - gRPC framework
  - `github.com/99designs/gqlgen` - GraphQL server
  - `github.com/lib/pq` - PostgreSQL driver
  - `github.com/elastic/go-elasticsearch/v9` - Elasticsearch client
  - `github.com/segmentio/ksuid` - Unique ID generation

## 📋 Prerequisites

- Go 1.24.2 or later
- Docker and Docker Compose
- Git

## 🚀 Quick Start

### 1. Clone the Repository
```bash
git clone https://github.com/suryanshvermaa/microservices.git
cd microservices/golang
```

### 2. Start Services with Docker Compose
```bash
docker-compose up -d
```

This will start:
- PostgreSQL databases for Account and Order services
- Elasticsearch (v8.15.0) for Catalog service  
- All microservices (account, catalog, order)
- GraphQL gateway (accessible on port 8000)

### 3. Verify Services
Check that all services are running:
```bash
docker-compose ps
```

### 4. Test the GraphQL Gateway
Access the GraphQL playground: `http://localhost:8000/playground`

## 📡 API Usage

### GraphQL Endpoint
The GraphQL gateway provides a unified API at: `http://localhost:8000/graphql`

### GraphQL Playground
Access the interactive GraphQL playground at: `http://localhost:8000/playground`

### Sample Queries

#### Create an Account
```graphql
mutation {
  createAccount(account: {name: "John Doe"}) {
    id
    name
  }
}
```

#### Create a Product
```graphql
mutation {
  createProduct(product: {
    name: "MacBook Pro"
    description: "High-performance laptop"
    price: 1999.99
  }) {
    id
    name
    price
  }
}
```

#### Create an Order
```graphql
mutation {
  createOrder(order: {
    accountId: "2G8UdqGCWJ9ZCJHeLM4s1Z8DKfP"
    products: [{
      id: "2G8UeA1K5L2NxJ4QxBj8sK9DfR7"
      quantity: 2
    }]
  }) {
    id
    totalPrice
    products {
      name
      quantity
      price
    }
  }
}
```

#### Query Accounts
```graphql
query {
  accounts(pagination: {skip: 0, take: 10}) {
    id
    name
    orders {
      id
      totalPrice
      createdAt
    }
  }
}
```

#### Search Products
```graphql
query {
  products(query: "laptop", pagination: {skip: 0, take: 10}) {
    id
    name
    description
    price
  }
}
```

## 🐳 Docker Configuration

The project includes Dockerfiles for each service:

- `services/account/app.dockerfile` - Account service container
- `services/account/db.dockerfile` - Account PostgreSQL container
- `services/catalog/app.dockerfile` - Catalog service container
- `services/order/app.dockerfile` - Order service container  
- `services/order/db.dockerfile` - Order PostgreSQL container
- `services/graphql/app.dockerfile` - GraphQL gateway container

## 🗄️ Database Schema

### Account Service (PostgreSQL)
Database: `account_db` with user `suryansh`
```sql
CREATE TABLE accounts(
    id CHAR(27) PRIMARY KEY,
    name VARCHAR(24) NOT NULL
);
```

### Order Service (PostgreSQL)  
Database: `order_db` with user `suryansh`
```sql
CREATE TABLE orders(
    id CHAR(27) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    account_id CHAR(27) NOT NULL,
    total_price MONEY NOT NULL
);

CREATE TABLE ordered_products(
    order_id CHAR(27) REFERENCES orders (id) ON DELETE CASCADE,
    product_id CHAR(27),
    quantity INT NOT NULL,
    PRIMARY KEY (product_id, order_id)
);
```

### Catalog Service (Elasticsearch)
Elasticsearch version: 8.15.0
Products are stored as documents with fields:
- `name` (string)
- `description` (string)  
- `price` (float)

## 🔧 Development

### Recent Updates
- **Fixed**: Tautological condition error in order repository (`respository.go` line 137)
- **Updated**: Docker-compose configuration with proper service networking
- **Improved**: Database connection configuration and error handling

### Building Services Locally

1. **Install dependencies**:
```bash
go mod tidy
```

2. **Generate Protocol Buffers**:
- Ensure `protoc` and Go plugins are installed:
- `protoc-gen-go` and `protoc-gen-go-grpc`
```bash
# Install protoc (if not already installed)
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
```bash
# For each service, regenerate protobuf files
protoc --go_out=./services/account --go-grpc_out=./services/account services/account/account.proto
protoc --go_out=./services/catalog --go-grpc_out=./services/catalog services/catalog/catalog.proto
protoc --go_out=./services/order --go-grpc_out=./services/order services/order/order.proto
```

3. **Generate GraphQL Code**:
```bash
cd services/graphql
go generate ./...
```

4. **Run individual services**:
```bash
# Account service
go run services/account/cmd/account/main.go

# Catalog service  
go run services/catalog/cmd/catalog/main.go

# Order service
go run services/order/cmd/order/main.go

# GraphQL gateway
go run services/graphql/main.go
```

### Environment Variables

Each service can be configured via environment variables:

#### Account Service
- `DATABASE_URL` - PostgreSQL connection string (default: `postgres://suryansh:123456@account_db/suryansh?sslmode=disable`)

#### Catalog Service  
- `DATABASE_URL` - Elasticsearch endpoint (default: `http://catalog_db:9200`)

#### Order Service
- `DATABASE_URL` - PostgreSQL connection string (default: `postgres://suryansh:123456@order_db/suryansh?sslmode=disable`)
- `ACCOUNT_SERVICE_URL` - Account service gRPC endpoint (default: `account:8080`)
- `CATALOG_SERVICE_URL` - Catalog service gRPC endpoint (default: `catalog:8080`)

#### GraphQL Gateway
- `ACCOUNT_SERVICE_URL` - Account service gRPC endpoint (default: `account:8080`)
- `CATALOG_SERVICE_URL` - Catalog service gRPC endpoint (default: `catalog:8080`)  
- `ORDER_SERVICE_URL` - Order service gRPC endpoint (default: `order:8080`)

## 🧪 Testing

### Testing GraphQL Queries
You can test the GraphQL API using:
- GraphQL Playground (if enabled)
- Postman
- curl commands
- Any GraphQL client

Example curl request:
```bash
curl -X POST http://localhost:8000/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ accounts(pagination: {skip: 0, take: 5}) { id name } }"}'
```

## 📁 Project Structure

```
.
├── docker-compose.yml          # Docker orchestration
├── go.mod                     # Go module definition
├── go.sum                     # Go module checksums
├── README.md                  # This file
└── services/
    ├── account/               # Account microservice
    │   ├── account.proto      # gRPC service definition
    │   ├── service.go         # Business logic
    │   ├── respository.go     # Data access layer
    │   ├── server.go          # gRPC server
    │   ├── client.go          # gRPC client
    │   ├── up.sql            # Database schema
    │   ├── app.dockerfile    # Service container
    │   ├── db.dockerfile     # Database container
    │   ├── cmd/account/main.go # Service entry point
    │   └── pb/               # Generated protobuf code
    ├── catalog/              # Catalog microservice
    │   ├── catalog.proto     # gRPC service definition
    │   ├── service.go        # Business logic
    │   ├── respository.go    # Elasticsearch integration
    │   ├── server.go         # gRPC server
    │   ├── client.go         # gRPC client
    │   ├── app.dockerfile    # Service container
    │   ├── cmd/catalog/main.go # Service entry point
    │   └── pb/               # Generated protobuf code
    ├── order/                # Order microservice
    │   ├── order.proto       # gRPC service definition
    │   ├── service.go        # Business logic
    │   ├── respository.go    # Data access layer
    │   ├── server.go         # gRPC server
    │   ├── client.go         # gRPC client
    │   ├── up.sql           # Database schema
    │   ├── app.dockerfile   # Service container
    │   ├── db.dockerfile    # Database container
    │   ├── cmd/order/main.go # Service entry point
    │   └── pb/              # Generated protobuf code
    └── graphql/             # GraphQL API Gateway
        ├── schema.graphql   # GraphQL schema definition
        ├── main.go         # Service entry point
        ├── graph.go        # GraphQL server setup
        ├── *_resolvers.go  # GraphQL resolvers
        ├── models*.go      # GraphQL models
        ├── generated.go    # Generated gqlgen code
        ├── gqlgen.yml     # gqlgen configuration
        └── app.dockerfile  # Service container
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🔗 Related Resources

- [gRPC Documentation](https://grpc.io/docs/)
- [GraphQL Documentation](https://graphql.org/)
- [gqlgen Documentation](https://gqlgen.com/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Elasticsearch Documentation](https://www.elastic.co/guide/)

## 🐛 Known Issues

- GraphQL gateway implementation may have incomplete resolver functions  
- Missing authentication/authorization system
- Database connection pooling not optimized for production
- Missing comprehensive logging and monitoring
- Error handling could be improved in some service methods

### Recently Fixed Issues
- ✅ **Fixed**: Tautological condition error in order repository (September 2025)
- ✅ **Fixed**: Docker build configuration for GraphQL gateway

## 🚀 Future Enhancements

- Complete GraphQL resolver implementations
- Add authentication and authorization
- Implement comprehensive logging
- Add monitoring and health checks
- Kubernetes manifests for deployment