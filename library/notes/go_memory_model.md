---
id: 20260522-go_memory_model
type: concept
status: draft
tags:
  - go
  - programing_language
created: 2026-05-22
---
## TL;DR
Largely automatic through garbage collection. Runtime decides stack (fast, auto-cleaned) vs heap (slower, GC required) allocation via escape analysis. Understanding allocation patterns and avoiding memory leaks helps write efficient, scalable Go programs.
# Go Memory Model
Specifies conditions under which reads of a variable in one [goroutine]() can be guaranteed to observe values produced by writes to the same variable in a different one.

Data race is defined as a write to a memory location happening concurrently (separate goroutine) with another read or write to that same location, unless the accesses involved are atomic data accesses as provided by the [`sync/atomic`](https://pkg.go.dev/sync/atomic) package.
In the absence of data races, Go programs behave as if all the goroutines were multiplexed (taking turns very quickly) onto a single processor. This makes programs data-race-free and execute in a sequentially consistent manner (DRF-SC).

>Programmers should write Go programs without data races. Go expects you to use [Channels](), [Mutexes](), [Atomics]() and [Synchronization]() to avoid this.

Even if a race occurs a program is guaranteed to see an actual value. It's not like in C where you might get random garbage memory or messed values. This is only guaranteed for small memory values like `int32`, `pointers`, `booleans`, or `bytes`, for larges structures, torn reads can still happen.

Go chooses to act like this during races to make programs easier to debug. If it had chosen to behave like C with undefined behavior on races, debugging would be a nightmare.
## Memory Model
goroutine executions are made up of memory operations which are modeled by four details:
- Kind: Whether it is an ordinary data read or write, or a synchronizing operation such as an atomic data access, mutex or channel operation.
- Location: In the program.
- Memory location: The address of the memory being accessed.
- Value: The values read or written by the operation.

Some operations are read-like, others are write-like and some such as atomic compare-and-swap, are both read-like and write-like.

>A goroutine execution is modeled as a set of memory operations executed by a single goroutine.

Bellow we have the 3 requirements for data-race-free program execution which are implemented by the Go programming language.
### Requirement 1
Each goroutine in a Go program should behave as if it is running normally, one step at a time, following the rules of the Go language.

For example:
```go
var x int = 0

go func() {
    x = 1
}()

go func() {
    fmt.Println(x)
}()
```

The read of `x` in `fmt.Println(x)` must get its value from some actual write to `x`.
Possible results:
- `0` → the read saw the original value
- `1` → the read saw the concurrent write
But the read cannot produce a made-up value like `42`.

Different runs of the same program may produce different valid executions depending on timing and scheduling.
### Requirement 2
Defines the expected concurrent memory access model in Go.

All synchronizing operations must fir into one imaginary global order that everyone agrees on. Synchronizing operations are:
- [mutex lock/unlock]()
- [channel send/receive]()
- [atomic operations]()

Synchronizing operations creates a guaranteed ordering. Example:
```go
// goroutine 1
x = 1
mu.Unlock()

// goroutine 2
mu.Lock()
print(x)

// Unlock synchronized before Lock.
```
#### Synchronized before
If a synchronizing read sees a synchronizing write, then the write is synchronized before the read.
```go
// goroutine 1
x = 1
ch <- true //send

// goroutine
<- ch      //receive
print(x)
```

Since the receive observes the send (because of [channel locking]()), so: `send synchronized before receive`.
#### Sequenced before
Normal order inside one goroutine:
```go
x = 1
y = 2

// x is sequenced before y.
```
#### Happens before
Combines **sequenced-before** and **synchronized-before** to follow chains transitively.

So if we look at the channels example above we get:
```plaintext
x=1 sequenced before send

send synchronized before receive

receive sequenced before print(x)

```

Therefore: `x=1` happens before `print(x)`.
### Requirement 3
Determines the types of data races and which write a read is allowed to see based on **happens-before** ordering.

For a normal read `r` of variable `x`, the value must come from a write `w` that's visible to that read, where visible means:
- The write happened before the read.
- There isn't a newer write (that also happened before the read).

>A read sees the most recent write that happened-before it.
#### Data Race
Occurs when two operations touch the same memory location and Go cannot establish an ordering between them. Neither operation is known to happen before the other.

**read-write race**
```go
var x int

go func() {
	x = 42
}()

fmt.Println(x)
```

Since there's no synchronizing operations, Go cannot determine if read happened before write, or if write happened before read.

**write-write race**
```go
var x int

go func () {
	x = 1
}()

go func() {
	x = 2
}()
```
Same as before but since the writes are not synchronized, this is a write-write race.

If there are no data races, then any read can only see one possible write. If your program is DRF, you can pretend that all goroutines took turns executing one operation at a time in some global order. Race-free programs behaves as if execution happened in a single interleaving.

```go
var x int
done := make(chan struct{})

go func() {
	x = 1
	close(done)
}()

<-done
fmt.Println(x)
```
In this example, write is synchronized before read by using a [channel](), making the program DRF-SC.

Any implementation can, upon detecting a data race, report the race and halt execution of the program. Implementations using `ThreadSanitizer` (accessed with `go build -race`) do exactly this.
### Torn data in large data structures
For performance reasons, Go might break a large write or read operation into a bunch of smaller, independent, word-sized (Natural amount of data a CPU can handle in a single action, default to 8 bytes in modern CPUs) steps, and the order isn't guaranteed.

If one goroutine is trying to write a large piece of data, and another goroutine tries to read it at the exact same millisecond without safety gear (a race condition), the reader might catch the writer halfway through. This results in **"torn" data**.
### Example: The Multi-Word Race
```go
// goroutine A
name := []string{"John", "Doe"}

// goroutine B
go func() {
	name[0] = "Jane"
	name[1] = "Smith"
}()

// goroutine A continuation
fmt.Println(name)
```

Because Go doesn't guarantee that this happens all at once, **Goroutine A** might read the memory right in the middle of the swap and end up with:
- `("Jane", "Doe")` or `("John", "Smith")`

Neither of those combined values was ever a valid "single write."
### Example: Disastrous Slice Race
Go strings, slices and maps are actually descriptors that contain many parts. In the case of a Slice, it's a 3-word structure consisting of:
- A pointer to the actual data.
- The length.
- The capacity.

```go
// goroutine A
nums := []int{1, 2}

// goroutine B
go func() {
	nums = append(nums, make([]int, 1000))
}()

fmt.Println(nums[500])
```

If Goroutine B updates the Slice while Goroutine A tries to read it, Goroutine A might mix up the pieces. It might read the new length, but the old pointer.

The computer tries to read item 500 that's way after the end of the 2 items original array. It blindly reads random computer memory. This leads to arbitrary memory corruption which results in a immediate crash (segmentation fault) or silently ruins other unrelated data in your app.
## Synchronization
### Initialization
Program initialization runs in a single goroutine, which can create other goroutines, which run concurrently.

Each package that is imported has its init functions completed before any of the importing package's.

The completion of all init functions is synchronized before the start of the main function.
### Goroutine creation
The `go` statement that starts a new goroutine is synchronized before the start of the goroutine's execution.

This means that before a function executed with the `go` statement is executed, the function must be placed in a new goroutine, meaning that the function execution itself will happen at some point in the future and not necessarily right after the `go` statement.

```go
var a string

func f() {
	print(a)
}

// Caling hello() will start a new goroutine, but the print statement will be executed at some point in the future, perhaps after hello() has returned.
func hello() {
	a = "hello, world!"
	go f()
}
```
### Channel communication
Main method of synchronization between goroutines. Each send is matched to a corresponding receive from that channel, usually in a different goroutine.

A send on a channel is synchronized before the completion of the corresponding receive from that channel.
```go
var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world" // sequenced before send on c
	c <- 0             // synchronized before corresponding receive on c
}

func main() {
	go f()
	<-c                // sequenced before print
	print(a)           // guaranteed to print "hello, world"
}
```

Closing a channel is synchronized before a receive that returns a zero value because the channel was closed (when a channel is closed, any goroutine waiting on it will immediately stop waiting and receive a zero value).
In the previous example, replacing `c <- 0` with `close(c)` yields the same guarantees.

*a receive on an unbuffered channel is synchronized before the completion of the corresponding send on that channel*
By swapping `c <- 0` and `<-c` and making `c = make(chan int)`, the same guarantees applies.
This happens because unbuffered channels cannot hold data, so, for a transfer to happen, the sender and receiver must communicate at the same instant (one freezes until there's response from the other).
If the channel was buffered, there wouldn't have any need to wait and the guarantee of printing "hello, world" wouldn't be enforced.

The `kth` receive on a channel with capacity *C* is synchronized before the completion of the `k+Cth` send on that channel.
So, a receive (`k = 1`) on an unbuffered channel (`C = 0`)  is synchronized before the completion of the (`k+C -> 1+0`) send on that channel.

This rule allows a counting semaphore (rate limiting) to be modeled by a buffered channel. The number of items in the channel correspond to the number of active uses, the capacity is the maximum number of simultaneous uses. Sending an item acquires the semaphore, and receiving an item releases the semaphore. It's a common idiom for limiting concurrency.
### Locks
For any `sync.Mutex` or `sync.RWMutex` variable *l* a call *n* of `l.Unlock` is *synchronized before* call *m* of `l.Lock()` returns. This is true for `n < m`.

```go
var l sync.Mutex
var a string

func f() {
	a = "hello, world"
	l.Unlock()
}

func main() {
	l.Lock()
	go f()
	l.Lock()
	print(a)
}
```

*"hello, world"* is guaranteed to be printed. First call to `l.Unlock()` is *synchronized before* the second call to `l.Lock()` returns, which is sequenced before `print(a)`.
This happens because you can have a call to `l.Lock()` before `l.Unlock()` is called.

For any call to `l.RLock`, there is an *n* such that the *nth* call to `l.Unlock` is *synchronized before* the return from `l.RLock` (writes are always synchronized before reads). And the matching call to `l.RUnlock` is synchronized before the return from call `n+1` to `l.Lock` (All reads must be finished before a write can occur).

>A successful call to `l.TryLock` (or `l.TryRLock`) is equivalent to a call to `l.Lock` (or `l.RLock`). An unsuccessful call has no synchronizing effect at all.
### Once
Multiple threads can execute `once.Do(f)`, but only one will run `f()`, and the other calls block until `f()` has returned.

The completion of a single call of `f()` from `once.Do(f)` is *synchronized before* the return of any call of `once.Do(f)`.

```go
var a string
var once sync.Once

func setup() {
	a = "hello, world"
}

func doprint() {
	once.Do(setup)
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}
```

calling `twoprint` will call `setup()` exactly once. Other calls are going to block until `setup()` has returned. This results in *"hello, world"* to be printed twice.
### Atomic Values
If the effect of an atomic operation *A* is observed by atomic operation *B*, then *A* is *synchronized before B*. All atomic operations executed in a program behave as though executed in some sequentially consistent order.
### Finalizer
The `runtime` package provides a `SetFinalizer` function that adds a callback when a particular object is no longer reachable by the program. A call to `SetFinalizer(x, f)` is synchronized before the finalization call `f(x)`.
### Incorrect synchronization
Note that a read *r* may observe the value written by any write *w* that executes concurrently with *r*. Even if this occurs, it doesn't imply that reads happening after *r* will observe writes that happened before *w*.
This happens because, to make your code execute as fast as possible, your computer's CPU and the Go compiler work together to optimize instructions. If two lines don't seem to depend on each other, the CPU is allowed to *reorder them* or change when they are flushed to main memory.

```go
var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func doprint() {
	if !done {
		once.Do(setup)
	}
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}
```

There's no guarantee that, in `doprint`, observing the write to `done` implies observing the write to `a`. This version can print an empty string instead of *"hello, world"*.
## Understanding
- explanation of the concept, using your own words.
- Focus on cause and effect.
Ex:
- This pattern exists because systems are likely to couple business rules and external details...
- The separation allows changing interfaces without having to rewrite central rules...
## When to Use
- Situations where this is useful
## When NOT to Use
- Situations where this is overkill or harmful
## Trade-offs
- Limitations
- Costs
- Complexity 
## Examples
## References
### Connects with
Add link to relative notes
### Contrasts with
- Add link to alternatives that tries to solve the same problem
- Always add relation definition like "expands", "contrasts", "depends"
## Questions
- Points that are still not clear.
## Iterate on
- Sections of the document that can be iterated and have it's quality 
improved but need more knowledge to do so.
## Flashcards
- Q: Some question about the notes.
- A: The answer for the question above.
