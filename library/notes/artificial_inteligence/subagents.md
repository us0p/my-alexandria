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
Subagents run in an isolated context window with their own system prompt, tools, and permissions, invoked via natural language, mention, or as the session's main agent. They can spawn nested subagents up to a fixed max depth of 5, and forks (full conversation copies run as separate instances) cannot spawn other forks. Each subagent session persists a transcript file, resumable via its sessionId. Subagents are created via /agents or manual frontmatter files in project, user, or plugin directories with defined priority, inheriting available tools and permission modes. They support scoped memory, preloaded skills, and configurable hooks.
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
## Forks
A fork is just a full copy of the current conversation which is operated as a sub agent with specific instructions and **always run as a separate instance**.

You should fork the conversation when you want to share the existing context with the next task. The spawned agent will continue the work with the existing context and once finished will report back only the final output to the main conversation.

You can use `/fork <directive>` to specify what needs to be done by the new agent in a separate context, but you cannot specify what agent is going to be used and there's no way spawn a agent with the context default to fork. **This is a capability of [skills](agent_skills.md#Run%20skills%20in%20sub-agents) only**.

>There's a default subagent used for fork workflows which is used when you fork a conversation.

Since a subagent always runs within its own context, using a subagent in a fork is creating a subagent within another.
## Understanding
Agents are specialized, isolated actors which don't have access to the main conversation context. They receive a task and work until completion, returning back only the final output.

Even though an agent don't have access to the main context, it can produce **memory files** that allows the agent to learn and evolve as it gains more knowledge on the task at hand. Memory files are scoped differently, allowing you to create specific memories for specific scopes.

Since each agent runs in isolation, this means that you can run several agents in parallel.

