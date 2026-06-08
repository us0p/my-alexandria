---
id: 20260531-vibe_coding
type: concept
status: draft
tags:
  - AI
  - vibe_coding
created: 2026-05-31
---
## TL;DR
Very short resume with only the essential information needed.
# Vibe Coding
It's a term used for when a person create an application by providing natural language instructions and letting the AI model generating the code.

It's a good idea to start with git early because Claude uses your git history and .gitignore as part of the project context.

The `CLAUDE.md` file is a document that Claude reads at the start of every session. It tells Claude what the project is about, how it works, and what rules to follow while generating code. Without it, every session starts from scratch. The file works similarly to a briefing document you would give a new team member on their first day.

>You can use Claude's `/init` command to automatically setup a `CLAUDE.md` file based on your current project conventions, build commands and project structure.

Example `CLAUDE.md` file:
```markdown
# task-manager

## Description
Short explanation of what this project does.

## Tech stack
...

## Run commands
...

## Conventions
Include things like:
- Coding Style
- Folder Structure
- Naming Conventions
- Testing Expectatios
- Architectural Patterns

## Guardrails
Include things like:
- Files that should never be deleted
- Rules about adding new dependencies
- Requirements like input validation for API handlers.
```

Use plan mode before every feature. By doing this you can catch gaps before implementation and iterate on the plan until you think it's good enough. It saves time by reducing rewrites and incorrect implementations while keeping the agent focused on what matters (it's common for an agent to get lost after many rewrites).

You can enter plan mode in Claude with `/claude`.

Plan mode can also help by catching errors on AI assumptions after vague prompts.

For a good plan describe the feature and ask Claude to outline:
1. What files it will create or edit.
2. Function signatures it will introduce.
3. Any edge cases or error handling.
4. List any assumptions.

Every plan should include a goal, a list of constraints, and instructions for verifying success. Avoid open-ended prompts. It's fundamental that you define the scope of the prompt and it becomes even more important in complex projects where Claude's output can cascade across multiple files.
## Vibe Code guidelines
Good prompts define the expected behavior, the files involved, and the constraints. The more specific you are, the fewer corrections you'll need.

```plaintext
Implement the add command handler in a task.ts.
It should accept a task title as a string argument.
Append a new task to tasks.json with these fields:
id (auto-incremented), title, done: false, createdAt (ISO timestamp).
If tasks.json doesn't exist, create it with an empty array first.
```

How to fix Claude mistakes:
- Interrupt mid-run: If you notice that the model is heading in the wrong direction like touching files you didn't asked for or by running a command you did not intend, you can stop by pressing the `esc` key. It halts the process, shows what it has done so far, and wait for next instructions. It's easier to correct things from that point than wait for it to finish and undo a larger set of changes.
- Undo with `/rewind`: If Claude changed several interconnected files, you can use this command to restore the project to a previous state so you can try a different approach.
- Give corrective prompts: When the just a small part of the output is not exactly what you wanted. It's usually faster to give clear instructions to what to improve and what not to change than rewind.

Managing context in Claude:
- `/clear`: Use it when you finish one feature and want to start another without older context influencing the next task. Do not assume Claude remembers anything from the previous session. Re-feed with goal current file state, and constraints only.
- `/compact`: Use when Claude warns against long context or when responses start to drift. It summarizes conversation, while keeping decisions and context. Use when context is almost full but you need to continue working. After running, skim the summary before you continue and look for anything that is off or missing
.
>`/rewind` Is not a substitute for `/clear`. The first undoes actions, while the former resets context.
## Vibe coding best practices
Use it before and during you vibe coding sessions.

|        | **Best practice**                               | **What to do**                                                                             |
| ------ | ----------------------------------------------- | ------------------------------------------------------------------------------------------ |
| **1**  | Keep your context config file lean and current. | Remove outdated content from **CLAUDE.md**, **GEMINI.md**, **.cursorrules**, or equivalent |
| **2**  | Reset context between features.                 | Have one feature per session. Use your tool’s reset command.                               |
| **3**  | Plan before every feature.                      | Ask for a plan before any code is written.                                                 |
| **4**  | Write scoped, constrained prompts.              | Name the goal, list constraints, specify what not to touch.                                |
| **5**  | Give the AI an example file.                    | Provide code examples instead of abstract descriptions.                                    |
| **6**  | Review every diff before accepting.             | Check deletions, API changes, new deps, and off-limits files.                              |
| **7**  | Run tests after every accepted change.          | Do a typecheck + test + lint every time.                                                   |
| **8**  | Declare off-limits zones.                       | Document where sensitive data resides in your config file and related prompts.             |
| **9**  | Require migration plans.                        | Read summary of schema changes, rollback, before generating code.                          |
| **10** | Build skills for repetitive tasks.              | Save best prompt patterns, reuse, remove variance.                                         |
Most AI-generated code failures come from stale or overloaded context.

