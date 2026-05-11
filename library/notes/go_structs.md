---
id: 20260509-go_structs
type: concept
status: draft
tags:
  - go
  - programing_language
created: 2026-05-09
---
## TL;DR
Structs in Go are custom data types that group related fields, similar to classes but without attached methods. They support object-oriented design through composition, especially via struct embedding, which lets one struct include another and access its fields directly. Struct tags add metadata (like JSON serialization rules) to fields, making structs essential for APIs and data handling.
# Go Structs
It's a collection of fields, in which fields are accessed using a dot.
```go
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v.X) // 1
}
```

You can also use the `Name:` syntax in which case the order of the fields is irrelevant.
## Pointer to structs
To access a field when you have a pointer to a struct you could use **pointer dereferencing** `(*p).field`. However since that can become cumbersome the language allow use to simply write `p.field`.
## Methods
Methods are specified with a receiver which specifies the type the method belongs to.

The receiver can be either a **value** or a **pointer** receiver.
### Value Receiver
Creates a **copy** of the struct when the method is called so **modifications to fields do not affect the original struct**.
```go
type User struct {
	Username string
	Email    string
}

// u is the receiver
func (u User) PrintInfo() {
	fmt.Printf("Username: %s, Email: %s\n", u.Username, u.Email)
}
```
### Pointer Receiver
Allows the method to modify the original struct fields directly:
```go
func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}
```
## Memory Alignment
Means that data in memory must be located at addresses that are multiple of certain values. The size of a data type determines its alignments requirement. For example, `int32` requires alignment to 4 bytes.

Memory alignment is critical for CPU performance. If a variable is not properly aligned, the CPU may need multiple memory accesses, leading to performance degradation.

In Go, structs are aligned in memory. Different data types have different alignment requirements, and the compiler may insert padding bytes between struct fields to meet requirements.

**Rules for struct memory alignment**:
- **Field alignment**: Each field address must meet its type's alignment requirements. The compiler may insert padding bytes.
- **Struct alignment**: The size of a struct must be a multiple of the **largest alignment** requirement among its fields.

```go
import (
	unsafe
	fmt
)

type Example struct {
	a int8  // 1 byte
	b int32 // 4 bytes
	c int 8 // 1 byte
}

fmt.Println(unsafe.Sizeof(Example{})) // 12
```

Analysis:
- `a` is `int8` occupying 1 byte, aligned to 1.
- `b` is `int32` requiring alignment to  4 bytes. The compiler insert 3 padding bytes between `a` and `b` to align `b`'s address to 4.
- `c` is `int8` requiring 1 byte, but the struct's total size must be a multiple of 4 (the largest alignment requirement). The compiler adds 3 padding bytes at the end.
### Optimizing memory alignment
You can rearrange struct fields to minimize padding and reduce memory usage.
```go
type Optimized struct {
	b int32
	a int8
	a int8
}

fmt.Println(unsafe.Sizeof(Optimized{})) // 8
```

Since `b` is placed first, it's aligned to 4 bytes. `a` and `c` are placed consecutively (5th and 6th bytes), the compiler adds 2 padding bytes at the end which is more compact.

>Adjusting the order of fields can minimize padding and optimize memory usage.
## Nested structs
Enable one struct to include another one as a **field**. Methods aren't promoted.

```go
// Define the Address struct  
type Address struct {  
	City    string  
	Country string  
}  

func (a Address) Full() string {
	return a.City + " " + a.Country
}
  
// Define the User struct, which includes the Address struct  
type User struct {  
	Username string  
	Email    string  
	Address  Address // Nested struct  
}  
  
// Initialize the nested struct  
user := User{  
	Username: "alice",  
	Email:    "alice@example.com",  
	Address: Address{  
	  City:    "New York",  
	  Country: "USA",  
	},  
}  

user.Full()         // INVALID
user.Address.Full() // VALID
```
## Struct composition (embedding)
It's often used to share behavior.
```go
type User struct {
	Name string
	Address // note no type specification here, Address is embedded.
}

u := User{
	Name: "alice",
	Address: Address{
		City: "New York",
		Country: "USA",
	},
}

// Fields and methods are promoted
u.City     // instead of u.Address.City
u.Full()   // instead of u.Address.Full()
```

>In Go, there's no inheritance. Composition is encouraged instead.

