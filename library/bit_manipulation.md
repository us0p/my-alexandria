---
id: 260316-bit_manipulation
tags:
  - algorithms
  - computer_theory
created: 2026-03-16
status: draft
type: concept
---
## TL;DR
Bit manipulation is a performant and memory efficient way of handling underlying bits.

The operators are:
- **AND(&)**: Keep bits that are equal in both numbers
- **OR(|)**: Sets a bit if they're different in both numbers
- **XOR(^)**: Sets a bit if they're different, if they're equal, unsets it.
- **Left Shift(<<)**: Moves all the bits to the left n times.
- **Right Shift(>>)**: Moves all the bits tot he right n times.
# Bit Manipulation
Are operations we do that have direct effect on top of the underlying bits.
## Operators
### AND (&)
Maintain bits that are 1 in both numbers
```plaintext
22: 00010110
27: 00011011
22 & 27: 00010010 -> 18
```
### OR (|)
Sets a bit if they're different
```plaintext
22: 00010110
27: 00011011
22 | 27: 00011111 -> 31
```
### XOR (^)
Sets a bit if they're different in both numbers, if they're equal, unsets it
```plaintext
22: 00010110
27: 00011011
22 ^ 27: 00001101 -> 13
```
### Left Shift (<<)
Moves bits in the original number to left n times (start counting from 0)
```plaintext
1: 00000001
1 << 7: 10000000 -> 128

13: 00001101
13 << 7: 11010000000 -> 1664
```
### Right Shift (>>)
Moves bits in the original number to the right n times (start counting from 0)
```plaintext
128: 10000000
128 >> 7: 00000001 -> 1

13: 00001101
13 >> 7: 00000000 -> 0
```
## Why we would use bit manipulation?
It's extremely useful for performance, memory efficiency and hardware-level control.

The performance gain comes from the fact that CPUs implement them as single machine instructions. In fact bit manipulation operations can be faster than arithmetic operations.

It also commonly used on compression algorithms.
## Examples
- [Bit Flag]: Used to store many booleans inside a single integer
## References
### Connects with
- [Two's Complement](two_complement): Sign number representation system.
- [Bit manipulation patterns](bit_manipulation_patterns.md): Techniques and patterns about bit manipulation.
## Flashcards
- Q: What is bit manipulation
- A: Operations made on numbers down to individual bits
- Q: What are some applications of bit manipulation?
- A: Bit manipulation can be used to perform fast calculations and it's commonly used to create compact structures, like bit flags or in compression algorithms.
- Q: What is the AND operator used for?
- A: The AND(&) operator, keeps bits that are set on both numbers.
- Q: What is the OR operator used for?
- A: The OR(|) operator, sets a bit if it's different on both numbers.
- Q: What is the XOR operator used for?
- A: The XOR(^) operator, sets a bit if they're different, but if they're not, it unsets it.
- Q: What is the Left Shift operator used for?
- A: The Left Shift(<<) operator moves a set of bits n bits to the left.
- Q: What is the Right Shift operator used for?
- A: The Right Shift(>>) operator moves a set of bits n bits to the left.