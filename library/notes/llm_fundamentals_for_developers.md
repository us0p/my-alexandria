---
id: 20260627-llm_fundamentals_developers
type: concept
status: refined
tags: [AI, LLM, agents, development]
created: 2026-06-27
---
## TL;DR
# Large Language Models
## 1. How LLMs work
An LLM is built on the **transformer architecture**. You don't need the math, but two ideas explain most behavior you'll observe:
- **Encoder vs. Decoder.** A transformer has two block types. The *encoder* builds an understanding of input (good for classification, search, entity extraction). The *decoder* generates text one token at a time (good for chat, completion, code). **Most modern LLMs (GPT, Claude, Gemini, Llama) are decoder-only** — they are optimized to continue text, which is why they feel like autocomplete on steroids.
- **Attention.** A layer that lets the model weigh which earlier words matter when interpreting the current one. This is why context matters so much: the model literally "looks back" at everything in the prompt to decide the next token. The practical consequence — *every token in your prompt influences the output, and longer context costs more compute.*
**Why this matters to you:** the model is **autoregressive** — it generates sequentially, token by token, each new token conditioned on all previous ones. That single fact explains latency (long outputs are slow), statefulness (the model has no memory beyond what you resend), and why prompt order and content change results.
## 2. Tokens
A **token** is a chunk of text — a word, a sub-word, or punctuation — that the model actually processes. Text is split into tokens by a **tokenizer** before the model ever sees it. The dominant method is **Byte Pair Encoding (BPE)**: it starts from characters and merges the most common pairs into single tokens. Common words become one token; rare words, typos, and slang get split into several sub-word tokens (so the model can still handle them).

Token-counting rules of thumb (English):
- **1 word ≈ 1.3 tokens**, or **1 token ≈ 0.75 words**. A 1,000-word document ≈ 1,300 tokens.
- **Code ≈ 2× more tokens** per word than prose.
- **Non-English text ≈ 2–4× more tokens** per word.
- A single high-resolution **image can cost 1,000+ tokens** once converted for a multimodal model.

