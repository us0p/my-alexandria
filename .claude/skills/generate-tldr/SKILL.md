---
name: generate-tldr
description: Generates a objective and concise summary of a note's body.
disable-model-invocation: true
---

# Generate TL;DR
For the following notes in "library/notes": $ARGUMENTS

Generate a 100 word max paragraph including the most important therms and 
concepts present in the note's body (the content between the "TL;DR" and 
"Understanding" section).

Insert it under the "TL;DR" section of the note.

## Guardrails
1. Do not use any information that's not in the note's body.
2. If any of the notes do not exists, make sure to make no changes and 
report back to the user with the note name.
3. Do not append a new line at the end of the paragraph.
