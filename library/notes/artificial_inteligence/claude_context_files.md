---
id: 20260628-claude_context_files
type: concept
status: refined
tags:
  - AI
  - claude_code
created: 2026-06-28
---
## TL;DR
`CLAUDE.md` is the per-session briefing Claude reads every time: project context (what), commands/workflows (how), and behavioral rules (what never to do). Keep it lean and current — it's a briefing, not documentation. Use `@path` imports and `.claude/rules/` to keep the main file small; sub-directory `CLAUDE.md` files load only when working in that area.
# Claude Context Files
Files (chiefly `CLAUDE.md`) that supply Claude with persistent project context, conventions, and rules applied to every session. Treating the file as exhaustive documentation is the common failure — it should hold only session-relevant content, or the agent starts with stale/contradictory context and produces wrong implementations.
## What to keep vs. remove
Keep: build/run commands, folder-structure overview, naming/style conventions, off-limits files, and how to run tests. Remove: last week's bug context, sprint notes, abandoned experiments, reversed decisions, and long rationale.
## Structure
- **What**: architecture, stack, code style.
- **How**: build/test/lint commands, review checklists.
- **Behave**: workflow preferences, and what Claude must never do.
## Placement and precedence
Root or `.claude/`; `CLAUDE.local.md` (gitignored) for personal notes; `~/.claude/CLAUDE.md` for user-wide defaults. Sub-directory `CLAUDE.md` files load only when Claude works in that directory.
## Imports and rules
`@path/to/file` pulls content in on demand (relative, absolute, or user paths); recursive, but use sparingly. All files in `.claude/rules/` load at the same priority as `CLAUDE.md` with no imports needed — good for large teams (fewer merge conflicts).
## Maintenance
Keep under ~300 lines (ideally <50 for a config file), every line earning its place. Mark only truly critical rules with `IMPORTANT`/`YOU MUST` (casing matters, but the agent may still cross them). Add rules as you work and from PR-review findings; review periodically for stale or conflicting rules. `/install-github-action` lets you tag `@claude` in PRs to update it. Most AI-generated failures come from stale or overloaded context, so guard against context bleed across long multi-feature sessions. When a section becomes a procedure rather than a fact, turn it into a Skill — see [[agent_skills]]. Connects with [[vibe_coding]] and the off-limits zones in [[ai_guardrails]].
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
