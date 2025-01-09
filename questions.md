Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> Concurrency is being able to handle multiple tasks at the same time, but not necessarly simultaniously. Parallelism is being able to compute the tasks simultaniously.

What is the difference between a *race condition* and a *data race*? 
> Race conditions: the program/result depends on timing the threads, access on data without proper synchronization --> unpredictable result
  Data race: threads access shared data concurrently, occurs when there is no synchronization (specific race condition)
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> Handles when and how to run threads


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> Multiple threads are used to allow a program to perform multiple tasks concurrently

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> Fibers or green threads are lightweight, user-space abstractions for managing concurrency, while coroutines are programming constructs that allow suspending and resuming functions at specific points. They are chosen for their efficiency 

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> Concurrent programming can make the programmer's life both easier and harder, depending on the context. It improves application performance, but introduces complexity

What do you think is best - *shared variables* or *message passing*?
> I think message passing is best, because it is easier and more efficient than shared variables.


