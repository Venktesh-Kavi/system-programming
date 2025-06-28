# Temporal References [WIP]

### Dependency Injection
* [Why avoid DI in workflows and activities?](https://community.temporal.io/t/temporal-with-springboot/2734/14)
  * Workflow state must be deterministic
  * External dependencies can cause non-determinism
  * Best practices for Spring integration

### Workflow Management
* [Managing Workflows with Same IDs](https://community.temporal.io/t/execute-a-workflow-with-the-same-workflowid-with-another-one-which-is-having-workflowtasktimeout/4497)
  * Handling duplicate workflow IDs
  * Analyzing workflow history
  * Timeout management

## Performance Monitoring

### Metrics and Latency
* Schedule-to-Start Latency:
  * [High latency troubleshooting](https://community.temporal.io/t/workflow-task-schedule-to-start-latency-high/5238)
  * [Contributing factors](https://community.temporal.io/t/what-factors-count-for-workflow-schedule-to-start-latency/5829)
  * [Java SDK latency issues](https://community.temporal.io/t/very-big-schedule-to-start-workflow-latency-java-sdk/2733)

## Advanced Features

### Context Propagation
* [Spring Sleuth Integration](https://community.temporal.io/t/java-sdk-passing-spring-sleuth-traces/3318/3)
  * Tracing in Temporal workflows
  * Sleuth context propagation

### Thread Context Management
* [Workflow Thread Context](https://community.temporal.io/t/proper-way-to-set-workflow-thread-context/10055)
  * Best practices
  * Implementation examples

### Implementation Examples
* [Java SDK Samples](https://github.com/Quinn-With-Two-Ns/samples-java/commit/c99e24af5f20d1965827c5e178f6541a9b5ddf12#diff-41c2e5762a1afb781edd4dda97a55825dd23846364addcf5f7b9ff7c165186bf)
  * Context propagation
  * Interceptor implementation
  * Thread management
