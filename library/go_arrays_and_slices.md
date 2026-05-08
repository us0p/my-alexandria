---
id: 20260507-go_arrays_and_slices
type: concept
status: draft
tags:
  - go
  - programing_language
created: 2026-05-07
---
## TL;DR
[Arrays]() in Go have fixed size and are copied on assignment/pass-by-value. Slices are dynamic views over arrays that share the same underlying data and grow automatically when needed.
# Go Arrays and Slices
- Arrays in Go have a fixed size that is part of their [type](go_data_types.md).
- Slices are dynamic views over arrays and are the primary way to work with collections in Go.
## Understanding
### Arrays
An array definition follows this format:
```go
var arr [10]int
```

The array size is part of the [type](go_data_types.md):
```go
var a [5]intvar b [6]int
```

`a` and `b` are different types because their sizes differ.

Array literals can be created like this:
```go
[3]bool{true, true, false}
```

Or with automatic size inference:
```go
[...]bool{true, true, false}
```

Arrays are value types. Assigning or passing an array copies all elements.
```go
b := a // copies the entire array
```

This behavior exists because arrays represent the full data structure itself, not a reference to its first element.

To avoid copying, pass a pointer to the array.
### Slices
A slice is a lightweight structure composed of:
- A pointer to an underlying array
- Length
- Capacity
```plaintext
graph TD    
- A[Slice Header] --> B[Pointer to Array]    
- A --> C[Length]    
- A --> D[Capacity]
```

Slices do not store data themselves. They describe a section of an array.

Because slices share the same underlying array:
- Modifying a slice modifies the array
- Other slices referencing the same array see the changes

This design allows slices to be efficient and flexible without copying data unnecessarily.

The zero value of a slice is `nil`.
```go
var s []int
```

A `nil` slice has:
- length = 0
- capacity = 0
### Creating Slices
There are three main ways to create slices:
1. Slice literals
2. Slicing existing arrays or slices
3. Using `make`

Slice type definition:
```go
[]T
```

Example literal:
```go
[]bool{true, true, false}
```
### Slice Ranges
Slice ranges use the format:
```plaintext
[start:end]
```

Examples:
```go
// second element to nintha[:10]  
// beginning to nintha[0:]   
// beginning to enda[:]    
// entire slice
a[1:10] 
```

Slices can be re-sliced as long as the new range stays within capacity.

Slices cannot be re-sliced below zero to access earlier elements.
### Length and Capacity
- `len(slice)` returns the number of visible elements
- `cap(slice)` returns how much of the underlying array is available starting from the slice pointer

The length can grow through re-slicing if enough capacity exists.

Re-slicing does not copy the underlying array. Because of this, a small slice can keep a large array alive in memory.
### Slice Growth
When appending beyond capacity, Go allocates a new underlying array and copies the data.
#### Small slices
Capacity usually doubles:
```plaintext
2 → 4 → 8 → 16
```

This reduces frequent allocations and copying.
#### Large slices
Growth becomes more conservative, usually around 25%:
```plaintext
1024 → 1280 → 1600
```

This reduces memory waste and large copy costs.
## When to Use
### Arrays
Use arrays when:
- The size is fixed
- Copy semantics are desired
- Working with small collections
### Slices
Use slices when:
- The collection size changes dynamically
- Sharing underlying data is useful
- You need flexible and efficient collection handling
## When NOT to Use
### Arrays
Avoid arrays when:
- The collection size changes frequently
- Large copies would hurt performance
### Slices
Avoid relying heavily on automatic growth when:
- Large reallocation become expensive
- Excess capacity wastes memory
- Holding a small slice unintentionally keeps a large array in memory
## Trade-offs
### Arrays
- Fixed size
- Full copies on assignment/pass
- Strong type safety
### Slices
- Share underlying memory
- Mutations affect shared data
- Automatic growth may trigger allocations and copies
- More flexible than arrays
## Examples
### Comparing Arrays
```go
a := [3]int{1, 2, 3}b := [3]int{1, 2, 3}c := [3]int{3, 2, 1}fmt.Println(a == b) // truefmt.Println(a == c) // false
```

Two arrays are equal when:
- They have the same type
- Same size
- Same element order
### Creating a Slice from an Array
```go
arr := [3]int{1, 2, 3}// slice generated from the underlying arrayslc := arr[0:2] // []int{1, 2}// re-slicingslc = slc[:3] // []int{1, 2, 3}
```

Slices are views over arrays, so re-slicing can extend the visible range if capacity allows.
### Creating Slices with `make`
```go
a := make([]int, 5) // len=5 cap=5
```

Specifying capacity:
```go
a := make([]int, 0, 5) // len=0 cap=5
```

`make` allocates the underlying array and returns a slice referencing it.
### Appending to a Slice
```go
a := make([]int, 1)a = append(a, 1, 2)
```

If capacity is insufficient, Go allocates a larger array and copies the elements.
### Copying Between Slices
```go
arr1 := [4]int{1, 2, 3, 4}arr2 := [4]int{5, 6, 7, 8}n := copy(arr2[:], arr1[:])fmt.Println(arr2) // [1 2 3 4]fmt.Println(n)    // 4
```

`copy`:
- Copies up to the smaller length
- Works with overlapping slices safely
- Returns the number of copied elements
### Using Interfaces for Mixed Types
```go
type Describable interface {	Describe() string}type Person struct{ Name string }func (p Person) Describe() string {	return "Person: " + p.Name}type Car struct{ Model string }func (c Car) Describe() string {	return "Car: " + c.Model}items := []Describable{	Person{"Alice"},	Car{"Tesla"},}
```

This approach provides type-safe polymorphism.
### Preallocating Slices
Use `make` with capacity when you know the expected size:
```go
nums := make([]int, 0, 500)
```

This reduces reallocations and copying.
### Heterogeneous Slices
Using `interface{}` allows mixed types but sacrifices type safety:
```go
items := []interface{}{"hello", 42, 3.14, true}
```

Using interfaces provides polymorphism but may introduce unnecessary abstractions.

Using [Generics]() is preferable when the slice only needs to support different concrete types at different usages rather than mixed types simultaneously.
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