A common mistake on config files (`CLAUDE.md`, `GEMINI.md`, etc) is treating the config file like documentation. It shouldn't contain instructions for every possible decision; rather, it should contain instructions relevant to that particular session.

| **Keep in your config file**           | **Remove from your config file**             |
| -------------------------------------- | -------------------------------------------- |
| Build and run commands                 | Bug context from a session last week         |
| Folder structure overview              | Temporary deadline or sprint notes           |
| Naming and style conventions           | Experiments you’re about to remove           |
| Off-limits files and folders           | Decisions that were reversed                 |
| Testing framework and how to run tests | Long explanations of why decisions were made |
Not following these rules means the AI starts with stale or contradictory context, leading to incorrect implementations.

A good rule of thumb is to keep the file under 50 lines if possible.

The best claude.md
What - project context
- High-level architecture (brief, non-obvious structure only)
- tech stack and key libraries
- code style and standards
How - Commands and workflows
- Build / dev/ typecheck commands
- Test commands and rules
- Lint / format commands
- Review checklists / pre-PR steps
Behave - behavioral instructions
- workflow preferences
- what claude should never do in this repo

Context bleed, occurs when you work on multiple features in a single long sessions. The model carries context from earlier conversations into subsequent ones and tends to follow previous patterns and references files that have been worked on before. Output looks correct but contain incorrect assumptions.

> Giving an example file to show the patterns your code base follows is effective, specially in Front-End applications where consistent styling classes (Tailwind CSS) usage and component structure matter throughout the code base.

>Accepting AI's output without reviewing the diff means you’re relying on the AI’s judgment over your own, which is risky.

>After implementation, prompt for an explicit risk review before you accept.

For security sensitive areas in your code base, always add them in the "off-limits" section in your configuration file. This must also be repeated in every relevant prompt, specially when working near those files.
```markdown
## Off-limits 
- /src/auth/**
- /src/payments/**
- /src/middleware/rbac.ts
- Any file ending in .migration.ts
- .env and any file that handles environment variables

## For off-limits areas, always:
- Propose a plan and get explicit approval before touching
- Include a rollback strategy in the plan
- Require test coverage for any change
```

Require a migration plan before any schema change. Avoid letting the model write a schema migration without first producing a written plan that includes the migration itself, a rollback path, and the tests.

Typically, the plan should define:
- The forward change (up SQL)
- The reversal path (down SQL)
- The application code updates required to support the new schema
- A rollback strategy, if something fails in production

Build skills for repetitive, high-stakes tasks. Skills are reusable prompt patterns saved as MD file to ensure consistent output for a recurring task. For tasks where small changes in how you prompt affects the AI quality, this is where Skills are strong.

In Claude, a skill is created by adding a directory with a `SKILL.md` to `.claude/skills`:
```plaintext
# .claude/skills/write-tests/SKILL.md
 
---
name: write-tests
description: Writes tests for a module following project Vitest
  conventions. Use when asked to write unit tests.
---
 
When asked to write tests for a module:
1. Read the module and identify all exported functions
2. For each function, write tests that cover:
   - Happy path with valid input
   - Edge cases (empty input, null values, boundaries)
   - Error cases (invalid input, missing dependencies)
3. Use Vitest. Follow the pattern in tests/example.test.ts
4. Run pnpm test when done and report results
5. Do not modify the module under test
```

Can also be used to define a repeated workflow, like fixing a GitHub issue:
```plaintext
# .claude/skills/review-code/SKILL.md
 
---
name: review-code
description: Review code for quality and issues
disable-model-invocation: true
---

When asked to review code: 
1. Read and understand the code
2. Check for bugs or logical errors
3. Identify performance issues
4. Suggest improvements for readability
5. Check for security concerns
6. Ensure coding standards are followed
7. Summarize findings clearly
```

You can run your skill with `/skill-name` directly, or it can also load the skill automatically when the context makes it relevant.

