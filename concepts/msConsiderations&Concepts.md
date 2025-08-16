# Some More Concepts and Considerations in Event-Driven Architecture
Beyond the basics of producers, consumers, and brokers, building robust and effective event-driven systems requires understanding a set of advanced concepts, processing strategies, and practical design trade-offs. This document covers these key considerations.

## 1. Fundamental Concepts of Asynchronous Systems
The asynchronous nature of EDA gives rise to several powerful but sometimes counter-intuitive concepts.

**Deferred Execution**: Unlike a synchronous API call, when you publish an event, you do not wait for a response. The execution of the logic that consumes the event is deferred—it will happen at some point in the future. This decouples the producer from the consumer's logic and performance.

**Eventual Consistency**: This is a direct consequence of deferred execution. Since you don't wait for confirmation that an event has been processed, you cannot guarantee that all parts of your system have the exact same state at the exact same time. You can only be sure that, given a period of time, the system will eventually become consistent.

**Event Mesh**: An Event Mesh is a configurable and dynamic infrastructure layer of interconnected event brokers. It allows for the routing of events between decoupled applications, regardless of where they are deployed (cloud, on-premises, or IoT devices). It creates a "nervous system" for the enterprise.

## 2. Advanced Event Processing Strategies
Handling a continuous flow of events in a distributed system requires specific strategies.

**Event Scheduling**: This involves controlling the timing of events. You can either delay the emission of events until certain conditions are met or schedule events to be processed based on predefined criteria (e.g., using cron jobs or message broker delay mechanisms to trigger price-checking events at the end of each trading day).

**Deterministic Event Processing**: The goal is to ensure that a given set of events always produces the same result. This is challenging due to network latency and concurrent processing. Key techniques to achieve this include using sequence numbers (like Kafka's offsets) to guarantee order and designing consumers to be idempotent (so that reprocessing the same event multiple times has no additional effect).

**Handling Late Events**: In real-time systems, events can arrive after their expected time due to network delays or system failures. Strategies to manage this include:

- **Watermarks**: A mechanism that defines a point in time after which late events are either ignored or processed differently.

- **Windows**: Using time-based windows (e.g., a 5-minute window) can accommodate events that are slightly late, as they are aggregated based on their event timestamp rather than their arrival time.

**Reprocessing vs. Real-Time Processing**: This is a critical trade-off:

- **Real-Time Processing**: Involves processing events as they arrive for near-instantaneous feedback (e.g., fraud detection). It ensures low latency but is resource-intensive and may miss details if late events are dropped.

- **Reprocessing**: Useful for batch-oriented systems or when historical data needs to be recalculated (e.g., after a code bug fix in an analytics pipeline). It ensures accuracy and completeness but comes at the cost of time and computational resources.

## 3. A Balanced View: Pros and Cons of Event-Driven Microservices
Applying EDA to a microservices architecture offers significant benefits but also introduces considerable challenges.

- **Benefits**: The architecture is inherently event-driven, extremely available, highly efficient, loosely coupled, and supports technological flexibility and continuous delivery. Services are also highly maintainable and testable due to their small, focused nature.

- **Drawbacks**: The system can be super complex and very difficult to architect correctly. It is often very resource-intensive, and failure identification can be a major challenge in such a distributed environment. A centralized main event stream can also become a security threat if not properly secured.

## 4. Practical Design and Implementation Considerations
When building event-driven microservices, several practical factors must be considered:

- **Single Writer Principle**: To maintain data authority and consistency, only the service that owns a particular piece of data should be allowed to emit events about changes to that data. For example, Order events should only be emitted by the order-microservice.

- **Event Broker Selection**: The choice of event broker is critical and should be based on its ability to provide:

  - Scalability

  - Durability

  - High Availability

  - High Performance

**Additional Factors**: You must also consider the availability of support tooling, hosted (managed) services, quality of client libraries and frameworks (e.g., Go kit, Molecular), community support, and options for long-term and tiered storage.

**Performance Considerations**: To ensure high performance, consider strategies for isolating internal data models, ensuring schema compatibility as services evolve, using database triggers for capturing change data, and efficiently sinking event data to other data stores.