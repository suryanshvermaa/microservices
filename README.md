# 🚀 My Microservices Architecture Journey

A comprehensive personal learning repository where I explore microservices architecture from theory to practice. This repository documents my journey through distributed systems, combining extensive theoretical study with hands-on implementations in Go and Node.js.

## 📋 Table of Contents

- [🎯 About This Project](#-about-this-project)
- [📚 My Learning Structure](#-my-learning-structure)
- [🏗️ Architecture Implementations](#️-architecture-implementations)
- [📖 Personal Learning Path](#-personal-learning-path)
- [💻 My Implementations](#-my-implementations)
- [🛠️ Technologies I Used](#️-technologies-i-used)
- [🚀 How to Run My Projects](#-how-to-run-my-projects)
- [📡 API Examples](#-api-examples)
- [🐳 Deployment Setup](#-deployment-setup)
- [📚 My Study Notes](#-my-study-notes)
- [🎯 What I Learned](#-what-i-learned)
- [📄 License](#-license)

## 🎯 About This Project

This is my personal exploration of microservices architecture, created to deepen my understanding of distributed systems. The repository serves multiple purposes:

### 📚 **My Theoretical Foundation** (`/concepts`)
My comprehensive study notes covering microservices architecture from fundamentals to advanced patterns. These are my personal interpretations and learnings from various sources, organized for future reference.

### 🐹 **My Go Implementation** (`/golang`)
My production-ready e-commerce microservices platform built with Go. This demonstrates my understanding of gRPC communication, GraphQL API gateways, and containerized deployments.

### 🟨 **My Node.js Implementation** (`/nodejs`)
My modern code snippet management platform showcasing TypeScript microservices with a React frontend. This represents my exploration of web-based microservices patterns.

## 📚 My Learning Structure

```
my-microservices-journey/
├── 📖 concepts/                    # My microservices study notes
│   ├── readme.md                   # My complete learning guide
│   ├── TheArchitecturalEvolution.md # How I understand system evolution
│   ├── BuisinessRequirements.md    # Business drivers I studied
│   ├── Views.md                    # Stakeholder perspectives I learned
│   ├── HexagonalArc.md            # Clean architecture principles I applied
│   ├── APIGateways.md             # Gateway patterns I explored
│   ├── decomposition.md           # Decomposition strategies I learned
│   ├── EventdrivenArchitecture.md # Event-driven patterns I studied
│   ├── EventStreams.md            # Streaming concepts I mastered
│   ├── EventProcessing.md         # Processing patterns I implemented
│   ├── ScalingMS.md               # Scaling strategies I researched
│   ├── Queues.md                  # Queue patterns I explored
│   ├── ChallengesWithMS.md        # Challenges I discovered
│   ├── ModularMonolith.md         # Alternative approaches I considered
│   ├── EDATakeways.md             # Key insights I gained
│   ├── msConsiderations&Concepts.md # Advanced concepts I studied
│   └── images/                     # Diagrams I created/collected
│
├── 🐹 golang/                      # My Go microservices project
│   ├── docker-compose.yml          # My container orchestration
│   ├── go.mod & go.sum            # My Go dependencies
│   ├── README.md                   # My Go project documentation
│   └── services/
│       ├── account/                # My user account service
│       ├── catalog/                # My product catalog service
│       ├── order/                  # My order processing service
│       └── graphql/                # My GraphQL API gateway
│
├── 🟨 nodejs/                      # My Node.js microservices project
│   ├── client/                     # My React frontend application
│   ├── snippet/                    # My code snippet service
│   └── comments/                   # My comment service
│
├── LICENSE                         # MIT License
└── README.md                       # This personal documentation
```

## 🏗️ Architecture Implementations

### 📚 My Conceptual Learning Journey
I've documented my complete understanding of microservices evolution:

```
My Learning Path: Monolith → Macroservices → Microservices → Event-Driven
                      ↓           ↓              ↓              ↓
Understanding:    Tight        Some           Loose        Completely
                 Coupling    Independence    Coupling      Decoupled
```

**My Study Modules:**
- **Foundations**: How I learned architectural evolution and business requirements
- **Core Patterns**: My exploration of hexagonal architecture and API gateways  
- **Communication**: My study of inter-service communication and event-driven patterns
- **Advanced Topics**: My research into scaling strategies and operational concerns
- **Trade-offs**: My analysis of challenges and alternative approaches

### 🐹 My Go E-commerce Platform
Production-ready system I built to understand distributed microservices:

```
                    ┌─────────────────────┐
                    │   My GraphQL        │
                    │     Gateway         │
                    │    (Port 8080)      │
                    └─────────┬───────────┘
                              │ gRPC
        ┌─────────────────────┼─────────────────────┐
        │                     │                     │
┌───────▼───────┐    ┌────────▼────────┐    ┌──────▼───────┐
│ My Account     │    │ My Catalog      │    │ My Order     │
│ Service        │    │ Service         │    │ Service      │
│ (Port 8080)    │    │ (Port 8080)     │    │ (Port 8080)  │
└───────┬───────┘    └────────┬────────┘    └──────┬───────┘
        │                     │                    │
┌───────▼───────┐    ┌────────▼────────┐    ┌──────▼───────┐
│ PostgreSQL    │    │ Elasticsearch   │    │ PostgreSQL   │
│ Database      │    │ Search Engine   │    │ Database     │
└───────────────┘    └─────────────────┘    └──────────────┘
```

**What I Built:**
- **gRPC Communication**: High-performance inter-service communication that I implemented
- **GraphQL Gateway**: Unified API with data aggregation that I designed
- **Database per Service**: Separate data stores with technologies I chose
- **Docker Containerization**: Complete containerized deployment I configured

### 🟨 My Node.js Snippet Platform
Modern web application I created to explore TypeScript microservices:

```
                    ┌─────────────────────┐
                    │   My React Client   │
                    │  (Vite + TypeScript)│
                    └─────────┬───────────┘
                              │ HTTP/REST
        ┌─────────────────────┼─────────────────────┐
        │                     │                     │
┌───────▼───────┐    ┌────────▼────────┐    ┌──────▼───────┐
│ My Snippet     │    │ My Comments     │    │ My Future    │
│ Service        │    │ Service         │    │ Services     │
│ (Port 3000)    │    │ (Port 3001)     │    │              │
└───────┬───────┘    └────────┬────────┘    └──────────────┘
        │                     │
┌───────▼───────┐    ┌────────▼────────┐
│ In-Memory     │    │ In-Memory       │
│ Storage       │    │ Storage         │
└───────────────┘    └─────────────────┘
```

**What I Implemented:**
- **TypeScript Services**: Type-safe backend development that I practiced
- **React Frontend**: Modern UI with libraries I learned
- **RESTful APIs**: Standard HTTP-based communication I designed
- **Modular Architecture**: Clean separation of concerns I applied

## � Learning Path

### 🎓 **For Beginners** (Start Here)
**Goal**: Understand why and when to use microservices

1. **📊 [Business Requirements](./concepts/BuisinessRequirements.md)** - Learn the business drivers for microservices adoption
2. **🔄 [Architectural Evolution](./concepts/TheArchitecturalEvolution.md)** - See the complete journey from monolith to microservices
3. **👥 [Views & Perspectives](./concepts/Views.md)** - Understand how different stakeholders view microservices

### 🔧 **For Intermediate Developers**
**Goal**: Master core architectural patterns and design principles

1. **🏗️ [Hexagonal Architecture](./concepts/HexagonalArc.md)** - Learn ports and adapters pattern for clean architecture
2. **⚡ [Decomposition Strategies](./concepts/decomposition.md)** - Master techniques for breaking down monoliths
3. **🚪 [API Gateway Patterns](./concepts/APIGateways.md)** - Design effective entry points and routing

### 🚀 **For Advanced Practitioners**
**Goal**: Build resilient, event-driven distributed systems

1. **📡 [Event-Driven Architecture](./concepts/EventdrivenArchitecture.md)** - Design decoupled, reactive systems
2. **🌊 [Event Streaming](./concepts/EventStreams.md)** - Handle high-throughput data flows
3. **📈 [Scaling Microservices](./concepts/ScalingMS.md)** - Optimize for performance and reliability

### 🧠 **For Experts & Architects**
**Goal**: Navigate complex scenarios and understand trade-offs

1. **⚙️ [Advanced Concepts](./concepts/msConsiderations&Concepts.md)** - Master complex distributed system patterns
2. **⚠️ [Challenges & Solutions](./concepts/ChallengesWithMS.md)** - Navigate pitfalls and anti-patterns
3. **🏢 [Modular Monolith](./concepts/ModularMonolith.md)** - Know when NOT to use microservices

### 🛠️ **Hands-On Practice**
**Goal**: Apply concepts with real implementations

1. **🐹 [Go Implementation](./golang/README.md)** - Build production-ready gRPC microservices
2. **🟨 [Node.js Implementation](./nodejs/)** - Create modern TypeScript services
3. **🔄 Compare Approaches** - Understand different technology choices and trade-offs

## 💻 Practical Implementations

### 🐹 Go Microservices Platform

**E-commerce system with production-grade features:**

#### 🔧 Core Services
- **👤 Account Service**: User management with PostgreSQL storage
  - gRPC API for account CRUD operations
  - User authentication and validation
  - Pagination support for account listings

- **📦 Catalog Service**: Product management with Elasticsearch
  - Full-text search capabilities
  - Product CRUD operations
  - Advanced search and filtering

- **🛒 Order Service**: Order processing with PostgreSQL
  - Multi-product order creation
  - Automatic price calculation
  - Order history and tracking

- **🌐 GraphQL Gateway**: Unified API interface
  - Schema stitching across all services
  - Type-safe query resolution
  - Data aggregation and optimization

#### 🛠️ **Technology Stack**
- **Language**: Go 1.24.2
- **Communication**: gRPC with Protocol Buffers
- **API Layer**: GraphQL with gqlgen
- **Databases**: PostgreSQL + Elasticsearch
- **Containerization**: Docker & Docker Compose

#### 🚀 **Getting Started**
```bash
cd golang
docker-compose up -d
```

Access GraphQL Playground: `http://localhost:8080/graphql`

### 🟨 Node.js Microservices Platform

**Code snippet management platform with modern web stack:**

#### 🔧 Core Services

- **💬 Comments Service**: Comment system
  - Comment CRUD operations
  - Snippet-comment associations
  - User-based commenting

- **🎨 React Client**: Modern frontend application
  - React 19 with TypeScript
  - Tailwind CSS for styling
  - Radix UI components
  - Vite for fast development

#### 🛠️ **Technology Stack**
- **Backend**: Node.js + TypeScript + Express.js
- **Frontend**: React 19 + TypeScript + Vite
- **Styling**: Tailwind CSS + Radix UI
- **HTTP Client**: Axios for API communication
- **Development**: Hot reload with Nodemon

#### 🚀 **Getting Started**
```bash
# Start Snippet Service
cd nodejs/snippet
npm install && npm run dev

# Start Comments Service  
cd nodejs/comments
npm install && npm run dev

# Start React Client
cd nodejs/client
npm install && npm run dev
```

Access Frontend: `http://localhost:5173`

## 🛠️ Technologies

### 📚 Concepts & Documentation
- **Markdown**: Comprehensive documentation with diagrams
- **Visual Diagrams**: Architecture patterns and system evolution
- **Learning Modules**: Structured progression from basics to advanced topics

### 🐹 Go Implementation Stack
- **Language**: Go 1.24.2 (high-performance, concurrent programming)
- **Communication**: gRPC (HTTP/2-based RPC framework)
- **API Gateway**: GraphQL with gqlgen (type-safe API layer)
- **Databases**: 
  - PostgreSQL (ACID-compliant relational data)
  - Elasticsearch (full-text search and analytics)
- **Containerization**: Docker + Docker Compose
- **Libraries**:
  - `google.golang.org/grpc` - gRPC framework
  - `github.com/99designs/gqlgen` - GraphQL server generator
  - `github.com/lib/pq` - PostgreSQL driver  
  - `github.com/elastic/go-elasticsearch/v9` - Elasticsearch client
  - `github.com/segmentio/ksuid` - Unique ID generation

### 🟨 Node.js Implementation Stack
- **Runtime**: Node.js 18+ (JavaScript runtime)
- **Language**: TypeScript 5.8+ (type-safe JavaScript)
- **Backend Framework**: Express.js 5.x (web application framework)
- **Frontend Framework**: React 19 (component-based UI library)
- **Build Tool**: Vite (fast frontend build tool)
- **Styling**: 
  - Tailwind CSS 4.x (utility-first CSS framework)
  - Radix UI (accessible component primitives)
- **HTTP Client**: Axios (promise-based HTTP client)
- **Development Tools**:
  - ESLint (code linting)
  - Prettier (code formatting)
  - Nodemon (auto-restart development server)

## 🚀 Quick Start

### 📋 Prerequisites
- **Go**: Version 1.24.2 or higher
- **Node.js**: Version 18 or higher  
- **Docker**: Docker Desktop with Docker Compose
- **Git**: For repository cloning

### ⚡ Fast Setup (All Services)

```bash
# Clone the repository
git clone https://github.com/suryanshvermaa/microservices.git
cd microservices

# Option 1: Start Go E-commerce Platform
cd golang
docker-compose up -d
# Access GraphQL: http://localhost:8080/graphql

# Option 2: Start Node.js Snippet Platform
# Terminal 1: Snippet Service
cd nodejs/snippet
npm install && npm run dev

# Terminal 2: Comments Service  
cd nodejs/comments
npm install && npm run dev

# Terminal 3: React Client
cd nodejs/client
npm install && npm run dev
# Access Frontend: http://localhost:5173
```

### 📚 Learning-First Approach

```bash
# Start with concepts for theoretical foundation
cd concepts
# Read readme.md for complete learning path

# Then explore implementations
cd ../golang && cat README.md
cd ../nodejs && ls -la
```

### 🔧 Development Setup

```bash
# Go Development
cd golang
go mod tidy                    # Install dependencies
# Generate protobuf: protoc --go_out=. --go-grpc_out=. --proto_path=. *.proto
go run services/graphql/main.go # Run individual service

# Node.js Development  
cd nodejs/snippet
npm install                    # Install dependencies
npm run dev                    # Start development server
npm run build                  # Build for production
```

### 🌍 Environment Configuration

Create `.env` files in each service directory with appropriate configuration:

```bash
# Go Services Environment
DATABASE_URL=postgres://user:password@localhost:5432/dbname
ACCOUNT_SERVICE_URL=localhost:8080
CATALOG_SERVICE_URL=localhost:8080
ORDER_SERVICE_URL=localhost:8080

# Node.js Services Environment  
PORT=3000
NODE_ENV=development
```

## 📡 API Documentation

### 🐹 GraphQL API (Go Platform)

**Endpoint**: `http://localhost:8080/graphql`

The GraphQL gateway provides a unified interface with type-safe queries:

#### **Sample Queries**

```graphql
# Create an Account
mutation CreateAccount {
  createAccount(account: {name: "John Doe"}) {
    id
    name
  }
}

# Create a Product
mutation CreateProduct {
  createProduct(product: {
    name: "MacBook Pro M3"
    description: "High-performance laptop for developers"
    price: 2499.99
  }) {
    id
    name
    price
  }
}

# Create an Order
mutation CreateOrder {
  createOrder(order: {
    accountId: "2G8UdqGCWJ9ZCJHeLM4s1Z8DKfP"
    products: [
      { id: "2G8UeA1K5L2NxJ4QxBj8sK9DfR7", quantity: 2 }
    ]
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

# Query Accounts with Orders
query GetAccounts {
  accounts(pagination: {skip: 0, take: 10}) {
    id
    name
    orders {
      id
      totalPrice
      createdAt
      products {
        name
        quantity
      }
    }
  }
}

# Search Products
query SearchProducts {
  products(query: "MacBook", pagination: {skip: 0, take: 5}) {
    id
    name
    description
    price
  }
}
```

### 🟨 REST APIs (Node.js Platform)

#### **Snippet Service** (`http://localhost:3000`)

```bash
# Create a snippet
POST /api/v1/snippet
Content-Type: application/json
{
  "title": "React Hook Example",
  "description": "Custom hook for API calls",
  "code": "const useApi = () => { ... }",
  "userId": 1
}

# Get all snippets
GET /api/v1/snippet

# Health check
GET /health
```

#### **Comments Service** (`http://localhost:3001`)

```bash
# Create a comment
POST /api/v1/comment
Content-Type: application/json
{
  "comment": "Great snippet!",
  "userId": 1,
  "snippet_id": "abc123"
}

# Get comments by snippet ID
GET /api/v1/comment/:snippetId

# Health check
GET /health
```

#### **Frontend Client** (`http://localhost:5173`)
- Modern React application with TypeScript
- Snippet creation and viewing interface
- Responsive design with Tailwind CSS
- Real-time interaction with backend services

## 🐳 Deployment

### 📦 Containerized Deployment (Go)

```bash
# Quick start with Docker Compose
cd golang
docker-compose up -d

# Individual service builds
docker build -f services/account/app.dockerfile -t account-service .
docker build -f services/catalog/app.dockerfile -t catalog-service .
docker build -f services/order/app.dockerfile -t order-service .
docker build -f services/graphql/app.dockerfile -t graphql-gateway .

# Production deployment
docker-compose -f docker-compose.prod.yml up -d
```

### 🚀 Manual Deployment (Node.js)

```bash
# Production build
cd nodejs/snippet && npm run build
cd nodejs/comments && npm run build  
cd nodejs/client && npm run build

# Start production services
npm start  # For each service
```

### ☁️ Cloud Deployment Considerations

- **Container Orchestration**: Kubernetes, Docker Swarm
- **API Gateway**: NGINX, Kong, AWS API Gateway
- **Database**: Managed PostgreSQL, Elasticsearch Service  
- **Monitoring**: Prometheus, Grafana, ELK Stack
- **Security**: HTTPS, JWT authentication, service mesh

## 📚 Comprehensive Learning Resources

### 🎓 **Complete Course Structure** (`/concepts`)

The theoretical foundation provides a complete microservices education:

#### **🏗️ Module 1: Foundations & Evolution**
- **[Architectural Evolution](./concepts/TheArchitecturalEvolution.md)** - Journey from monoliths to microservices
- **[Business Requirements](./concepts/BuisinessRequirements.md)** - When and why to adopt microservices
- **[Stakeholder Views](./concepts/Views.md)** - Different perspectives on microservices architecture

#### **🔧 Module 2: Core Architectural Patterns**  
- **[Hexagonal Architecture](./concepts/HexagonalArc.md)** - Ports and adapters for clean architecture
- **[API Gateway Evolution](./concepts/APIGateways.md)** - From centralized to distributed gateways
- **[Decomposition Strategies](./concepts/decomposition.md)** - Breaking down monoliths effectively

#### **🔄 Module 3: Communication & Integration**
- **[Inter-Service Communication](./concepts/images/IntercommunicationService.md)** - Synchronous to event-driven evolution
- **[Event-Driven Architecture](./concepts/EventdrivenArchitecture.md)** - Design principles and best practices
- **[Event Streaming](./concepts/EventStreams.md)** - High-throughput data platforms
- **[Event Processing](./concepts/EventProcessing.md)** - Real-time vs batch processing

#### **📊 Module 4: Advanced Patterns & Scaling**
- **[Scaling Microservices](./concepts/ScalingMS.md)** - Performance and reliability optimization
- **[Message Queues](./concepts/Queues.md)** - Asynchronous messaging patterns
- **[Queues vs Streams](./concepts/images/QueueVSStreams.md)** - Choosing the right approach

#### **⚖️ Module 5: Challenges & Alternatives**
- **[Challenges with Microservices](./concepts/ChallengesWithMS.md)** - Common pitfalls and solutions
- **[Modular Monolith](./concepts/ModularMonolith.md)** - When simpler approaches are better
- **[EDA Key Takeaways](./concepts/EDATakeways.md)** - Event-driven architecture summary
- **[Advanced Concepts](./concepts/msConsiderations&Concepts.md)** - Complex scenarios and patterns

### 🛠️ **Key Concepts Covered**

**Architectural Patterns:**
- Hexagonal Architecture (Ports & Adapters)
- API Gateway Evolution
- Database per Service
- Strangler Fig Pattern
- Event Sourcing

**Communication Patterns:**
- Synchronous vs. Asynchronous
- Request-Response vs. Event-Driven
- Publish-Subscribe
- Message Queues vs. Event Streams

**Scaling Strategies:**
- Horizontal vs. Vertical Scaling
- Partitioning and Sharding
- Consumer Groups
- Backpressure Handling
- Service Discovery

**Operational Concerns:**
- Distributed Tracing
- Centralized Logging
- Circuit Breakers
- Retry Mechanisms
- Idempotent Processing

## 🎯 What I Learned

This personal project has been instrumental in deepening my understanding of distributed systems:

### 👨‍� **Technical Skills I Developed**
- **Go Programming**: Advanced concurrent programming and gRPC implementation
- **TypeScript/Node.js**: Modern backend development with type safety
- **React Development**: Component-based UI development with modern tooling
- **Database Design**: Polyglot persistence with PostgreSQL and Elasticsearch
- **Container Orchestration**: Docker and Docker Compose for microservices deployment

### �️ **Architectural Patterns I Mastered**
- **Hexagonal Architecture**: Clean separation of business logic from infrastructure
- **Event-Driven Design**: Asynchronous communication patterns and event streaming
- **API Gateway Patterns**: Unified interfaces and cross-cutting concerns
- **Database per Service**: Data ownership and consistency patterns
- **gRPC Communication**: High-performance inter-service communication

### � **Concepts I Internalized**
- **When to use microservices** vs simpler alternatives like modular monoliths
- **Trade-offs** between different architectural approaches
- **Scaling strategies** for distributed systems
- **Operational complexity** of microservices and mitigation strategies
- **Business alignment** of technical architecture decisions

### � **Personal Development Goals Achieved**
- Built production-ready microservices from scratch
- Created comprehensive documentation for complex systems
- Learned to balance theoretical knowledge with practical implementation
- Developed skills in multiple programming languages and frameworks
- Gained experience with modern development and deployment tools

## 🤝 About This Repository

This is my personal learning project, created to explore microservices architecture comprehensively. While it's designed for my own education and reference, others are welcome to:

### � **Study My Work**
- Review my concept notes and learning materials
- Examine my implementations and architectural decisions
- Use my documentation as a reference for your own projects

### 💻 **Run My Projects**
- Try out my Go e-commerce platform
- Experiment with my Node.js snippet management system
- Test different scenarios and configurations

### 🔄 **Build Upon My Work**
- Fork the repository for your own learning
- Extend my implementations with additional features
- Create your own microservices based on my patterns

### 🙏 **Acknowledgments for My Learning**
- **Martin Fowler & Sam Newman**: For foundational microservices concepts
- **Go Community**: For excellent tooling and libraries
- **React & TypeScript Teams**: For modern development experiences
- **Open Source Contributors**: For the amazing tools I used throughout this journey

## 📄 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

### 📜 **What this means:**
- ✅ **Commercial use** - Use in commercial projects
- ✅ **Modification** - Adapt and modify the code
- ✅ **Distribution** - Share and redistribute
- ✅ **Private use** - Use for personal projects
- ❗ **Include license** - Include original license in copies

## 🙏 Acknowledgments

### 🌟 **Community & Contributors**
- **Open Source Community**: For the incredible tools and libraries
- **Microservices Pioneers**: Martin Fowler, Sam Newman, and other thought leaders
- **Go Community**: For excellent concurrent programming patterns
- **Node.js Ecosystem**: For rapid development capabilities
- **React Team**: For modern frontend development patterns

### 📚 **Educational Resources**
- **Building Microservices** by Sam Newman
- **Microservices Patterns** by Chris Richardson  
- **Designing Data-Intensive Applications** by Martin Kleppmann
- **Go Concurrency Patterns** and **Node.js Best Practices**

## 📞 Support & Community

### 🆘 **Getting Help**
- **📖 Documentation**: Start with the [concepts guide](./concepts/readme.md)
- **🐛 Issues**: Report bugs on [GitHub Issues](https://github.com/suryanshvermaa/microservices/issues)
- **💬 Discussions**: Join conversations on GitHub Discussions
- **📧 Direct Contact**: Reach out for collaboration opportunities

### 🎯 **Next Steps**
1. **🚀 Start Learning**: Begin with [Business Requirements](./concepts/BuisinessRequirements.md)
2. **🛠️ Get Hands-On**: Try the [Go Implementation](./golang/README.md)
3. **🎨 Build UIs**: Explore the [Node.js Platform](./nodejs/)
4. **🤝 Contribute**: Share your improvements and learnings

---

**🌟 Star this repository if it helped you learn microservices architecture!**

> *"The best way to learn microservices is to build them. The best way to master them is to understand when not to use them."*

