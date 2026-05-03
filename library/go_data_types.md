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
Very short resume with only the essential information needed.
# Go data types
Go is a **statically typed** language with types determined at compile time. Types are defined after [variable](go_variable_and_constants.md) names `myVarName <type>`.

Heres a list of Go types per category:
- `bool`
- `string`
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
## Type inference
When declaring a variable without specifying an explicit type, the variable type is inferred from the **value**.

When the value of the declaration is typed, the new variable is of that same type.
```go
var i int
j := i // j is an int
```

But when the value is an **untyped numeric constant**, the new variable may be an `int`, `float64` or `complex128` depending on the precision of the constant.
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
## Raw strings literals
Are strings literal sequences between **back quotes**. Within the quotes, any character will appear just as it is displayed between the back quotes.

```go
a := `Say "hello" to Go!\n`
fmt.Println(a) // Prints: Say "hello" to Go!\n

b := `Raw strings
can also be used
to create multiline strings`
```

>Backslashes have no special behavior in raw strings literals. If you want escaping to work, you must use **Interpreted String Literals** with double quotes `""`.
## Strings with [UTF-8]() characters
Go supports [UTF-8]() characters out of the box. Go uses the `rune` alias type for `UTF-8` data.
```go
a :=  "Hello, 世界"
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
