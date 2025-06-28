## Observability


### Observability Defenition

* How well we can understand the internals of a system based on its outputs.
* Today’s systems are big ball of mud, death star architecture.


### What is Telemetry data?

Telemetry data constitutes the following:
    * Logs
    * Metrics
    * Traces
* Application is instrumented and ships telemetry data to an observability back end (eg.., prometheus).
* From the observability backend query alerting, visualization and performance (eg.., grafana)


### What is OpenTelemetry?

* OpenTelemetry provides standards, protocols and SDK’s from CNCF
* Earlier every vendor had their own instrumentation tools, instrument sdk’s and agents. (Even though in java it is solved via Micrometer). Open telemetry solves this (OpenCencus + OpenTracing merged)


Critical Signals in Observability

* Logging: What happened (Why?) - Emitting events
* Metrics: What is the context? - Aggregated data (is it bad, is it good, how bad/good?)
* Distributed tracing: Why happened? - Recording causal ordering of events


### Micrometer in Java

* Micrometer is a facade for metrics similar to slf4j to logging.
* Multiple vendors have provided their MeterRegistry implementations.

Micrometer by default ships the following metrics from the java application:

* JVM utilization
* Memory and buffer pools
* Garbage collection stats
* Thread Utilization
* Number of classes loaded/unloaded
* CPU Usage
* Spring MVC and WebFlux request latencies
* RestTemplate latencies
* Cache Utilization
* Datasource utilization, including HikariCP
* RabbitMQ
* File Descriptor Usage
* Logback: record number of events logged to logback at each level
* Uptime: report a gauge for uptime and a fixed gauge representing the application’s absolute start time
* Tomcat usage



### Metric and Alerting Criterias

#### Apdex Score

- Ref: apedex_score.md
- Current Apdex score threshold across all web transactions is 0.5s
- Apdex score can be sent an individual web txn level.

Possible Metrics and Alerts

Temporal Alerting Conditions

* temporal_workflow_schedule_to_start_latency
    * Monitored at worklow and activity level, dimension - worker_type
    * Alert if passes beyound 2x/3x threshold.
* :


Database
* JDBC connection wait time
* JDBC connection usage time (query timing)
