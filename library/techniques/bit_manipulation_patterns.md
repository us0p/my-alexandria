---
id: 260316-bit_manipulation_patterns
tags:
  - computer_theory
  - algorithms
created: 2026-03-16
status: draft
---
# Bit Manipulation patterns
```plaintext
1 << n: Sets the n bith
x & (1<<n): Checks if n bit of x is set
x | (1<<n): Sets n bit of x
x ^ (1<<n): Toggle n bit of x
x << n: Multiply by 2^n
x >> n: Divide by 2^n
```
## References
- [Bit Manipulation](bit_manipulation.md): Bit manipulation operators.
## Flashcards
- Q: How do we set the nth bit?
- A: 1 << n
- Q: How to check if the nth bit of x is set?
- A: x & (1<<n)
- Q: How to set the nth bit of x?
- A: x | (1<<n)
- Q: How to toggle the nth bit of x?
- A: x ^ (1<<n)
- Q: How to multiply x by a 2^n using bitwise operators?
- A: x << n
- Q: How to divide x by 2^n using bitwise operators?
- A: x >> n