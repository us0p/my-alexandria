---
id: 20260522-go_pointers
type: concept
status: draft
tags:
  - go
  - programing_language
created: 2026-05-22
---
## TL;DR
Variables storing memory addresses of other variables. Enable efficient memory usage and allow functions to modify values. Declared with `*Type`, address obtained with `&`. No pointer arithmetic for safety.
# Go pointers
A pointer holds the memory address of a value. The type `*T` is a pointer to a `T` value. It's zero value is `nil`.

The `&` operator generates a pointer to its operand. The `*` operator denotes the pointer's underlying value, this is known as **dereferencing** or **indirecting**.
```go
i := 42
p = &i
fmt.Println(*p) // read i throught the pointer p. Prints 42
```

>Go has no pointer arithmetic.

[Slices](go_arrays_and_slices.md), [Maps](go_maps.md), [Channels]() are reference types, so they use a **descriptor** which contains a **pointer** with the address of the structure with the actual data in memory.

Pointers can also be used with [Structs pointer receivers](go_structs.md#Pointer%20Receiver).
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
