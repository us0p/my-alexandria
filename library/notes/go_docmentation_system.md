---
id: "20260507"
type: concept
status: draft
tags:
  - go
  - programing_language
created: 2026-05-07
---
## TL;DR
A tool that uses simple comments to generate a web page for a code's documentation.
# Go Documentation
The [`godoc`](https://pkg.go.dev/golang.org/x/tools/cmd/godoc) tool, is the default tool used to document all standard library in Go.

It parses Go source code with comments to generate human-readable documentation in a HTML site.

>Only exported types, variables, constants, and functions are visible in the documentation.

If your package is **public**, it will be indexed automatically in [`pkg.go.dev`](https://pkg.go.dev/github.com/nirdosh17/go-sdk-template) once you perform `go get`. Which means that your package documentation is instantly available to everyone there!
To install:
```bash
go install golang.org/x/tools/cmd/godoc@latest
```
## How documentation is extracted from source code
In the image bellow you'll see how each comment section is extracted to create the documentation page.
![[go_documentation_generation.png]]

Some other documentation are:
- Top-level comments that begin with the word `"BUG(who)”` are recognized as known bugs, and included in the “Bugs” section of the package documentation. The “who” part should be the user name of someone who could provide more information.
- To signal that an identifier should not be used, add a paragraph to its doc comment that begins with “Deprecated:” followed by some information about the deprecation.
## Manage lengthy documentation using `doc.go`
If your package documentation is very long, you can create a `doc.go` file. It only contains comments that are written following [`godoc`](https://pkg.go.dev/golang.org/x/tools/cmd/godoc) convention.

You can also add that file inside sub-packages if necessary.
```plaintext
my-sdk  
├── client  
│   └── client.go   
│   └── doc.go # documentation for subpackage 'client'  
├── doc.go     # root doc.go | documentation for the whole package 'my-sdk'  
└── go.mod
```
## Documentation good practices and details
- The comment is a complete sentence that begins with the name of the element it describes.
- Comments on package declarations should provide general package documentation. Keep in mind that its first sentence will appear in `godoc’s package list`.
-  Subsequent lines of text are considered part of the same paragraph; you must leave a blank line to separate paragraphs.
- URLs will be converted to HTML links; no special markup is necessary.
## Understanding
`godoc` is a tool that uses lightweight comments to generate a documentation page for easy access and discovery. It keeps documentation close to code so it's easier to developers to keep it up to date and it uses a simple format so that it's easy to create them.
## When to Use
- Use it when you want to publish a public package. The source code becomes the documentation itself.
- When you need to add more context to a component. Comments are displayed in code references, so it's easy to get more info when you're or someone else is coding.
## When NOT to Use
- When you won't publish the package.
- When your application is really simple and don't need further explanation on what or why it's doing something.
## Trade-offs
- **Clarity x Density**: Adding too much comments can make your code more clear but it can also make your code difficult to read and navigate.
- **Reliability x Maintenance**: New documentation also mean more stuff to maintain. You need to keep every documentation up to date.
## Examples
```go
// Package level intro
//
// Package long intro
//
// # Title example
//
//     code := "sample"
```
## References
- [Go Package Definition]()
## Flashcards
Q: What is `godoc`?  
A: A tool that generates human-readable documentation from Go source code comments.
Q: What is the main purpose of `godoc`?  
A: To create documentation pages directly from source code and comments.
Q: How does `godoc` generate documentation?  
A: By parsing Go source files and extracting specially formatted comments.
Q: What kinds of elements appear in generated documentation?  
A: Exported types, variables, constants, and functions.
Q: Why are only exported identifiers documented?  
A: Because exported identifiers define the public interface of a Go package.
Q: What happens when a public Go package is fetched with `go get`?  
A: Its documentation becomes automatically indexed and available on `pkg.go.dev`.
Q: Why is keeping documentation close to the code beneficial?  
A: It makes documentation easier to maintain and keep synchronized with implementation changes.
Q: What is the purpose of a `doc.go` file?  
A: To store package-level or lengthy documentation separately from implementation files.
Q: When should `doc.go` be used?  
A: When package documentation becomes large or requires dedicated explanations.
Q: Can sub-packages have their own `doc.go` files?  
A: Yes, each sub-package can define its own package-level documentation.
Q: What is a documentation convention in Go comments?  
A: Comments should begin with the name of the element they describe.
Q: Why should package comments be carefully written?  
A: Because the first sentence appears in package listings and summaries.
Q: How are paragraphs separated in Go documentation comments?  
A: By leaving a blank line between sections.
Q: How are URLs handled in Go documentation comments?  
A: They are automatically converted into HTML links.
Q: How are known bugs documented in Go?  
A: Using top-level comments starting with `BUG(who)`.
Q: How are deprecated identifiers documented?  
A: By adding a paragraph starting with `Deprecated:` in the comment.
Q: What is the main philosophy behind Go documentation?  
A: Lightweight, code-centric documentation that is easy to generate and maintain.
Q: When should `godoc` be used?  
A: When publishing public packages or providing contextual explanations for components.
Q: When might detailed documentation not be necessary?  
A: In very small or simple applications that are self-explanatory.
Q: What is a trade-off of extensive documentation?  
A: It improves clarity but can reduce readability and increase maintenance effort.
Q: Why can excessive comments become problematic?  
A: They may clutter the codebase and make navigation more difficult.
Q: What is a maintenance challenge of documentation?  
A: Documentation must stay updated as the code evolves.
Q: What is a key benefit of Go’s documentation approach?  
A: The source code itself becomes the primary documentation source.
Q: How does `godoc` improve discoverability?  
A: By generating searchable, structured HTML documentation from comments.
Q: What kind of formatting can Go documentation comments support?  
A: Titles, paragraphs, code examples, and automatic links.
Q: Why is Go documentation considered lightweight?  
A: Because it relies on simple comments instead of separate documentation systems.
Q: What is a practical use case for `godoc`?  
A: Publishing SDKs or libraries with accessible online documentation.