# Temporal vs Kafka vs Job Queue [WIP]

## Distributed Queue Features

Producers can enqueue tasks and queues will accumulate tasks if consumers are slow or dead. Each task is delivered at least once (determined by visibility timeout at consumer). Consumers can report ack/nack, and nacked tasks are placed back in the queue.

If tasks are not ack'ed/nack'ed within a configured timeout (visibility timeouts in SQS), they are considered failed. Some queues support extending the running task timeout (aka heartbeat). Nacked tasks are re-delivered after some backoff. Some queues support DLQ for tasks that are nacked multiple times.

## Limitations of Distributed Queues

- No transactions between the queue and other data storages
- Maximum task execution time is limited even with heartbeat support
- Retry duration is limited (can't retry a task for hours)
- No task cancellation support
- Primitive error handling (DLQ is the only mechanism)
- No status tracking for specific tasks
- At-least-once execution semantics

## When to Use Task Queues

Task queues are good when:
- Tasks are stateless (no DB updates needed with queue operations)
- Tasks are idempotent and short-running
- Limited retry requirements
- Basic error handling with DLQ is sufficient
- Manual intervention for DLQ is acceptable
- No need for task status or cancellation
- Tasks are independent (no dependencies or chaining)

## Temporal vs Kafka Guidelines

Use Temporal when you need orchestration between microservices for complex business cases. The event-driven pattern with Kafka is more imperative.

### Temporal Advantages
- Clear workflow visibility and monitoring
- Easy to modify event sequences
- Better production monitoring and debugging

### Temporal Constraints
- Workflows must be deterministic
- Input/Output must be serializable
- Higher latency due to state management

### Kafka Long-Running Task Issues
- Default 5-minute poll timeout (broker reassigns partition after this)
- Zombie consumer problem
- Partition rebalancing challenges with long-running tasks
- Need to complete processing before partition revocation

## Why Choose Temporal Over Kafka

Kafka is better for simple payload transformations and computations. When dealing with complex stateful operations:

1. **State Management**:
   Temporal persists and manages workflow state, with automatic checkpointing and failure recovery.

2. **Execution Guarantees**:
   Exactly-once execution semantics, preventing duplicates or skipped operations.

3. **Fault Handling**:
   Built-in fault tolerance and automatic retries, continuing from last known state.

4. **Transactional Consistency**:
   Maintains consistency across external systems (databases, gRPC services) as a single unit of work.

## References
[When to use SQS vs Temporal](https://community.temporal.io/t/when-to-use-sqs/3689)