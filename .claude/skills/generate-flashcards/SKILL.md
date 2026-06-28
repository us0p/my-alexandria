---
name: generate-flashcards
description: Generates questions and answers in the "Flashcards" section of
a note following project standard.
disable-model-invocation: true
---

# Generate Flashcards
Follow the instructions bellow for the following notes in the 
'library/notes' directory: $ARGUMENTS

If a note doesn't exist, don't perform any change and make sure to report 
back to the user with the note name.

## Instructions
1. Read note's content (everything between TL;DR and Understanding sections).
2. Generate a Markdown list of questions (- Q:) and answers (- A:) using 
only the most important, not negligible information.
3. Insert them under the note's "Flashcards" section. If there already 
content under this section make sure to override with the question and 
answers you created. You must make sure that there's no information loss

## Formatting instructions
Since the questions and answers is going to be processed and used as 
content for flashcards generation in Anki, you must make sure that there's 
no unsupported character or formatting in either a question or answer.

## Example
Note content example:
```markdown
## TL;DR
...
# Go Variables
A variable is a memory space with an address that stores a value of a given
type.

...
## Understanding
...
```

Example of a relevant "Flashcards" section based on note's content:
```markdown
- Q: What is a variable?
- A: It's a space in memory that stores a value of a given type.
```