There are four important skills you should build in your projects:
- Generate unit tests
- Security review pass: Use this before merging code related to APIs, authentication, or user input. Checks for exposed sensitive data, unvalidated user input, missing access controls checks, injects attacks, and unsafe code patterns around API keys and environment variables.
- Safe refactor: Use this when restructuring existing code. Preservers public API, updates callers across the entire codebase, adds a chagelog entry and run tests.
- Generate safe migration: Use this before every schema change. Requires up SQL, down SQL, schema changes, application code updates and a rollback plan before code creation.

Skills best practices:
- Build skills after you've run the same prompt pattern three or four times and know what the best output actually looks like.
- Commit your skills directory to version control, to make your best prompt patterns available to your team members and sub-agents in more complex workflows and for follow-up prompts that build on previous work.

Start-to-finish feature building with Claude Code
![[Pasted image 20260531220816.png]]
## Vibe Coding Mindset
1. Don't use vague, vast prompts, you need to feed the model with enough context (functional requirements, design decisions, tests, constraints) so that a Junior developer would be able to implement the feature with only what you gave. You can also work alongside the model to iterate and improve this documentation before giving it to an coding agent.
2. Do not ask for a full project from the start. Let the agent implement a working core and than add feature to it separately and iterate on edge cages, refactoring and security issues.
3. Outsourcing all the implementation to the model can be fast but it also becomes unmanageable. It's hard to understand how a codebase with 5k+ lines works, imagine how hard would it be if you haven't even written a single line of code yourself? Always ask the model to add comments and create documentation of your code so that your future self won't be mad at you.
## Vibe Planning
The planning mode of your model can be used in two ways:
1. Plan -> Build: You get the model to analyze the codebase, understand the problem, formulate a plan and then build the implementation direclty from that plan.
2. Plan -> Spec -> Build: You add a intermediate step and then it writer a detailed specification (a natural language instruction that defines what should be built, a.k.a. meta-prompting). This approach should be used for more complex implementations.

You can tell the agent what you are trying to build, ask it to help refine the idea, establish different phases, and once done, ask it to document everything so that you can refer to that when actually building the product.
## Defining the "what" and the "how" (PRD & Plan)
**Product Requirement Doc (PRD)** is just a detailed guide for how the app should look and behave with some guidelines of how it should be implemented.

After generating the PRD, we ask the model to generate a setp-by-setp actionable plan that will implement the app in phases using a modified **vertical slice method** suitable for LLM-assisted development in full-stack frameworks.

Vertical slices instructs the model to develop the app in full-stack "slices" (from DB to UI) in increasingly complexity.

Rather than trying to define all your database models from the start, for example, this approach tackles the simplest form of a full-stack feature individually, and then builds upon them in later phases. This means, in an early phase, we might only define the database models needed for Authentication, then its related server-side functions, and the UI for it like Login forms and pages.

![[Pasted image 20260601201740.png]]

if you realize there is a feature set you want to add on later that didn't already exist in the plan, You can ask the LLM to review the plan and find the best time/phase within it to implement it.

After completing a significant feature. You should make an habit of tasking the AI with documenting what was just built. You can even create a Skill for that:
- Gather the key files related to the implementation feature.
- Provide the relevant sections of the PRD and the Plan that described the feature.
- Reference the rule file with the Doc creation task.
- Have it review the Doc for breadth and clarity.

The important is to to focus on the core logic, how the different parts connect and any key decisions made, referencing specific files where the implementation details can be found.

The model would then generate a MD file in a particular directory which is nice because:
- It create a clear decision document that humans can easily understand.
- It builds a knowledge base within the project that could be fed back into the AI's context in later stages, helping maintain consistency and reducing context losses.

