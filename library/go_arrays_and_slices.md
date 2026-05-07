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
Very short resume with only the essential information needed.
# Go Arrays
Go [Array]() definition takes the following form `var varName[10]int`.

The array size is part of the type so an `varName[5]` do not have the same type as `varName2[6]`.

An array literal follows this structure: `[3]bool{true, true, false}`.
## Slices
It's a **dynamically-sized**, flexible view **into the elements of an array**. Since a slice is nothing more than a view into the elements of an array, a slice doesn't store any data, it just describes a section of an underlying array. If we change elements in the slice, we change elements in the underlying array and therefore, other slices that share the same array are going to see the changes.

A slice definition takes the form: `[]T` where `T` is the type. There's no size specification.

An array literal follows this structure: `[]bool{true, true, false}`.

Slicing range are `[start:end[` and can be specified as:
- `a[1:10]`: creates a slice from the second element to the 9th.
- `a[:10]`: creates a slice from the first element to the 9th.
- `a[0:]`: creates a slice from the beginning to the last element.
- `a[:]`: creates a slice with all array elements.
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
**Creating a slice from an array**:
```go
var arr[3]int = []int{1, 2, 3}
// slice is generated from the underlying array
// range operator is [start:end[
var slc []int = arr[0:4]
```
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
