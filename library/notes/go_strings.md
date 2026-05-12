---
id: 20260508-go_strings
type: concept
status: draft
tags:
  - go
  - programing_language
created: 2026-05-08
---
## TL;DR
Strings are a base type of go that were designed for easy international text representation.
## Strings
Immutable sequences of bytes representing UTF-8 encoded text.
### Raw strings literals
Are strings literal sequences between **back quotes**. Within the quotes, any character will appear just as it is displayed between the back quotes.

```go
a := `Say "hello" to Go!\n`
fmt.Println(a) // Prints: Say "hello" to Go!\n

b := `Raw strings
can also be used
to create multiline strings`
```

>Backslashes have no special behavior in raw strings literals. If you want escaping to work, you must use **Interpreted String Literals** with double quotes `""`.
### UTF-8 and string literals
In Go a string is a **read-only [slice]() of [bytes](binary_system.md)**. A string holds arbitrary bytes. It's not required to hold [Unicode text, UTF-8](ascii_unicode_and_utf8.md) or any other predefined format.

As implied up front, **indexing a string accesses individual bytes, not characters**. 
```go
func main() {
	const placeOfInterest = `⌘` 
	
	fmt.Printf("plain string: ") 
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n") 
	
	fmt.Printf("quoted string: ") 
	// This verb escapes not only non-printable sequences, but also any non-ASCII bytes while interpreting UTF-8.
	fmt.Printf("%+q", placeOfInterest) 
	fmt.Printf("\n") 
	
	fmt.Printf("hex bytes: ") 
	for i := 0; i < len(placeOfInterest); i++ { 
		fmt.Printf("%x ", placeOfInterest[i]) 
	} 
	fmt.Printf("\n")
}
```

The output is:
```plaintext
plain string: ⌘
quoted string: "\u2318"
hex bytes: e2 8c 98 
```

We can see that the "Place of interest" symbol ⌘, is represented by three [bytes](binary_system.md) and that those bytes are the [UTF-8](ascii_unicode_and_utf8.md) encoding of the hexadecimal value `2318`.

>The UTF-8 representation of the string was created when the source code was written. **Source code in Go is defined to be UTF-8 text**. This means that any string literal (raw or not) is always going to be valid UTF-8, but during runtime, a string **value** is just a slice of bytes, and therefore, it's not guaranteed to be composed of valid UTF-8 bytes. It's also possible to create invalid UTF-8 string literals using escape sequences like `\xff`.
### Bytes, characters and runes
Since the definition of [character is ambiguous in computing](ascii_unicode_and_utf8.md#Code%20points%20and%20characters), the correct term to refer to individual characters into a string would be **code point**, but since it's a bit of a mouthful, Go introduces a shorter term for the concept: `rune`.

A `rune` means the same as **code point** but it's also an alias for the type `int32` so programs can be clearer when an integer value represents a code point.

Therefore, a character constant (individual characters in a string) are a `rune` constant. Individual runes are represented by single quotes: `'⌘'`, this is a rune with an integer value `0x2318`.
### Range loops
We've seen what happens with a regular [`for loop`]() when we iterate a string. A [`for range loop`]() decodes one [UTF-8 encoded](ascii_unicode_and_utf8.md) `rune` on each iteration. **Each time around the loop, the index of the loop is the starting position of the current rune, measured in bytes**.
```go
const nihongo = "日本語" 
for index, runeValue := range nihongo {
	fmt.Printf("%#U starts at byte position %d\n", runeValue, index) 
}
```

The output shows how each code point occupies multiple bytes:
```plaintext
U+65E5 '日' starts at byte position 0
U+672C '本' starts at byte position 3
U+8A9E '語' starts at byte position 6
```

>If a `for range` loop isn't sufficient, you can try the [unicode/utf8](https://go.dev/pkg/unicode/utf8/) standard library.
## Understanding
- Strings are **byte sequences**, so Unicode text requires `rune` for correct character handling.
## When to Use
- You should worry about runes whenever your application is expected to receive multi-national text or emoji and other types of non-usual text.
## When NOT to Use
- Usually you don't want to avoid creating applications that don't take go string's nature into account. But if you don't want to bother you can skip it in prototypes, small applications, or applications that are only ever going to use standard ASCII characters.
## Trade-offs
- **Learning curve**: Understanding nuances like `rune` vs `byte`, or UTF-8 handling, takes time.
## Examples
- A chat app that supports emojis and international languages (`rune` usage).
## References
### Connects with
- [Binary System](binary_system.md)
- [ASCII, Unicode and UTF-8](ascii_unicode_and_utf8.md)
- [Go Slices](go_arrays_and_slices.md)
- [Go loops](go_loops.md)
## Flashcards
Q: Why are `rune` types important?  
A: They allow correct handling of Unicode characters in strings.
Q: What is a string in Go?  
A: A read-only slice of bytes.
Q: Why is string indexing potentially misleading?  
A: Because it accesses bytes, not characters.
Q: How are characters properly represented in Go strings?  
A: Using runes, which represent Unicode code points.
Q: What is the difference between raw and interpreted string literals?  
A: Raw strings preserve characters exactly, while interpreted strings process escape sequences.
Q: How does a `for range` loop iterate over strings?  
A: It decodes UTF-8 and returns one rune at a time.
Q: What does the index in a `for range` string loop represent?  
A: The byte position where the rune starts.
Q: Why are Go strings not guaranteed to be valid UTF-8 at runtime?  
A: Because they are just byte slices and can contain arbitrary data.
Q: What is a common failure mode related to strings?  
A: Treating each byte as a character instead of handling runes properly.
Q: What is a real-world use case for runes?  
A: Handling multilingual text or emojis correctly.