**The six core areas:** GitHub’s analysis of [over 2,500 agent configuration files](https://github.blog/ai-and-ml/github-copilot/how-to-write-a-great-agents-md-lessons-from-over-2500-repositories/) revealed a clear pattern: the most effective specs cover six areas. Use this as a checklist for completeness:
**1. Commands:** Put executable commands early - not just tool names, but full commands with flags: `npm test`, `pytest -v`, `npm run build`. The agent will reference these constantly.
**2. Testing:** How to run tests, what framework you use, where test files live, and what coverage expectations exist.
**3. Project structure:** Where source code lives, where tests go, where docs belong. Be explicit: “`src/` for application code, `tests/` for unit tests, `docs/` for documentation.”
**4. Code style:** One real code snippet showing your style beats three paragraphs describing it. Include naming conventions, formatting rules, and examples of good output.
**5. Git workflow:** Branch naming, commit message format, PR requirements. The agent can follow these if you spell them out.
**6. Boundaries:** What the agent should never touch - secrets, vendor directories, production configs, specific folders. “Never commit secrets” was the single most common helpful constraint in the GitHub study.

**Be specific about your stack:** Say “React 18 with TypeScript, Vite, and Tailwind CSS” not “React project.” Include versions and key dependencies. Vague specs produce vague code.

**Use a consistent format:** Clarity is king. Many devs use Markdown headings or even XML-like tags in the spec to delineate sections, because AI models handle well-structured text better than free-form prose.

remember, “minimal does not necessarily mean short” - don’t shy away from detail in the spec if it matters, but keep it focused.
## Spec-Driven Development
Is much closer to traditional engineering practices. Instead of jumping straight into implementation, we start by doing the hard thinking ourselves: making architectural decisions, defining requirements, and documenting them in a structured markdown specification stored in the repository and updated alongside the project. This creates an important shift: we decouple the specification (what we are building and why) from the implementation (the actual code).

SDD addresses many of the core issues of vibe coding by preserving context across sessions and different ai agents, while aligning both humans and agents around the project's main non-negotiable.
### SDD Stages
- Constitution: Agreement of key decisions for the project, it usually includes several documents: Mission (explains the why), tech stack (documents technical decisions as well as deployment), road map (outline project phases, planned features, this document is continuously updated with the project evolution).
- Development: understand what we want to build and writing detailed specification. Implementing the changes. Validating that the implementation works as expected.
- Re-planning: dedicated phase for revisiting the constitution and reviewing previous feature decisions and plans to make sure they still align with the project goals.

>You can use AI to generate all the documents in each specific phase.

E.g.: Constitution documents:
```markdown
We are building Trainlytics, a personal fitness tracking web app built
for people who want more control, flexibility, and insights than standard
fitness apps provide. Find the full requirements in README.md.

Let's create a "constitution" in a specs directory that consists of 
the following parts:
- mission.md - what and why we are building; the main mission of the product
- tech-stack.md - core technical decisions
- roadmap.md - project phases broken down in implementation order

IMPORTANT: You must use your AskUserQuestion tool to get my feedback.
```

E.g.: Task planning phase
```plaintext
Find the next phase in specs/roadmap.md and create a new branch, 
ask me about any steps in the specs that are not fully clear.

Then create a new directory in the format YYYY-MM-DD-feature-name under specs/ 
for this feature, with the following files:
- plan.md - a structured list of numbered task groups
- requirements.md - scope, key decisions, and context
- validation.md - how we define success and confirm the implementation can 
be merged

Use specs/mission.md and specs/tech-stack.md as guidance.
```

E.g.: Development phase
```markdown
Take the next task group from 2026-05-04-phase-1-mvp/plan.md and implement it.
Use requirements.md and validation.md for guidance.
Once done, update the status in both the plan and validation documents.
```

>A good practice is to make all changes through the agent rather than patching documents yourself to maintain consistency across the project. For example, you might require a change and the agent might update more than one related document.

There are evidences that placing the output of an agent in another and asking for critiques improves output quality.

In theory, spec-driven development suggests that the feature phase ends with validation. In practice, it rarely works that cleanly. You will likely find that some parts of the implementation don’t work as expected. At that point, you have two options:
- Add a couple more iterations to your `plan.md` and continue refining the feature (this works well for smaller changes), or
- If the issues are more substantial, treat them as part of the next feature phase and handle them during re-planning.

>One important thing to watch out for: it can be tempting to simply explain the issue to the LLM agent and ask for fixes, instead of updating the specs and reworking the implementation. Try to resist that shortcut. Keeping the specification as the source of truth is what makes the approach robust.

>In the current AI era, the main value of a human lies in thinking and architecture.

**Demand Multiple Options**: Counter AI's tendency toward sophisticated solutions by explicitly requesting alternatives. Try "Give me three approaches to this problem; the simplest possible solution, a moderate approach, and a full-featured version. Explain the trade offs of each.". You can apply this for errors as well, ask for multiple causes of why a particular error is happening.

The progression should be natural: establish expertise → get the plan → evaluate options → make informed decisions → implement with confidence. This approach transforms AI from an unpredictable code generator into a reliable development partner who understands both your technical needs and your constraints.

AI models have knowledge cutoffs and may not be familiar with the latest versions of frameworks or your specific project requirements. Providing context prevents frustration and reduces iterations. You can even ask the AI what its knowledge cutoffs are, and supplement its data accordingly.

A good safe net for you vibe coded application is to use pre-commit. Before any changes reaches a shared branch, you can run tests, linting, formatting and security checks to validate that the generated code meet your application standards.

Traceability and provenance are crucial when models contribute code. Simple record keeping reduces uncertainty and eases audits.
- Store prompt versions and AI outputs in the PR description or a linked artifact store.
- Tag commits that include AI-generated text with a consistent marker, e.g., `AI:generated`.
- Include the model version, prompt, and timestamp with any generated snippet.
- Use a lightweight governance document that outlines acceptable uses and approval workflows.

Traceability supports accountability and resolves disputes quickly when regressions occur. It also provides data for continuous improvement of prompts and validation steps.

Teams that adopt provenance practices have an easier time demonstrating compliance and understanding root causes when incidents arise.

>SDD is ideal for medium-sized features.

>Always provide constraints about what the agent mustn't do, this helps the agent to be more focused and objective.

A spec is a structured, behavior-oriented artifact - or set of related artifacts - written in natural language that expresses software functionality.

Specs aren't the same as the general context documents in a codebase. That general context are things like rule files, or high level descriptions of the product and the codebase. Some times it's referenced as a **memory bank**. These files are relevant across all AI coding sessions in the codebase, whereas specs are only relevant to the task that actually create or change that particular functionality.
![[Pasted image 20260608083657.png]]
### SDD implementation levels
- **Spec-first**: Spec is written first, and then used in the AI-assisted development workflow for the task at hand.
- **Spec-anchored**: The spec is kept even after the task is complete, to continue using it for evolution and maintenance of the respective feature.
- **Spec-as-source**: Spec is the main source file over time, and only the spec is edited by the human, the human never touches the code.

The implementation levels are not required by each other. In fact, there seems to be no standard strategy about spec maintenance over time.
### SDD Tools
These tools are more focused on creating an application from scratch. If you need to fix a simple bug or add a new simple feature, these tools can feel "overkill" and considering the amount of context and files you'll need to review and validate it can make the progress even slower than a simple "plan" section with a coding agent.

It's important to highlight that even with all of these requirements, the agent can still ignore them. The best way to stay in control of what we're building is to keep small, iterative steps instead of trying to build all at once.
Adding a lot of up-front spec design might not be a good idea, especially when it's overly verbose.

Because of all this upfront planning and design, SDD isn't reliable for problems that are large or that aren't clear. The amount of documents that would be necessary for a large problem isn't negligible and if a problem isn't clear enough, you can't do much planning.

Model-Driven-Development (MDD) was a past initiative to cast specifications into code using a custom language. SDD seems to be going towards that same direction while leveraging AI to take the code implementation heavy work. It's worth to notice that even tough MDD added some flexibility it was, most of the time, inflexible and non-deterministic. This can also be a problem in SDD specially with AI assisted coding standards. We must be careful to keep this practice relevant without falling in the same pitfalls we experienced in the past.
#### Kiro
Lightweight, spec-first SDD tool. Used for tasks or a user story (there's no mention of it being used with spec-anchored strategies).
##### Workflow
###### Requirement
Structured as a list of requirements, where each requirement represents a user story using [Gherkin Syntax](). 

![[Pasted image 20260608090755.png]]- 
###### Design
Consists of sections  describing technical considerations of the task.

![[Pasted image 20260608091100.png]]
###### Tasks
A list of the tasks that map to each requirement specification.

Kiro's memory bank is called "steering" and it's composed of the following documents:
- `product.md`
- `teach.md`
- `structure.md`

Each workflow step is represented by one markdown document.

This framework produces a lot less files for review but it can also be very verbose even for small tasks.
#### Spec-kit
GitHub's version of SDD. It's a CLI that can create workspace setups for a wide range of common coding assistants. Once that structure is set up, you interact with spec-kit via slash commands in your coding assistant. It's highly customizable.

Spec-kit's memory bank is a prerequisite for the spec-driven approach. It's called **Constitution**. It's supposed to contain high level principles that are "immutable" and should always be applied, to every change.

In each of the workflow steps (specify, plan, tasks), spec-kit instantiates a set of files and prompts with the help of a bash script and some templates. The workflow makes heavy use of checklists inside of the files, to track necessary user clarifications, constitution violations, research tasks, etc.

Bellow is an overview that illustrates the topology in spec-kit. Note how one spec is made up of many files.

![[Pasted image 20260608101629.png]]

Spec-kit seems to be aspiring to a spec-anchored approach. However, spec-kit creates a branch for every spec that gets created, which seems to indicate that they see a spec as a living artifact for the lifetime for a change request, not the lifetime of a feature.

Note that this framework can generate a LOT of files to review and they can be repetitive and redundant.
#### Tessl Framework (beta)
Distributed as a CLI. The CLI command also doubles as an MCP server. It's the only one of these three tools that explicitly aspires to a spec-anchored approach, and is even exploring the spec-as-source level os SDD.
In this framework the spec maintainer can tag parts of the specification to make sure that more crucial parts of the generated component are fully under the control of the maintainer.
Putting the specs for spec-as-source at a quite low abstraction level, per code file, probably reduces amount of steps and interpretations the LLM has to do, and therefore the chance of errors.
![[Pasted image 20260608103323.png]]
### How to write good specifications for SDD
a good spec doesn’t just tell the AI what to build, it also helps it self-correct and stay within safe boundaries. By baking in verification steps, constraints, and your hard-earned knowledge, you drastically increase the odds that the agent’s output is correct on the first try (or at least much closer to correct).

1. Kick off your project with a concise high-level spec, then have the AI expand it into a detailed plan. Instead of over-engineering upfront, begin with a clear goal statement and a few core requirements. Treat this as a “product brief” and let the agent generate a more elaborate spec from it. This leverages the AI’s strength in elaboration while you maintain control of the direction. This works well unless you already feel you have very specific technical requirements that must be met from the start. **Why this works:** LLM-based agents excel at fleshing out details when given a solid high-level directive, but they need a clear mission to avoid drifting off course. By providing a short outline or objective description and asking the AI to produce a full specification (e.g. a `spec.md`), you create a persistent reference for the agent. **Keep it goal-oriented:** A high-level spec for an AI agent should focus on what and why, more than the nitty-gritty how (at least initially).
2. Design for Agent Experience (AX): Just as we design APIs for developer experience (DX), consider designing specs for “Agent Experience.” This means clean, parseable formats: OpenAPI schemas for any APIs the agent will consume, llms.txt files that summarize documentation for LLM consumption, and explicit type definitions. The Agentic AI Foundation (AAIF) is standardizing protocols like MCP (Model Context Protocol) for tool integration - specs that follow these patterns are easier for agents to consume and act on reliably.
3. **Make the spec a “living document”:** Don’t write it and forget it. Update the spec as you and the agent make decisions or discover new info. If the AI had to change the data model or you decided to cut a feature, reflect that in the spec so it remains the ground truth.
4. Avoid context overload: Don’t mix authentication tasks with database schema changes in one go, as the [DigitalOcean AI guide](https://docs.digitalocean.com/products/gradient-ai-platform/concepts/context-management/) warns. Keep each prompt tightly scoped to the current goal.
5. have the agent build an extended Table of Contents with summaries for the spec. This is essentially a “spec summary” that condenses each section into a few key points or keywords, and references where details can be found. For example, if your full spec has a section on “Security Requirements” spanning 500 words, you might have the agent summarize it to: “Security: use HTTPS, protect API keys, implement input validation (see full spec §4.2)”. By creating a hierarchical summary in the planning phase, you get a bird’s-eye view that can stay in the prompt, while the fine details remain offloaded unless needed. This extended TOC acts as an index.
6. Utilize sub-agents or “skills” for different spec parts. Each subagent is configured for a specific area of expertise and given the portion of the spec relevant to that area. The main agent (or an orchestrator) can route tasks to the appropriate subagent automatically. The benefit is each agent has a smaller context window to deal with and a more focused role, which can [boost accuracy and allow parallel work](https://10xdevelopers.dev/structured/claude-code-with-subagents/) on independent tasks. Each subagent has a specific purpose and expertise area, uses its own context window separate from the main conversation, and has a custom system prompt guiding its behavior,” as their docs describe. When a task comes up that matches a subagent’s domain, Claude can delegate that task to it, with the subagent returning results independently.
7.  **Use three-tier boundaries:** The [GitHub analysis of 2,500+ agent files](https://github.blog/ai-and-ml/github-copilot/how-to-write-a-great-agents-md-lessons-from-over-2500-repositories/) found that the most effective specs use a three-tier boundary system rather than a simple list of don’ts. This gives the agent clearer guidance on when to proceed, when to pause, and when to stop. **Always do (proceed without asking)**: Run test, follow style guide, log errors. **Ask first (pause for human approval)**: Schema changes, new dependencies, CI config. **Never do (hard stop -no exceptions)**: Commit secrets, edit vendor, remove failing tests.
8. **Encourage self-verification:** One powerful pattern is to have the agent verify its work against the spec automatically. e.g. “After implementing, compare the result with the spec and confirm all requirements are met. List any spec items that are not addressed.” This pushes the LLM to reflect on its output relative to the spec, catching omissions. It’s a form of self-audit built into the process.
9. **Leverage testing in the spec:** If possible, incorporate a test plan or even actual tests in your spec and prompt flow. The agent can be prompted to run through those cases in its head or actually execute them if it has that capability. In an AI coding context, writing a bit of pseudocode for tests or expected outcomes in the spec can guide the agent’s implementation. Additionally, you can use a dedicated “[test agent](https://10xdevelopers.dev/structured/claude-code-with-subagents/)” in a subagent setup that takes the spec’s criteria and continuously verifies the “code agent’s” output.
10. Bring your domain knowledge: Your spec should reflect insights that only an experienced developer or someone with context would know. For example, if you’re building an e-commerce agent and you know that “products” and “categories” have a many-to-many relationship, state that clearly (don’t assume the AI will infer it - it might not). Essentially, pour your mentorship into the spec. The spec can contain advice like “If using library X, watch out for memory leak issue in version Y (apply workaround Z).” This level of detail is what turns an average AI output into a truly robust solution, because you’ve steered the AI away from common traps.
11. **Minimalism for simple tasks:** While we advocate thorough specs, part of expertise is knowing when to keep it simple. Don’t under-spec a hard problem (the agent will flail or go off-track), but don’t over-spec a trivial one (the agent might get tangled or use up context on unnecessary instructions).
12. **Utilize context-management and memory tools**. For instance, [retrieval-augmented generation (RAG)](https://addyosmani.com/agentic-engineering/rag/) is a pattern where the agent can pull in relevant chunks of data from a knowledge base (like a vector database) on the fly. If your spec is huge, you could embed sections of it and let the agent retrieve the most relevant parts when needed, instead of always providing the whole thing. There are also frameworks implementing the Model Context Protocol (MCP), which automates feeding the right context to the model based on the current task.
13. Commit the spec file itself to the repo. This not only preserves history, but the agent can even use git diff or blame to understand changes (LLMs are quite capable of reading diffs).
14. use model selection and batching smartly. If using multiple agents, maybe not all need to be top-tier; a test-running agent or a linter agent could be a smaller model. Also consider throttling context size: don’t feed 20k tokens if 5k will do.

### SDD pitfalls
- **Vague prompts:** Be specific about inputs, outputs, and constraints. “You are a helpful coding assistant” doesn’t work. “You are a test engineer who writes tests for React components, follows these examples, and never modifies source code” does.
- **Overlong contexts without summarization:** Use hierarchical summaries or RAG to surface only what’s relevant. Context length is not a substitute for context quality.
- Ignoring the “lethal trifecta”: There are three properties that make AI agents dangerous: speed (they work faster than you can review), non-determinism (same input, different outputs), and cost (encouraging corner-cutting on verification). Your spec and review process must account for all three. Don’t let speed outpace your ability to verify.

**Single vs. multi-agent: when to use each**

| Aspect         | Single Agent                                                                | Parallel/Multi-Agent                                                               |
| -------------- | --------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| **Strengths**  | Simpler setup; lower overhead; easier to debug and follow                   | Higher throughput; handles complex interdependencies; specialists per domain       |
| **Challenges** | Context overload on big projects; slower iteration; single point of failure | Coordination overhead; potential conflicts; needs shared memory (e.g., vector DBs) |
| **Best For**   | Isolated modules; small-to-medium projects; early prototyping               | Large codebases; one codes + one tests + one reviews; independent features         |
| **Tips**       | Use spec summaries; refresh context per task; start fresh sessions often    | Limit to 2-3 agents initially; use MCP for tool sharing; define clear boundaries   |
## Task, Context, Elements, Behavior, Constraints (TC-EBC)
It's a prompt structure designed to given the model only the information it needs in the most clear and explicit manner.
- **Task** defines what you’re building.
- **Context** frames why and for whom. It prevents drift.
- **Constraints** set the guardrails, keeping the system controlled and consistent.

Specially useful in UI design, it's clear that it sets the necessary constraints and goals clearly with the purpose of the feature. This clarity can drastically increase output quality.

A well-structured prompt reads like a recipe card: short, direct, and instructive. Every “maybe,” “just,” or “please” dilutes intention and adds noise. The goal isn’t to be verbose or polite—it’s to be clear.
### Examples
```markdown
Vague prompt with too much noise:

"Please build a new app that allows home cooks to take a picture of their pantry or freezer to suggest recipes. Remember any allergies or preferences. Thanks"

TC-EBC version:
- `Task: Build an AI-powered meal suggestion app using pantry/fridge photo inputs`
- `Context: Home cooking assistant for households with dietary restrictions`
- `Elements: Camera input, pantry scanner, dietary settings form, meal suggestions list, recipe cards`
- `Behavior: User uploads photos; app scans inventory, filters by diet prefs, suggests recipes`
- `Constraints: Mobile-first, iOS/Android, accessible UI, supports multiple household profiles`
```

```markdown
Vague prompt missing intent, constraints and adding noise with incertainty:
"Write a description for this feature. Keep it simple but also exciting. Maybe like how Apple does it?"

TC-EBC version:
- `Task: Write a short product feature description.`
- `Context: For a new “One-Click Export” feature in a design tool.`
- `Elements: Headline (max 7 words), subheadline, single-sentence body copy.`
- `Behavior: Body should imply speed, simplicity, and trust.`
- `Constraints: No jargon. Match the brand tone of Duolingo or Notion. Total length: under 200 characters.`
```

>The more direct the language, the more efficient the exchange.

It's important to define what belongs in your context or not. Curating what the model sees, remembers and weights is what we call **context engineering**. The goal is to keep context focused so that the intention is pure.

>You can use a separate model to assist you building you prompts following the TC-EBC framework.
## Show, don't tell
In this prompt strategy, you're encouraging your generative AI model to create outputs that evoke emotion, build atmosphere, and reveal meaning in each specific task:

```markdown
## Scenario 1
- Telling: "Write a story about a man whos very sad"
- Showing: "Write a story about a man sitting alone in a dark room, turning a photograph over his hands while rain taps against the window"

## Scenario 2 
- Telling: "Describe a woman who is nervous about giving a speech"
- Showing: "Describe a woman backstage, pacing in small circles, palms slick with sweat as she rehearses her opening line under her breath"
```

This technique is intended to guide the model to craft more emotionally resonant and compelling responses.

By demonstrating what you want—rather than just explaining it—you give the machine learning model a clear template to follow, resulting in more accurate and consistent outputs.

 Instead of simply instructing your language models to “write professionally” or “sound poetic,” you show it what that looks like by providing a sample. This technique of reinforcement learning is especially powerful when you’re aiming for consistency,
## Prompting Frameworks
- PRD
- SDD
- TC-EBC
- Show, don't tell
## Guardrails
constraints you set up to limit what an AI agent can do wrong. They're not instructions - they're boundaries.
In software, guardrails come in many forms: type checkers that catch incorrect data shapes, test suites that catch regressions, linters that enforce code style, file access restrictions that prevent agents from touching production configs, and mandatory human review before code gets merged.

The key insight is that guardrails are _automated_. They don't require you to watch the agent constantly. They fire automatically when something goes wrong, giving the agent feedback in its observe phase or blocking a bad change before it lands.

Without guardrails, you're doing vibe coding - letting the AI do whatever it wants and hoping for the best. Guardrails are what make agentic engineering a disciplined practice. They let you give agents more autonomy without proportionally increasing risk.
## In practice
- **Type systems**: TypeScript's compiler, Python's mypy, Rust's borrow checker - catch bugs at build time before they reach runtime
- **Test suites**: If the agent's changes break existing tests, the agent knows immediately and can self-correct
- **Linting**: Enforces code style and catches common mistakes (unused variables, missing error handling)
- **File access restrictions**: Limit which directories the agent can read or write - keep it out of secrets, configs, and infrastructure code
- **Iteration limits**: Cap the number of times an agent can retry before escalating to a human
- **Sandboxing**: Run agents in isolated environments so mistakes don't affect production
- **Code review**: The ultimate guardrail - a human reviews every change before it ships
- **Scope limits**: Restrict agents to specific tasks rather than giving them free rein over the entire codebase

The best guardrails are the ones you'd want in place anyway, even without AI.
## Meaningful Links
- [SDD GitHub Repo](https://github.com/github/spec-kit)
- [GitHub Copilot VSCode Config Doc](https://code.visualstudio.com/docs/copilot/overview)
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
