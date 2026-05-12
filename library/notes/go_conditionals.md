---
id: 20260511-go_conditionals
type: concept
status: draft
tags:
  - go
  - programing_language
created: 2026-05-11
---
## TL;DR
Control program flow based on conditions. `if` for basic logic, `if-else` for binary decisions, `switch` for multiple conditions. `if` supports optional initialization, no parentheses needed but braces required. `switch` supports expressions, type switches, fall-through. Fundamental for business logic.
# Go conditionals
In Go, parentheses aren't needed in `if/else` conditionals.
```go
if x % 2 == 0 {
	fmt.Printf("%d is multiple of 2", x)
} else if x % 5 {
	fmt.Printf("%d is multiple of 5", x)
} else {
	fmt.Printf("%d is not a multiple of 2 or 5", x)
}
```

Logical operators available are:
- **AND**:`&&`
- **OR**: `||` 
- **NOT**: `!`
Comparison operators available are:
- **EQUAL**: `==`
- **Bigger/Bigger or equal**: `>/>=`
- **Smaller/Smaller or equal**: `</<=`

>There's no ternary in Go, you'll need to use full `if` statements even for basic conditions.
## Statement precedence in conditionals
It's possible to declare [variables](go_variable_and_constants.md) in conditionals. Those variables are available in the **current and all subsequent branches**.
```go
if num := 9; num < 0 {
	fmt.Println(num, "is negative")
} else if num < 10 {
	fmt.Println(num, "has 1 digit")
} else {
	fmt.Println(num, "has multiple digits")
}
```
## Switch
It's a shorter way to write a sequence of `if/else` statements. It **runs only the first case** whose value **and [type](go_data_types.md)** are equal to the condition.
```go
switch t := time.Now(); {
// values don't need to be constant
case t.Hour() < 12:
	fmt.Println("It's before noon")
default:
	fmt.Println("It's after noon")
}
```

In the example above, we initiate a [variable](go_variable_and_constants.md) `t` but the don't use it after the colon. This means that we're using a `switch true{}` statement **(good for long if-then-else chains)**, therefore, **case branches must match the expression type**, in this case, `booleans`.

If we have used the variable in the switch, the branches would have to match it's type like bellow:
```go
// Switches on time.Weekday type
switch t := time.Now().Weekday(); t {
// Correctly branches on multiple values of time.Weekday type
case time.Saturday, time.Sunday:
	fmt.Println("It's weekend")
// ERROR, branch has boolean type and not time.Weekday type.
// case values must match switch expression
case t == time.Monday:
	fmt.Println("It's Monday")
default:
	fmt.Println("It's a weekday")
}
```

