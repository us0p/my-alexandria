---
id: 20260628-subagents
type: concept
status: refined
tags:
  - AI
  - claude_code
  - subagents
created: 2026-06-28
---
## TL;DR
Sub-agents are specialized agents with their own context window, system prompt, tools, and permissions. Use them to keep verbose output (tests, logs, searches) out of your main context, to enforce tool restrictions, or to parallelize independent work. They start fresh with no conversation history — except a fork, which inherits the full conversation.
# Sub-agents
Task-specific agents Claude can delegate to, each running in an isolated context window with its own prompt, tools, and permissions. Delegating keeps the main conversation clean and lets specialists handle focused work; the cost is latency and coordination, since a non-fork sub-agent starts with no history and must gather context.
## Built-in sub-agents
- **Explore**: Haiku, read-only, skips CLAUDE.md and git status; thoroughness levels quick / medium / very thorough.
- **Plan**: read-only research during plan mode.
- **General-purpose**: all tools, for multi-step explore-and-act work.
## Definition and scope
Defined in Markdown with YAML frontmatter (`name`, `description`, optional `tools`/`disallowedTools`, `model`, `mcpServers`, `skills`, `memory`, `hooks`); create manually or via `/agents`. The body is the system prompt. Precedence: managed > `--agents` flag > project `.claude/agents/` > user `~/.claude/agents/` > plugin.
## Tools and permissions
Inherit the main conversation's tools by default; `tools` is an allowlist and `disallowedTools` a denylist (applied first). `AskUserQuestion`, plan-mode tools, `WaitForMcpServers`, and `ScheduleWakeup` are never available. `skills` preloads skill content (see [[agent_skills]]); `memory` (user/project/local) gives a persistent directory.
## Invocation and execution
Invoke via natural language (Claude decides), `@-mention` (guaranteed once), or session-wide with `claude --agent name`. Run in the foreground (blocks; prompts pass through) or background (concurrent; auto-denies prompts).
## Patterns
Isolate high-volume operations, run parallel research, or chain sub-agents in sequence. Nesting is allowed up to depth 5. Resume a general-purpose or custom sub-agent to continue with full history (Explore and Plan are one-shot and can't be resumed).
## Forks
A fork inherits the entire conversation (system prompt, tools, model, and history) instead of starting fresh; its own tool calls stay out of your context and only the result returns. Start one with `/fork <directive>`; Claude can pass `isolation: "worktree"` to write the fork's edits to a separate git worktree.
## When to use vs. main conversation
Use a sub-agent for verbose output you won't reference, to enforce tool restrictions, or for self-contained work that returns a summary. Stay in the main conversation for frequent back-and-forth, shared multi-phase context, quick targeted changes, or latency-sensitive work (and use `/btw` for a quick question over existing context). Best practices: design focused sub-agents, write detailed descriptions, limit tool access, and check project sub-agents into version control. Used heavily in [[spec_driven_development]] to parallelize and specialize.
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
