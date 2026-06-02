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
## Vibe Coding Mindset
1. Don't use vague, vast prompts, you need to feed the model with enough context (functional requirements, design decisions, tests, constraints) so that a Junior developer would be able to implement the feature with only what you gave. You can also work alongside the model to iterate and improve this documentation before giving it to an coding agent.
2. Do not ask for a full project from the start. Let the agent implement a working core and than add feature to it separately and iterate on edge cages, refactoring and security issues.
3. Outsourcing all the implementation to the model can be fast but it also becomes unmanageable. It's hard to understand how a codebase with 5k+ lines works, imagine how hard would it be if you haven't even written a single line of code yourself? Always ask the model to add comments and create documentation of your code so that your future self won't be mad at you.
## Vibe Planning
The planning mode of your model can be used in two ways:
1. Plan -> Build: You get the model to analyze the codebase, understand the problem, formulate a plan and then build the implementation direclty from that plan.
2. Plan -> Spec -> Build: You add a intermediate step and then it writer a detailed specification (a natural language instruction that defines what should be built, a.k.a. meta-prompting). This approach should be used for more complex implementations.

You can tell the agent what you are trying to build, ask it to help refine the idea, establish different phases, and once done, ask it to document everything so that you can refer to that when actually building the product.
## Defining the "what" and the "how" (PRD & Plan)
Product Requirement Doc (PRD) is just a detailed guide for how the app should look and behave with some guidelines of how it should be implemented.

After generating the PRD, we ask the model to generate a setp-by-setp actionable plan that will implement the app in phases using a modified **vertical slice method** suitable for LLM-assisted development in full-stack frameworks.

Vertical slices instructs the model to develop the app in full-stack "slices" (from DB to UI) in increasingly complexity.

Rather than trying to define all your database models from the start, for example, this approach tackles the simplest form of a full-stack feature individually, and then builds upon them in later phases. This means, in an early phase, we might only define the database models needed for Authentication, then its related server-side functions, and the UI for it like Login forms and pages.

![[Pasted image 20260601201740.png]]

if you realize there is a feature set you want to add on later that didn't already exist in the plan, You can ask the LLM to review the plan and find the best time/phase within it to implement it.

After completing a significant feature. You should make an habit of tasking the AI with documenting what was just built. You can even create a Skill for that:
- Gather the key files related to the implementation feature.
- Provide the relevant sections of the PRD and the Plan that described the feature.
- Reference the rule file with the Doc creation task.
- Have it review the Doc for breadth and clarity.

The important is to to focus on the core logic, how the different parts connect and any key decisions made, referencing specific files where the implementation details can be found.

The model would then generate a MD file in a particular directory which is nice because:
- It create a clear decision document that humans can easily understand.
- It builds a knowledge base within the project that could be fed back into the AI's context in later stages, helping maintain consistency and reducing context losses.
## Spec-Driven Development
Is much closer to traditional engineering practices. Instead of jumping straight into implementation, we start by doing the hard thinking ourselves: making architectural decisions, defining requirements, and documenting them in a structured markdown specification stored in the repository and updated alongside the project. This creates an important shift: we decouple the specification (what we are building and why) from the implementation (the actual code).

SDD addresses many of the core issues of vibe coding by preserving context across sessions and different ai agents, while aligning both humans and agents around the project's main non-negotiables.
### SDD Stages
- Constitution: Agreement of key decisions for the project, it usually includes several documents: Mission (explains the why), tech stack (documents technical decisions as well as deployment), roadmap (outline project phases, planned features, this document is continuously updated with the project evolution).
- Development: understand what we want to build and writing detailed specification. Implementing the changes. Validating that the implementation works as expected.
- Replanning: dedicated phase for revisiting the constitution and reviewing previous feature decisions and plans to make sure they still align with the project goals.

>You can use AI to generate all the documents in each specific phase.

E.g.: Constitution documents:
```markdown
We are building Trainlytics, a personal fitness tracking web app built
for people who want more control, flexibility, and insights than standard
fitness apps provide. Find the full requirements in README.md.

Let's create a "constitution" in a specs directory that consists of 
the following parts:
- mission.md - what and why we are building; the main mission of the product
- tech-stack.md - core technical decisions
- roadmap.md - project phases broken down in implementation order

IMPORTANT: You must use your AskUserQuestion tool to get my feedback.
```

E.g.: Task planning phase
```plaintext
Find the next phase in specs/roadmap.md and create a new branch, 
ask me about any steps in the specs that are not fully clear.

Then create a new directory in the format YYYY-MM-DD-feature-name under specs/ 
for this feature, with the following files:
- plan.md - a structured list of numbered task groups
- requirements.md - scope, key decisions, and context
- validation.md - how we define success and confirm the implementation can 
be merged

Use specs/mission.md and specs/tech-stack.md as guidance.
```

E.g.: Development phase
```markdown
Take the next task group from 2026-05-04-phase-1-mvp/plan.md and implement it.
Use requirements.md and validation.md for guidance.
Once done, update the status in both the plan and validation documents.
```

>A good practice is to make all changes through the agent rather than patching documents yourself to maintain consistency across the project. For example, you might require a change and the agent might update more than one related document.

There are evidences that placing the output of an agent in another and asking for critiques improves output quality.

In theory, spec-driven development suggests that the feature phase ends with validation. In practice, it rarely works that cleanly. You will likely find that some parts of the implementation don’t work as expected. At that point, you have two options:
- Add a couple more iterations to your `plan.md` and continue refining the feature (this works well for smaller changes), or
- If the issues are more substantial, treat them as part of the next feature phase and handle them during replanning.

>One important thing to watch out for: it can be tempting to simply explain the issue to the LLM agent and ask for fixes, instead of updating the specs and reworking the implementation. Try to resist that shortcut. Keeping the specification as the source of truth is what makes the approach robust.

>In the current AI era, the main value of a human lies in thinking and architecture.

[SDD GitHub Repo](https://github.com/github/spec-kit)
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
