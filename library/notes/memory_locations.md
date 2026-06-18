---
id: 20260618-memory_locations
type: concept
status: draft
tags:
  - computer_theory
  - programing_language
created: 2026-06-18
---
## TL;DR
Very short resume with only the essential information needed.
# Memory Locations
There are two memory locations in which programs store transient data during execution.
## Heap
Contains values that are referenced outside of a function. Usually by explicitly allocating memory to store complex data structures like arrays. The heap is a [graph]() where objects are represented as nodes which are referenced by other objects in the heap. If the heap is not cleared the memory will continue to grow. Thin process is manual in some languages like C and automatic in languages with **garbage collector** like JavaScript, Go, etc.
## Stack
[LIFO](keywords.md#Last%20In%20First%20Out%20(LIFO)) structure which stores the values and results of function calls. Calling a new function within a function pushes a new **frame** onto the stack. When the called function returns its stack frame is popped from the stack. Most programming languages returns a **stack trace** which displays the functions that have been called leading up to that point.
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
