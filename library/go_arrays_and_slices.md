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
[Arrays]() in Go have fixed size and are copied on assignment/pass-by-value. **Slices** are dynamic views over arrays that share the same underlying data and grow automatically when needed.
# Go Arrays and Slices
## Arrays
- [Arrays]() in Go have a fixed size that is part of their [type](go_data_types.md).
- Arrays with different sizes are essentially of different types.

Arrays are value types. Assigning or passing an array copies all elements.
```go
b := a // copies the entire array
```

This behavior exists because arrays represent the full data structure itself, not a reference to its first element.

>To avoid copying, pass a pointer to the array.
## Slices
It's a lightweight structure composed of:
- A pointer to an underlying array
- Length
- Capacity
```plaintext
- A[Slice Header] --> B[Pointer to Array]    
- A --> C[Length]    
- A --> D[Capacity]
```

**Slices do not store data themselves. They describe a section of an array.**

Because slices share the same underlying array:
- Modifying a slice modifies the array
- Other slices referencing the same array see the changes
```go
a := [...]int{1, 2, 3, 4, 5}
b := a[:2]    // []int{1, 2}
c := a[:2]    // []int{1, 2}
a[0] = 6      // [...]int{6, 2, 3, 4, 5}
fmt.Prinln(b) // []int{6, 2}
fmt.Prinln(c) // []int{6, 2}
```

>This design allows slices to be efficient and flexible without copying data unnecessarily.

