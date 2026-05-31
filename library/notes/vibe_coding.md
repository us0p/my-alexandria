---
id: 20260531-vibe_coding
type: concept
status: draft
tags:
  - AI
  - vibe_coding
created: 2026-05-31
---
## TL;DR
Very short resume with only the essential information needed.
# Vibe Coding
It's a term used for when a person create an application by providing natural language instructions and letting the AI model generating the code.

It's a good idea to start with git early because Claude uses your git history and .gitignore as part of the project context.

The `CLAUDE.md` file is a document that Claude reads at the start of every session. It tells Claude what the project is about, how it works, and what rules to follow while generating code. Without it, every session starts from scratch. The file works similarly to a briefing document you would give a new team member on their first day.

>You can use Claude's `/init` command to automatically setup a `CLAUDE.md` file based on your current project conventions, build commands and project structure.

Example `CLAUDE.md` file:
```markdown
# task-manager

## Description
Short explanation of what this project does.

## Tech stack
...

## Run commands
...

## Conventions
Include things like:
- Coding Style
- Folder Structure
- Naming Conventions
- Testing Expectatios
- Architectural Patterns

## Guardrails
Include things like:
- Files that should never be deleted
- Rules about adding new dependencies
- Requirements like input validation for API handlers.
```

Use plan mode before every feature. By doing this you can catch gaps before implementation and iterate on the plan until you think it's good enough. It saves time by reducing rewrites and incorrect implementations while keeping the agent focused on what matters (it's common for an agent to get lost after many rewrites).

You can enter plan mode in Claude with `/claude`.

Vibe Code guidelines:
Good prompts define the expected behavior, the files involved, and the constraints. The more specific you are, the fewer corrections you'll need.

```plaintext
Implement the add command handler in a task.ts.
It should accept a task title as a string argument.
Append a new task to tasks.json with these fields:
id (auto-incremented), title, done: false, createdAt (ISO timestamp).
If tasks.json doesn't exist, create it with an empty array first.
```

How to fix Claude mistakes:
- Interrupt mid-run: If you notice that the model is heading in the wrong direction like touching files you didn't asked for or by running a command you did not intend, you can stop by pressing the `esc` key. It halts the process, shows what it has done so far, and wait for next instructions. It's easier to correct things from that point than wait for it to finish and undo a larger set of changes.
- Undo with `/rewind`: If Claude changed several interconnected files, you can use this command to restore the project to a previous state so you can try a different approach.
- Give corrective prompts: When the just a small part of the output is not exactly what you wanted. It's usually faster to give clear instructions to what to improve and what not to change than rewind.

Managing context in Claude:
- `/clear`: Use it when you finish one feature and want to start another without older context influencing the next task.
- `/compact`: Use when Claude warns against long context or when responses start to drift. It summarizes conversation, while keeping decisions and context.
## Understanding
- explanation of the concept, using your own words.
- Focus on cause and effect.
Ex:
- This pattern exists because systems are likely to couple business rules and external details...
- The separation allows changing interfaces without having to rewrite central rules...
## When to Use
- Situations where this is useful
## When NOT to Use
- Situations where this is overkill or harmful
## Trade-offs
- Limitations
- Costs
- Complexity 
## Examples
## References
### Connects with
Add link to relative notes
### Contrasts with
- Add link to alternatives that tries to solve the same problem
- Always add relation definition like "expands", "contrasts", "depends"
## Questions
- Points that are still not clear.
## Iterate on
- Sections of the document that can be iterated and have it's quality 
improved but need more knowledge to do so.
## Flashcards
- Q: Some question about the notes.
- A: The answer for the question above.
