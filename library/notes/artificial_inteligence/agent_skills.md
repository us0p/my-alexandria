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
Agent skills are modular, filesystem-based packages (`SKILL.md` with optional scripts/resources) that specialize agent behavior through progressive disclosure: frontmatter metadata loads at startup for discovery, instructions load on trigger, and extra files/scripts load only when referenced. A `SKILL.md` requires `name` (≤64 chars) and `description` (≤1024 chars). Skills are stored at enterprise, personal, project, or plugin scope — higher levels override lower. Invoked via `/skill-name` or auto-triggered by description match. Advanced features include string substitutions, dynamic context injection via `!` commands, and forked execution in isolated sub-agents.
# Agent Skills
Modular, file-based packages of instructions, metadata, and optional resources that an agent loads automatically when relevant, to specialize its behavior. 

Because Claude runs with filesystem access, a skill can live as a directory it reads only as needed instead of consuming context upfront.

>If a section of your [CLAUDE.md](claude_context_files.md) becomes a procedure rather than a fact, you should make it a skill.

Claude ships with a collection of commands and skills, check more details of the available functionality at [their doc](https://code.claude.com/docs/en/commands)

Use for repetitive, high-stakes tasks where small prompt changes affect quality. Avoid for one-off tasks or guidance better expressed as a plain fact in [Agent context files](claude_context_files.md).
## File content loading
- **Metadata**: file frontmatter, `name` and `description` are required, loaded at startup with the system prompt, used for skill discovery by the agent.
- **Instructions**: The body of the file, loaded when the skill is triggered.
- **Resources/code**: Skill can bundle additional materials like other markdown files with extra instructions, code the agent can run via bash and extra resources.

Extra resources are used only when referenced:
- **Extra markdown instructions**: Stored as top level files with the skill. Contains extra detailed instructions. Conventions, patterns, style guides, domain knowledge, etc.
- **Code**: Stored under the `/scripts` directory in the skill repository. Contains scripts the agent can run via bash like Python, JavaScript or Bash files, useful when deterministic actions must be taken. Used by the agent when necessary.
- **Resources**: Stored as top level files with the skill. Contains materials like database schemas, API documentations etc. Not always necessary files that would take too much context on every skill invocation.
### Live change detection
Changes under:
- `~/.claude/skills`
- `.claude/skills`
- `.claude/skills` inside `--add-dir`
Takes effect within the current session. 

Creating a top-level skills directory requires restarting Claude.

>Live change detection works for `SKILL.md` files only.
### Content loading sequence
When a skill is triggered:
1. Skill body is read from the filesystem using bash and populating the context window with it.
2. If skill body references other files, these files are read too with bash.
3. If scripts are mentioned, the agent runs them via bash and receives only the output. The code never enters context.

When the agent invokes a skill, it enters the conversation as a single message and stays there for the rest of the session.

When a conversation is summarized (auto-compaction), the Agent re-attaches the most recent invocations of each skill, keeping the first `5K` tokens of each. Re-attached skills share a combined budget of `25K` tokens, filled from the most recent skill calls. Because of this, at some point in long conversations, a skill might stop influencing behavior, to remediate this issue:
- Strengthen the skill's `description` and instructions so the model keeps preferring it.
- Use [hooks](agent_hooks.md) to enforce behavior deterministically.
- If the skill is large or you invoked several other after it, re-invoke it after compaction.

>Since skill content is kept in context you should write guidance as standing instructions rather than one-time steps.
## Skill Structure
Every skill requires a `SKILL.md` file with a frontmatter that must contain at least:
- `name` (≤64 chars; lowercase letters, numbers, hyphens; no `anthropic`/`claude`).
- `description` (≤1024 chars, stating what it does and when to use it). Should include what the skill does and when Claude should use it.

>Keep `SKILL.md` under 500 lines.
## Skill storage locations
Where you store a skill determines who can use it:

| Location   | Path                                                                            | Applies to                     |
| ---------- | ------------------------------------------------------------------------------- | ------------------------------ |
| Enterprise | See [managed settings](https://code.claude.com/docs/en/settings#settings-files) | All users in your organization |
| Personal   | `~/.claude/skills/<skill-name>/SKILL.md`                                        | All your projects              |
| Project    | `.claude/skills/<skill-name>/SKILL.md`                                          | This project only              |
| Plugin     | `<plugin>/skills/<skill-name>/SKILL.md`                                         | Where plugin is enabled        |
When skills share the same name across levels, enterprise overrides personal, and personal overrides project. Plugin skills use a `plugin-name:skill-name` namespace, so they cannot conflict.

If a skill and a command share the same name, the skill takes precedence.
## Advanced features
### String substitutions
Allow for dynamic values in the skill content. You can check available substitutions in the [oficial documentation](https://code.claude.com/docs/en/skills#available-string-substitutions)

```plaintext
---
name: session-logger
description: Log activity for this session
---

Log the following to logs/${CLAUDE_SESSION_ID}.log:

$ARGUMENTS
```

Indexed arguments use shell-style quoting, so wrap multi-word values in quotes to pass them as single argument.
### Support Files
You can reference **extra instructions** using markdown wiki-link formatting to include files in the skill directory.

The Agent access detailed reference material only when needed.
```plaintext
## Additional resources
- For complete API details, see [reference.md](reference.md)
- For usage example, see [examples.md](examples.md)
```
### Inject Dynamic Context
You can add dynamic commands to a skill so that the content is populated before the skill is sent to the Agent:

```plaintext
---
name: pr-summary
description: Summarize changes in a pull request
---

## Pull request context
- PR dif: !`gh pr diff`
- PR comments: !`gh pr view -- comments`
- Changed files: !`gh pr diff --name -only`

## Your task
Summarize this pull request...
```

>The agent only sees the final result. Commands are pre-processed.

Output is inserted as plain text. Output will not be scanned for command invocation.

For multi-line commands, use a fenced code block \`\`\`! instead of the inline form.
### Run skills in sub-agents
A skill with a `context: fork` frontmatter is executed in isolation in a [separate agent](subagents.md) and don't have access to your conversation history.

A forked skill using `agent: Explore` skip `CLAUDE.md and git status` to keep their context small, so this agent sees only the `SKILL.md` content and the agent's own system prompt.

```plaintext
---
name: deep-research
description: Research a topic thoroughly
context: fork
agent: Explore
---

Research $ARGUMENTS thoroughly
```
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
## Invocation and scope
Run with `/skill-name`, or let the model auto-trigger from the description; 
### Restricting skill access
Your permission settings govern approval behavior for all tools that aren't mentioned by `disable-model-invocation` and `allowed-tools` frontmatters.

- You can disable all skills in `/permissions`.
- Allow or deny specific skills using [permission rules](https://code.claude.com/docs/en/permissions).
- Hide individual skills with `disable-model-invocation: true`.

>The `user-invocable` frontmatter only controls menu visibility, not skill tool access.
### Override skill visibility from settings
The `skillOverrides` setting controls skill visibility from your settings instead of the skill's on frontmatter. Use it for skills you don't want to edit, such as ones checked into a shared project repo or provided by MCP server. You can access it in the `/skills` menu.
## Understanding
Skills are a useful way to package repetitive tasks or input sensitive prompts so that they are repeatable and reusable. Can produce deterministic output by using scripts and can be shared which creates standard practices across ai coding agents in a particular team.
## When to Use
- When you keep repeating the same prompt over and over.
- When you have tasks in which slight differences in prompt can produce a very different output.
- When your context file contains a lot of procedural instructions instead of facts. Those are better represented as skills.
## When NOT to Use
- When you don't have a repetitive task or prompt.
- When you wan't to enforce guidelines and consistent behavior across a session. This is better represented in a [Context file](claude_context_files.md) or [Agent](subagents.md).
## Trade-offs
- **Context Window**: Adding too much information to a skill can give more background but can also make it less accurate as the skill content is loaded in your conversation history.
- **Quality x Quantity**: Having too many skills in your agent can confuse the agent as it tries to determine which skill to use. Skills that do the same task but have different name/description can cause the agent to invoke different skill on the same problem which can produce non-deterministic output.
## Examples
```markdown
---
name: sample-skill
description: demonstrate how a skill is created, use it when user requests an explanation of what is a skill
---

# Sample Skill
Give explanations on what is an Agent Skill, how to use it and when to create one.
```
## References
- [Context Files](claude_context_files.md)
- [Agent](subagents.md)
- [Hooks](agent_hooks.md)
### External Links
- [Claude Bundled Commands and Skills](https://code.claude.com/docs/en/commands)
- [Permission Rules](https://code.claude.com/docs/en/permissions)
- [String Substitution](https://code.claude.com/docs/en/skills#available-string-substitutions)
- [Settings File](https://code.claude.com/docs/en/settings#settings-files)
## Flashcards
- Q: What are agent skills?
- A: Modular, file-based packages of instructions, metadata, and optional resources that an agent loads automatically when relevant to specialize its behavior.
- Q: What are the three content types in an agent skill and when is each loaded?
- A: Metadata (frontmatter): loaded at startup for skill discovery. Instructions (body): loaded when the skill is triggered. Resources/code: loaded only when explicitly referenced.
- Q: What three kinds of extra resources can a skill bundle?
- A: Extra markdown instructions (conventions, patterns, style guides), executable scripts stored under a /scripts directory, and reference resources like database schemas or API docs.
- Q: What file does every skill require and what must it contain?
- A: A SKILL.md file with frontmatter that must include at least a name and a description field.
- Q: What are the constraints on a skill name?
- A: Max 64 characters, lowercase letters, numbers, and hyphens only; cannot contain "anthropic" or "claude".
- Q: What are the four skill storage locations and their scopes?
- A: Enterprise (all users in the organization), Personal (~/.claude/skills, all your projects), Project (.claude/skills, this project only), Plugin (where the plugin is enabled).
- Q: How does precedence work when skills share the same name across levels?
- A: Enterprise overrides personal, and personal overrides project.
- Q: How can a skill be invoked?
- A: Explicitly with /skill-name, or automatically when the model matches the task to the skill description.
- Q: What happens to skill content during conversation auto-compaction?
- A: The agent re-attaches the most recent invocations of each skill, keeping the first 5K tokens per skill, with a combined re-attachment budget of 25K tokens across all skills.
- Q: What is the content loading sequence when a skill is triggered?
- A: 1. Skill body is read from the filesystem. 2. Any files referenced in the body are read. 3. Any scripts mentioned are executed and only their output enters the context; the code itself does not.
- Q: What does the "context: fork" frontmatter do in a skill?
- A: The skill runs in isolation inside a separate sub-agent that has no access to the current conversation history.
- Q: What network access do skills have when running via the Claude API?
- A: No network access. Skills cannot make external API calls or access the internet.
- Q: What does the "disable-model-invocation: true" frontmatter do?
- A: Prevents the model from auto-triggering the skill; it can still be invoked explicitly by the user.
- Q: How does live change detection work for skills?
- A: Edits to SKILL.md files under recognized skill directories take effect within the current session without restarting Claude. Creating a new top-level skills directory still requires a restart.
- Q: What is the inline dynamic context syntax in a skill?
- A: Prefix a shell command with !` ` (backtick) to inject its output inline, or use a fenced code block opened with ```! for multi-line commands. The agent sees only the output, not the command.
- Q: What does the $ARGUMENTS substitution do in a skill?
- A: Inserts the arguments passed when the skill is invoked. Multi-word values must be quoted to be treated as a single argument.
- Q: How do you include support files in a skill?
- A: Reference them using markdown wiki-link formatting inside the skill body; they are loaded only when that reference is reached.
- Q: What does the skillOverrides setting control?
- A: Skill visibility configured from settings files rather than from the skill frontmatter, useful for shared or MCP-provided skills you cannot edit directly.
- Q: What is the recommended maximum length for a SKILL.md file?
- A: Under 500 lines.
- Q: When should you prefer a skill over a fact in a context file?
- A: When a section of your context file describes a procedure rather than a fact, or when you find yourself running the same prompt pattern three or four times.
