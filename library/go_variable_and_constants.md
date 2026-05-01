---
id: 20260501-go_variable_and_constants
type: practice
status: draft
tags:
  - go
  - programing_language
created: 2026-05-01
---
## TL;DR
There are three ways of creating variables in Go, `var`, `:=` and `const`. Variables are always [typed]() but the compiler can **infer** and add the type to a variable with only a value. Variables that aren't initialized will received its type's **zero value**. You can generate sequences of counts using **Iota**. Variables with names starting with an **upper case** character are exported.
# Go Variables and Constants
Go is a static [typed]() language so variables must always have a predefined type associated with it.

Variables can be defined with one of the following sintaxes:
- `var`: It's the full form of variable declaration, can be used anywhere. Commonly used as global variables alongside constants.
- `:=`: Short form for variable **initialization only**. Can only be used inside functions and don't require type definitions, the type is inferred from the value.
- `const`: It's evaluated at compile time, so there's no runtime overhead. It stores immutable values. Must always be initialized and can be untyped allowing then to be used with different compatible types.

>You don't necessarily need to provide the type for the variable as the compiler will provide the necessary type for the variable during compile time. Hence we have short forms for variable declarations.
## Iota
Used with constants to simplify definitions of incrementing numbers.

The default behavior is to start at 0 **in every new const block** and increment by one after each line.
```go
const (
	 = iota // skipts initial 0 value
	Sunday   // 1
	Monday   // 2
	Tuesday  // 3
	//...
)
```
## Naming rules and style
- Variable names must not contain spaces.
- Variable names must be made up of only letters, numbers, and underscores.
- Variable names cannot begin with a number.
- Go uses `CamelCase` naming convention.
- Acronyms should be capitalized, `serveHTTP` instead of `serveHttp`.
## Zero values
When a variable is created but not initialized (not value is assigned), it assumes the zero value for its corresponding [type]():

| Type                                           | Zero Value |
| ---------------------------------------------- | ---------- |
| `int`, `byte`, `rune`                          | 0          |
| `string`                                       | ""         |
| `float64`                                      | 0.0        |
| `bool`                                         | false      |
| `map`, `chan`, `interface`, `slice`, `pointer` | `nil`      |
## Variable scope
- Variables that starts with an **upper case** character are public and therefore visible outside their package.
## Shadowing
Occurs when a variable declared within a certain scope has the same name as a variable in an outer scope. The inner variable "shadows" the outer variable, meaning that the inner variable is accessible in that scope while the other is temporarily hidden.

```go
func main() {
	x := 10
	if true {
		x := 5 // Shadows outer xx
		fmt.Println(x) // 5
	}
	fmt.Println(x) // 10
}
```
## Understanding
Variables in go depends heavily on **scope** and **naming** conventions to determine visibility and access. The language is strongly typed but the compiler provides flexibility so that you don't have to add the type every time.
## When to Use
- Use it whenever you want to store a value.
- Use **Iota** when you want to create **sequential codes** for you application.
- Use constant for default values that aren't going to change.
- Use global variables when you need values that should be available in the whole application and that aren't going to be changed often.
- Export only variables that are part of your package public interface, this allows you to [encapsulate]() your code and keep your package [interface]() consistent.
## When NOT to Use
- When you want to define a behavior. Should create a [function]().
## Patterns
- Intermediate data storage between calculations and results.
- System state definition.
## Trade-offs
- **Readability x Complexity**: Using variable initialization block can save space but can make your code complex if they're not separated. Keep block codes concise and group related variables.
- **Size x Clarity**: **Shadowing** variables can reduce the size of your code, but it can impact clarity, specially in nested functions.
## Failure Modes
- Using Iota for one use only variables.
- Exporting everything in a package without consideration.
- Making all variables global.
## Implementation (Practical)
### Variable declaration
**Multiple variables with the same type in one 'var' statement:**
```go
var a, b, c int
```

**Multiple variables with different types:**
```go
var a, b = 10, "hello"
```

**Grouped variable declaration block:**
```go
var (
	x  int
	y  string
	z  float64
)
```

**Or with initialization:**
```go
var (
	x = 10
	y = "test"
	z = true
)
```

**It's not possible to declared grouped var blocks with a singles shared type:**
```go
var (
	a, b, c int // ERROR!
)
```

