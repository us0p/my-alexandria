---
id: 20260618-go_garbage_collector
type: concept
status: draft
tags:
  - programing_language
  - go
created: 2026-06-18
---
## TL;DR
Very short resume with only the essential information needed.
# Go Garbage Collector
Automatic process that clears memory in Go programs.

Go prefers to allocate memory on the [stack](memory_locations.md#Stack). This means that Go has a stack per goroutine and when possible it will allocate variables on this stack. For that, the compiler attempts to prove that a variable is not needed outside the function by performing **escape analysis** to see if an object "escapes" the function.

If the compiler can determine a variable's lifecycle, it will be allocated in the stack, otherwise it'll be allocated on the heap.

Generally, if you have a pointer to an object then that object is stored on the heap.
## Collector
Executes garbage collection logic and finds objects that should have their memory freed.
## Mutator
Executes application code and allocates new objects in the heap. It also updates objects on the heap as the program runs, which includes making some objects unreachable when they're no longer needed.
## Go's garbage collector implementation
Go's garbage collector is a **non-generational concurrent**, **tri-color mark and sweep garbage collector**

The generational hypothesis assumes that short lived objects, like temporary variables, are reclaimed most often. These variables are usually defined inside functions and as mentioned before, if a variable lifecycle is associated with a function, Go can allocate space for it in the stack. This means that there will be fewer objects in the heap that needs to be collected. This is the reason why go uses a non-generational garbage collector.

Concurrent means that collector runs at the same time as mutator threads.

Mark and sweep is the type of garbage collector. In this type of garbage collector there are two phases:
- **Mark**: The collector traverses the heap and mark objects that are no longer needed.
- **Sweep**: Removes unused objects.

Garbage collector flow:
- **Stop The World (STW)**: Stops your application to do two quick **setup** tasks:
	- Turns on the **write barrier**, a smart logging system for memory. Once turned on, whenever a goroutine tries to update or move a pointer on the heap, the write barrier intercepts it and takes a quick note of it.
	- It identifies the starting points of memory, like global variables and pointers currently on the goroutine stacks.
- **Start The World**: Right after identifying the root memory your application starts running again and the background workers, starting from the roots identified in the STW, marks the pieces of memory that it touches as "in use". If your app modifies pointers during this time, the write barrier catches it so marking phase stays accurate.
- **Second STW**: Once the background workers think they're done marking, the application is briefly stopped and cleans any leftover items. It turns off the write barrier.
- **Second start**: Application starts again and the garbage collector enters the **sweeping phase**. Because everything that is currently being used was successfully "marked" in step 2, the GC workers simply walk through memory and look for anything that _doesn't_ have a mark. It reclaims that unmarked space and clears the marks on the survivors so they are ready for the next GC cycle. This sweeping happens completely in the background while your app is active.

**The Reality:** Freeze (Setup) → Unfreeze → **Mark Concurrently** → Freeze (Finish Mark) → Unfreeze → **Sweep Concurrently**.

By moving the heavy "marking" and "sweeping" work to when the application is awake, Go keeps its STW pauses down to microseconds.
## How Go determines it's time to to trigger GC
Go determines it is time to trigger a garbage collection cycle using a mix of dynamic pacing and a safety time limit.

The runtime uses two primary triggers:
## GOCG Pacer
Go determines the next GC target based on a percentage of your currently active memory. This behavior is controlled by a environment variable called `GOGC` which defaults to `100`.

$$
Target Heap Memory=Live Heap Memory×(1+\frac{GOGC}{100​})
$$

Example:
- Imagine that a GC cycle just finished, and the memory being used is `10MB`.
- Applying the formula above $10MB × (1 + 1.00) = 20MB$
- Go sets the target at `20MB`. Your program resumes running, allocating memory along the way.
- As the heap grows closer to `20MB`, the Go runtime triggers the GC. It paces the background worker threads so that the marking and sweeping steps complete just as the total memory usage touches that `20MB` target.

You can adjust the `GOGC` variable to change how often this GC is triggered, affecting your system RAM and CPU usage accordingly. Increasing it forces your program to save CPU and consume more RAM as you'll let the memory to grow more. If you decrease, you save RAM but will use more CPU.
## 2-Minute time limit
If your memory doesn't grow for some period time your application your application can sit well bellow the `GOGC` target threshold for hours.

To prevent dead memory from idling inside your application indefinitely. Go has a hardcoded time-based safety net. If a garbage collection has not run for 2 minutes, Go will force one to start anyway.

This ensures that idle applications periodically release unused memory back to the operating system.
## Memory Limit
Starting from Go version `1.19+`, there's an additional tool called `GOMEMLIMIT`. Go will ignore the normal `GOGC` mathematical pacing the moment memory usage gets close to your configured `GOMEMLIMIT`, the runtime stop everything and triggers a GC cycle immediately.
## Marking with tri-color algorithm
When marking begins, all objects are white except for the root objects which are grey. The garbage collector begins marking by scanning stacks, globals and heap pointers to understand what is in use.
When scanning a stack, the worker stops the goroutine and marks all found objects grey by traversing downwards from the roots. It then resumes the goroutine.

The grey objects are then enqueued to be turned black, which indicates that they're still in use. Once all grey objects have been turned black, the collector will stop the world again and clean up all the white nodes that are no longer needed. The program can now continue running until it needs to clean up more memory again.

and tri-color is the algorithm used to implement it.
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
