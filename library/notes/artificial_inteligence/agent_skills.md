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
Modular, file-based packages of instructions, metadata, and optional resources that an agent loads automatically when relevant, to specialize its behavior. 

Because Claude runs with filesystem access, a skill can live as a directory it reads only as needed instead of consuming context upfront.

>If a section of your [CLAUDE.md](claude_context_files.md) becomes a procedure rather than a fact, you should make it a skill.
## File content loading
- **Metadata**: file frontmatter, `name` and `description` are required, loaded at startup with the system prompt, used for skill discovery by the agent.
- **Instructions**: The body of the file, loaded when the skill is triggered.
- **Resources/code**: Skill can bundle additional materials like other markdown files with extra instructions, code the agent can run via bash and extra resources.

Extra resources are used only when referenced:
- **Extra markdown instructions**: Stored as top level files with the skill. Contains extra detailed instructions. It abstracts  information that are not always required away to be accessed only when necessary.
- **Code**: Stored under the `/scripts` directory in the skill repository. Contains scripts the agent can run via bash like Python, JavaScript or Bash files, useful when deterministic actions must be taken. Used by the agent when necessary.
- **Resources**: Stored as top level files with the skill. Contains materials like database schemas, API documentations etc. Not always necessary files that would take too much context on every skill invocation.
### Content loading sequence
When a skill is triggered:
1. Skill body is read from the filesystem using bash and populating the context window with it.
2. If skill body references other files, these files are read too with bash.
3. If scripts are mentioned, the agent runs them via bash and receives only the output. The code never enters context.
## Skill Structure
Every skill requires a `SKILL.md` file with a frontmatter that must contain at least:
- `name` (≤64 chars; lowercase letters, numbers, hyphens; no `anthropic`/`claude`).
- `description` (≤1024 chars, stating what it does and when to use it). Should include what the skill does and when Claude should use it.

>Keep `SKILL.md` under 500 lines.
## Invocation and scope
Run with `/skill-name`, or let the model auto-trigger from the description; `disable-model-invocation: true` makes it manual-only. Stored at enterprise / personal (`~/.claude/skills`) / project (`.claude/skills`) / plugin scope, with precedence in that order; a skill beats a same-named command. Live change detection reloads an edited `SKILL.md` without restart; nested `.claude/skills` are discovered on demand.
## Advanced features
`$ARGUMENTS` receives passed arguments; inline commands prefixed with `!` and fenced `!` code blocks inject dynamic context (pre-processed, not re-scanned); `context: fork` runs the skill in an isolated sub-agent with a chosen `agent` type (see [[subagents]]). Once loaded, a skill stays in context; after compaction, recent invocations are re-attached (first 5000 tokens each, 25K combined budget).
## Limitations and constraints
The exact runtime environment available to your skill depends on the product surface where you use it.
- **Web Interface**:
	- Varying network access: Depending on settings, skills may have full, partial or no network access.
- **Claude API**:
	- **No network access**: Skills cannot make external API calls or access the internet.
	- **No runtime package installation**: Only pre-installed packages are available.
	- **Pre-configured dependencies only**
- **Claude CLI**:
	- **Full network access**: Skills have the same network access as any other program on the computer.
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
