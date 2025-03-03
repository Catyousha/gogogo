new things learned from each project:
* create
    * goroutine run along with `main` process. delay or waiting command in `main()` is necessary for holding on goroutine completion.

* multiple
    * multiple goroutine results order often vary.

* varGoroutines
    * `WaitGroup` is useful for tracking the number of active goroutines and ensuring each one completes before proceeding.
        * increase the counter with `Add(1)`.
        * mark each goroutine as completed using `Done()` (`Add(-1)`).
        * block main process until all forked goroutines finish (counter == 0) with `Wait()`.

* channels
    * channel (`make(chan int)`) used to communicate between goroutines
    * `<-` operator to sent & receive message through channel
        * sent: `aChannel <- value`
        * receive: `var1 := <- aChannel`
    * channel can be closed
        * send message to closed channel will throw panic err.
        * iterate (`for range`) channel to listen for message
            * listen on sleeping channel will cause deadlock.
    * channel in func param can be set to read or write only:
        * read only: `func (ch <- chan int)`
        * write only: `func (ch chan <- int)`

* select
    * use `select` to listen to multiple channels and act accordingly.
        * set timeout on listener to prevent deadlock

* wpools
    * channels and goroutine can be utilized to create worker pools
        * something something like divide and conquer

* mutex
    * lock shared var with `Mutex.Lock()` to prevent race condition
    * don't forget to unlock with `Mutex.Unlock()` to prevent deadlock

* rwmutex
    * locked value can be protected but allow reading by using `RWMutex.Lock()`
        * and set to writing only / prevent read by `RWMutex.RLock()`

* atomic
    * `atomic` can be used as alternative to mutex
    * all r/w operations with multiple goroutines cannot be interrupted.
    * write example: `atomic.AddInt64(&var, 1)`
    * read example: `atomic.LoadInt64(&var)`

* useContext
    * `context` package provides ways to control goroutine execution:
        * `WithCancel`: allows manual cancellation of context
        * `WithTimeout`: automatically cancels AFTER specified duration
        * `WithDeadline`: automatically cancels AT specified time
    * contexts can be cancelled in multiple ways:
        * manually via cancel function
        * timeout expiration 
        * deadline reached
    * use `select` with context's `Done()` channel to handle cancellation
    * always `defer cancel()` to prevent resource leaks

* keyVal
    * store context key in `ctx := context.WithValue(context.Background(), key, value)`
    * load with `ctx.Value(key)`

* semaphore
    * semaphore controls how many goroutines can execute simultaneously
    * unlike WaitGroup which just waits for all goroutines to complete
    * create a semaphore with `semaphore.NewWeighted(int64(maxWorkers))`
    * acquire semaphore units before starting a goroutine with `sem.Acquire(ctx, 1)`
    * release semaphore when goroutine completes with `sem.Release(1)`
    * wait for all workers to complete by acquiring all semaphore units `sem.Acquire(ctx, int64(maxWorkers))`

* wc
    * exercise 1-4