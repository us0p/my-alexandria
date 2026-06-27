---
id: 20260626-linked_list
type: concept
status: draft
tags:
  - computer_theory
  - algorithms
  - data_structures
created: 2026-06-26
---
## TL;DR
Simple Data structure that wires elements in each other by referencing their container's (Node) address. It has a `O(n)` runtime for scans and `O(1)` runtime for writes.
# Linked List
It's a data structure used to store a collection of data. This structure is composed of a list of `Node` in which each `Node` points to the next in the list.

A `Node` is nothing more than a container which encapsulates the node's value and the next node's address:

```go
type Node struct {
	Val int
	Next *Node
	Prev *Node // might exists, creates a double-linked list, not common
}
```

>The first element of a linked list is usually referred to be the **head**, the last is the **tail**.
## Writing to the list
To add elements to a list you must wire them up by referencing a new node by a previous node in the list using the `Next` attribute.

```go
func buildLL() *Node {
	head := Node{}
	n1 := Node{}
	n2 := Node{}
	
	head.Next = &n1
	n1.Next = &n2
	
	// if the list is doubly linked:
	n1.Prev = &head
	n2.Prev = &n1
	
	return &head
}
```
## Traversing the list
Since each element references the adjacent ones, you must walk each element at the time to know where the next element is.

This means that unlike [Arrays](array.md), you can't access an element in the middle of the list without walking all the previous elements.

```go
func printList(ll *Node) {
	if ll == nil {
		return
	}
	fmt.Prinln(ll.Val)
	printList(ll.Next)
}
```

>Normally, linked list algorithms are [recursive](recursion.md), but there's always a iterative version, you must consider [when to recurse rather to iterate](recursion.md).
## List operation runtime
| Operation                        | Runtime | Reason                                                                                                    |
| -------------------------------- | ------- | --------------------------------------------------------------------------------------------------------- |
| Search                           | `O(n)`  | Must traverse each element in order to find a given `Node`.                                               |
| Writes to the heads              | `O(1)`  | Since you usually stores a reference to the list's heads, any write operations to them are constant time. |
| Writes to the middle of the list | `O(n)`  | Must traverse to the given node before executing operation, same constraint as a search operation.        |
## Linked List x Array
So you can use a `Node` to encapsulate a value and then create a list of this containers to store a group of values, which is basically what an [Array](array.md) do, so when should you use one rather the other? When are Linked Lists preferred over Arrays?

Let's look at a comparison between an Array and a Linked List runtime operations:

| Operation                   | Array Runtime | Linked List Runtime |
| --------------------------- | ------------- | ------------------- |
| Search                      | `O(n)`        | `O(n)`              |
| Insert/Delete to head       | `O(n)`        | `O(1)`              |
| Update to head              | `O(1)`        | `O(1)`              |
| Write to tail               | `O(1)`        | `O(1)`              |
| Insert/Delete in the middle | `O(n)`        | `O(n)`              |
| Update in the middle<br>    | `O(1)`        | `O(n)`              |
Looking at this comparison table it's clear that any insert or delete operation in the array that's not in the **tail** is going to have a runtime worst than it's counterpart in the Linked List.

It's reasonable to think that a Linked List is a good choice when you need to perform a lot of operations at the start of the structure. But, this is completely wrong because of [L1 and L2 caches](l1l2_cache.md).

Since Arrays are a contiguous block of memory, the CPU can perform operations a lot faster if they fall within the same cache (when data is close together). But in Linked List, data is sparse in memory because of pointer reference, each iteration of a list traversal is a cache miss, which makes the algorithm performance even worse as the data structure size increases.

This means that even when the runtime is the same (`O(n)`) an Array is **probably** going to have a faster execution in reality because of L1 and L2 caches.

Therefore, you must **almost never** prefer a Linked List over an Array even when dealing with "best case" operations scenario. **Even arrays `O(n)` shifts can be faster** due to cache optimizations.