The zero value of a slice is `nil`. And it has length and capacity as 0:
```go
var s []int
```
### Creating Slices
There are three main ways to create slices:
1. [Slice literals](#Array%20and%20Slice%20literals)
2. [Slicing existing arrays or slices](#Creating%20a%20Slice%20from%20an%20Array)
3. [Using `make`](#Creating%20Slices%20with%20`make`)
### Length and Capacity
- `len(slice)` returns the number of visible elements
- `cap(slice)` returns how much of the underlying array is available **starting from the slice pointer**. 

This means that the capacity of a slice is equal to the number of elements counting from it's referenced start index. It also means that the length can grow through re-slicing if enough capacity exists.

```go
a := [...]int{1, 2, 3, 4, 5}
b := a[:2] // []int{1, 2}, len(2), cap(5)
// Re-slicing b to change length and capacity
b := b[2:] // []]int{3, 4, 5}, len(2), cap(3)
```

>Re-slicing does not copy the underlying array. Because of this, a small slice can keep a large array alive in memory.
### Slice Ranges
- Slices can be re-sliced as long as the new range stays within **capacity**.
- Slices cannot be re-sliced below zero to access earlier elements.

Slice ranges use the format:
```go
//[start:end]
a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
b := a[1:10] 
```
### Slice Growth
When appending beyond capacity, Go allocates a new underlying array and copies the data.
- **Small slices**:  Capacity usually doubles. This reduces frequent allocations and copying.
- **Large slices**:  Growth becomes more conservative, usually around 25%. This reduces memory waste and large copy costs.
## Best practices
### Preallocating Slices
Use `make` with capacity when you know the expected size:
```go
nums := make([]int, 0, 500)
```

This reduces reallocations and copying.
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
- The collection size changes frequently
- Large copies would hurt performance
### Slices
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
### Array and slice creation
#### Array and Slice literals
```go
// Array literal
[3]bool{true, true, false}

// Array with automatic size inference
[...]bool{true, true, false}

// Slice literal
[]bool{true, true, false}
```
#### Creating a Slice from an Array
Slices are views over arrays, so re-slicing can extend the visible range if capacity allows.
```go
arr := [3]int{1, 2, 3} 
slc := arr[0:2] // []int{1, 2}
```
#### Creating Slices with `make`
`make` allocates the underlying array and returns a slice referencing it.
```go
a := make([]int, 5) // len=5 cap=5
```

Specifying capacity:
```go
a := make([]int, 0, 5) // len=0 cap=5
```
### Comparing Arrays
Two arrays are equal when:
- They have the same type
- Same size
- Same element order
```go
a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}
c := [3]int{3, 2, 1}
fmt.Println(a == b) // true
fmt.Println(a == c) // false
```
### Appending to a Slice
```go
a := make([]int, 1)a = append(a, 1, 2)
```

If capacity is insufficient, Go allocates a larger array and copies the elements.
### Copying Between Slices
```go
arr1 := [4]int{1, 2, 3, 4}
arr2 := [4]int{5, 6, 7, 8}
n := copy(arr2[:], arr1[:])
fmt.Println(arr2) // [1 2 3 4]
fmt.Println(n)    // 4
```

`copy`:
- Copies up to the smaller length
- Works with overlapping slices safely
- Returns the number of copied elements
### Mixing types in an Array or Slice
#### Using Interfaces
This approach provides type-safe polymorphism. But it's fragile as it requires type assertions for accessing element data.
```go
type Describable interface {	
	Describe() string
}
type Person struct{ 
	Name string 
}
func (p Person) Describe() string {	
	return "Person: " + p.Name
}
type Car struct{
	Model string 
}
func (c Car) Describe() string {
	return "Car: " + c.Model
}
items := []Describable{	
	Person{"Alice"},	
	Car{"Tesla"},
}
```
#### Empty interface
Using `interface{}` allows mixed types but sacrifices type safety:
```go
items := []interface{}{"hello", 42, 3.14, true}
```
Using interfaces provides polymorphism but may introduce unnecessary abstractions.

Using [Generics]() is preferable when the slice only needs to support different concrete types at different usages rather than mixed types simultaneously.
## References
### Connects with
- [Arrays]()
- [Go Generics]()
- [Go Interfaces]()
- [Go Data Types](go_data_types.md)
## Flashcards
Q: What is the main difference between arrays and slices in Go?  
A: Arrays have fixed size and are copied by value, while slices are dynamic views over arrays that share underlying data.
Q: Why are arrays with different sizes considered different types?  
A: Because the array length is part of the array’s type definition.
Q: What happens when an array is assigned or passed to a function?  
A: The entire array is copied.
Q: Why do arrays copy all elements on assignment?  
A: Because arrays represent the complete data structure itself rather than a reference.
Q: How can array copying be avoided?  
A: By passing a pointer to the array.
Q: What is a slice in Go?  
A: A lightweight structure describing a section of an underlying array.
Q: What are the components of a slice?  
A: A pointer to the array, a length, and a capacity.
Q: Do slices store their own data?  
A: No, they only reference data stored in an underlying array.
Q: What happens when multiple slices reference the same array?  
A: Changes made through one slice are visible to all slices sharing that array.
Q: Why are slices efficient?  
A: Because they avoid unnecessary copying by sharing underlying memory.
Q: What is the zero value of a slice?  
A: `nil`.
Q: What are the length and capacity of a nil slice?  
A: Both are zero.
Q: What are the main ways to create slices?  
A: Slice literals, slicing existing arrays/slices, and using `make`.
Q: What does `len(slice)` return?  
A: The number of visible elements in the slice.
Q: What does `cap(slice)` return?  
A: The number of elements available in the underlying array starting from the slice pointer.
Q: Why can a slice’s length grow without allocation?  
A: Because re-slicing can expose more of the existing capacity.
Q: Does re-slicing copy the underlying array?  
A: No, it only changes the visible range.
Q: Why can small slices cause memory issues?  
A: Because they may keep a large underlying array alive in memory.
Q: What is the syntax for slice ranges?  
A: `[start:end]`.
Q: What limits re-slicing operations?  
A: The new range must stay within the slice’s capacity.
Q: What happens when appending exceeds slice capacity?  
A: Go allocates a new array and copies the existing elements.
Q: How do small slices typically grow?  
A: Capacity usually doubles.
Q: How do large slices typically grow?  
A: Growth becomes more conservative, usually around 25%.
Q: Why does Go use different growth strategies for slices?  
A: To balance allocation cost, memory usage, and performance.
Q: Why is preallocating slices considered a best practice?  
A: It reduces reallocations and copying when the expected size is known.
Q: When should arrays be used?  
A: When collection size is fixed, copy semantics are useful, or collections are small.
Q: When should slices be preferred?  
A: When collections are dynamic or shared data access is needed.
Q: Why are slices more flexible than arrays?  
A: Because they support dynamic sizing and efficient shared memory usage.
Q: What is a key trade-off of slices sharing memory?  
A: Mutating one slice can unintentionally affect others sharing the same array.
Q: What is a key trade-off of automatic slice growth?  
A: Growth may trigger allocations and data copying.
Q: What does `make` do when creating slices?  
A: It allocates the underlying array and returns a slice referencing it.
Q: How are arrays compared for equality?  
A: They must have the same type, size, and element order.
Q: What does the `copy` function do with slices?  
A: It copies elements between slices and returns the number of copied elements.
Q: How many elements does `copy` transfer?  
A: Up to the smaller length between source and destination.
Q: Why are interfaces useful in slices?  
A: They allow polymorphism by storing different types implementing the same behavior.
Q: What is a drawback of interface-based slices?  
A: Accessing concrete data may require type assertions.
Q: What is the drawback of using `interface{}` slices?  
A: They sacrifice type safety.
Q: When are generics preferable to interface-based mixed slices?  
A: When supporting different concrete types separately rather than mixing types simultaneously.
Q: What is the purpose of slice literals?  
A: To create slices directly with predefined values.
Q: What is the purpose of slicing an array?  
A: To create a slice view over part of the array without copying data.