# 🚀 Microservices Architecture Project

A comprehensive microservices implementation showcasing both theoretical concepts and practical implementations using Go and Node.js. This project demonstrates modern microservices architecture patterns, inter-service communication, and scalable system design.

## 📋 Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Services](#services)
- [Technologies](#technologies)
- [Getting Started](#getting-started)
- [API Documentation](#api-documentation)
- [Deployment](#deployment)
- [Contributing](#contributing)
- [License](#license)

## 🎯 Overview

This project serves as both a learning resource for microservices concepts and a practical implementation of microservices architecture. It includes:

- **Theoretical Foundation**: Comprehensive documentation on microservices concepts, patterns, and best practices
- **Go Implementation**: Microservices built with Go, featuring gRPC, GraphQL, and containerized deployment
- **Node.js Implementation**: RESTful microservices with modern TypeScript and Express.js
- **Frontend Client**: React-based user interface for interacting with the microservices

## 🏗️ Architecture

The project follows a distributed microservices architecture with the following key components:

### Core Services
- **Account Service**: User management and authentication
- **Catalog Service**: Product catalog and inventory management
- **Order Service**: Order processing and management
- **GraphQL Gateway**: Unified API interface for frontend clients

### Communication Patterns
- **Synchronous**: REST APIs and gRPC for direct service-to-service communication
- **Asynchronous**: Event-driven architecture with message queues
- **API Gateway**: Centralized routing and request handling

## 🔧 Services

### Go Microservices (`/golang`)

#### Account Service
- **Purpose**: User account management and authentication
- **Technology**: Go, gRPC, PostgreSQL
- **Features**: User CRUD operations, authentication, account validation

#### Catalog Service
- **Purpose**: Product catalog and inventory management
- **Technology**: Go, gRPC, PostgreSQL
- **Features**: Product CRUD operations, search, inventory tracking

#### Order Service
- **Purpose**: Order processing and management
- **Technology**: Go, gRPC, PostgreSQL
- **Features**: Order creation, status tracking, order history

#### GraphQL Gateway
- **Purpose**: Unified API interface for frontend clients
- **Technology**: Go, GraphQL (gqlgen), gRPC client
- **Features**: Schema stitching, query optimization, data aggregation

### Node.js Microservices (`/nodejs`)

#### Snippet Service
- **Purpose**: Code snippet management and sharing
- **Technology**: Node.js, Express.js, TypeScript
- **Features**: Snippet CRUD operations, user association, search

#### Comments Service
- **Purpose**: Comment system for snippets
- **Technology**: Node.js, Express.js, TypeScript
- **Features**: Comment CRUD operations, user association, moderation

#### Frontend Client
- **Purpose**: User interface for snippet management
- **Technology**: React, TypeScript, Tailwind CSS, Vite
- **Features**: Modern UI components, responsive design, dark mode support

## 🛠️ Technologies

### Backend Technologies
- **Go 1.24.2**: High-performance microservices
- **Node.js**: JavaScript runtime for web services
- **gRPC**: High-performance RPC framework
- **GraphQL**: Query language and runtime
- **PostgreSQL**: Relational database
- **Docker**: Containerization and orchestration

### Frontend Technologies
- **React 19**: Modern UI framework
- **TypeScript**: Type-safe JavaScript
- **Tailwind CSS**: Utility-first CSS framework
- **Vite**: Fast build tool and dev server
- **Radix UI**: Accessible component primitives

### Development Tools
- **ESLint**: Code linting and formatting
- **Prettier**: Code formatting
- **Husky**: Git hooks for code quality
- **Nodemon**: Development server with auto-reload

## 🚀 Getting Started
### Protobuf generate command
```
protoc --go_out=. --go-grpc_out=. --proto_path=.  *.proto
```
### Prerequisites
- Go 1.24.2 or higher
- Node.js 18 or higher
- Docker and Docker Compose
- PostgreSQL (optional, can use Docker)

### Quick Start

#### 1. Clone the Repository
```bash
git clone <repository-url>
cd microservices
```

#### 2. Start Go Microservices
```bash
cd golang
docker-compose up -d
```

#### 3. Start Node.js Microservices
```bash
# Start Snippet Service
cd nodejs/snippet
npm install
npm run dev

# Start Comments Service
cd ../comments
npm install
npm run dev
```

#### 4. Start Frontend Client
```bash
cd nodejs/client
npm install
npm run dev
```

### Environment Configuration

Create `.env` files in each service directory with appropriate configuration:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME=service_name
DB_USER=username
DB_PASSWORD=password

# Service Configuration
PORT=3000
JWT_SECRET=your-secret-key
```

## 📚 API Documentation

### GraphQL API (Go)

The GraphQL gateway provides a unified interface for all services:

```graphql
# Query accounts with pagination
query {
  accounts(pagination: { skip: 0, take: 10 }) {
    id
    name
    orders {
      id
      totalPrice
      createdAt
    }
  }
}

# Create a new order
mutation {
  createOrder(order: {
    accountId: "user123"
    products: [
      { id: "prod1", quantity: 2 }
    ]
  }) {
    id
    totalPrice
    createdAt
  }
}
```

### REST APIs (Node.js)

#### Snippet Service
- `GET /api/v1/snippets` - List snippets
- `POST /api/v1/snippets` - Create snippet
- `GET /api/v1/snippets/:id` - Get snippet by ID
- `PUT /api/v1/snippets/:id` - Update snippet
- `DELETE /api/v1/snippets/:id` - Delete snippet

#### Comments Service
- `GET /api/v1/comments` - List comments
- `POST /api/v1/comments` - Create comment
- `GET /api/v1/comments/:id` - Get comment by ID
- `PUT /api/v1/comments/:id` - Update comment
- `DELETE /api/v1/comments/:id` - Delete comment

## 🐳 Deployment

### Docker Deployment

All services include Docker configurations for easy deployment:

```bash
# Build and run all services
docker-compose up --build

# Run specific service
docker-compose up service-name
```

### Production Considerations

- Use environment variables for configuration
- Implement proper logging and monitoring
- Set up health checks and circuit breakers
- Configure load balancing and auto-scaling
- Implement proper security measures (HTTPS, authentication)

## 📖 Learning Resources

The `/concepts` directory contains comprehensive documentation on microservices architecture:

- **Architectural Evolution**: From monoliths to microservices
- **Decomposition Strategies**: How to break down applications
- **Inter-Service Communication**: Patterns and best practices
- **Event-Driven Architecture**: Asynchronous communication patterns
- **Scaling Strategies**: Horizontal and vertical scaling approaches
- **Challenges and Solutions**: Common pitfalls and mitigation strategies

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow the existing code style and conventions
- Write tests for new functionality
- Update documentation as needed
- Ensure all services build and run correctly
- Follow microservices best practices

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **Go Community**: For excellent microservices libraries and tools
- **Node.js Community**: For robust web development ecosystem
- **Microservices Community**: For architectural patterns and best practices

## 📞 Support

For questions, issues, or contributions:

- Open an issue on GitHub
- Review the documentation in `/concepts`
- Check the individual service READMEs for specific guidance

---
