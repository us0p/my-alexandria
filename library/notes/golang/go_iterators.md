---
id: 20260709-go_iterators
type: concept
status: draft
tags:
  - go
created: 2026-07-09
---
## TL;DR
Iterators (Go 1.23+) enable lazy evaluation, generating sequence values one at a time instead of allocating them all upfront like slices. `for/range` now supports ranging over functions taking a single `yield` argument (signatures `func(yield func() bool)`, `func(yield func(T) bool)`, or `func(yield func(K, V) bool)`). `yield` returns `true` to continue or `false` to stop (via `break`/`return`/`goto`), and panics if called again after returning `false`. Some iterators are single-use, walking a stream only once. The standard library provides `iter.Seq`/`iter.Seq2` types for single/two-value push iterators, and `iter.Pull` to convert a push iterator into a pull iterator returning `next`/`stop` functions.
# Go Iterators
Iterators provide a memory-efficient way of processing sequences of data by generating values one at the time. This is also known as a **"Lazy Evaluation"**.

With a [Slice](go_arrays_and_slices.md) approach, we need to wait and allocate all slice elements in memory before starting to process them.

>Introduced in 1.23.
## Changes to `for/range` statements
With the addition of iterators feature, `for` and `range` statements now support ranging over functions that **take a single argument**. The single argument **must itself be a function that takes zero to two arguments and returns a bool**. Conventionally called `yield`.

An iterator is a **function that returns the next element in the sequence**.

The signature of an iterator function is one of
```go
func (yield func() bool)
func (yield func(T) bool)
func (yield func(K, V) bool)
```

Where `T, K, V` are the types of elements in the sequence and `yield` is a function that returns the next element in the sequence. This function stops either when the sequence is finished or when yield returns false, **indicating to stop the iteration early**.

>Iterators are first-class citizen in Go.

