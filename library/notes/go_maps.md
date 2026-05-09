---
id: 20260508-go_maps
type: concept
status: draft
tags:
  - go
  - programing_language
created: 2026-05-08
---
## TL;DR
Built-in associative data type mapping keys to values. Reference types created with `make(map[KeyType]ValueType)` or map literals. Keys must be comparable types. Support insertion, deletion, lookup operations. Check existence with comma ok idiom: `value, ok := map[key]`.
# Go maps
Go's [Hash Map]() implementation. It offers fast lookups, adds, and deletes.
## Declaration and initialization
A Go map looks like this:
```go
map[KeyType]ValueType
```

Where `KeyType` may be any type that is **comparable**, and `ValueType` may be anything.
```go
// A map of strings to ints
var m map[string]int // initial value is nil
```

>A **comparable** types are boolean, numeric, [string](go_strings.md), pointer, [channel](), and [interface]() types, and [structs]() or [arrays](go_arrays_and_slices.md) that contain **only** those types.

A `nil` map behaves like an empty map when **reading**, but attempts to **write to a `nil` map will case a runtime panic**.

Use `make` to initialize a map:
```go
m = make(map[string]int)
```

The `make` function allocates and initializes a [Hash Map]() and returns a map value that points to it.

You can also use a map literal:
```go
m = map[string]int{
	"x": 0,
	"y": 0,
}
```
## Working with maps
- If you try to access a key that doesn't exist, you get the **value type's zero value**.
- The `len()` function returns the number of items in the map.
- The `delete()` function removes an entry from the map: `delete(map, "key")`.
- To iterate over the contents of a `map`, use the [`range`]() keyword.

>Note that when iterating over a map, the iteration order **is not specified and is not guaranteed to be the same from one iteration to the next**. If you require a stable iteration order you must maintain a separate structure that specifies that order.
## Comma-ok idiom
It's a language construct that allows you to access the value of something while asserting if the operation was successful or not. It can be used in maps, [type assertion](go_data_types.md#Type%20Assertions) and [channel receive](). You can think of it to be similar to error handling pattern in go.

When you try to access the value of a map,  the first value returned is the map's value, and the second is a [boolean](go_data_types.md) that tells you if the if the key existed in the map.
```go
// comma-ok idiom to test if key "route" exist in map m
i, ok := m["route"]
```
## Concurrency
**Maps are not safe for concurrent use**. If you need to interact with a map from concurrently executed [goroutines](), the access must be mediated by some kind of [mutex]().
## Understanding
- Allows for constant time CRUD operations.
- Keys must be **comparable**.
- It's **not safe** for concurrency.
- Must be initialized with `make` or using map literal syntax.
## When to Use
- Use it when you want to have access to a value based on another element. For example if you need to map a person's email to an name.
- It can be used as a simple cache, for faster operations if you have an unique key.
## When NOT to Use
- When you need linear access (use a [slice](go_arrays_and_slices.md) instead).
- When you don't have a unique way of differentiating keys.
## Trade-offs
- **Performance**: It's usually a performance optimization technique but needs to be used with care, underlying array might become full what can lead to memory costs of new allocation and copying.
- **Complexity**: Using it in concurrent application needs careful synchronization with [mutexes]()
## Examples
The following example uses a `map` to count the number of access for a specific route of a website per country.
```go
type Key {
	Path string
	Country string
}
func main() {
	hits := map[Key]int
	
	// key initialization
	
	// Increases count for existing entries
	// Creates new one for new keys (since zero value of int is 0) and sets to 1
	hits[key]++
}
```

If we had used a `map[string]map[string]int`, we would have to take care of verification, map initialization etc. The example shows clearly the benefits of type's zero values and why we would want to use a composite type like [`struct`]() as a map key.
## References
- [Go structs]()
- [Go types](go_data_types.md)
- [Go Blank identifier](go_variable_and_constants.md)
- [Go arrays and slices](go_arrays_and_slices.md)
- [Go concurrency]()
- [Hash Maps]()
- [Go strings](go_strings.md)
- [Go loops]()
## Flashcards
Q: What is a Go map?  
A: A Go map is a built-in associative data type that maps keys to values and provides fast lookups, inserts, and deletes.
Q: What is the syntax of a Go map type?  
A: A Go map type uses the syntax `map[KeyType]ValueType`.
Q: What types can be used as map keys in Go?  
A: Map keys must be comparable types.
Q: What are examples of comparable types in Go?  
A: Comparable types include booleans, numbers, strings, pointers, channels, interfaces, and structs or arrays containing only comparable fields.
Q: What is the initial value of an uninitialized map?  
A: The initial value of an uninitialized map is nil.
Q: How does a nil map behave when reading values?  
A: A nil map behaves like an empty map when reading values.
Q: What happens when writing to a nil map?  
A: Writing to a nil map causes a runtime panic.
Q: How do you initialize a map in Go?  
A: A map can be initialized using make or a map literal.
Q: What does the make function do for maps?  
A: The make function allocates and initializes the underlying hash map and returns a map value pointing to it.
Q: What happens when accessing a non-existent key in a map?  
A: Accessing a non-existent key returns the zero value of the map's value type.
Q: What does the `len()` function return for a map?  
A: The `len()` function returns the number of items in the map.
Q: How do you remove an entry from a map?  
A: You remove an entry using the delete() function.
Q: How do you iterate over a map in Go?  
A: You iterate over a map using the range keyword.
Q: Is map iteration order guaranteed in Go?  
A: No, map iteration order is not specified and is not guaranteed to remain the same between iterations.
Q: What should be done if stable iteration order is required?  
A: A separate structure must be maintained to preserve stable ordering.
Q: What is the comma-ok idiom in Go?  
A: The comma-ok idiom is a language construct that returns a value along with a boolean indicating whether an operation succeeded.
Q: How is the comma-ok idiom used with maps?  
A: It returns the value associated with a key and a boolean indicating whether the key exists.
Q: What does the boolean value in the comma-ok idiom represent?  
A: The boolean indicates whether the key existed in the map.
Q: Are Go maps safe for concurrent use?  
A: No, Go maps are not safe for concurrent use.
Q: How should concurrent access to maps be handled?  
A: Concurrent access to maps should be synchronized using mechanisms like mutexes.
Q: What is a major advantage of using maps?  
A: Maps allow constant time CRUD operations in most cases.
Q: When should maps be used?  
A: Maps should be used when values need to be accessed efficiently through unique keys.
Q: What is an example use case for a map?  
A: A map can associate a person's email with their name or act as a cache for quick lookups.
Q: When should maps not be used?  
A: Maps should not be used when linear access is needed or when keys are not uniquely identifiable.
Q: What performance trade-off exists when using maps?  
A: Map growth can cause additional memory allocations and copying when the underlying structure becomes full.
Q: What complexity trade-off exists when using maps concurrently?  
A: Concurrent use requires careful synchronization, increasing complexity.
Q: Why are zero values useful in maps?  
A: Zero values allow automatic initialization-like behavior for missing keys during operations such as incrementing counters.
Q: Why can structs be useful as map keys?  
A: Structs can combine multiple comparable fields into a single unique key.