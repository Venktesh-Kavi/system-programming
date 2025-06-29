## Count Down Latch

### What is Count Down Latch

- Count down latch is mechanism by which the calling thread/client is made to await till all the threads performing some set of operation are complete. The threads performing the operations reduce the latch count, the waiting thread/client call await() waiting for completion.
- No return type can be captured in count down latch.
- Typically used for a client to wait for varied type of parallel operations to complete.

### When to use?

- Use when we want to wait for hetergoneous tasks to complete.
- Use when we want to coordinate a thread which we didn't create or control.
- Acts as a one-time barrier that multiple threads must reach before processing.

### Differences from futures from executor service:

- Executor Service managed thread creation and lifecycle management of a thread for you.
- Futures via executorService provides way to wait for a result.
- The executorService also provides constructs to wait for multiple futures till all of them resolve or any of them resolve. (invokeAll)

### Choose CountDownLatch when:

1. You're coordinating with external systems or threads you don't control
2. You need a simple one-time synchronization point
3. You're dealing with heterogeneous tasks that don't fit a uniform execution model
4. You need fine-grained control over thread creation and management

### Choose ExecutorService when:

1. You need to process a collection of similar tasks
2. You need to collect and process results from tasks
3. You want managed thread pools and lifecycle management
4. You need advanced features like timeouts, scheduling, or task dependencies

- Task Homogeneity: invokeAll() works best with a collection of similar tasks (all Callable with the same return type). CountDownLatch is more flexible for heterogeneous operations.
- Task Creation: With invokeAll(), you need to create all tasks upfront. CountDownLatch can be used with threads/tasks that are created at different times or by different components.
- Granular Control: CountDownLatch gives you more fine-grained control over when each task signals completion. With invokeAll(), tasks are considered complete when they return or throw an exception.
- External Thread Integration: If you need to coordinate with threads you don't create (like callback handlers from external libraries), CountDownLatch is easier to integrate.
- Partial Success Handling: With CountDownLatch, you can implement custom logic for what happens when some but not all tasks complete. With invokeAll(tasks, timeout), if the timeout expires, you get the current state of all futures, but some may not be done.


### Reference Code

[CountDownLatchPoc.java](/Users/venktesh.k/workspace/personal/system-programming/pocs/countDownLatch/CountDownLatchPoc.java)
