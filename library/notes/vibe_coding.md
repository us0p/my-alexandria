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

Plan mode can also help by catching errors on AI assumptions after vague prompts.

For a good plan describe the feature and ask Claude to outline:
1. What files it will create or edit.
2. Function signatures it will introduce.
3. Any edge cases or error handling.
4. List any assumptions.

Every plan should include a goal, a list of constraints, and instructions for verifying success. Avoid open-ended prompts. It's fundamental that you define the scope of the prompt and it becomes even more important in complex projects where Claude's output can cascade across multiple files.
## Vibe Code guidelines
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
- `/clear`: Use it when you finish one feature and want to start another without older context influencing the next task. Do not assume Claude remembers anything from the previous session. Re-feed with goal current file state, and constraints only.
- `/compact`: Use when Claude warns against long context or when responses start to drift. It summarizes conversation, while keeping decisions and context. Use when context is almost full but you need to continue working. After running, skim the summary before you continue and look for anything that is off or missing
.
>`/rewind` Is not a substitute for `/clear`. The first undoes actions, while the former resets context.
## Vibe coding best practices
Use it before and during you vibe coding sessions.

|        | **Best practice**                               | **What to do**                                                                             |
| ------ | ----------------------------------------------- | ------------------------------------------------------------------------------------------ |
| **1**  | Keep your context config file lean and current. | Remove outdated content from **CLAUDE.md**, **GEMINI.md**, **.cursorrules**, or equivalent |
| **2**  | Reset context between features.                 | Have one feature per session. Use your tool’s reset command.                               |
| **3**  | Plan before every feature.                      | Ask for a plan before any code is written.                                                 |
| **4**  | Write scoped, constrained prompts.              | Name the goal, list constraints, specify what not to touch.                                |
| **5**  | Give the AI an example file.                    | Provide code examples instead of abstract descriptions.                                    |
| **6**  | Review every diff before accepting.             | Check deletions, API changes, new deps, and off-limits files.                              |
| **7**  | Run tests after every accepted change.          | Do a typecheck + test + lint every time.                                                   |
| **8**  | Declare off-limits zones.                       | Document where sensitive data resides in your config file and related prompts.             |
| **9**  | Require migration plans.                        | Read summary of schema changes, rollback, before generating code.                          |
| **10** | Build skills for repetitive tasks.              | Save best prompt patterns, reuse, remove variance.                                         |
Most AI-generated code failures come from stale or overloaded context.

A common mistake on config files (`CLAUDE.md`, `GEMINI.md`, etc) is treating the config file like documentation. It shouldn't contain instructions for every possible decision; rather, it should contain instructions relevant to that particular session.

| **Keep in your config file**           | **Remove from your config file**             |
| -------------------------------------- | -------------------------------------------- |
| Build and run commands                 | Bug context from a session last week         |
| Folder structure overview              | Temporary deadline or sprint notes           |
| Naming and style conventions           | Experiments you’re about to remove           |
| Off-limits files and folders           | Decisions that were reversed                 |
| Testing framework and how to run tests | Long explanations of why decisions were made |
Not following these rules means the AI starts with stale or contradictory context, leading to incorrect implementations.

A good rule of thumb is to keep the file under 50 lines if possible.

The best claude.md
What - project context
- High-level architecture (brief, non-obvious structure only)
- tech stack and key libraries
- code style and standards
How - Commands and workflows
- Build / dev/ typecheck commands
- Test commands and rules
- Lint / format commands
- Review checklists / pre-PR steps
Behave - behavioral instructions
- workflow preferences
- what claude should never do in this repo

Context bleed, occurs when you work on multiple features in a single long sessions. The model carries context from earlier conversations into subsequent ones and tends to follow previous patterns and references files that have been worked on before. Output looks correct but contain incorrect assumptions.

> Giving an example file to show the patterns your code base follows is effective, specially in Front-End applications where consistent styling classes (Tailwind CSS) usage and component structure matter throughout the code base.

>Accepting AI's output without reviewing the diff means you’re relying on the AI’s judgment over your own, which is risky.

>After implementation, prompt for an explicit risk review before you accept.

For security sensitive areas in your code base, always add them in the "off-limits" section in your configuration file. This must also be repeated in every relevant prompt, specially when working near those files.
```markdown
## Off-limits 
- /src/auth/**
- /src/payments/**
- /src/middleware/rbac.ts
- Any file ending in .migration.ts
- .env and any file that handles environment variables

## For off-limits areas, always:
- Propose a plan and get explicit approval before touching
- Include a rollback strategy in the plan
- Require test coverage for any change
```

Require a migration plan before any schema change. Avoid letting the model write a schema migration without first producing a written plan that includes the migration itself, a rollback path, and the tests.

Typically, the plan should define:
- The forward change (up SQL)
- The reversal path (down SQL)
- The application code updates required to support the new schema
- A rollback strategy, if something fails in production

Build skills for repetitive, high-stakes tasks. Skills are reusable prompt patterns saved as MD file to ensure consistent output for a recurring task. For tasks where small changes in how you prompt affects the AI quality, this is where Skills are strong.

In Claude, a skill is created by adding a directory with a `SKILL.md` to `.claude/skills`:
```plaintext
# .claude/skills/write-tests/SKILL.md
 
---
name: write-tests
description: Writes tests for a module following project Vitest
  conventions. Use when asked to write unit tests.
---
 
When asked to write tests for a module:
1. Read the module and identify all exported functions
2. For each function, write tests that cover:
   - Happy path with valid input
   - Edge cases (empty input, null values, boundaries)
   - Error cases (invalid input, missing dependencies)
3. Use Vitest. Follow the pattern in tests/example.test.ts
4. Run pnpm test when done and report results
5. Do not modify the module under test
```

Can also be used to define a repeated workflow, like fixing a GitHub issue:
```plaintext
# .claude/skills/review-code/SKILL.md
 
---
name: review-code
description: Review code for quality and issues
disable-model-invocation: true
---

When asked to review code: 
1. Read and understand the code
2. Check for bugs or logical errors
3. Identify performance issues
4. Suggest improvements for readability
5. Check for security concerns
6. Ensure coding standards are followed
7. Summarize findings clearly
```

You can run your skill with `/skill-name` directly, or it can also load the skill automatically when the context makes it relevant.

There are four important skills you should build in your projects:
- Generate unit tests
- Security review pass: Use this before merging code related to APIs, authentication, or user input. Checks for exposed sensitive data, unvalidated user input, missing access controls checks, injects attacks, and unsafe code patterns around API keys and environment variables.
- Safe refactor: Use this when restructuring existing code. Preservers public API, updates callers across the entire codebase, adds a chagelog entry and run tests.
- Generate safe migration: Use this before every schema change. Requires up SQL, down SQL, schema changes, application code updates and a rollback plan before code creation.

Skills best practices:
- Build skills after you've run the same prompt pattern three or four times and know what the best output actually looks like.
- Commit your skills directory to version control, to make your best prompt patterns available to your team members and sub-agents in more complex workflows and for follow-up prompts that build on previous work.

Start-to-finish feature building with Claude Code
![[Pasted image 20260531220816.png]]
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
