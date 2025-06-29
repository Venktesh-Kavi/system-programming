# Spring Kafka POC

This POC is intended to understand the following concepts:

1. Offset management by spring kafka consumer - (auto commit mode)
2. Consumer poll operations behavior
3. Idempotency during I/O operations in kafka consumer during failures.
4. Failure handling and retries

## Setup

Kafka Docker - Has kafka set up using KRaft protocol and kafka-ui is bundled.
```bash
docker-compose up
```

`make build`: Build the project
`make run`: Run the project

Testing with multiple events:

```bash
ab -n 1 -c 1 -p kafka_payload.json -T 'application/json' \                                                                                 
   -H "Accept: application/json" \
   -H "Origin: http://localhost:8080" \
   "http://127.0.0.1:8080/api/clusters/local/topics/test/messages"
```