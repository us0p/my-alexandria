---
id: 20260627-recursion
type: concept
status: draft
tags:
  - computer_theory
  - algorithms
created: 2026-06-27
---
## TL;DR
Recursion happens when a function calls itself, storing intermediate state in the runtime environment instead of variables. Every recursive function needs a base case, the non-reducible stopping point, and a recursive case that reduces the problem toward it. Variants include mutual recursion (two functions calling each other) and tree recursion (branching into multiple calls). Any recursive function can be converted to iteration by identifying the state it must carry, weighing performance against simplicity. List recursion can use first/rest decomposition or a helper function to avoid copying. Tail-call optimization reuses stack frames but is not available in Go.
# Recursion
It happens when a function call itself directly or indirectly.

A recursive function breaks a problem in simple steps that when repeated over and over again, will reach a solution.

The core idea of recursion and also its contrast with iteration is that all the intermediate states that would be stored in variables on a iterative solution, are stored on the function runtime environment, as function parameters or function return values, this way the environment itself takes care of the state management, the function only cares about the computation.

>It's often clearer to think about recursive calls as functional abstractions. That is, we should not care about how the function is implemented, we should simply **trust** that it computes the output correctly. This has been called a *recursive leap of faith*.
## Base Case
It's the simplest non-reducible case in which the function should simply return. Basically, it's what constrain the function of going on indefinitely.

The base case is the most important point of a recursive function and it often makes a good start point to creating one, since it makes it simpler to reason about how the function is executed.

The base case can be used to verify the correctness of a recursive function as well. Since for any *N* step in the function recursive step we must assume that the recursive call is going to work correctly (*leap of faith*), then it all comes down to whether the base case is implemented correctly or not.
## Recursive Case
Set of rules that reduce all cases toward base cases. In other words, the recursive step should always make the problem smaller or simpler in some way.
## Space Complexity
The stack stores the variables and references for each function call, thus, memory must be allocated to keep track of the current stack.

The space complexity of recursive solutions is usually related to the time complexity of the recursive function, given that for each recursive call, a new stack is created.
## Mutual Recursion
When a recursive procedure is divided among two functions that call each other.
```go
func is_even(n int) bool {
	if n == 0 {
		return true
	}
	return is_odd(n - 1)
}

func is_odd(n int) bool {
	if n == 0 {
		return False
	}
	return is_even(n - 1)
}

fmt.Printf("Is 4 even? %t\n", is_even(4))
```
## Tree Recursion
A recursive function that calls itself more than once. It's said to be tree recursive because each call branches into multiple smaller calls.

```go
func fib(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return fib(n - 2) + fib(n - 1)
}

fmt.Printf("Fib(6): %d\n", fib(6))
```
## Converting Recursion to Iteration
*Iteration is a special case of recursion*.

To convert a recursive function to an iterative version we must discover **what state must be maintained by the iterative function**. This state is going to become the control variables the iterative version is going to use to walk through the solution.

Converting a iterative function to a recursive one it's often simpler. **The state of an iteration can be passed as arguments to the recursive function**.

```go
// notice that recursive parameters becomes iterative state through variables
func sumDigitsIter(n int) int {
	var sum int
	
	for n > 0 {
		n, last := int(math.Floor(n / 10)), n % 10
		sum += last
	}
	
	return sum
}

// state becomes parameters to the recursive call
func sumDigitsRec(n, sum int) int {
	if n == 0:
		return sum
	n, last := int(math.Floor(n / 10)), n % 10	
	return sum_digits(n, sum + last)
}

func sumDigitsVariation(n int) int {
	if n == 0:
		return n
	n, last := int(math.Floor(n / 10)), n % 10
	// can also be represented as a compound change to the recursive step output
	return sumDigitsVariation(n) + last 
}
```
### Converting Tree-Like Recursion to Iteration
Usually, to convert a Tree-like recursive function to an iterative one we must implement a [Stack](stack.md) data structure, that will hold the parameters of each recursive call so that the loop can pick them up and execute it in order.
## When to use recursion instead of iteration
Any recursive function can be written iteratively. But some problems lend themselves *naturally* to a recursive solution.

There are two reasons that need considering when choosing between a recursive and a iterative implementation:
1. **Performance**: Using recursion instead of iteration will lead to a better or worst performance? Is there a lot of difference between each?
2. **Simplicity**: Even if the iterative solution is faster, if the performance doesn't makes a big difference and there's a lot of complexity involved. It's better to use a straightforward recursive solution, even if the performance is a little worst.

Another factor that must be considered during this decision is whether the data itself is defined recursively (like a [Tree](tree.md)).

>In many programming languages, recursion is limited by a maximum call-stack. When a function tries to call deeper than that, the runtime usually produces an error. In Go, there's no [call-stack limit inside a goroutine](goroutine_infinite_stack.md)
## Different data structures recursive patterns
### Lists
- **First/Rest decomposition**: Decomposes the list into it's first and rest elements at each iteration and them process them individually on each step. Has a characteristic of performing the action backwards as it recombines the values. The drawback of this implementation is that each recursive call needs to copy the rest of the list which means that `O(n²)` time is spent copying, also, the amount of space that needs to be consumed is a lot more than keeping a single instance of the list in memory.
- **Helper function**: If you need the computation to be performed in order as the recursive step goes down the list, you can add a helper function that carries the local variables that would be needed in the iterative solution on each iteration as parameters.
## Optimizations
### First-Rest List Copying
Since in each call we only care about the first element and the rest is only used to walk to the end of the list.

