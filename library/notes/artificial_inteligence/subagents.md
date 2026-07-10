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
Specialized agents that can be used for task-specific workflows.

Subagents have a **separate context window, system prompt, tool access and permissions**. They are best utilized when you don't want to flood your main conversation.

Good use cases are:
- Running tests
- Searching the code base
- Reading log files
- Repetitive workflows

Once the agent completes its task, its output is returned to the main session and the context remains clean and small.
## Invoking subagents
There are three patterns for subagent invoking:
- **Natural Language**: The main Agent can delegate tasks to specific subagents when task matches agent's description.
- **Mention**: Guarantees the subagent runs for one task.
- **Session-Wide**: The specified agent is used as the **main** agent in the session. Invoked with `claude --agent code-reviewer`.

>You can run multiple agents in the background. You can either ask the Main agent to do so, or manually push a task to background with `CTRL+B`.
### Spawning nested subagents
A subagent can spawn its own subagents. This is used when a delegated task itself splits into parallel subtasks, such as a reviewer subagent that dispatches a verifier per finding, so the intermediate output never reaches your main conversation. Only the top-level subagent's summary returns to you.

The **max depth** of nested subagents  is **5**, this is **fixed and non-configurable**.

Depth is counted as the number of subagents levels below the main conversation, regardless of whether each level runs in the foreground or background.

>A fork cannot spawn another fork. It can spawn other subagent types, and those count toward the depth limit.
## Subagent Context
Each subagent starts with a fresh, isolated context window. It does not see your conversation history, the skills you’ve already invoked, or the files Claude has already read. Claude composes a delegation message that summarizes the task, and the subagent works from there. The exception is a fork, which inherits the parent conversation instead of starting fresh.
## Transcript files
A transcript file captures the entire interaction of a subagent session.

When a subagent completes, Claude receives its agent ID and stores a transcript file at:

`~/.claude/projects/{project}/{sessionId}/subagents/`

>The built-in Explore and Plan agents are one-shot and return no agent ID, so they can’t be resumed.

This transcript files are persisted independently of the main conversation and are cleaned up based on settings which defaults to 30 days.
### Resuming subagents
Based on this `sessionId` you can ask the main agent to resume a previous interaction with a subagent by providing its `sessionId`.

In this case, the subagent is not going to start fresh and will have in its context all the previous interactions it performed.
## Creating Subagents
Can be created via the `/agents`(check provider support), or manually by adding a markdown file with YAML frontmatter in:

| Location                                | Scope                   | Priority   | How to create          |
| --------------------------------------- | ----------------------- | ---------- | ---------------------- |
| `.claude/agents/`, `.github/agents/`    | Current Project         | 3          | Interactive or manual  |
| `~/.claude/agents/`, `~/.github/agents` | All projects            | 4          | Interactive or manual  |
| Plugin's `agents/` directory            | Where plugin is enabled | 5 (lowest) | Installed with plugins |
For a list of the supported frontmatter fields, check the [official documentation](https://code.claude.com/docs/en/sub-agents#supported-frontmatter-fields).

>Subagents are loaded at session start. If you add or edit a subagent, restart your session to load it. Subagents created through the `/agents` interface take effect immediately.
## Available Tools
Subagents inherit the internal tools and MCP tools available in the main conversation by default. For a list of available internal tools, check the [official documentation](https://code.claude.com/docs/en/tools-reference).

There are tools that are not available even when listed, you can check the complete list in the [official documentation](https://code.claude.com/docs/en/sub-agents#available-tools).

>`disallowedTools` is applied first, `tools` is resolved against the remaining pool. A tool listed in both is removed.
## Permission modes
Control how the subagent handles permission prompts. Subagents inherit the permission context from the main conversation and can override the mode.
## Agent memory scope
You can scope agent memory files to specific locations to make the memory available to specific areas:

| Scope     | Location                                      | Use when                                                                                    |
| --------- | --------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `user`    | `~/.claude/agent-memory/<name-of-agent>/`     | the subagent should remember learnings across all projects                                  |
| `project` | `.claude/agent-memory/<name-of-agent>/`       | the subagent’s knowledge is project-specific and shareable via version control              |
| `local`   | `.claude/agent-memory-local/<name-of-agent>/` | the subagent’s knowledge is project-specific but should not be checked into version control |
## Agent Skills
You can provide a list of skills that are pre-loaded into the agent context with the `skills` frontmatter.

This field controls which skills are preloaded, not which skills the subagent can access.

To prevent a subagent from invoking skills, manage it through [Available Tools](#Available%20Tools).
## Agent hooks
There are two ways to configure [hooks](agent_hooks.md):
1. In the subagent's frontmatter: Define hooks that run only while that subagent is active and are cleared up when it finishes.
2. In `settings.json`: Define hooks that run in the main session when subagents start or stop.

All [hook events](https://code.claude.com/docs/en/hooks#hook-events) are supported.
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
### Creating a subagent
```markdown
---
name: code-reviewer
description: Reviews code for quality and best practices
tools: Read, Glob, Grep
model: sonnet
---

You are a code reviewer. When invoked, analyze the code and provide specific, actionable feedback on quality, security, and best practices.
```
## References
### Connects with
Add link to relative notes
### Contrasts with
- Add link to alternatives that tries to solve the same problem
- Always add relation definition like "expands", "contrasts", "depends"
## Questions
- What's the ground distinction between skills, forks and agents?
- When should a skill become an agent?
- How to determine if a prompt is better suited as a skill or as an agent?
- When to fork and when to use a subagent?
## Iterate on
- Sections of the document that can be iterated and have it's quality 
improved but need more knowledge to do so.
## Flashcards
- Q: Some question about the notes.
- A: The answer for the question above.
