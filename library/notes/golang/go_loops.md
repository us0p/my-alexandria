---
id: 20260512-go_loops
type: concept
status: draft
tags:
  - go
  - programing_language
created: 2026-05-12
---
## TL;DR
Go has only one looping construct: the flexible `for` loop. Basic form has initialization, condition, post statement. Supports `for range` for arrays, slices, maps, strings, channels. Can create infinite loops or while-style loops. Control with `break` and `continue`.
# Go Loops
In Go we have only one loop keyword: `for`. We can represent `for`, `while` and `do while` loops with it. Check the [examples](#Examples) section.

Go's `loop` has three components separated by **semicolons**:
- **Init statement (optional)**: Executed **before** the **first** iteration.
- **Condition**: Evaluated **before** every iteration.
- **Post statement (optional)**: Executed at the end of **every iteration**.

Since the **init** and **post** statements are optional, this is how we can create many loops in Go with just a single structure.
## Loop control statements
Statements that change the execution from its normal sequence.
- `continue`: Skip to the next loop iteration.
- `break`: Terminate loop at that point.
## The `range` keyword
This keyword can be used to iterate over [maps](go_maps.md), [strings](go_strings.md), [array and slices](go_arrays_and_slices.md) and [channels](). It returns the index and the value for each iteration:
```go
for i, c := range "Go is awesome" {
	fmt.Prinftf("%c position is %d\n", c, i)
}
```

Each data type has it's considerations with the range keyword.
- Strings: The range returns the index and position of the **start** of the [UTF-8](ascii_unicode_and_utf8.md) encoded byte. It's the preferred way to iterate over strings as index manipulation only works for one-byte ASCII characters.
- Maps: The range returns the **key** and the **key's value** instead of index and value as in slices.
## Labeled statements
A label is an identifier followed by a **semicolon**: e.g. `LABEL:` 

>It's a good practice to use uppercase letters for labels to improve readability.

Some characteristics are:
- Must always be used. Unused labels are considered errors by the compiler.
- A label is not block scoped and does not conflict with identifiers that are not labels.
- The scope of a label is the body of the [function]() in which it's declared and excludes the body of any nested function.
- Labels can be used with the `break`, `continue` and `goto` statements.

>A label will be executed if your regular code flow does not return early and traverses them.
## The `goto` statement
 It's used to transfer control to the labeled statement. The referenced label can be both forward and backward in the code. Check [examples](#Examples)

```go
for i := 0; i < 10; i++ {
	fmt.Printf("Index: %d\n", i)
	switch i {
	case 5:
		goto exit
	case 6:
		// never reached
		goto second
	}
}

// Executed because there's no early return
fmt.Println("Skip this line here")

exit:
	// Executed once by case 5 
	fmt.Println("We are now exiting the program")

second:
	// Executed once because previuos label didn't returned
	fmt.Println("This is a second label executed because it comes after exit")
	return
```

In a `goto` statement you must refer to labels defined in the same **block** that the `goto` statement.
### `goto` best practices
In general the use of the `goto` statement is **discouraged**. But, in certain occasions it might prove useful:

1. Avoiding conditional flags for storing states and jumping over code blocks like in the `go/scanner`:
```go
func doSomething() {
	// internal logic
	if state == "someting" {
		// do something
		goto exit
	}
	
	// else logic for state
	
	exit:
		// final logic for all processes
}
```

By using `goto` here we were able to simplify the code instead of adding unnecessary boolean flags or adding separation of concerns here making the function concise and still readable.

2. Standardize an exit plan, in cases where you have a **switch** or many **for** loops that all can cause the same standard exit like in `go/types/expr`:
```go
func doSomething() {
	switch {
		case "z":
			// case handling
			if err != nil {
				goto Error
			}
		// many more cases that repeat the same error handling logic
		default:
			// some default handling.
	}
	
	Error:
		// common error handling for all cases.
}
```

Here we were able to provide a clean exit plan for the many [switch cases](go_conditionals.md) that had the same exit process, without duplication while keeping our code readable.

3. Exiting nested loops
```go
for i := 0; i < 5; i++ {
	for j := 0; j < 5; j++ {
		if j == 3 {
			goto END
		}
		fmt.Println(i, j)
	}
}

END:
```

Exiting nested loops is usually cumbersome and almost always relies on boolean flags. Using `goto` can make this task easier while keeping the code readable.
## Understanding
- Go `for` statement has the capability to represent `for`, `while` and `do while` loops with just a single keyword.
- We can use it in conjunction with the `range` keyword to iterate over structs and get a pair of values out of it according to the structure's type.
- We can use `break`, `continue` and `goto` statements to control execution.
- We can move execution to **labels** to simplify loop logic.
- **labels** aren't block scoped and do not conflict with block scoped identifications.
- Labeled blocks are executed normally, you must make sure to exit program before reaching then to avoid code re-execution.
## When to Use
- Use it whenever you need logic to be repeated based on a certain condition.
- Use loop control statements when you need to break or stop processing based on specific conditions.
- Use labels when you want to signal a specific piece of code that can be targeted by other parts of your program.
## When NOT to Use
- Don't use when you don't need to repeat action.
- Don't overly use `goto` statements.
- Don't over use label statements as it can impact code readability.
- Avoid infinite loops as they need delicate control.
- Avoid nested loops as they can impact program performance.
## Trade-offs
- Flexibility x Complexity: Using labeled statements can add flexibility to your code but can make it overly complex and impact readability.
- Performance: You need to carefully design your loop to not create loops with expensive [time complexities]().
## Examples
### `while` and `do while` loop representations
```go
i := 0
// While loop representation
for i < 5 {
	// do something
}

// Do while loop representation
for {
	// do something
	if i >= 5 {
		break
	}
}
```
### `goto` statement
```go
for i := 1; i <= 5; i++ {
	if i == 3 {
		// jump to label when i is 3
		goto label
	}
	fmt.Println(i)
}

label:
fmt.Println("Jumped to label")
```
## References
- [Go maps](go_maps.md)
- [Go strings](go_strings.md)
- [Go arrays and slices](go_arrays_and_slices.md)
- [Go concurrency]()
- [Go functions](go_functions.md)
- [Go conditionals](go_conditionals.md)
- [ASCII and Unicode](ascii_unicode_and_utf8.md)
- [Time Complexities]()
## Flashcards
- Q: What looping construct does Go provide?
- A: Go provides only the `for` loop.
- Q: Which types of loops can Go's `for` statement represent?
- A: It can represent `for`, `while`, and `do while` loops.
- Q: What are the three components of a Go `for` loop?
- A: Init statement, condition, and post statement.
- Q: When is the init statement executed in a Go `for` loop?
- A: Before the first iteration.
- Q: When is the loop condition evaluated in a Go `for` loop?
- A: Before every iteration.
- Q: When is the post statement executed in a Go `for` loop?
- A: At the end of every iteration.
- Q: Are the init and post statements required in a Go `for` loop?
- A: No, both are optional.
- Q: Why can Go represent multiple loop styles with a single `for` keyword?
- A: Because the init and post statements are optional.
- Q: What does the `continue` statement do in a loop?
- A: It skips to the next loop iteration.
- Q: What does the `break` statement do in a loop?
- A: It terminates the loop immediately.
- Q: What is the purpose of the `range` keyword in Go?
- A: It is used to iterate over maps, strings, arrays, slices, and channels.
- Q: What values does `range` return during iteration?
- A: It returns the index and value for each iteration.
- Q: What does `range` return when iterating over maps?
- A: It returns the key and the key’s value.
- Q: Why is `range` the preferred way to iterate over strings in Go?
- A: Because it correctly handles UTF-8 encoded characters.
- Q: What does `range` return for strings?
- A: The index and position of the start of the UTF-8 encoded byte.
- Q: What is a label in Go?
- A: An identifier followed by a colon used with control statements.
- Q: Which statements can use labels in Go?
- A: `break`, `continue`, and `goto`.
- Q: Are labels block scoped in Go?
- A: No, labels are not block scoped.
- Q: What happens if a label is declared but never used?
- A: The compiler reports it as an error.
- Q: What is the scope of a label in Go?
- A: The body of the function where it is declared, excluding nested functions.
- Q: What does the `goto` statement do in Go?
- A: It transfers control to a labeled statement.
- Q: Can `goto` jump both forward and backward in code?
- A: Yes, it can jump in both directions.
- Q: What restriction exists when using `goto` with labels?
- A: The label must be defined in the same block as the `goto`.
- Q: Why is the use of `goto` generally discouraged?
- A: Because excessive use can reduce readability and increase complexity.
- Q: What is one valid use case for `goto` in Go?
- A: Avoiding unnecessary conditional flags.
- Q: How can `goto` help with common error handling?
- A: By standardizing a shared exit path for multiple branches.
- Q: Why can `goto` be useful for nested loops?
- A: It simplifies exiting multiple nested loops without extra flags.
- Q: What happens to labeled code if program flow naturally reaches it?
- A: The labeled code executes normally.
- Q: What should be done to avoid unintentionally executing labeled code?
- A: Ensure the program exits or returns before reaching it.
- Q: When should loops be used in Go?
- A: When logic must repeat based on a condition.
- Q: When should loop control statements be used?
- A: When loop execution needs to be skipped or terminated based on conditions.
- Q: When should labels be used?
- A: When specific code locations need to be targeted by control statements.
- Q: What are the risks of overusing `goto` and labels?
- A: Reduced readability and increased code complexity.
- Q: Why should infinite loops be used carefully?
- A: Because they require delicate control to avoid unintended endless execution.
- Q: Why should nested loops be avoided when possible?
- A: Because they can negatively impact performance.
- Q: What trade-off exists when using labels?
- A: They increase flexibility but can also increase complexity.
- Q: What must be considered when designing loops?
- A: Their time complexity and performance impact.
- Q: How can a `while` loop be represented in Go?
- A: By using a `for` loop with only a condition.
- Q: How can a `do while` loop be represented in Go?
- A: By using an infinite `for` loop with a conditional `break`.