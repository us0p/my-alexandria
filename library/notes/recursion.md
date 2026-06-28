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
Very short resume with only the essential information needed.
# Recursion
It happens when a function call itself directly or indirectly.

A recursive function breaks a problem in simple steps that when repeated over and over again, will reach a solution.

Example: sum all digits of a number.

The sum of all digits of a number can be represented by the sum of the last digit with the sum of all other digits. This can be broken into a simple step: Summing the last digit to the sum of all the digits previous to it.
We can create a function that keeps extracting the last digit of a number and summing it to all the other digits until there's a single digit, in which case its sum is itself.
```go
import "math"

func sum_digits(n int) int {
	if n < 10 {
		return n
	}
	
	all_but_last, last := int(math.Floor(n / 10)), n % 10
	return sum_digits(all_but_last) + last
}
```

We broke the problem in a single repeatable step and executed it until the problem got resolved.

Note that the result of each call depends on the next until the base case is reached.

>It's often clearer to think about recursive calls as functional abstractions. That is, we should not care about how the function is implemented, we should simply **trust** that it computes the output correctly. This has been called a *recursive leap of faith*.
## Base Case
It's the simplest non-reducible case in which the function should simply return. Basically, it's what impedes the function of going on indefinitely.
In the example above, the base case was `n < 10`, this is obvious since we can't extract the last digit of a number when the number has a single digit on it.

The base case is the most important point of a recursive function and it often makes a good start point to creating one, since it makes it simpler to reason about how the function is executed.

The base case can be used to verify the correctness of a recursive function as well. Since for any *N* step in the function recursive step we must assume that the recursive call is going to work correctly (*leap of faith*), then it all comes down to whether the base case is implemented correctly or not.
## Mutual Recursion
When a recursive procedure is divided among two functions that call each other.

Example: recursive even/odd number evaluation.

Two functions `is_even` and `is_odd` call each other with `n - 1` to determine if a number is even or odd.
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

Example: Fibonacci
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

When converting a iterative function to a recursive one it's often simpler. **The state of an iteration can be passed as arguments to the recursive function**.

Example: sum digits of a number.
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

// can also be represented as a compound change to the recursive step output
func sumDigitsVariation(n int) int {
	if n == 0:
		return n
	n, last := int(math.Floor(n / 10)), n % 10
	return sumDigitsVariation(n) + last 
}
```
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