>`switch` evaluate cases from top to bottom, stopping when a case succeeds.
### Type switches
`switch` statements can also be used to branch code based on the [variable's](go_variable_and_constants.md) [type](go_data_types.md):
```go
switch t := i.(type) {
case bool:
	fmt.Println("It's a bool")
case int:
	fmt.Println("It's an int")
default:
	fmt.Println("Can't determine type")
}
```
### Fallthrough
You can use a `fallthrough` statement to **transfer controls to the first statement immediately after the case which has been executed**.
```go
switch num: getNum(); {
case num < 50:
	fmt.Printf("%d is lesser than 50\n", num)
	fallthrough
case num > 100:
	fmt.Printf("%d is greather than 100\n", num)
}

// if num is 25, outputs:
// 25 is leesser than 50
// 25 is greather than 100
```

Considerations:
- `fallthrough` should be the last statement in a `case`, otherwise compiler fails.
- `fallthrough` **happens even when the case evaluates to `false`**.
- `fallthrough` cannot be used in the last case of a switch.
### Breaking switch
The `break` statement can be used to terminate a switch early before it completes.
```go
switch {
case num < 50:
	if num < 0 {
		break
	}
//...
}
```
#### Breaking the outer [loop]()
When the `switch` is inside a [for loop](), there might be a need to terminate the loop early. This can be done by labeling the loop and breaking the loop using that label.
```go
randloop:
	for {
		switch i := rang.Intn(100); {
		case i % 2 == 0:
			fmt.Printf("Generated even number %d", i)
			break randloop // breaks out the loop itself.[:w
		}
	}
```

>If `break` was used without the label it would only break the `switch` statement and not the loop.
## Understanding
- `if/else` statements are a way to add conditional execution into your program.
- You can declare variables scoped to the conditional block.
- `switch` statements can be used to:
	- Represent long `if/else` chains.
	- Execute conditional logic based on value type.
	- Match possible exact values.
- `switch` statements `case` blocks must have the same type as the value being switched upon.
- `switch` statements execute only the first top to bottom case that matches.
- `switch fallthrought` can be used to execute more than one case at the time.
- `switch fallthrought` will not check the condition of the case it's going to execute next.
- `switch fallthrought` can only be used as the last statement in a case block and cannot be used in the last case block of a `switch` statement.
- You can use `break` to terminate a `switch` execution early or to break a labeled loop.
## When to Use
- Use it when you need to check for conditions.
- When you need a variable with a life cycle associated with the condition only.
- When you need to check on various conditions, possible types and exact values like [enums](go_variable_and_constants.md#Iota).
## When NOT to Use
- Don't use when you don't need conditional branching on your code.
- Avoid nesting conditional blocks as it makes your code less readable.
- Avoid using `switch` statements where a simple `if/else` would suffice and vise versa.
## Trade-offs
- Complexity: Using too much conditions or nested conditions can increase your code complexity by damaging readability.
## Examples
Switching on enum values
```go
type CarMake int

const (
	UNDEFINED_MAKE CarMake = iota
	FORD
	FIAT
)

switch getMake() {
case FORD:
	fmt.Println("Make is Ford")
case FIAT:
	fmt.Println("Make is Fiat")
default:
	fmt.Println("Make is undefined")
}
```
## References
 - [Go variables](go_variable_and_constants.md)
 - [Go data types](go_data_types.md)
 - [Go loops](go_loops.md)
## Flashcards
- Q: What are Go conditionals used for?
- A: They are used to control program flow based on conditions.
- Q: Which conditional structures are available in Go?
- A: Go provides `if`, `if-else`, and `switch`.
- Q: What is the purpose of `if` statements in Go?
- A: They are used for basic conditional logic and execution.
- Q: What is the purpose of `if-else` statements in Go?
- A: They are used for binary or chained conditional decisions.
- Q: What is the purpose of `switch` statements in Go?
- A: They provide a shorter way to write long `if/else` chains and can branch on values or types.
- Q: Are parentheses required in Go `if` statements?
- A: No, parentheses are not required.
- Q: Are braces required in Go conditionals?
- A: Yes, braces are required.
- Q: Does Go support ternary operators?
- A: No, Go does not have ternary operators.
- Q: Which logical operators are available in Go?
- A: `&&` for AND, `||` for OR, and `!` for NOT.
- Q: Which comparison operators are available in Go?
- A: `==`, `>`, `>=`, `<`, and `<=`.
- Q: Can variables be declared inside Go conditionals?
- A: Yes, variables can be declared in conditionals.
- Q: What is the scope of variables declared in a conditional statement?
- A: They are available in the current and all subsequent branches.
- Q: How does a Go `switch` statement evaluate cases?
- A: It evaluates cases from top to bottom and stops at the first matching case.
- Q: What must match in a Go `switch` case?
- A: Both the value and the type must match the switch expression.
- Q: What is `switch true {}` commonly used for?
- A: It is useful for long `if-then-else` chains.
- Q: In a `switch true {}` statement, what type must case expressions return?
- A: They must return boolean values.
- Q: What are type switches in Go used for?
- A: They are used to branch logic based on a variable’s type.
- Q: What does the `fallthrough` statement do in a Go `switch`?
- A: It transfers control to the next case immediately after the current one.
- Q: Does `fallthrough` check the next case condition before executing it?
- A: No, it executes the next case even if its condition is false.
- Q: Where must `fallthrough` appear inside a case block?
- A: It must be the last statement in the case block.
- Q: Can `fallthrough` be used in the last case of a switch?
- A: No, it cannot be used in the last case.
- Q: What is the purpose of `break` in a Go `switch`?
- A: It terminates the switch execution early.
- Q: What happens if `break` is used inside a `switch` within a loop?
- A: It only breaks the `switch`, not the outer loop.
- Q: How can you break out of an outer loop from inside a `switch`?
- A: By using a labeled loop and `break` with the label.
- Q: What are common uses of `switch` statements in Go?
- A: Representing long `if/else` chains, checking types, and matching exact values.
- Q: What must the type of `switch` case expressions match?
- A: They must match the type of the switch expression.
- Q: How many matching cases are executed in a normal Go `switch`?
- A: Only the first matching case is executed.
- Q: When should conditional statements be used in Go?
- A: When conditional branching or condition-specific variables are needed.
- Q: When should deeply nested conditionals be avoided?
- A: When they reduce readability and increase complexity.
- Q: Why can excessive conditional logic be problematic?
- A: It increases code complexity and harms readability.
- Q: When should `switch` be preferred over `if/else`?
- A: When handling many conditions, exact values, or type-based branching.
- Q: When should `if/else` be preferred over `switch`?
- A: When the logic is simple and does not require multiple cases.