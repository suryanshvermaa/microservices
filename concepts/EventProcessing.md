# A Guide to Event-Driven Processing: Real-time Streams and Batching

Once events are flowing through a system, they need to be processed to be useful. Event processing is how a system responds to and acts upon the events it receives. This document explores the fundamental concepts of real-time event stream processing and contrasts this with the batch processing model.

## Part 1: Real-time Stream Processing
Real-time stream processing involves acting on events as they flow through the system, one by one or in small micro-batches.

### Core Stream Processing Operations
**Stateless Topologies**: This is the simplest processing model, where each event is processed independently without any information being retained from past events. This is highly efficient for parallel processing of simple, unrelated events (e.g., filtering events that don't match a specific criterion).

**Transformations**: These are operations that modify events in-flight to fit the needs of downstream consumers. Common transformations include:

- **Filter**: Selects events based on a specific criterion.

- **Map**: Transforms each event by applying a function.

- **Map Value**: Specifically targets and adjusts the value portion of a key-value event.

**Branching and Merging**:

- **Branching**: Splitting a single event stream into multiple streams based on certain conditions (e.g., branching a transaction stream into "high-value" and "regular" streams).

- **Merging**: Combining multiple separate streams into a single output stream (e.g., merging a payment stream and an order fulfillment stream to create a "completed transaction" stream).

### Partitioning Strategies for Scalability and Consistency
In distributed streaming systems, partitions are key to managing how data is processed.

**Repartitioning**: The process of redistributing events across partitions according to a new key. For example, if a stream was originally partitioned by transaction ID but now requires processing grouped by user ID, repartitioning allows for more efficient processing.

**Copartitioning**: A technique to ensure that events from different streams that are related (e.g., an order event and a payment event sharing the same order ID) are sent to the same partition. This is critical for simplifying correlated or "joint" processing.

**Consumer Partition Assignment**: To balance the workload, partitions need to be assigned to the available consumer instances. Common strategies include:

- **Round-Robin Assignment**: Evenly distributes partitions across all consumers for balanced load.

- **Custom Assignment**: Assigns partitions based on specific rules, such as data locality or consumer workload, which is useful when processing affinity is important.

- **Static Assignment**: Permanently assigns specific partitions to specific consumers. This is beneficial in stable environments where consistent processing by the same consumer is required.

### Common Reasons for Transforming Events
Events are rarely used "as-is" and are often transformed for a variety of reasons:

- **Data Enrichment**: Adding additional context to an event (e.g., adding customer details to an order event) to make it more useful for consumers.

- **Format Conversion**: Standardizing data formats, such as converting all timestamps to UTC for consistency.

- **Filtering, Reduction & Aggregation**: Removing unnecessary fields from an event to reduce its size, or summarizing a series of individual events into a single, aggregated event.

- **Schema Evolution**: Converting events to a new format or schema to maintain backward compatibility as the system evolves.

- **Security and Privacy**: Obfuscating or redacting sensitive information (like PII) before sending events to analytics services.

## Part 2: The Batch Processing Model
Batch processing is an alternative model where events are grouped together and processed as a single unit.

### What is Batch Processing?
Instead of processing events in real time as they arrive, this approach collects events and processes them together in a "batch." This is a common and highly efficient pattern in data processing pipelines where real-time responses are not a primary requirement.

### How it Works
Collecting Events: Events are gathered over a specified period or until a certain quantity is reached. They are typically stored in a temporary buffer or queue.

- **Batch Triggering**: The processing is triggered based on a defined rule, which can be time-based (e.g., every 10 minutes) or volume-based (e.g., every 1000 events).

- **Batch Processing**: Once triggered, the entire batch of events is passed to a processing system where transformations or other operations are applied to the group as a whole.

### Key Benefits and When to Use It
Batch processing offers several advantages, particularly related to efficiency and cost.

- **Key Benefits**: It reduces the overhead of handling each event separately, which improves overall performance, lowers the strain on resources (CPU, memory, I/O), and reduces the number of network requests.

### When to Use:

- When real-time responses are not required.

- When there is a large volume of data that can be processed at set intervals.

- When reducing costs and optimizing resource usage is a primary concern, especially in cloud-based systems with transaction-based pricing.