# Server Concurrency, JVM and Tomcat Trivia

## Concurrency

* Multi-processes get time slices of execution time, managed by the CPU scheduler through context switching
* Core-Process relationship:
  * A CPU core can handle multiple processes
  * A process typically runs on a single core/CPU
  * A process can spawn multiple threads across multiple cores

### Time Slicing vs Context Switching
* **Context Switching**: Process of saving and restoring the state of a running process
  * Saves current process state for later restoration
  * Loads state of another process/thread
  * Enables efficient multitasking

### CPU Process Handling
* **Dual Core CPU Example**:
  * Base level: 2 cores = 2 simultaneous processes
  * With SMT/Hyperthreading: 4 threads (2 per core)
    * Each core still runs one instruction set
    * Thread 1 executes while Thread 2 prefetches instructions
  * OS Level:
    * Can manage arbitrary number of active processes
    * Limited by available memory
    * Only 4 can be in CPU, 2 actually running

> [Reference](https://www.quora.com/How-many-processes-can-run-at-once-on-one-CPU-core)

### JVM and Multi-core Usage
* [JVM on Single Core](https://softwareengineeringexperiences.quora.com/How-does-a-JVM-running-on-a-single-core-of-a-multi-core-processor-make-effective-use-of-other-cores)
* [Java Threads Overview](https://medium.com/@ja.m.arunkumar/java-threads-part-1-7855b11ddb6)

## Parallelism
* Achieved when tasks are subdivided and distributed across multiple cores
* Enables true concurrent execution

## Web Server Types

### Blocking Web Server (Tomcat)
* Characteristics:
  * Assigns one thread per request from thread pool
  * Thread handles entire request-response lifecycle
  * Default 200 threads in Tomcat
* Challenges:
  * Blocking operations cause thread waiting
  * Heavy IO can degrade system performance
  * Request queuing under high load
  * Higher memory requirements

### Non-Blocking Web Servers (Netty, Node, Go)
* Features:
  * Supports TCP, UDP with back pressure compatibility
  * Single non-blocking thread for request processing
  * Channel-based client-server communication
  * Event-loop architecture
* Benefits:
  * Efficient resource utilization
  * Better scalability for request handling
  * Delegated I/O operations

#### Netty Server Architecture
* **Thread Groups**:
  1. Boss Group:
     * Handles channel establishment
     * Manages event loop (typically single-threaded)
  2. Worker Group:
     * Decodes socket data
     * Forwards to executor threads
  3. Application Executor Group:
     * Performs IO operations
     * Thread Pool Configuration:
       * Default: CachedThreadPool (not recommended for production)
       * Recommended: FixedThreadPool

### Further Reading
* [Spring WebFlux Netty Internals](https://www.linkedin.com/pulse/spring-webflux-under-hood-diego-lucas-silva/)
* [Netty Server Tuning](https://java.msk.ru/experience-in-webflux-netty-highload-optimization/)
* [Netty Server Builder Details](https://groups.google.com/g/grpc-io/c/LrnAbWFozb0)

## Hardware Architecture

### MultiCore vs MultiProcessor
* **MultiCore**:
  * Common in modern home PCs
  * Single integrated chip with multiple cores
  * Each core is an execution unit

* **MultiProcessor**:
  * Found in servers and workstations
  * Multiple CPU chips
  * Each chip has multiple cores
