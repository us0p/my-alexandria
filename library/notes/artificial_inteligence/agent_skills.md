---
id: 20260628-agent_skills
type: concept
status: refined
tags:
  - AI
  - claude_code
  - skills
created: 2026-06-28
---
## TL;DR
Skills are reusable, filesystem-based capabilities (a `SKILL.md` with YAML frontmatter, plus optional scripts/resources) that turn a general agent into a specialist. They use progressive disclosure: metadata is always loaded, instructions load when triggered, and extra files/scripts load only when referenced — saving context. Build one after running the same prompt pattern three or four times.
# Agent Skills
Modular, file-based packages of instructions, metadata, and optional resources that an agent loads automatically when relevant, to specialize its behavior. Because Claude runs with filesystem access, a skill can live as a directory it reads only as needed instead of consuming context upfront.
## Three content tiers (progressive disclosure)
- **Metadata** (`name` + `description`): always loaded into the system prompt for discovery.
- **Instructions** (the `SKILL.md` body): loaded when the request matches the description.
- **Resources/code**: extra markdown, scripts run via bash (only their output returns), and reference material — loaded only when referenced.
## Structure
Requires a `SKILL.md` with `name` (≤64 chars; lowercase letters, numbers, hyphens; no "anthropic"/"claude") and `description` (≤1024 chars, stating what it does and when to use it). Keep `SKILL.md` under 500 lines.
## Invocation and scope
Run with `/skill-name`, or let the model auto-trigger from the description; `disable-model-invocation: true` makes it manual-only. Stored at enterprise / personal (`~/.claude/skills`) / project (`.claude/skills`) / plugin scope, with precedence in that order; a skill beats a same-named command. Live change detection reloads an edited `SKILL.md` without restart; nested `.claude/skills` are discovered on demand.
## Advanced features
`$ARGUMENTS` receives passed arguments; inline commands prefixed with `!` and fenced `!` code blocks inject dynamic context (pre-processed, not re-scanned); `context: fork` runs the skill in an isolated sub-agent with a chosen `agent` type (see [[subagents]]). Once loaded, a skill stays in context; after compaction, recent invocations are re-attached (first 5000 tokens each, 25K combined budget).
## Where it fits
Use for repetitive, high-stakes tasks where small prompt changes affect quality — generate tests, security review, safe refactor, safe migration. Build after the pattern is proven and commit it to share. Avoid for one-off tasks or guidance better expressed as a plain fact in [[claude_context_files]]. Caveats: Skills don't sync across surfaces (claude.ai / API / Claude Code are separate), runtime and network access vary by surface, and you should only use Skills from trusted sources — a malicious skill can direct the agent to misuse tools. A core [[vibe_coding]] best practice is "build skills for repetition".
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
