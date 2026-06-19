# Regexp `\b` — Word Boundary Assertion

## What it is

`\b` is a **zero-width word boundary assertion**. It matches a *position* in the string, not a character — it consumes no input, it just asserts that a boundary exists at that point.

A boundary exists at any position where one side is a **word character** (`[A-Za-z0-9_]`) and the other side is either a **non-word character** (space, punctuation, symbol) or the **start/end of the string**.

## Practical use: matching a complete word

Placing `\b` on both sides of a pattern means "match this as a complete word, not as part of a larger one."

```
\bHELLO\b
```

| Input | Match? | Reason |
|---|---|---|
| `"HELLO world"` | ✓ | Left: start of string, right: space |
| `"say HELLO!"` | ✓ | Left: space, right: `!` |
| `"HELLOWORLD"` | ✗ | No boundary between `O` and `W` — both are word chars |
| `"HELLO123"` | ✗ | No boundary between `O` and `1` — both are word chars |

## Key nuance: what counts as a "word character"

In regex, a word character is strictly `[A-Za-z0-9_]`. Everything else — spaces, punctuation, symbols — acts as a boundary. So `\b` fires at the edge between:

- a letter and a space → boundary ✓
- a letter and `!` → boundary ✓
- a letter and another letter (different case) → **no** boundary ✗
- a letter and a digit → **no** boundary ✗

This last point matters: `\b[A-Z]{5,}\b` does **not** match the `HELLO` inside `HELLOworld`, because the transition from `O` (uppercase) to `w` (lowercase) is letter-to-letter — no boundary.

## Real-world example: screaming text detector

Without `\b`, the pattern `[A-Z]{5,}` would match the uppercase run inside strings like `https://SOMEURL/path` or `ABCDEFsomething`, producing false positives.

With `\b`:

```go
screamingRegexp := regexp.MustCompile(`\b[A-Z]{5,}\b`)
matches := screamingRegexp.FindAllString(story, -1)
isScreaming := len(matches) >= 3
```

This only matches uppercase sequences that stand alone as full tokens — which is the correct definition of ALL-CAPS screaming in prose.

## Further reading

- [Regular-Expressions.info — Word Boundaries](https://www.regular-expressions.info/wordboundaries.html) — thorough reference on `\b`, `\B`, and edge cases
- [Go regexp package docs](https://pkg.go.dev/regexp/syntax) — Go's RE2 syntax, including boundary assertions
- [Regex101](https://regex101.com/) — interactive tester; use the Go flavor to experiment with `\b` patterns in real time