If names are ambiguous, you must provide specification:
```go
type A struct {
	Name string
}

type B struct {
	Name string
	A	
}

B.Name   // B's Name field
B.A.Name // A's Name field
```
## When to compose vs when to nest
Use nested structs when:
- relationship should be explicit
- you want clear ownership
- field promotion would hurt readability
Use embedding when:
- you want method promotion
- you want composition-style reuse
- the embedded type is conceptually part of the outer type
## Empty Structs
- It's a struct with no fields.
- It occupies **zero bytes** in memory.
- Its memory address may or may not **be equal**. When memory escape to the [heap]() occurs, the addresses are equal, pointing to `runtime.zerobas`.
- Occupy no space when it's the only field in the struct and when it's the first or an intermediary field. If it's the last field, it occupies space equal to the previous field.
- They also occupy no spaces in [arrays or slices](go_arrays_and_slices.md): `[...]struct{}`.

Empty structs are well suited in scenarios where you don't need a value, just a representation. Since they don't occupy space in memory, they can improve the performance of your application. 

A good example is [channel]() signal transmission. Since some times the value is irrelevant and we just need a signal to unlock the channel's other end.
## Struct Tags
A field declaration may be followed by an optional string literal **tag**, which becomes an attribute for all the field.

The tags are made visible through [reflection interface]() and take part in the type identity (meaning same fields with different tags aren't equal).

```go
struct {
	x, y float64 "" // an empty string is like an absent tag
	name string "any string is permitted as a tag"
}
```

Tags can be used to add metadata for specific tags, like for example, [JSON serializing]().
## Understanding
- Provide a way of grouping values under the same identifier.
- It's possible to add behavior by adding methods to your structs using either a pointer (read/write) or value (read-only) receiver.
- Memory is padded for efficient access which can cause program to use more memory than necessary if care is not taken.
- Tags provide metadata that can be used for other packages to perform specific tasks like JSON serialization.
- Empty structs can be used to represent presence without filling memory with unused data.
## When to use
- You need to group related data under a single type.
- You want object-oriented style design through composition instead of inheritance.
- Use pointer receivers when:
	- Methods must modify the original struct.
	-  Struct is large and copying would be expensive.
- Use value receivers when:
	- Methods are read-only.
	- Struct is small and inexpensive to copy.
	- Immutability-like semantics are preferred.
- Use struct embedding when sharing behavior between types.
- Use struct nesting when ownership relationships should remain explicit.
## When NOT to Use
- Data has no meaningful grouping.
- A simpler primitive type is enough.
- Dynamic schemas are required (maps may fit better).
- Excessive nesting makes code hard to read.
- Multiple embedded structs create naming conflicts.
- Avoid value receivers when:
	- The struct is large.
	- Methods need to mutate state.
	- Copying introduces unnecessary allocations or performance costs.
## Trade-offs
- Clear data modeling and type safety. But can become verbose with deep nesting.
- Pointer receivers introduce mutability but can create bugs when state becomes too widespread.
- Value receivers are good for small read only structs, but can become expensive is struct becomes too large.
- Embedding allows for code reuse but can lead to hidden coupling and ambiguity.
## Examples
### Empty Structs
Preventing unkeyed struct initialization
```go
type MustKeyedStruct struct {
	Name string
	Age int
	_ struct{}
}

// OK
MustKeyedStruct{Name: "hello", age: 10}

// ERROR: Too few values
MustKeyedStruct{"hello", 10}
```

Used to represent something without identity:
```go
// Set data structure

// The base is a hash map, but we don't need the value associated with the key
// Just the key uniqueness
type Set struct {
	items map[interface{}]emptyItem
}

type emptyItem struct{}

var itemExists = emptyItem{}

// We keep only unique items using the unique key constraint from a hash map
// while we do not waste memory with the value using an empty struct.
func (s *Set) Add(item interface{}) {
	s.items[item] = itemExists
}
```
## References
- [Heap and Stack]()
- [Go Arrays and Slices](go_arrays_and_slices.md)
- [Go Concurrency]()
- [Go reflection interface]()
- [Go `encoding/json`package]()
## Flashcards
- Q: What are structs in Go?
- A: Structs are custom data types that group related fields under the same identifier, with fields accessed using dot notation.
- Q: How can struct fields be initialized in Go?
- A: Struct fields can be initialized positionally or using the Name: syntax, where field order becomes irrelevant.
- Q: How are fields accessed through a pointer to a struct in Go?
- A: Although pointer dereferencing syntax `(*p).field` can be used, Go allows the shorthand `p.field`.
- Q: What is a method receiver in Go?
- A: A receiver specifies the type a method belongs to.
- Q: What is the difference between value receivers and pointer receivers?
- A: Value receivers create a copy of the struct, while pointer receivers allow direct modification of the original struct.
- Q: When should value receivers be used?
- A: Use value receivers when methods are read-only, structs are small and inexpensive to copy, or immutability-like semantics are preferred.
- Q: When should pointer receivers be used?
- A: Use pointer receivers when methods must modify the original struct or when copying a large struct would be expensive.
- Q: What is memory alignment in Go structs?
- A: Memory alignment ensures data is stored at addresses matching alignment requirements of their types, improving CPU performance.
- Q: Why does the Go compiler insert padding bytes in structs?
- A: Padding bytes are inserted to satisfy alignment requirements of struct fields and the struct itself.
- Q: What are the rules for struct memory alignment in Go?
- A: Each field must meet its type alignment requirements, and the total struct size must be a multiple of the largest alignment requirement among its fields.
- Q: How can struct memory usage be optimized?
- A: Rearranging fields to reduce padding can minimize memory usage.
- Q: What are nested structs in Go?
- A: Nested structs are structs included as fields inside another struct.
- Q: Are methods promoted in nested structs?
- A: No, methods are not promoted in nested structs and must be accessed through the nested field.
- Q: What is struct composition or embedding in Go?
- A: Embedding allows one struct to include another directly, promoting its fields and methods.
- Q: How are promoted fields and methods accessed in embedded structs?
- A: Promoted fields and methods can be accessed directly from the outer struct without referencing the embedded type.
- Q: What should be done when embedded structs create naming ambiguities?
- A: The specific embedded struct must be explicitly referenced.
- Q: Why does Go encourage composition instead of inheritance?
- A: Go does not support inheritance and encourages composition for code reuse and behavior sharing.
- Q: When should nested structs be preferred over embedding?
- A: Use nested structs when relationships should remain explicit, ownership should be clear, or field promotion would hurt readability.
- Q: When should embedding be preferred over nesting?
- A: Use embedding when method promotion, composition-style reuse, or conceptual inclusion of the embedded type is desired.
- Q: What is an empty struct in Go?
- A: An empty struct is a struct with no fields that occupies zero bytes in memory.
- Q: Why are empty structs useful in Go?
- A: They represent presence without storing data and avoid unnecessary memory usage.
- Q: How do empty structs behave in memory?
- A: They occupy no space in arrays or slices and generally occupy no space in structs unless placed as the last field.
- Q: What happens to empty struct addresses when escaping to the heap?
- A: Their addresses may become equal and point to `runtime.zerobas`.
- Q: What is a common use case for empty structs?
- A: Empty structs are commonly used for signaling in channels or as values in set-like map implementations.
- Q: How can empty structs prevent unkeyed struct initialization?
- A: Adding an anonymous empty struct field forces initialization using named fields.
- Q: What are struct tags in Go?
- A: Struct tags are optional string literals attached to fields that provide metadata.
- Q: How are struct tags accessed?
- A: Struct tags are accessed through reflection.
- Q: Do struct tags affect type identity?
- A: Yes, structs with identical fields but different tags are considered different types.
- Q: What is a common use case for struct tags?
- A: Struct tags are commonly used for metadata such as JSON serialization rules.
- Q: What are the main benefits of structs in Go?
- A: Structs provide clear data modeling, type safety, object-oriented design through composition, and support metadata through tags.
- Q: What are some trade-offs of deep struct nesting?
- A: Deep nesting can make code verbose and harder to read.
- Q: What are the risks of pointer receivers?
- A: Pointer receivers introduce mutability, which can create bugs if state becomes too widespread.
- Q: What are the drawbacks of value receivers on large structs?
- A: Copying large structs can become expensive and introduce unnecessary performance costs.
- Q: What are the trade-offs of embedding?
- A: Embedding enables code reuse but can create hidden coupling and naming ambiguities.
- Q: When should structs not be used?
- A: Structs should not be used when data lacks meaningful grouping, simpler primitive types are enough, dynamic schemas are required, or excessive nesting hurts readability.