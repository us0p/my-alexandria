---
id: 20260628-spec_driven_development
type: concept
status: refined
tags:
  - AI
  - vibe_coding
  - SDD
created: 2026-06-28
---
## TL;DR
Spec-Driven Development (SDD) decouples what/why (a structured natural-language specification kept in the repo) from how (the code). You do the hard thinking upfront — architecture, requirements — and document it so the spec becomes the source of truth, preserving context across sessions and agents. Ideal for medium-sized, well-understood features; overkill for trivial fixes and unreliable for large or unclear problems.
# Spec-Driven Development (SDD)
A methodology where a structured, behavior-oriented spec written in natural language drives AI-assisted implementation, with the spec kept and updated alongside the project. It is the disciplined counter to ad-hoc [[vibe_coding]], which loses context between sessions and lets agents drift.
## Stages
- **Constitution**: durable key decisions — mission (the why), tech stack, and a continuously updated roadmap.
- **Development**: write a detailed spec, implement it, validate it works.
- **Re-planning**: revisit the constitution and prior decisions to keep them aligned with project goals.
## Implementation levels
- **Spec-first**: spec written first, then used in the workflow.
- **Spec-anchored**: spec kept after the task for evolution and maintenance.
- **Spec-as-source**: the human edits only the spec, never the code.
Levels don't depend on each other, and there is no standard yet for maintaining specs over time.
## Writing good specs
Start high-level and let the AI expand it (goal-oriented: what/why over how); design for "Agent Experience" with clean, parseable formats (OpenAPI, llms.txt, explicit types); keep it a living document; avoid context overload (one concern per prompt); build a hierarchical table-of-contents summary; delegate spec areas to sub-agents or skills; use three-tier boundaries (always-do / ask-first / never-do); have the agent self-verify against the spec; embed a test plan; bring domain knowledge and known pitfalls; stay minimal for simple tasks; use RAG/MCP for large specs; commit the spec to the repo.
## Pitfalls
Vague prompts; overlong contexts without summarization; and ignoring the "lethal trifecta" that makes agents dangerous — speed (faster than you can review), non-determinism (same input, different output), and cost (encourages cutting verification corners).
## Specs vs. memory bank
A spec is scoped to the task that creates or changes a feature; a memory bank (rule files, product/codebase descriptions — see [[claude_context_files]]) is relevant across all sessions.
## Tools
- **Kiro**: lightweight, spec-first; requirements as Gherkin user stories → design → tasks; "steering" memory bank.
- **Spec-kit**: GitHub CLI; "Constitution" memory bank; checklist-heavy, many files, a branch per spec.
- **Tessl**: CLI + MCP; explicitly spec-anchored, exploring spec-as-source; per-file low-abstraction specs.
## Single vs. multi-agent
Single agent: simpler, easier to debug — best for isolated modules and small/medium projects. Multi-agent: higher throughput and per-domain specialists (one codes, one tests, one reviews) — best for large codebases. Start with 2–3 agents and clear boundaries; coordinate with [[subagents]] and [[agent_skills]].
## When it fits
Best for medium-sized, well-understood features and multi-session work where context continuity matters. Poor fit for trivial fixes (document overhead slows you down) and large/unclear problems (you can't plan what isn't clear). Agents can still ignore the spec — small iterative steps keep you in control. Watch for repeating Model-Driven-Development's past inflexibility and non-determinism.
![[Pasted image 20260608083657.png]]
![[Pasted image 20260608090755.png]]
![[Pasted image 20260608091100.png]]
![[Pasted image 20260608101629.png]]
![[Pasted image 20260608103323.png]]
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