**The short hand variable definition syntax can be used to reassign values to existing variables. This can only be done if there are new values on the left side of the definition:**
```go
func main() {
	x := 10
	y, x := 10, 5
	fmt.Println(x, y) // 5, 10
	x, y := 10, 5 // ERROR, no new variables on the left side
}
```
### Variable patterns
**Go prefers shorter variable names**
```go
func main() {
	userName := "Garou" // NOT IDIOMATIC
	user := "Garou"     // IDIOMATIC
}
```

**The smaller the scope the variable exists in, the smaller the variable name**
```go
func main() {
	name := "Luan"
	for i, n := range name {
		// ...
	}
}
```

**Group related variables together and possibly add a descriptive comment**
```go
// HTTP Constants
const (
	// HTTP Status Codes
	StatusOK           = 200
	StatusCreated      = 201
	
	// Error Codes
	StatusBadRequest   = 400
	StatusUnauthorized = 401
)

// Config Constants
const (
	MaxRetries         = 3
	RetryDelay         = 5 * time.Second
)
```
### Iota patterns
```go
// skip values
const (
	_ = iota // 0
	a        // 1
	_        // 2
	b        // 3
)

// offset and multiplier
const (
	Offset = 2 * iota + 1 // 1
	_                     // 3
	Value                 // 5
	_                     // 7
	Result                // 9
)
```
## Real-world Usage
- Creating permission flags with Iota.
- Creating configuration options with global variables.
- Creating system default codes with constants.
## Relationships
### Depends
- [Go Types]()
### References
- [OOP - Encapsulation]()
- [OOP - Interfaces]()
- [Go Functions]()
## Flashcards
Q: What are the three ways to declare variables in Go?  
A: `var`, `:=`, and `const`.
Q: What is a key characteristic of variables in Go?  
A: They are statically typed, meaning every variable has a defined type.
Q: Can Go infer variable types?  
A: Yes, the compiler can infer the type based on the assigned value.
Q: What happens if a variable is declared but not initialized?  
A: It is assigned the zero value of its type.
Q: What is the `var` keyword used for?  
A: It is the full form of variable declaration and can be used anywhere, including global scope.
Q: What is the `:=` syntax used for?  
A: It is a shorthand for variable initialization inside functions, with type inference.
Q: What is a limitation of the `:=` syntax?  
A: It can only be used inside functions and requires at least one new variable in assignment.
Q: What is a `const` in Go?  
A: A constant is an immutable value evaluated at compile time with no runtime overhead.
Q: Must constants always be initialized?  
A: Yes, constants must always be assigned a value.
Q: Can constants be untyped in Go?  
A: Yes, allowing them to be used with different compatible types.
Q: What is `iota` used for?  
A: To generate sequential values in constant blocks.
Q: How does `iota` behave in a const block?  
A: It starts at 0 and increments by 1 for each line.
Q: When does `iota` reset?  
A: It resets to 0 in each new const block.
Q: What are zero values in Go?  
A: Default values assigned to uninitialized variables based on their type.
Q: What is the zero value for numeric types like `int`?  
A: 0.
Q: What is the zero value for a `string`?  
A: An empty string `""`.
Q: What is the zero value for a `bool`?  
A: `false`.
Q: What is the zero value for reference types like slices or maps?  
A: `nil`.
Q: How does Go determine variable visibility?  
A: Variables starting with an uppercase letter are exported and visible outside the package.
Q: What is variable shadowing?  
A: When an inner scope declares a variable with the same name as an outer scope, hiding the outer variable.
Q: Why can shadowing be problematic?  
A: It can reduce code clarity, especially in nested scopes.
Q: What are the basic naming rules for variables in Go?  
A: Names must not contain spaces, cannot start with numbers, and can include letters, numbers, and underscores.
Q: What naming convention does Go use?  
A: CamelCase, with acronyms capitalized.
Q: What is the role of variable scope in Go?  
A: It determines where a variable can be accessed and how long it exists.
Q: When should constants be used?  
A: For values that do not change and are known at compile time.
Q: When should global variables be used?  
A: When values need to be shared across the application and change infrequently.
Q: When should variables NOT be used?  
A: When defining behavior, where functions should be used instead.
Q: What is a common use of `iota` in real-world applications?  
A: Creating sequential codes such as enums or permission flags.
Q: What is a trade-off of grouping variable declarations?  
A: It improves compactness but can reduce readability if overused.
Q: What is a trade-off of using shadowing?  
A: It reduces code size but can make the code harder to understand.
Q: What is a common failure mode when using variables?  
A: Making too many variables global, increasing coupling and reducing maintainability.
Q: Why should not all variables be exported?  
A: To maintain encapsulation and keep the package interface clean.