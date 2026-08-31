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
Guardrails are automated constraints bounding an AI agent's actions, catching or blocking mistakes without constant supervision — without them it's pure Vibe Coding. In practice they include type systems, test suites, linting, file-access restrictions, iteration limits, sandboxing, code review, and scope limits, firing automatically during the agent's observe phase or blocking bad changes before they land. Implementation happens across four layers: training/fine-tuning (baked into model weights), prompt/system instructions, runtime filtering (external checks on generated output), and tool/capability restrictions in agent systems.
# AI Guardrails
Automated constraints that bound an AI agent's actions so mistakes are caught or blocked without constant supervision. Without them you're doing pure [Vibe Coding](vibe_coding.md) — hoping for the best; with them, AI assistance becomes a disciplined practice.
## In practice
- **Type systems** (TypeScript, Rust borrow checker): catch bugs at build time.
- **Test suites**: break on regressions so the agent self-corrects.
- **Linting**: enforces style and catches common mistakes.
- **File-access restrictions**: keep the agent out of secrets, configs, and infra (see off-limits zones in [Context Files](claude_context_files.md)).
- **Iteration limits**: cap retries before escalating to a human.
- **Sand-boxing**: isolate execution from production.
- **Code review**: the ultimate guardrail — a human reviews every change.
- **Scope limits**: restrict the agent to specific tasks, not the whole codebase.

>Because guardrails are automated, they don't require watching the agent: they fire on failure during its observe phase or block a bad change before it lands. 
## Guardrails implementation
Guardrails are implemented as layers on an [LLM](llm_fundamentals.md)/[Agent](llm_agent.md) according to requirements.
- **Layer 1 (Training & Fine-Tuning):** Guardrails are baked into the model's weights during training.
- **Layer 2 (Prompt/System Instructions):** The system prompt might include guidelines the model follows.
- **Layer 3 (Runtime Filtering):** After the model generates a response, external system checks it.
- **Layer 4 (Tool/Capability Restrictions):** In agent systems, guardrails can appear as tool restrictions.
## Understanding
Guardrails are like [CI/CD](cicd.md) pipelines that are applied to AI output. It's a deterministic safe-guard added to the generated output so that we can make sure that no "hard" errors are missed.
## When to Use
- Use it whenever you have to make sure that AI output must follow deterministic metrics like a type system, tests output, linting, etc.
- When you have an AI loop and your agent is working mostly alone. Guardrails can be used as a step in the loop to make sure that the output of each iteration is working, before next iteration.
## When NOT to Use
- If your use-case is simple and don't need strict validation, like question and answers. Although it would be interesting to add a [Prompt Injection Guardrail](prompt_injection.md)
## Trade-offs
- **Safety x Capability**: Stronger restrictions reduce harmful outputs but can limit legitimate uses and helpfulness.
- **Transparency x Security:** Explaining how guardrails work makes them easier to circumvent.
- **Flexibility x Consistency**: Context aware guardrails (LLM powered) handles nuance but create inconsistency and unpredictability.
- **Upstream x Dowstream Controls:** If guardrail is implemented upstream (during training) it prevents certain knowledge from existing. It affects all users uniformly but it's costly to implement. If the guardrail is implemented downstream (at inference time), responses become slower, can be bypassed, but allows targeted control.
- **Speed x Accuracy**: Real-time guardrails must be fast but sophisticated detection is computationally expensive.
- **Cost x Coverage**: Comprehensive guardrails are expensive; limited ones leaves gaps.
## Examples
## References
- [Prompt Injection Guardrail](prompt_injection.md)
- [CI/CD](cicd.md)
- [Vibe Coding](vibe_coding.md)
- [Context Files](claude_context_files.md)
- [LLM](llm_fundamentals.md)
- [Agent](llm_agent.md)
## Flashcards
- Q: What are AI guardrails?
- A: Automated constraints that bound an AI agent's actions so mistakes are caught or blocked without constant supervision.
- Q: What happens without AI guardrails?
- A: You are doing pure Vibe Coding, hoping for the best, instead of a disciplined practice.
- Q: What do type systems do as a guardrail?
- A: Catch bugs at build time.
- Q: What do test suites do as a guardrail?
- A: Break on regressions so the agent self-corrects.
- Q: What does linting do as a guardrail?
- A: Enforces style and catches common mistakes.
- Q: What do file-access restrictions do as a guardrail?
- A: Keep the agent out of secrets, configs, and infra.
- Q: What do iteration limits do as a guardrail?
- A: Cap retries before escalating to a human.
- Q: What does sandboxing do as a guardrail?
- A: Isolates execution from production.
- Q: What is the ultimate guardrail?
- A: Code review, where a human reviews every change.
- Q: What do scope limits do as a guardrail?
- A: Restrict the agent to specific tasks, not the whole codebase.
- Q: Why don't automated guardrails require watching the agent?
- A: Because they fire on failure during its observe phase or block a bad change before it lands.
- Q: On what are guardrails implemented as layers?
- A: An LLM or Agent, according to requirements.
- Q: What happens in Layer 1, Training and Fine-Tuning, of guardrail implementation?
- A: Guardrails are baked into the model's weights during training.
- Q: What happens in Layer 2, Prompt or System Instructions, of guardrail implementation?
- A: The system prompt might include guidelines the model follows.
- Q: What happens in Layer 3, Runtime Filtering, of guardrail implementation?
- A: After the model generates a response, an external system checks it.
- Q: What happens in Layer 4, Tool or Capability Restrictions, of guardrail implementation?
- A: In agent systems, guardrails can appear as tool restrictions.
