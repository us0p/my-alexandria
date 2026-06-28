---
id: 20260531-vibe_coding
type: concept
status: refined
tags:
  - AI
  - vibe_coding
created: 2026-05-31
---
## TL;DR
Vibe coding is building software by describing intent in natural language and letting an AI agent write the code. Quality is bounded by the context you supply: set up the project (git, CLAUDE.md), plan before each feature, prompt with scoped and structured prompts, review every diff, and build a working core then iterate — never ask for a whole app blind.
# Vibe Coding
Creating an application by giving an AI model natural-language instructions and letting it generate the code. Because agents start every session with no memory, output quality is bounded by the context and structure you give them.
## Setup for success
Start git early — Claude uses your git history and `.gitignore` as project context. Keep a `CLAUDE.md` briefing (generate it with `/init`); see [[claude_context_files]]. Use plan mode before each feature to catch gaps cheaply before any code is written.
## Planning a feature
Describe the feature and ask the agent to outline the files it will touch, function signatures, edge cases, and assumptions. Every plan needs a goal, constraints, and a way to verify success. Avoid open-ended prompts.
## Prompting well
Define expected behavior, the files involved, and constraints; specificity reduces corrections. Useful frameworks:
- **PRD + vertical slices**: a Product Requirement Doc describes how the app looks and behaves; from it, generate a phased plan that builds full-stack "slices" (DB→UI) of increasing complexity, instead of defining all models upfront.
- **TC-EBC** (Task, Context, Elements, Behavior, Constraints): give the model only what it needs, clearly — reads like a recipe card; drop "maybe/just/please". Strong for UI work.
- **Show, don't tell**: demonstrate the desired output with examples rather than describing it ("write professionally"), for more consistent results.
## Working with the agent
Interrupt a wrong direction with `esc`; undo file changes with `/rewind`; give corrective prompts for small fixes. Manage context with `/clear` between features and `/compact` when context grows long or responses drift (`/rewind` undoes actions, `/clear` resets context — not interchangeable).
## Mindset
Give context complete enough that a junior could implement the feature. Build a core and iterate rather than requesting a whole project at once. Require comments and docs so the codebase stays understandable.
## Best practices
Keep the context file lean and current · reset context between features · plan before every feature · write scoped, constrained prompts · give example files · review every diff · run tests after each accepted change · declare off-limits zones · require migration plans · build skills for repetition. Use [[ai_guardrails]] (types, tests, linters, reviews) to grant autonomy without proportional risk.
## Agent coding guidelines
Agents don't share "common knowledge" — make rules explicit, demonstrative, and obvious. Specify naming conventions, tabs vs. spaces, error/logging behavior, and comment style. Write docs that are clear, consistent, and boring (no idioms). Show patterns with correct, incorrect, and "gold standard" examples. Treat errors and PR-review findings as feedback to improve context files. The goal of standards is predictable code.
## Design principles the agent should follow
High cohesion (one purpose per module), loose coupling (depend via abstractions), separation of concerns, and encapsulation.
## AI refactoring capabilities
Variable renaming across scope, function extraction/decomposition, dead-code elimination, documentation generation, and style consistency. Keep refactor scope small to limit blast radius.
## Related notes
Contrasts with [[spec_driven_development]] (the disciplined, spec-first alternative). Connects with [[claude_context_files]], [[ai_guardrails]], [[agent_skills]], [[subagents]].
![[Pasted image 20260531220816.png]]
![[Pasted image 20260601201740.png]]
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
- `.claude/rules` don't require imports to be referenced. Do Claude reference this files when necessary or are they always in context? What's the impact of having many files in here?
## Iterate on
- Sections of the document that can be iterated and have it's quality 
improved but need more knowledge to do so.
## Flashcards
- Q: Some question about the notes.
- A: The answer for the question above.
