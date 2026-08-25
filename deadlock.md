# Deadlock Scenarios in Go Concurrency
https://medium.com/@AlexanderObregon/deadlock-scenarios-in-go-concurrency-2f628d5e4d15

Alexander Obregon, Sep 3, 2025

Go makes it possible to launch massive numbers of lightweight goroutines that run independently without the heavy cost of threads. These goroutines talk to one another through channels and shared synchronization tools. While the model feels simple, situations can arise where every goroutine ends up stuck waiting. That state is called a deadlock, and it means nothing can move forward, leaving execution frozen. Go stands out because its runtime has built-in checks for deadlocks. If all goroutines are blocked and no work is left to be done, the runtime stops the process and reports the deadlock.

I publish free articles like this daily, and I have a [Golang section on Substack](https://alexanderobregon.substack.com/p/browse-my-go-articles-by-topic) where you’ll find my whole series. My Substack also includes weekly recaps if you’d like to keep up with everything I’m publishing.
## How Deadlocks Form in Go
Deadlocks appear when goroutines end up waiting in ways that make progress impossible. Go’s concurrency model is designed so that goroutines naturally pause when they need input or when a lock is unavailable. This blocking behavior is normally what allows them to coordinate, but it also means there’s a risk of every goroutine freezing in place. To see how this happens, it helps to look at some common scenarios where blocking turns into a full deadlock.
### Goroutines Blocking States
A goroutine can be running, runnable, or blocked. Running means it’s currently executing, runnable means it’s ready but waiting for a processor, and blocked means it’s stuck on something like a channel read, channel send, or synchronization primitive. Blocking is normal and temporary when another goroutine is active to complete the operation, but if all active goroutines end up in blocked states, nothing can continue.
### Single Channel Waits
One of the most direct cases is a receive from a channel that never gets a value. If a goroutine expects something to arrive but no one is scheduled to send, the process halts right there.
``` Go
func main() {
    ch := make(chan int)
    value := <-ch
    fmt.Println(value)
}
```
This short example halts because no sender exists. The receive operation never completes, so the main goroutine never exits. Since there are no other goroutines running, Go’s runtime reports a deadlock. Deadlocks can also appear with buffered channels when the buffer is already full and no reader is present:
``` Go
func main() {
    ch := make(chan int, 1)
    ch <- 1
    ch <- 2 // blocked here because the buffer is full
}
```
The second send never goes through because the buffer has no free space and there’s no receiver.
### Mutual Waiting Between Goroutines
Deadlocks often emerge when goroutines wait for each other in a cycle. One blocks waiting for data that another plans to send later, but that second goroutine is itself waiting for the first to act. Both are stuck.
``` Go
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        msg := <-ch1
        ch2 <- msg
    }()

    msg := <-ch2
    ch1 <- msg
}
```
Neither side can progress here. The spawned goroutine needs a value from `ch1` first, while the main goroutine waits on `ch2`. Both are frozen at their first channel operations.
### WaitGroups Without Producers
`sync.WaitGroup` is a convenient way to pause until a set of goroutines finish. The pattern works only if the `Add` and `Done` calls match correctly. Forgetting to call `Done` leaves the counter stuck above zero, which means the `Wait` call blocks forever.
``` Go
func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        fmt.Println("worker 1 done")
    }()

    // worker 2 never calls Done
    go func() {
        fmt.Println("worker 2 stuck")
    }()

    wg.Wait()
    fmt.Println("all workers finished")
}
```
The second worker fails to reduce the counter. The main goroutine waits forever, and if no other goroutines remain runnable, the runtime declares a deadlock. A small oversight like this is enough to freeze everything. It’s also possible to create a deadlock by mismatched increments, such as calling `Add(1)` too late after a worker has already completed. This causes the counter to rise and never fall back to zero.
### Locks That Never Release
Mutexes are a core tool for controlling shared memory access. They’re simple to use but easy to misuse in ways that lock out progress. If a mutex is locked and never unlocked, any future attempt to acquire it will block forever.
``` Go
func main() {
    var mu sync.Mutex
    mu.Lock()

    // Forget to call mu.Unlock()

    go func() {
        mu.Lock() // stuck forever
        fmt.Println("never reached")
    }()

    time.Sleep(2 * time.Second)
}
```
The goroutine in this code will never print because the lock was never released. This leaves every subsequent lock call blocked permanently. A subtle case where this also happens is when a goroutine tries to re-lock the same mutex it already holds:
``` Go
func main() {
    var mu sync.Mutex
    mu.Lock()
    fmt.Println("locked once")
    mu.Lock() // deadlock because Mutex is not reentrant
}
```
Because Go’s `Mutex` isn’t reentrant, trying to acquire it again without unlocking causes the goroutine to block itself. This is a guaranteed deadlock, even with only one goroutine involved.
### Cycles Across Multiple Resources
The most complex deadlocks appear when multiple resources are locked in different orders by different goroutines. Each goroutine holds one resource and waits for another, creating a cycle that none can escape.
``` Go
func main() {
    var mu1, mu2 sync.Mutex

    go func() {
        mu1.Lock()
        time.Sleep(500 * time.Millisecond)
        mu2.Lock()
        fmt.Println("goroutine 1 got both locks")
    }()

    go func() {
        mu2.Lock()
        time.Sleep(500 * time.Millisecond)
        mu1.Lock()
        fmt.Println("goroutine 2 got both locks")
    }()

    time.Sleep(2 * time.Second)
}
```
The first goroutine locks `mu1` and waits for `mu2`, while the second locks `mu2` and waits for `mu1`. Both remain stuck forever with no release path. This classic lock-ordering problem is a textbook deadlock case.

Cycles can also happen across more than two resources. Imagine three goroutines each holding a different mutex and each waiting for the next one in a chain. None can progress, and the deadlock scales with the number of participants.
## How the Go Runtime Detects and Reports Deadlocks
Detection of deadlocks in Go is part of the runtime scheduler itself. The runtime keeps constant track of every goroutine and its state. That awareness is what makes it possible to declare that no progress is left to be made. When all user goroutines are blocked and there are no external events like timers or system calls still active, Go ends execution with a clear diagnostic message.
### Scheduler Tracking of Goroutine States
Each goroutine is managed through an internal structure called a G, which holds fields describing its status. A goroutine can be marked as running, runnable, waiting, or blocked. The scheduler maintains queues of runnable goroutines and decides which ones get CPU time. When a goroutine blocks, it’s removed from the runnable queue and recorded so the runtime doesn’t waste effort trying to schedule it.

A channel receive is a clear way to see this. When a goroutine calls `<-ch` on an empty channel, the runtime marks it as waiting. If another goroutine sends into that channel, the blocked goroutine is moved back into the runnable queue.
``` Go
func main() {
    ch := make(chan int)

    go func() {
        fmt.Println(<-ch) // blocked until a value arrives
    }()

    time.Sleep(2 * time.Second)
}
```
With no sender present, the spawned goroutine never changes state. The scheduler still tracks it, but it remains in waiting status. If every goroutine ends up in that condition, the runtime detects the freeze.

The same rules apply to mutexes and wait groups. Any blocked state is recorded, and the scheduler only moves those goroutines forward when an event makes them ready again.
### Reporting Deadlocks
When Go concludes that no progress is possible, it prints an error that begins with `fatal error: all goroutines are asleep - deadlock!`. This message is followed by stack traces for each goroutine, showing the exact lines where they became stuck. That output makes debugging practical because it directs attention to the source of the freeze.
``` Go
func main() {
    ch := make(chan int)
    <-ch
}
```
Running this short case produces output similar to:
``` bash
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [chan receive]:
main.main()
    /tmp/deadlock.go:6 +0x39
```
The report shows the call stack of the main goroutine, which is blocked on a channel read. For more complex situations, multiple goroutines and their stacks appear together. That gives a wider view of which routines were waiting and on what resources.

Sometimes the output reveals mistakes that weren’t obvious, like a missing `Done` call on a wait group or a forgotten `Unlock` on a mutex. Seeing the exact file and line number makes it easier to identify where progress stopped.
### Special Case of Background Goroutines
Not every goroutine comes from user code. The runtime starts its own in the background to handle things like garbage collection, timers, and network polling. Their activity doesn’t prevent a deadlock from being reported. The runtime only calls a deadlock when no user goroutines can make progress and nothing is left to wake them up. That’s why not every freeze ends with a fatal error. A common case is when the `main` goroutine finishes while others are still active. The process stops right away because Go treats the return from `main` as the end of execution, even if background workers are still running. No deadlock message appears because the stop was caused by `main` ending, not by blocking.
``` Go
func main() {
    go func() {
        for {
            fmt.Println("worker still active")
            time.Sleep(1 * time.Second)
        }
    }()
    // main exits right away
}
```
This doesn’t report a deadlock. The worker goroutine runs for a short while, but the process ends as soon as the main function returns. Runtime goroutines like those for garbage collection don’t change that outcome because they’re not part of the pool that needs to make forward progress in user code.
## Practical Debugging with GODEBUG
The `GODEBUG` environment variable offers a way to see scheduler activity in real time. Setting `GODEBUG=schedtrace=1000` tells the runtime to print scheduling information every second. The output shows the number of goroutines that are runnable, waiting, or blocked, along with details about processors and queues. It provides a look into how the runtime itself sees the state of execution.

    GODEBUG=schedtrace=1000 go run main.go
The output looks like this:
``` bash
SCHED 1001ms: gomaxprocs=4 idleprocs=4 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0]
```
These lines repeat every second, showing how many goroutines are waiting and whether the queues are empty. If the run queues stay at zero while goroutines are stuck, it’s clear that no forward progress will happen. There are other flags such as `GODEBUG=scheddetail=1` that show more detail on scheduling choices. Although more cluttered, that view can help confirm why certain goroutines never resume.

When combined with the runtime’s deadlock error message, the `GODEBUG` traces form a complete picture. You can see the gradual slide toward a freeze and then get the exact lines of code where everything stopped.
## Conclusion
Deadlocks in Go happen through blocking mechanics that leave no path forward. Every channel wait, lock request, or wait group hold is tracked by the runtime. When every goroutine ends up stuck and no event remains to trigger progress, the scheduler steps in and reports the freeze. That built-in check, together with stack traces and tools like `GODEBUG`, shows how Go not only runs concurrent work but also gives a clear window into why it stopped.
