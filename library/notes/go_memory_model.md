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
