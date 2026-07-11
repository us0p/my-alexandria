---
id: 20260711-goroutine_infinite_stack
type: concept
status: draft
tags:
  - go
created: 2026-07-11
---
## TL;DR
Goroutines start with a tiny stack, far cheaper than the 1-8 MB of POSIX threads, because the linker inserts a check at each function's start that grows or shrinks the stack as needed: if space is insufficient, a new stack page is allocated, arguments are copied over, and on return the process is undone. This makes the stack effectively infinite unless stack splitting occurs. The catch is that new stack pages come from the heap, so an infinite recursive function keeps allocating heap pages until they exceed available physical memory, making the machine unresponsive. Unlike most languages, goroutines have no max stack frame limit, trading flexibility for the risk of exhausting OS resources.
# Goroutine Stack Infinite
[Goroutines](goroutine.md) are cheap to create in terms of initial memory footprint, opposed to the 1 to 8 MB with traditional POSIX threads) and **their stack grows and shrinks as necessary**.

To implement this, the linker inserts a small introduction at the start of each function which checks to see if the amount of stack required for the function is bellow the amount currently available. If not, a call to allocate a new stack page is made and then, the arguments of the caller are copied over and the control is returned to the original function. When the function exits, the process is undone, its return arguments are copied back to the stack frame of the caller and the stack space released.

By this process the stack is effectively infinite, and assuming that you’re not continually straddling the boundary between two stacks, colloquially known as _stack splitting_, is very cheap.

The gotcha is that when new stack pages are needed, they are allocated from the heap.

So if you have a infinite [recursive function](recursion.md), new stack pages are allocated from the heap, permitting the function to continue to call itself. The size of the heap will exceed the amount of free physical memory in your machine, at which point it will make it unusable.

This means that a goroutine doesn't have a max stack frame limit like most languages do but on the other hand, it creates a risk of stopping your machine. Other programming languages do have this limit, but infinite functions don't pose a risk to the OS.
## Understanding
Go goroutines are cheap and easy to create which allows the language to take advantage of this and create virtually infinite stacks that rely on the heap.

While this gives some flexibility when you're writing your programs, it can pose a serious risk to the parent OS as it can make the whole machine unresponsive.
## Trade-offs
- Flexibility: You can write functions without worrying about max depth limits of the language.
- Performance: As your heap keeps growing, your application performance keeps degrading with the OS resources exhaustion.
- Complexity: An unaware developer might create a infinite function and wouldn't be aware of why the machine itself is becoming slow. It gets even worse if the functions leaks and this behavior cannot be felt right away.
## Examples
### Infinite Recursive function
>This code will make your machine unresponsive if you let it run for some time.

```go
import "fmt"

type S struct {
    a, b int
}

// String implements the fmt.Stringer interface
func (s *S) String() string {
    return fmt.Sprintf("%s", s) // Sprintf will call s.String()
}

func main() {
    s := &S{a: 1, b: 2}
    fmt.Println(s)
}
```
## References
- [Goroutines](goroutine.md)
- [Recursive Function](recursion.md)
## Flashcards
- Q: Why are goroutines cheap to create compared to POSIX threads?
- A: Their initial memory footprint is much smaller than the 1 to 8 MB used by POSIX threads, and their stack grows and shrinks as necessary.
- Q: How does Go implement a growable goroutine stack?
- A: The linker inserts a check at the start of each function comparing the stack space required to the amount available. If not enough is available, a new stack page is allocated, the caller's arguments are copied over, and control returns to the original function. When the function exits, the process is undone and the stack space is released.
- Q: What is stack splitting?
- A: Continually straddling the boundary between two stack pages, which makes the stack growth process expensive.
- Q: Where do new goroutine stack pages come from?
- A: They are allocated from the heap.
- Q: What happens when a goroutine runs an infinite recursive function?
- A: New stack pages keep being allocated from the heap until the heap size exceeds the free physical memory on the machine, making it unusable.
- Q: Does a goroutine have a max stack frame limit like most languages?
- A: No, a goroutine has no max stack frame limit, unlike most other languages.
- Q: What risk does the lack of a max stack frame limit create for goroutines?
- A: An infinite or leaking function can exhaust heap and physical memory, risking making the whole machine unresponsive.
