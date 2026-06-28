---
name: summarize
description: Removes redundancy and organize information, effetivelly 
summarizing the node. Can recommend breaking the note into separate topics 
to keep notes concise and objective.
disable-model-invocation: true
---

# Summarize
For the following note in "library/notes": $ARGUMENTS

For this skill reference, when "note's body" is referenced, you must 
understand as the content between the "TL;DR" and "Understanding sections.

## Instructions
1. Read the note's body content.
2. Identify all the different topics that are mentioned in the note.
3. If all the topics are related, like going into more details about a 
single subject, keep the note unified.
4. If the topics diverge into different topics of a subject, you can 
suggest the user to split then.
5. If identified that notes must be splitted, you must design the new 
note's name and body following the [template file](atomic_note_template.md).
The note's body must follow the [Summarization Step](#Summarization%20Step).
6. Once you have te new note's design ready, you must present them to the 
user and ask for his approval, you must make sure that the user can provide
feedback on your design when he's approving it.
7. If the note is to be kept unified, simply apply the [Summarization Step](#Summarization%20Step).
8. Write all back to the file.

## Summarization Step
For each of the identified topics:
1. Order the topics so that one builds on top of the other.
2. Summarize each topic by keeping only the most important information.

## Guardrails
1. Do not make any changes if the provided note doesn't exist in the 
specified directory. Report back to user with note's name.
2. When generating new notes you must make sure to generate the frontmatter
as per the template file.
3. Do not add new lines on paragraphs that precedes headings.
