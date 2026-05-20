---
id: 20260130-hex-arch
type: concept
status: draft | refined
tags: [arquitetura, sistemas, teoria]
created: 2026-01-30
---
## TL;DR
First-class citizens in Go. Support parameters and return values. Can be assigned to variables, passed as arguments, returned from other functions. Support multiple return values, named returns, and variadic parameters.
# Go Functions
In Go, functions are **first-class citizens**, meaning they can be:
- Assigned to variables.
- Passed as arguments to other functions.
- Return from other functions.

A typical function in Go looks like this:
```go
func add(x int, y int) int {
	return x + y
}
```

When two or more named function parameters share a type, you can omit the type from all but the last: `func (x, y int)`.

A function can return any number of results:
```go
func swap(x, y string) (string, string) {
	return y, x
}
```

Return values may be named. If so, they are treated as [variables](go_variable_and_constants.md) defined at the top of the function.

>These names should be used to document the meaning of the return values.

A `return` statement without arguments, returns the named return values. This is known as **"naked"** return.

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

>Naked return should only be used in short functions.

Some conventions for functions with multiple return values are:
- `Value, Error`: For when a function might produce an error.
- `Value, Boolean`: When a function need to indicates success/failure e.g. type assertions.
## Variadic functions
Uses `...` to handle a variable number of arguments.
```go
func sum(numbers ...int) (total int) {
	for _, num := range numbers {
		total += num
	}
	return
}

// Call example
sum(1, 2, 3, 4, 5)
```
- There can be **only one** variadic parameter and it must be the **last parameter** in the function.
- Inside the function, the variadic parameter is treated as a [slice](go_arrays_and_slices.md) of the specified type. Note that this means that an slice is created internally leading to memory allocation.
- Since variadic arguments are converted to a slice, when zero arguments are passed, that parameter is `nil` and not an empty slice.

You can unpack slices directly into variadic functions using the `...` operator:
```go
nums := []int{1, 2, 3, 4, 5}

// Works same as above
sum(nums...)
```

