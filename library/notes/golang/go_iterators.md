---
id: 20260709-go_iterators
type: concept
status: draft
tags:
  - go
created: 2026-07-09
---
## TL;DR
Very short resume with only the essential information needed.
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

`yield` panics if called after it returns false.

>Iterators are first-class citizen in Go.

You can check iterator's naming convention in the [official docs](https://pkg.go.dev/iter@go1.26.5#hdr-Naming_Conventions).
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
	// s.All() returns an iterator which is a function
	// A iterator function must have a single argument which must be the yield function
	// Here shown that it can be manually specified.
	s.All()(func(v E) bool {
		fmt.Println(v)
		return true
	})
}
```
## References
### Connects with
Add link to relative notes
### Contrasts with
- Add link to alternatives that tries to solve the same problem
- Always add relation definition like "expands", "contrasts", "depends"
## Questions
- What causes the yield function to return false?
- How to stop an iterator early?
## Iterate on
- Sections of the document that can be iterated and have it's quality 
improved but need more knowledge to do so.
## Flashcards
- Q: Some question about the notes.
- A: The answer for the question above.