>Performance difference can be `50x` to `100x` faster due to cache optimizations.
## Understanding
- explanation of the concept, using your own words.
- Focus on cause and effect.
Ex:
- This pattern exists because systems are likely to couple business rules and external details...
- The separation allows changing interfaces without having to rewrite central rules...
## When to Use
- You hold long living references to nodes (so that you skip the traversing step and the cache misses).
- You need hierarchy in the way data is stored (like an undo history where each action comes after another).
- You have a well-defined way of accessing your data. [Queues](queue.md), [Stacks](stack.md) and other data structures are implementations of Linked Lists with specific data access rules, optimized for a specific condition.
## When NOT to Use
- You need random access to elements in the middle of the list. Prefer Arrays.
- Memory is tight. Prefer Arrays.
## Trade-offs
- Indirection: each nodes points to the other which makes it necessary a random memory access that causes a lot of cache misses. Traversing a Linked List is very expensive.
- Size: each container requires a pointer reference alongside the value which creates a lot of memory overhead because of pointers.
## Examples
### The Sentinel Pattern
A *sentinel* is a fake node you prepend to your result list that you **never return**. It exists purely to eliminate edge cases.

**The Problem**: Usually, without the sentinel, initializing the head results in a special case:
```go
for ... {
	if head == nil {
		// special first case: initializes list
		head = node
		curr = head
	} else {
		// other cases uses a different step
		curr.Next = node
		curr = curr.Next
	}
}
```

**The Solution**: Add a sentinel node that carries no value and it's not returned to skip edge case:
```go
head = &Node{}
curr = head

for ... {
	// skip branching completely since head is already initialized
	curr.Next = node
	cur = curr.Next
}

// skip the sentinel at the end by returning the next element
return head.Next 
```

This pattern is used whenever you're building a list from scratch and helps keeping the operations focused.
## References
- [Queue](queue.md)
- [Stack](stack.md)
- [Array](array.md)
- [L1/L2 cache](l1l2_cache.md)
- [Recursion](recursion.md)
## Questions
- What is the Fast/Slow pointer pattern and when to use it?
## Flashcards
- Q: What is a Linked List?
- A: A data structure that stores a collection of data as a list of Nodes, where each Node holds a value and a pointer to the next Node.
- Q: What is a Node in a Linked List?
- A: A container that encapsulates a value and the address of the next Node (and optionally the previous Node in a doubly-linked list).
- Q: What are the head and tail of a Linked List?
- A: The head is the first element and the tail is the last element of a Linked List.
- Q: What is the runtime for searching an element in a Linked List and why?
- A: O(n), because you must traverse each element in order to find a given Node.
- Q: What is the runtime for writing to the head or tail of a Linked List and why?
- A: O(1), because you usually store a reference to the list's head/tail, so write operations to them are constant time.
- Q: What is the runtime for writing to the middle of a Linked List and why?
- A: O(n), because you must traverse to the target Node first, same constraint as a search operation.
- Q: What is the key advantage of a Linked List over an Array for insert/delete operations?
- A: Insert/Delete at the head is O(1) in a Linked List vs O(n) in an Array, since Arrays must shift all elements.
- Q: Where does a Linked List perform worse than an Array in terms of runtime?
- A: Update in the middle: O(n) for Linked List vs O(1) for Array, because Linked Lists must traverse to the target Node first.
- Q: Why is an Array usually faster than a Linked List even when both have O(n) runtime?
- A: Arrays are a contiguous block of memory, so the CPU can leverage L1/L2 caches. Linked List nodes are sparse in memory, causing a cache miss on every traversal step.
- Q: By how much can cache optimizations make Arrays faster than Linked Lists?
- A: 50x to 100x faster.
- Q: When should you prefer a Linked List over an Array?
- A: When you hold long-living references to specific nodes (skipping traversal), need hierarchical data access (e.g., undo history), or use a well-defined access pattern like Queues or Stacks.
- Q: When should you NOT use a Linked List?
- A: When you need random access to elements in the middle (prefer Arrays) or when memory is tight (prefer Arrays, since each Node requires an extra pointer).
- Q: What are the two main trade-offs of Linked Lists?
- A: Indirection (pointer chasing causes cache misses, making traversal expensive) and size overhead (each Node stores a pointer alongside its value).
- Q: What is the Sentinel pattern in Linked Lists?
- A: Prepending a fake Node (the sentinel) to the result list to eliminate edge cases when initializing the head, allowing uniform handling of all nodes in the loop. The sentinel is never returned.
- Q: What is a doubly-linked list?
- A: A Linked List where each Node also holds a pointer to the previous Node (Prev), enabling traversal in both directions.