You can check iterator's naming convention in the [official docs](https://pkg.go.dev/iter@go1.26.5#hdr-Naming_Conventions).
## `yield` behavior in `range` loops
In range-over-function iteration pattern, the `yield` function returns a boolean that represents the **range loop body**.
- `yield` returns `true` if the body finished normally (it will keep iterating).
- `yield` returns `false` if the loop executed `break`, `return` or `goto` (it won't keep iterating).

Since `yield` panics if called after it returns false, we **always need to check for valid `yield` calls**.

```go
// If caller returns early, function stops iteration early as well.
if !yield(v) {
	return
}
```
## Single-Use Iterators
Most iterators when called does any setup necessary to start the sequence, and cleans up before returning. Calling the iterator again walks the sequence again.

Some iterators break that convention, providing the ability to walk a sequence only once. These iterators typically report values from a data stream that cannot be rewound to start over. Calling the iterator again after stopping early may continue the stream, but calling it again after the sequence is finished will yield no values at all.

Doc comments for these iterators should document this fact.
## Standard Library Iterators
### Single Value Iterators
If the iterator returns a single value, you can use `iter.Seq` to create an iterator function.

```go
type Seq[V any] func(yield func(V) bool)
```
### Two Value Iterators
If your iterator returns two values, you can use the `iter.Seq2` type:

```go
type Seq2[K, V any] func(yield func(K, V) bool)
```

>Usually the first value is the index of the loop, and the second value is the actual element, following convention of looping Maps or Slices.
## Pulling Values
The standard iterators (`Seq and Seq2`) can be thought as **"push iterators"**, which push values to the `yield` function.

If you don't need this behavior you can use `iter.Pull` to convert a standard **push iterator** to a **"pull iterator"**, which pull one value at a time from the sequence. Useful when the sequence you're looping is not naturally supported by `range` or `for` statements.

`iter.Pull` starts an iterator and returns a pair of functions `next and stop`.

If a client do not consume the sequence to completion, they must call stop, which allows the iterator function to finish and return.
## Understanding
Iterators are a memory efficient way to processing a sequence of items. They gives us a standard that we can apply across the language without having to implement it manually every time while also giving use the flexibility to implement only the necessary to fit our needs.

It relies heavily on [Functional Programming](functional_programming.md) paradigm and its patterns.
## When to Use
- Use when you need to treat container like data (Structs) or data that is not a linear collection of items (Trees, Graphs, etc) as a linear collection that can be iterated over.
- If you need memory efficiency and don't want to be keeping large elements into memory at once.
## When NOT to Use
- When the sequence modifies its own state during iteration.
- Concurrent data streaming, range-over-function iterators are synchronous, they pass control back and forth. Do not use them if your data producer doesn't run on the same goroutine as the data consumer, use [Channels](go_channels.md) instead.
- If you already have a plain Slice or Map. Avoids cognitive overhead.
## Trade-offs
### Pros
- Encapsulation: Container like data can now be easily represented as a sequence without exposing its internals.
- Resource safety: Each iterator can control how the data source is managed, this means acquiring locks to databases or opening files are all managed by the iterator function and thus their lifetime.
- Lazy Evaluation: You can represent infinitely large datasets without having to store it all in memory first.
### Cons
- Complexity: Event if iterators makes it easy to treat non-sequence data as a sequence, it adds a bit of complexity, specially if the developer is not familiar with functional programming paradigms.
- Loss of mutability: It becomes harder to mutate the underlying data while you're traversing it.
- Concurrency restrictions: Iterators are data generators that run in a synchronous loop, usually paired with a data consumer running in the same goroutine. You cannot easily split it across multiple CPU cores like you can with channels.
## Examples
### Iterative function
```go
func iterateAllUsers(yield func(User) bool) {
	for i := 0; i < 1_000_000; i++ {
		if !yield(User{ID: i, Name: fmt.Sprintf("User%d", i)}) {
			return
		}
	}
}

for user := range iterateAllUsers {
	// ...
}
```
### `iter.Seq` example
```go
import "iter"

func iterateAllUsers() iter.Seq[User] {
	return func(yield func(User) bool) {
		// User generation logic
	}
}

for user := range iterateAllUsers() {
	// ...
}
```
### `iter.Seq2` example
```go
import "iter"

func iterateAllUsers() iter.Seq2[int, User] {
	return func(yield func(int, User) bool) {
		for i := 0; i < 1_000_000; i++ {
			if !yield(i, User{}) {
				return
			}
		}
	}
}

for idx, user := range iterateAllUsers() {
	// ...
}
```
### Hand Written yield function
```go
func PrintAllElements[E comparable](s *Set[E]) {
	s.All()(func(v E) bool {
		fmt.Println(v)
		return true
	})
}
```
This example shows that an iterator function is just a function that is applied to a collection.

The `s.All()` function returns an iterator. Remember that an iterator is just **a function that receives an `yield` function as arguments that must have 0 - 2 arguments and return a boolean value**. This `yield` function is the function we define that only prints the element.

Since the iterator itself already walks every element in the sequence and calls the `yield` function on each, This just prints the whole collection, without needing a loop.
## References
- [Slice](go_arrays_and_slices.md)
- [Functional Programming](functional_programming.md)
- [Channels](go_channels.md)
### External Links
- [Iterator's Naming Convention](https://pkg.go.dev/iter@go1.26.5#hdr-Naming_Conventions)
## Flashcards
- Q: What are Go iterators and what benefit do they provide?
- A: A way to process sequences by generating values one at a time instead of allocating them all in memory upfront, also known as lazy evaluation.
- Q: When were iterators introduced to Go?
- A: In Go 1.23.
- Q: What single argument must a for and range statement accept to range over a function?
- A: A function that itself takes zero to two arguments and returns a bool, conventionally called yield.
- Q: What is an iterator in Go?
- A: A function that returns the next element in the sequence.
- Q: What are the three possible signatures of an iterator function?
- A: func(yield func() bool), func(yield func(T) bool), and func(yield func(K, V) bool).
- Q: When does an iterator function stop calling yield?
- A: When the sequence is finished or when yield returns false, indicating to stop iteration early.
- Q: In a range over function loop, what does it mean when yield returns true?
- A: The loop body finished normally and iteration continues.
- Q: In a range over function loop, what does it mean when yield returns false?
- A: The loop body executed break, return, or goto, and iteration should not continue.
- Q: What happens if yield is called after it has already returned false?
- A: It panics, so callers must always check for valid yield calls before calling it again.
- Q: What is a single-use iterator?
- A: An iterator that can only walk its sequence once, typically because it reports values from a data stream that cannot be rewound to the start.
- Q: What happens when a single-use iterator is called again after its sequence finished?
- A: It yields no values at all.
- Q: What Go standard library type represents an iterator that returns a single value?
- A: iter.Seq, defined as type Seq[V any] func(yield func(V) bool).
- Q: What Go standard library type represents an iterator that returns two values?
- A: iter.Seq2, defined as type Seq2[K, V any] func(yield func(K, V) bool).
- Q: In an iter.Seq2, what do the first and second values conventionally represent?
- A: The first value is usually the index of the loop, and the second is the actual element, following the convention of looping over maps or slices.
- Q: What is the difference between a push iterator and a pull iterator?
- A: A push iterator, like Seq and Seq2, pushes values to the yield function, while a pull iterator pulls one value at a time from the sequence.
- Q: How do you convert a standard push iterator into a pull iterator?
- A: Use iter.Pull, which starts the iterator and returns a pair of functions, next and stop.
- Q: When is a pull iterator useful?
- A: When the sequence being looped is not naturally supported by range or for statements.
- Q: What must a client do if it does not consume a pull iterator sequence to completion?
- A: It must call stop, which allows the iterator function to finish and return.
