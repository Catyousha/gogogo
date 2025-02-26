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