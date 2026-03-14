---
id: 260314-two_complement
tags:
  - algorithms
  - computer_theory
created: 2026-03-14
status: draft
---
# Two's complement
It's the standard way computers represent signed integers.

It's preferred over other methods since it ensures that addition works identically for positive and negative numbers. Other methods represent two types of zero (+0, -0) which makes addition complicated.
## How two's complement represent negative numbers
The most significant bit (MSB) determines the sign
- 0 -> positive
- 1 -> negative

Since the MSB must be set to determine a number as negative, this means that the binary representation of a negative number is equal to the binary representation of `2^n - |x|` where x is the negative number.

```plaintext
Example: Binary representation of -5 in a 8 bit signed integer
2^8 - |-5| = 256 - 5 = 251
251 = 11111011
```

This happens because once the MSB is set, we "wrap around" to the negative side.

| 8 Bit binary representation | Signed Integer | Unsigned integer |
| --------------------------- | -------------- | ---------------- |
| 00000000                    | 0              | 0                |
| 01111111                    | 127            | 127              |
| 10000000                    | -128           | 128              |
| 11111111                    | -1             | 255              |
| 11111011                    | -5             | 251              |
>Two's complement can be computed by inverting all bits and adding 1.

```plaintext
(Binary representation of 1)
00000001 = 1

(invert all bits, binary representation of 254)
11111110 = 254

(add 1, binary representation of 255)
11111111 = 255
```
## Representation range
Since negative numbers require MSB to be set, this means that we can only represent `2 ^ (n - 1)` negative numbers.

```plaintext
Example: Number os negative numbers in a 8 bit signed integer
2 ^ (8 - 1) = 2 ^ 7 = 128 (negative integers, starting from -128)
```

The maximum number of numbers is represented by `2 ^ n`, so, in a 8 bit number we have:
`2 ^ 8 = 256` possible numbers

```plaintext
256 -> available numbers
128 -> negative
127 -> positive
1 -> zero
```

The range of numbers we can represent with 8 bits therefore is: `-128..127 (256 numbers)`.
## Examples
- 0 -> 00000000
- 127 -> 01111111
- -128 -> 10000000
- -1 -> 11111111 (2 ^ 8 - |-1| = 256 - 1 = 255)

## TL;DR
To get the binary representation of a negative number you take the maximum number of available integers `2^n` minus the module of the number you want to represent. It's like counting backwards.

The amount of negative numbers that can be represented by an N bit integer is equal to `2 ^ (n - 1)` and the numbers goes from `-(2 ^ (n - 1))..(2^n-1)`, this drift in the number of numbers is because of the zero.
## Flashcards
- Q: What are two's complement used for in computing?
- A: It's used to represent negative number using binary notation.
- Q: How a number is marked as negative using two's complement?
- A: The most significant number (MSB) determines the sign.
- Q: What's the MSB bit for a positive number? And for a negative?
- A: 0 for positive, 1 for negative.
- Q: How do we know which negative number is represented by a binary notation using two's complement?
- A: The binary form of a negative number can be discovered by taking the binary representation of the biggest number subtracted by the module of the negative number we want to get.
- Q: How does computer usually calculate the two's complement of a number?
- A: They invert all the bits of the positive representation of that number and then add 1 to it.
- Q: Given a N bit signed integer, how we can discover the number of possible numbers? And how we can know the number of negative and positive integers?
- A: To discover the number of possible number we can use the formula `2 ^ n`. The number of negative numbers is equal to `2 ^ (n - 1)`, starting at `-(2^(n-1))` and the number of non-negative numbers is equal to `2 ^ n - 1` which is also the max positive range.
- Q: When we look at the range of numbers in a 8-bit signed integer we see that it goes from -128 to 127, why?
- A: This happens because of the 0, which is also a number that can be represented.