Each agent can be configured with the different models, accessible tools, permissions and hooks that it has access to.
## When to Use
- You're not interested in the middle output, only with the resulting operation.
- Tasks that require a structured approach or a specific expertise, like code review.
- Want to span several actors to work on the task simultaneously.
## When NOT to Use
- You want the middle output of an operation.
- You want to interact with the agent as it works through a solution.
- The necessary expertise is needed by more than one agent or conversation, it's best suited as a skill.
## Trade-offs
### Positive points
- **Specialization**: Each subagents can be optimized for a specific task rather than a generalist trying to do everything.
- **Scalability**: Can paralelize independent tasks across multiple agents running simultaneously.
- **Modularity**: Swap out agents without rebuilding the entire system.
- **Cost reduction (When done Right)**: Each subagent works on a smaller context window using focused cheaper models instead of big generalist models for the entire problem.
- **Flexibility**: Can dynamically route tasks to different agents based on context.
### Negative points
- **Coordination**: Managing communication between agents, synchronization, message passing, and orchestration adds significant complexity.
- **Error Propagation/Compounding**: Error from one subagent can propagate to dependent agents. Debugging becomes exponentially harder, you need to trace failures across multiple agents to find the root cause.
- **Inconsistencies**: Multiple agents working on the same problem may reach contradictory conclusions.
- **Cost**: Can be significantly more expensive than a single unified agent for the same task.
- **State management**: Race conditions, stale data, and synchronization issues become real concerns.
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
- [Skills](agent_skills.md)
- [Hooks](agent_hooks.md)
### External References
- [Agent Frontmatter Support](https://code.claude.com/docs/en/sub-agents#supported-frontmatter-fields)
- [Built-in Tools](https://code.claude.com/docs/en/tools-reference)
- [Agent tool access control and list of available tools](https://code.claude.com/docs/en/sub-agents#available-tools)
- [List of hook events](https://code.claude.com/docs/en/hooks#hook-events)
## Flashcards
- Q: What isolated components does a subagent have compared to the main conversation?
- A: Its own context window, system prompt, tool access, and permissions.
- Q: What are good use cases for delegating work to a subagent?
- A: Running tests, searching the codebase, reading log files, and repetitive workflows.
- Q: What happens to the main conversation's context after a subagent completes its task?
- A: Only the subagent's output is returned to the main session, so the main context stays clean and small.
- Q: What are the three patterns for invoking subagents?
- A: Natural Language, where the main agent delegates when a task matches an agent's description, Mention, which guarantees a specific subagent runs for one task, and Session-Wide, where the agent is used as the main agent for the whole session.
- Q: How do you invoke a subagent as the main agent for a session?
- A: With claude --agent code-reviewer.
- Q: How can you manually push a running task to the background?
- A: With CTRL+B.
- Q: Why would a subagent spawn its own nested subagents?
- A: When a delegated task splits into parallel subtasks, such as a reviewer dispatching a verifier per finding, so the intermediate output never reaches the main conversation.
- Q: What is the maximum depth of nested subagents?
- A: 5, and this limit is fixed and non-configurable.
- Q: How is subagent depth counted?
- A: As the number of subagent levels below the main conversation, regardless of whether each level runs in the foreground or background.
- Q: Can a fork spawn another fork?
- A: No, a fork cannot spawn another fork, though it can spawn other subagent types, which count toward the depth limit.
- Q: What does a subagent see from the main conversation when it starts?
- A: Nothing by default, it starts with a fresh isolated context window with no conversation history, invoked skills, or previously read files.
- Q: How does a subagent know what task to perform if it has no conversation history?
- A: Claude composes a delegation message that summarizes the task, and the subagent works from that.
- Q: Which type of subagent is the exception that inherits the parent conversation instead of starting fresh?
- A: A fork.
- Q: What does a transcript file capture?
- A: The entire interaction of a subagent session.
- Q: Where are subagent transcript files stored?
- A: At ~/.claude/projects/{project}/{sessionId}/subagents/
- Q: Why can't the built-in Explore and Plan agents be resumed?
- A: Because they are one-shot and return no agent ID.
- Q: How long are subagent transcript files kept by default?
- A: 30 days, based on settings.
- Q: How do you resume a previous subagent interaction?
- A: By providing the main agent with the subagent's sessionId, after which the subagent continues with all its previous interactions in context instead of starting fresh.
- Q: What are the two ways to create a subagent?
- A: Through the /agents interface, or by manually adding a markdown file with YAML frontmatter in a designated agents directory.
- Q: What is the priority order of subagent definition locations, from highest to lowest?
- A: Project level directories such as .claude/agents have priority 3, user level directories such as ~/.claude/agents have priority 4, and a plugin's agents directory has priority 5, the lowest.
- Q: When do manually created or edited subagents take effect?
- A: Only after restarting the session, unless they were created through the /agents interface, which takes effect immediately.
- Q: What tools does a subagent have access to by default?
- A: It inherits the internal tools and MCP tools available in the main conversation.
- Q: How do the disallowedTools and tools frontmatter fields interact?
- A: disallowedTools is applied first, then tools is resolved against the remaining pool, and a tool listed in both is removed.
- Q: What do permission modes control for a subagent?
- A: How the subagent handles permission prompts. Subagents inherit the permission context from the main conversation but can override the mode.
- Q: What are the three scopes for agent memory files?
- A: user scope at ~/.claude/agent-memory/name/ for learnings across all projects, project scope at .claude/agent-memory/name/ for project specific knowledge shareable via version control, and local scope at .claude/agent-memory-local/name/ for project specific knowledge not checked into version control.
- Q: What does the skills frontmatter field control for a subagent?
- A: Which skills are preloaded into the agent's context, not which skills the subagent is allowed to invoke.
- Q: How do you prevent a subagent from invoking skills?
- A: By managing it through Available Tools rather than the skills frontmatter field.
- Q: What are the two ways to configure hooks for a subagent?
- A: In the subagent's frontmatter, where hooks run only while it is active and are cleared when it finishes, or in settings.json, where hooks run in the main session whenever any subagent starts or stops.
- Q: What is a fork?
- A: A full copy of the current conversation operated as a subagent with specific instructions, always run as a separate instance.
- Q: When should you fork a conversation instead of using a regular subagent?
- A: When you want to share the existing context with the next task, since the spawned agent continues the work with that context and reports back only the final output.
- Q: How do you fork a conversation, and what limitation does this command have?
- A: With /fork directive, but you cannot specify which agent is used, since spawning an agent with the fork context by default is only possible through skills.
- Q: Why is running a subagent inside a fork considered a subagent within another?
- A: Because a subagent always runs within its own context, so a subagent invoked inside a fork is nested within the fork's own subagent context.