We walk the list one element at the time (like a [Linked-List](linked_lists.md)) and consume only the first element by implementing an index.
### Tail-Call
> Not available in [Go](https://en.wikipedia.org/wiki/Tail_call#Language_support).

When the runtime system encounters a tail-call (the last thing in the function's body is the recursive call), it deduces that it will no longer need the frame for the current call and can simply *reuse* it for the new recursive call, rather than creating a new frame.

With this optimization, the recursion depth never exceeds 1, and the performance is essentially like a loop.

Tail-Call only works when the recursive function is called and there's no dangling reference to the previous call stack, for example:

```go
// Tail-Call optimized, recursion is the last step in the function and there's no dangling reference to the previous stack.
func sum_list(x []int) int {
	func sum_helper(sum_so_far int, lst []int) int {
		if len(lst) == 0 {
			return sum_so_far
		}
		
		num, rest := lst[0], lst[1:]
		return sum_helper(sum_so_far + num, rest)
	}
	return sum_helper(0, x)
}

// Not tail-call optimized, the recursion is part of the last step but it keeps a reference to the previous stack in each iteration.
func sum_list(x []int) int {
	if len(x) == 0 {
		return 0
	}
	return x[0] + sum_list(x[1:])
}
```
## Understanding
Recursion is a function that calls itself. It's a standard practice in algorithms and some problems lead themselves naturally to this approach.

The base cases determines when the function should stop and the recursive step should always reduce the solution towards a base case.

Some data structures like Trees and Graphs are naturally recursive and have specific patterns of recursive iterations.
## When to Use
- When you need a simple and straight forward solution.
- When the data itself already follows a recursive hierarchy.
## When NOT to Use
- Performance is a must, iterative solutions are often faster as they don't have the stack overhead.
- When the problem itself doesn't have a recursive nature. Trying to force the implementation will lead you to implement a loop variation anyway.
## Trade-offs
- **Simplicity**: Recursive solutions are often simpler than iterative ones as they usually have lot less instructions and rely heavily on the environment.
- **Complexity**: Often recursive solutions are more difficult to debug and reason about specially to developers not familiar with it.
- **Performance**: Since recursive function pushes calls to the stack, there's a lot of stack overhead for each call and memory consumption as the call stack gets deeper, this can cause your program to run a lot slower than an iterative solution.
## Examples
### Helper Function and List Copying Optimization Example
```go
func sum_list(x []int) int {
	func sum_helper(i, sum_so_far int) int {
		if i >= len(x) {
			return sum_so_far
		}
		return sum_helper(i + 1, sum_so_far + x[i])
	}
	return sum_helper(0, 0)
}
```
## References
- [Linked-List](linked_lists.md)
- [Stack](stack.md)
- [Goroutine Infinite Call Stack](goroutine_infinite_stack.md)
### External links
- [Tail-Call Language Support](https://en.wikipedia.org/wiki/Tail_call#Language_support)
## Flashcards
- Q: What is recursion?
- A: When a function calls itself, directly or indirectly, breaking a problem into simple steps that when repeated reach a solution.
- Q: What is the recursive leap of faith?
- A: Treating a recursive call as a functional abstraction, trusting that it computes the correct output without worrying about how it is implemented.
- Q: What is a base case?
- A: The simplest non-reducible case in which a recursive function simply returns, constraining the function from running indefinitely and serving as the key point for verifying correctness.
- Q: What is the recursive case?
- A: The set of rules that reduce every case toward the base case, making the problem smaller or simpler at each step.
- Q: What is mutual recursion?
- A: When a recursive procedure is divided among two functions that call each other.
- Q: What is tree recursion?
- A: A recursive function that calls itself more than once per call, branching into multiple smaller calls, like a naive Fibonacci implementation.
- Q: How do you convert a recursive function into an iterative one?
- A: By discovering what state must be maintained, which becomes the control variables the iterative version uses to walk through the solution.
- Q: How do you convert an iterative function into a recursive one?
- A: By passing the state of the iteration as arguments to the recursive function.
- Q: How do you convert tree-like recursion into iteration?
- A: By implementing a stack data structure that holds the parameters of each recursive call so a loop can process them in order.
- Q: What two factors should guide choosing recursion over iteration?
- A: Performance, whether recursion causes a meaningful slowdown, and simplicity, whether the recursive solution is easier to write and reason about.
- Q: Why might recursion be the natural choice regardless of performance?
- A: When the data itself is defined recursively, such as a tree.
- Q: Does Go limit recursion depth inside a goroutine?
- A: No, unlike many languages that enforce a maximum call stack, Go has no call stack limit inside a goroutine.
- Q: What is the first and rest decomposition pattern for list recursion?
- A: Decomposing a list into its first and rest elements at each step and processing them individually, though it performs the action backwards, copies the rest of the list on each call, and costs quadratic time and extra space.
- Q: How does the helper function pattern improve list recursion?
- A: It carries the local variables needed for in order processing as parameters, avoiding the need to process the list backwards.
- Q: How can first-rest list copying be optimized?
- A: By walking the list one element at a time using an index, like a linked list, instead of copying the rest of the list on each call.
- Q: What is tail-call optimization?
- A: When the last action in a function is the recursive call, the runtime reuses the current stack frame instead of creating a new one, keeping recursion depth at one and performance close to a loop.
- Q: Is tail-call optimization available in Go?
- A: No, tail-call optimization is not available in Go.
- Q: When does tail-call optimization fail to apply?
- A: When the recursive call is not the last step in the function or there is a dangling reference to the previous call stack, such as adding a value to the result of the recursive call.
