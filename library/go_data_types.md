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
To get the underlying type of an [interface](), we can assert/convert the interface to an specific value. For that we use the [comma-ok idiom]():
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
### Type conversion and assertion best practices
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
## Strings
### Raw strings literals
Are strings literal sequences between **back quotes**. Within the quotes, any character will appear just as it is displayed between the back quotes.

```go
a := `Say "hello" to Go!\n`
fmt.Println(a) // Prints: Say "hello" to Go!\n

b := `Raw strings
can also be used
to create multiline strings`
```

>Backslashes have no special behavior in raw strings literals. If you want escaping to work, you must use **Interpreted String Literals** with double quotes `""`.
### UTF-8 and string literals
In Go a string is a **read-only [slice]() of [bytes](binary_system.md)**. A string holds arbitrary bytes. It's not required to hold [Unicode text, UTF-8](ascii_unicode_and_utf8.md) or any other predefined format.

As implied up front, **indexing a string accesses individual bytes, not characters**. 
```go
func main() {
	const placeOfInterest = `⌘` 
	
	fmt.Printf("plain string: ") 
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n") 
	
	fmt.Printf("quoted string: ") 
	// This verb escapes not only non-printable sequences, but also any non-ASCII bytes while interpreting UTF-8.
	fmt.Printf("%+q", placeOfInterest) 
	fmt.Printf("\n") 
	
	fmt.Printf("hex bytes: ") 
	for i := 0; i < len(placeOfInterest); i++ { 
		fmt.Printf("%x ", placeOfInterest[i]) 
	} 
	fmt.Printf("\n")
}
```

The output is:
```plaintext
plain string: ⌘
quoted string: "\u2318"
hex bytes: e2 8c 98 
```

We can see that the "Place of interest" symbol ⌘, is represented by three [bytes](binary_system.md) and that those bytes are the [UTF-8](ascii_unicode_and_utf8.md) encoding of the hexadecimal value `2318`.

>The UTF-8 representation of the string was created when the source code was written. **Source code in Go is defined to be UTF-8 text**. This means that any string literal (raw or not) is always going to be valid UTF-8, but during runtime, a string **value** is just a slice of bytes, and therefore, it's not guaranteed to be composed of valid UTF-8 bytes. It's also possible to create invalid UTF-8 string literals using escape sequences like `\xff`.
### Bytes, characters and runes
Since the definition of [character is ambiguous in computing](ascii_unicode_and_utf8.md#Code%20points%20and%20characters), the correct term to refer to individual characters into a string would be **code point**, but since it's a bit of a mouthful, Go introduces a shorter term for the concept: `rune`.

A `rune` means the same as **code point** but it's also an alias for the type `int32` so programs can be clearer when an integer value represents a code point.

Therefore, a character constant (individual characters in a string) are a `rune` constant. Individual runes are represented by single quotes: `'⌘'`, this is a rune with an integer value `0x2318`.
### Range loops
We've seen what happens with a regular [`for loop`]() when we iterate a string. A [`for range loop`]() decodes one [UTF-8 encoded](ascii_unicode_and_utf8.md) `rune` on each iteration. **Each time around the loop, the index of the loop is the starting position of the current rune, measured in bytes**.
```go
const nihongo = "日本語" 
for index, runeValue := range nihongo {
	fmt.Printf("%#U starts at byte position %d\n", runeValue, index) 
}
```

The output shows how each code point occupies multiple bytes:
```plaintext
U+65E5 '日' starts at byte position 0
U+672C '本' starts at byte position 3
U+8A9E '語' starts at byte position 6
```

>If a `for range` loop isn't sufficient, you can try the [unicode/utf8](https://go.dev/pkg/unicode/utf8/) standard library.
## Understanding
- Go uses **static typing** so errors are caught at compile time and performance is optimized early.
- It avoids **implicit conversions**, forcing explicit casting to prevent hidden bugs and precision loss.
- Type inference reduces verbosity while still preserving strict type safety.
- Custom types add **meaning and safety**, preventing misuse of similar-looking values.
- Strings are **byte sequences**, so Unicode text requires `rune` for correct character handling.
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
- **Learning curve**: Understanding nuances like `rune` vs `byte`, or UTF-8 handling, takes time.
- **Precision pitfalls**: Floating-point limitations can still cause subtle bugs.
- **Overflow behavior**: Silent wraparound at runtime can lead to unexpected results if not handled carefully.
- **Complexity with interfaces**: Type assertions require care (and checks) to avoid runtime panics.
## Examples
- A chat app that supports emojis and international languages (`rune` usage).
- Reading a file or handling HTTP request bodies (`byte` heavy).
- An e-commerce system distinguishing `UserID` from `ProductID`.  Both might be integers, but represent different concepts.
## References
### Connects with
- [Go variables](go_variable_and_constants.md)
- [Go interfaces]()
- [Go Comma-OK Idiom]()
- [Go Switches]()
- [Go Slices]()
- [Binary System](binary_system.md)
- [ASCII, Unicode and UTF-8](ascii_unicode_and_utf8.md)
- [Go loops]()
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
Q: Why are `rune` types important?  
A: They allow correct handling of Unicode characters in strings.
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
Q: What is a string in Go?  
A: A read-only slice of bytes.
Q: Why is string indexing potentially misleading?  
A: Because it accesses bytes, not characters.
Q: How are characters properly represented in Go strings?  
A: Using runes, which represent Unicode code points.
Q: What is the difference between raw and interpreted string literals?  
A: Raw strings preserve characters exactly, while interpreted strings process escape sequences.
Q: How does a `for range` loop iterate over strings?  
A: It decodes UTF-8 and returns one rune at a time.
Q: What does the index in a `for range` string loop represent?  
A: The byte position where the rune starts.
Q: Why are Go strings not guaranteed to be valid UTF-8 at runtime?  
A: Because they are just byte slices and can contain arbitrary data.
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
Q: What is a common failure mode related to strings?  
A: Treating each byte as a character instead of handling runes properly.
Q: What is a real-world use case for runes?  
A: Handling multilingual text or emojis correctly.
Q: What is a real-world use case for byte-level operations?  
A: Processing file data or HTTP request bodies.