>Unpacking, likewise variadic parameter definition, must be the last argument in the function call. This is wrong: `sum(nums..., 1)`.
## Anonymous functions
Functions without an identifier. Can be assigned to variables (which becomes the function's identifier) and be **immediately invoked (Immediately Invoked Function Expression - IIFE)**.
```go
func main() {
	sayHello := func(name string) {
		fmt.Prinln("Hello,",name)
	}
	sayHello("Alice")
}

// immediately invoked

	return fmt.Printf("Hello %s!\n", name)
}("Luan")
```

Common use cases include:
- Create [closures]().
- Create [concurrent tasks]().
- Creating [table driven tests]().
## Call by value
Copies the actual value of an **argument** into the **function parameter**. This means that changes to the **function parameter** don't affect the **function argument**.

In Go, this is the default and whenever you call a function you need to remember that the arguments are going to be copied to the function's scope.

Since [maps](go_maps.md) are a **pointer to a runtime hash table structure**. When we pass a map to a function, the pointer is copied, and thus, changes to the underlying value are reflected in the original variable.

For [slices](go_array_and_slices.md) we must remember that changes to the function argument **may or may not** be reflected in the original variable.

This happens because a Slice is a descriptor that defines:
- A pointer to an array.
- The length.
- The capacity.
So if you use `append()` inside the function and it allocates a new array because capacity was not enough. The slice in the function is going to have a new array reference while the outside variable will not.

In summary, every **reference type** is a type that contains a header (descriptor) that points to shared data. When passed to a function without a pointer, the descriptor is always going to be copied.
## Call by reference
Here the **function parameter** receives the **address of the argument**. Which means that changes to the **argument's value** can be made from within the function.

In this mode, there's no value copying which is good for saving memory. The drawback is that the function must take care with `nil` values.
## Understanding
- Functions can be assigned to variables, returned from or passed by to other functions.
- A function can receive a variadic number of arguments by using the variadic operator `...` .
- There can only be one variadic argument in a function.
- In go all functions are called by value.
- Reference types in function arguments have their descriptors copied, but the reference still makes the underlying data structure mutable.
- For mutability you can expect a pointer as a function argument and then you can mutate its original value. This kind of function is called by reference.
## When to Use
- Use when you need to encapsulate behavior or calculations.
- Use anonymous functions when you need to execute one time actions or when you need to have stateful functions.
- Use naked returns when the function is is short and simple.
## When NOT to Use
- Avoid naked returns when your function is long or when it's too complex.
- Avoid using functions without pointer arguments for types that are large. This will cause the data to be copied in memory leading to great memory inefficiency.
## Trade-offs
- Simplicity x Performance: Functions called by value are simple to use but can degrade application performance if big data keeps being copied.
- Flexibility x Complexity: Using functions with explicit references makes them more flexible and more performant, but adds the complexity of checking for `nil` edge cases.
## Examples
### Handling multiple types with variadic functions
```go
func logItems(prefix string, items ...any) {
	for i, item := range items {
		fmt.Printf("%s %d: %v\n", prefix, i+1, item)
	}
}
```
## References
- [Variables](go_variable_and_constants.md)
- [Maps](go_maps.md)
- [Slices](go_array_and_slices.md)
- [Closures]()
- [Go concurrency model]()
- [Table driven tests]()
## Questions
- What are Go [iterators](https://blog.alexoglou.com/posts/iterators-golang/)?
## Flashcards
- Q: What does it mean that functions are first-class citizens in Go?
- A: Functions can be assigned to variables, passed as arguments, and returned from other functions.
- Q: Can Go functions return multiple values?
- A: Yes, Go functions can return any number of values.
- Q: What shorthand can be used when multiple function parameters share the same type?
- A: The type can be omitted from all but the last parameter.
- Q: What are named return values in Go?
- A: They are return values declared as variables at the top of the function.
- Q: What is a naked return in Go?
- A: A `return` statement without arguments that returns named return values automatically.
- Q: When should naked returns be used?
- A: Only in short and simple functions.
- Q: What is a common convention for functions that may fail?
- A: Returning `Value, Error`.
- Q: What is a common convention for functions that indicate success or failure?
- A: Returning `Value, Boolean`.
- Q: What are variadic functions in Go?
- A: Functions that accept a variable number of arguments using `...`.
- Q: How many variadic parameters can a Go function have?
- A: Only one variadic parameter.
- Q: Where must a variadic parameter be placed in a function signature?
- A: It must be the last parameter.
- Q: How is a variadic parameter represented inside a function?
- A: As a slice of the specified type.
- Q: What happens internally when variadic arguments are passed?
- A: A slice is created internally, causing memory allocation.
- Q: What is the value of a variadic parameter when no arguments are passed?
- A: It is `nil`, not an empty slice.
- Q: How can a slice be passed to a variadic function?
- A: By unpacking it with the `...` operator.
- Q: Where must unpacked variadic arguments appear in a function call?
- A: As the last argument.
- Q: What are anonymous functions in Go?
- A: Functions without an identifier.
- Q: Can anonymous functions be assigned to variables?
- A: Yes, and the variable becomes the function identifier.
- Q: What is an IIFE in Go?
- A: An Immediately Invoked Function Expression.
- Q: What are common use cases for anonymous functions?
- A: Closures, concurrent tasks, and table-driven tests.
- Q: What does call by value mean in Go?
- A: Function arguments are copied into function parameters.
- Q: Do changes to function parameters affect the original arguments in call by value?
- A: No, changes affect only the copied values.
- Q: What is the default argument passing behavior in Go?
- A: Call by value.
- Q: Why do map modifications inside functions affect the original map?
- A: Because maps are pointers to runtime hash table structures and the pointer is copied.
- Q: Why may slice modifications inside functions not affect the original slice?
- A: Because `append()` may allocate a new underlying array.
- Q: What does a slice descriptor contain?
- A: A pointer to an array, length, and capacity.
- Q: What is a reference type in Go?
- A: A type containing a descriptor that points to shared data.
- Q: What happens when a reference type is passed to a function?
- A: Its descriptor is copied.
- Q: What is call by reference?
- A: Passing the address of an argument so the function can modify the original value.
- Q: What is an advantage of call by reference?
- A: It avoids copying large values and saves memory.
- Q: What is a drawback of call by reference?
- A: Functions must handle possible `nil` values.
- Q: Why should large types without pointers be avoided in function arguments?
- A: Because copying them can be memory inefficient.
- Q: Why are functions useful in Go?
- A: They encapsulate behavior and calculations.
- Q: When should anonymous functions be used?
- A: For one-time actions or stateful behavior.
- Q: What trade-off exists with call by value?
- A: Simplicity versus potential performance costs from copying large data.
- Q: What trade-off exists with explicit references?
- A: Better performance and flexibility versus added complexity handling `nil` values.
- Q: Can functions in Go be returned from other functions?
- A: Yes, functions can be returned like any other value.
- Q: Can functions in Go be passed as arguments to other functions?
- A: Yes, functions can be passed as arguments.