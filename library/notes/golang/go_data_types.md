---
id: 20260503-go_data_types
type: concept
status: draft
tags:
  - go
  - programing_language
created: 2026-01-03
---
## TL;DR
Go uses **strict, static typing**: every value has a type known at compile time, and you must explicitly convert between types. This improves safety and performance but requires more attention when mixing values, handling numbers, and working with strings (which are byte-based, not character-based).
# Go data types
Go is a **statically typed** language with types determined at compile time. Types are defined after [variable](go_variable_and_constants.md) names `myVarName <type>`.

Heres a list of Go types per category:
- `bool`
- [`string`](go_strings.md)
- `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr` (large enough to store uninterpreted bits of a pointer value)
- `byte` (alias for `int8`)
- `rune` (alias for `int32`. Represents a Unicode code point)
- `float32`, `float64`
- `complex64`, `complex128`

>`int`, `uint` and `uintptr` are 32 or 64 bits wide according to OS architecture. This makes them usually the best choice because they align with the machine's word size.
## Which type variation to choose
Unless you have a very specific reason to use a determined integer type, you should stick to the generic types `int` or `uint`.

With that said, there are a couple of considerations you need to take when choosing the right type:
- **Memory Optimization**: Using less memory make your program faster and smaller.
- **Consistency**: You should always follow the project structure.
## Custom data types
Using the `type` keyword we can create aliases/custom types based on our program needs. This helps with encapsulation of our program.

```go
type MyInt int

var mi MyInt = 1
var i int = int(mi) // Must convert, error otherwise.
```
## Type inference
When declaring a variable without specifying an explicit type, the variable type is inferred from the **value**.

When the value of the declaration is typed, the new variable is of that same type.
```go
var i int
j := i // j is an int
```

But when the value is an **untyped numeric constant**, the new variable may be an `int`, `float64` or `complex128` depending on the precision of the constant.
## Type conversion and compatibility
There's no automatic type or conversion.
```go
a := 80                    // int
b := 91.8                  // float64
sum := a + int(b)          // int + float64 (error), must convert to other component type

var c float64 = float64(a) // need explicit type conversion for assignment as well.
```

