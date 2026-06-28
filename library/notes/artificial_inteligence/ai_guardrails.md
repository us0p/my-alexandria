---
id: 20260628-ai_guardrails
type: concept
status: refined
tags:
  - AI
  - vibe_coding
created: 2026-06-28
---
## TL;DR
Guardrails are automated boundaries — not instructions — that limit what an AI agent can do wrong. Type checkers, tests, linters, file-access restrictions, sandboxing, and mandatory human review fire automatically, feeding the agent feedback or blocking bad changes, so you can grant more autonomy without proportionally more risk. The best guardrails are ones you'd want even without AI.
# AI Guardrails
Automated constraints that bound an AI agent's actions so mistakes are caught or blocked without constant supervision. Without them you're doing pure [[vibe_coding]] — hoping for the best; with them, AI assistance becomes a disciplined practice.
## In practice
- **Type systems** (TypeScript, mypy, Rust borrow checker): catch bugs at build time.
- **Test suites**: break on regressions so the agent self-corrects.
- **Linting**: enforces style and catches common mistakes.
- **File-access restrictions**: keep the agent out of secrets, configs, and infra (see off-limits zones in [[claude_context_files]]).
- **Iteration limits**: cap retries before escalating to a human.
- **Sandboxing**: isolate execution from production.
- **Code review**: the ultimate guardrail — a human reviews every change.
- **Scope limits**: restrict the agent to specific tasks, not the whole codebase.
Because guardrails are automated, they don't require watching the agent: they fire on failure during its observe phase or block a bad change before it lands. The three-tier boundaries used in [[spec_driven_development]] are spec-level guardrails.
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