**Why this matters to you:** APIs bill **per token**, and **output tokens usually cost 3–8× more than input tokens** — the model's *reply* is your main cost driver, not your prompt. Cutting average output from 500 → 200 tokens per response cuts output cost ~60%. Because each provider tokenizes differently, the *same text yields different token counts across models* — always count with the target model's tokenizer, and prefer common words (they're more often a single token).
## 3. Text generation flow
Inference runs in two phases:
1. **Prefill** — the model reads your *entire* prompt at once: tokenize → convert tokens to embeddings (vectors that capture meaning) → run them through the network. This is **compute-intensive** and dominates **Time To First Token (TTFT)**. Think "reading the whole question before answering."
2. **Decode** — the model generates the response **one token at a time**, each token attending to all previous ones. This is **memory-intensive** (it must keep prior tokens and their relationships) and determines **Time Per Output Token (TPOT)**.
Internally each step produces **logits** (raw, unnormalized scores per candidate token), which a **softmax** function turns into probabilities. The decoding strategy then picks the next token from that probability distribution.
Key performance metrics to design around:
- **TTFT** — responsiveness; driven by prefill and prompt size.
- **TPOT** — generation speed per token.
- **Throughput** — concurrent requests handled (scaling/cost).
- **VRAM usage** — GPU memory needed.

**Optimization you benefit from for free:** the **KV (Key-Value) cache** stores intermediate attention calculations so previously seen tokens aren't recomputed. It trades extra memory for much faster generation and is what makes long-context responses practical.
## 4. Agentic capabilities
Beyond plain text, most production models expose:

- **Tool calling (function calling).** The model can decide to call an external function you define (search, DB query, API call), receive the result, and fold it into its reasoning and final answer. *You* run the code; the model just chooses when and with what arguments. Handle the actual action in your code and grant least privilege — never let the model execute privileged operations directly.
- **Structured output.** Constrain the response to a defined schema (e.g., JSON). Essential for reliable parsing, classification, and chaining steps. Also useful as a cheap guardrail (e.g., a small model that returns `{"safe": true|false}`).
- **Multimodality.** Accept/return more than text — images, audio, files. Remember these convert into token-equivalents and consume your context budget.
- **Reasoning.** Multi-step planning before answering. This is the capability that turns an LLM into an **agent**: it drives the decide → act → observe loop — *which* tool to call, how to interpret results, and *when* to stop and answer.

An **agent** = LLM + planning/reflection + tool access + memory. **MCP (Model Context Protocol)** is an open standard ("USB-C for AI") for exposing tools and data to models in a uniform way.
## 5. Decoding & generation control
After logits are produced, these knobs shape token selection (set them per request via the API):

- **Temperature** — randomness. `<1` = focused, deterministic, repetitive-safe; `>1` = creative, diverse, more error-prone. Use low (0–0.3) for extraction/classification/code, higher (0.7–1.0) for brainstorming/creative writing. **Note: even `temperature=0` is not fully deterministic** across API calls.
- **Top-k** — only consider the *k* most likely tokens. Caps weirdness but a flat distribution can still admit odd tokens.
- **Top-p (nucleus sampling)** — consider the smallest set of tokens whose probabilities sum to *p* (e.g., 0.9). Adapts to the distribution; generally preferred over top-k.
- **Presence penalty** — fixed penalty for any token already used (discourages reusing words at all).
- **Frequency penalty** — scaling penalty that grows with how often a token was used (discourages loops). Penalties are applied early, before sampling.
- **Beam search** — explores several candidate sequences (typically 5–10) in parallel and keeps the most probable overall. More coherent, more compute; rarely exposed in chat APIs.

Control **output length** with token limits (min/max), **stop sequences**, or letting the model emit its end-of-sequence token naturally.
## 6. Context windows
The **context window** is the maximum number of tokens the model can handle **at once, input + output combined**. Modern models range from ~8K to 128K, with some reaching ~1M tokens. Practical realities:

- **Memory grows quadratically** and **processing speed drops roughly linearly** as context lengthens — bigger context is slower and pricier.
- **Input and output share the budget.** A 2K-token prompt leaves less room to reply than a 200-token one. More context = more guidance but less space to answer.
- **Conversation memory is your job.** The model is stateless; in a multi-turn chat you resend the whole history every turn, and tokens accumulate. When the total exceeds the window, chat platforms **silently drop the oldest messages** — early system instructions can vanish mid-conversation.
- **Hidden overhead.** A 500-token system prompt is subtracted from every request's usable budget before the user even types.

Mitigations: **server-side compaction** (automatic summarization of older turns), **context editing** (keep only the messages relevant to the current step), and RAG (retrieve just what's needed instead of stuffing everything in).
## 7. Limitations (brief and factual)
- **Hallucination** — fluent, confident output that is false. Inherent to the architecture; mitigate with RAG, tool calling, and citation grounding.
- **Bias** — reflects patterns in training data; can surface in outputs.
- **Context rot** — degraded recall for content in the middle of long inputs.
- **Weak exact logic/math** — the architecture approximates patterns; it's unreliable for precise arithmetic and float math. Offload to tools.
- **Cost & latency** — both scale with tokens; output tokens dominate cost, long outputs dominate latency.
## Understanding
## When to Use
- **Open-ended language tasks**: summarization, drafting, rewriting, translation, Q&A.
- **Conversational interfaces** needing context retention across turns (chatbots, support).
- **Agentic workflows**: tasks needing planning + external tools (research assistants, coding agents).
- **Flexible parsing/classification** where rules are fuzzy and structured output can constrain results.
- **Retrieval-grounded answers** over your own documents (RAG).
## When NOT to Use
- **Exact computation or deterministic logic** — use real code/a calculator; the model approximates.
- **Tasks with a guaranteed-correct algorithm** (sorting, validation, math) — code is cheaper, faster, and correct.
- **Hard latency/cost ceilings at scale** where a small specialized model or classic ML suffices.
- **Single, narrow, high-volume tasks** (e.g., sentiment on one domain) — a small fine-tuned/encoder model can be cheaper and more reliable than a general LLM.
- **When you can't tolerate hallucination and can't ground or verify** the output.
## Trade-offs
- **Quality vs. latency/cost** — bigger models and longer outputs are better but slower and pricier; output tokens cost 3–8× input.
- **Creativity vs. reliability** — high temperature/top-p = diverse but error-prone; low = consistent but dull.
- **Context size vs. performance** — more context gives more grounding but quadratic memory, linear slowdown, and context rot.
- **Input vs. output budget** — they share the window; richer prompts leave less room to answer.
- **Determinism** — not guaranteed even at `temperature=0`; design tests to tolerate variance.
- **Repetition vs. fluency** — penalties cut loops but, if too strong, strip necessary key terms.
## Examples
- **Customer-support chatbot** — needs multi-turn memory; you resend history, use a system prompt for tone/policy, watch for dropped early instructions, and use low temperature for consistent answers.
- **Coding assistant** — decoder model, low temperature, structured output for diffs/JSON, tool calling to run tests or read files; code-heavy prompts cost ~2× tokens.
- **RAG document Q&A** — retrieve only relevant chunks into the context window instead of dumping the whole corpus; ask for quotes + citations to curb hallucination.
- **Research agent** — reasoning model loops over tool calls (search, fetch), plans steps, and stops when confident; privileged actions run in your code with least privilege.
- **Cost estimation** — a chatbot at 10,000 conversations/day: trimming average reply 500→200 tokens cuts output spend ~60%.
## References
## Questions
- When is a fine-tuned small model the better economic choice over a large general LLM + good prompting?
- How do different providers' tokenizers diverge in practice, and how much does that shift real cost estimates?
- What are robust, measurable strategies for mitigating context rot in very long (>200K token) contexts?
- How should conversation memory be compacted without losing critical early instructions?
## Flashcards