Even different types of integers cannot be mixed in operations without explicit conversion:
```go
var a int8 = 10
var b int16 = 10
var c int32 = a + b // invalid operation, mismatched typed
```
### Type conversion cheat sheet
```go
// Struct conversions
type Point struct {
	X, Y int
}

type Coord struct {
	X, Y int
}

var p Point = Point{1, 2}
var c Coord = Coord(p) // same field types

// Function type conversions
type FuncType func(int) int

func square(x int) int {
	return x * x
}

var f FuncType = FuncType(square)
```
## Type Assertions
To get the underlying type of an [interface](), we can assert/convert the interface to an specific value. For that we use the [comma-ok idiom](go_maps.md#Comma-ok%20idiom):
```go
var a interface{} = 1
if v, ok := a.(int); !ok {
	fmt.Println("Interface value is not an integer")	
}
```

>The comma-ok idiom is optional but asserting the type of a value without it will cause the program to panic if the value type doesn't match the assertion.

Another kind of type assertion is with [switches]():
```go
switch v := a.(type) {
	case int:
		fmt.Println("...")
	case string:
		fmt.Println("...")
	// ...
}
```
## Type conversion and assertion best practices
- Always check and handle errors when performing type assertion with multiple values.
- Avoid using unsafe type conversion and assertion as much as possible.
- Use custom types to provide safety and clarity. Custom types can provide more safety and clarity in the code, and can also reduce the number of type assertions required.
- Keep type conversion and assertion logic separate from the main logic of the program. This makes it easier to test and modify the code when necessary.,
## Numeric overflow error and wrap around behavior
If the compiler can determine that a calculation will overflow the numeric type, the compilation will fail with an `overflow` error.

If the overflow happens at runtime, the number is going to wraparound to the initial number.
```go
var short1 int8 = 128  // overflow compile time error

var short2 int8 = 127

fmt.Println(short2 + 1) // it'll wraparound to -128 (initial number for int8) during runtime
```
## Precision limitations in floating-point calculations
floating-point numbers are represented using a finite number of bits. This can lead to rounding errors and unexpected behavior in certain calculations.
```go
a := 0.1
b := 0.2
fmt.Println(a + b) // Prints 0.30000000000000004
```

To mitigate this precision issues, you should consider the following best practices:
- **Use appropriate data types**: consider application requirements before choosing between `float32` and `float64` and stick to it.
- **Avoid direct comparisons**: Instead of `a == b`, consider using a small tolerance value to check if difference is within acceptable range.
- **Round values appropriately**: Consider rounding it to an appropriate number of decimal places.
- **Use specialized libraries**: Consider for applications that require precise floating-point calculations.
## Understanding
- Go uses **static typing** so errors are caught at compile time and performance is optimized early.
- It avoids **implicit conversions**, forcing explicit casting to prevent hidden bugs and precision loss.
- Type inference reduces verbosity while still preserving strict type safety.
- Custom types add **meaning and safety**, preventing misuse of similar-looking values.
- Type assertions enable flexibility with interfaces but require checks to avoid runtime errors.
- Limitations like overflow and floating-point precision come from **hardware constraints**.
- Overall, Go prioritizes **explicitness, safety, and predictability** over convenience.
## When to Use
- When you want **compile-time safety** and fewer runtime surprises.
- In systems where **performance matters**, since types are optimized and checked ahead of execution.
- When building **large or long-lived codebases**, where explicit types improve readability and maintainability.
- When you need **clear data modeling**, using custom types to express intent (e.g., domain-specific types).
- When working with **low-level data (bytes, memory, encoding)**, since Go gives precise control over representations.
## When NOT to Use
- When rapid prototyping is more important than strict correctness (dynamic languages may be faster to iterate).
- When your logic involves **frequent implicit conversions**, which Go does not support and can feel verbose.
- When dealing with **loosely structured or highly dynamic data** (e.g., heavy use of arbitrary JSON), where strict typing can add friction.
- When you don’t need fine control over numeric types or memory usage—Go’s strictness may feel like overkill.
## Trade-offs
- **Verbosity**: Explicit conversions (`int` ↔ `float64`, etc.) add extra code.
- **Rigidity**: No implicit casting means less flexibility when combining values.
- **Precision pitfalls**: Floating-point limitations can still cause subtle bugs.
- **Overflow behavior**: Silent wraparound at runtime can lead to unexpected results if not handled carefully.
- **Complexity with interfaces**: Type assertions require care (and checks) to avoid runtime panics.
## Examples
- Reading a file or handling HTTP request bodies (`byte` heavy).
- An e-commerce system distinguishing `UserID` from `ProductID`.  Both might be integers, but represent different concepts.
## References
### Connects with
- [Go variables](go_variable_and_constants.md)
- [Go interfaces]()
- [Go Comma-OK Idiom](go_maps.md)
- [Go Switches](go_conditionals.md)
- [Go Slices](go_arrays_and_slices.md)
- [Go Strings](go_strings.md)
## Iterate on
- Document is too dense, need to split into more manageable pieces.
## Flashcards
Q: What is the core typing model used in Go?  
A: Go uses strict, static typing where all types are known at compile time.
Q: Why does Go enforce explicit type conversions?  
A: To prevent hidden bugs and ensure type safety and predictability.
Q: What is type inference in Go?  
A: The compiler automatically determines a variable’s type based on its assigned value.
Q: When does type inference occur?  
A: When a variable is declared without an explicit type but initialized with a value.
Q: What are the main categories of basic data types in Go?  
A: Booleans, strings, integers (signed/unsigned), floats, complex numbers, bytes, and runes.
Q: What is the recommended default integer type in Go?  
A: `int`, as it matches the system’s architecture (32-bit or 64-bit).
Q: What is a `byte` in Go?  
A: An alias for `int8`.
Q: What is a `rune` in Go?  
A: An alias for `int32` representing a Unicode code point.
Q: What is a custom type in Go?  
A: A user-defined type created using the `type` keyword to add meaning and safety.
Q: Why use custom types?  
A: To improve clarity, enforce constraints, and avoid misuse of similar data.
Q: Can different numeric types be mixed in operations?  
A: No, explicit conversion is required even between similar numeric types.
Q: What happens if you try to mix incompatible types without conversion?  
A: The program will fail to compile.
Q: What is type assertion in Go?  
A: Extracting the concrete value from an interface type.
Q: What is the comma-ok idiom in type assertions?  
A: A safe way to check if a type assertion succeeded without causing a panic.
Q: What happens if a type assertion fails without checking?  
A: The program panics at runtime.
Q: How can type assertions be handled with multiple types?  
A: Using a type switch.
Q: What is numeric overflow in Go?  
A: When a value exceeds the limits of its type.
Q: What happens if overflow is detected at compile time?  
A: Compilation fails with an error.
Q: What happens if overflow occurs at runtime?  
A: The value wraps around to the minimum value of the type.
Q: What is a limitation of floating-point numbers?  
A: They can introduce precision errors due to finite representation.
Q: Why should floating-point values not be compared directly?  
A: Because small precision errors can lead to incorrect equality checks.
Q: What is a key benefit of Go’s static typing?  
A: Errors are caught at compile time, improving safety and performance.
Q: What is a key trade-off of strict typing in Go?  
A: Increased verbosity due to required explicit conversions.
Q: What is a challenge when working with Go interfaces?  
A: Type assertions require careful handling to avoid runtime errors.
Q: When is Go’s type system especially beneficial?  
A: In large systems, performance-critical applications, and clear data modeling.
Q: When might Go’s strict typing be a drawback?  
A: In rapid prototyping or when working with highly dynamic data.
Q: What is a real-world use case for byte-level operations?  
A: Processing file data or HTTP request bodies.