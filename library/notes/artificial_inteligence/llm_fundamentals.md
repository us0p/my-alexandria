---
id: 20260627-llm_fundamentals
type: concept
status: refined
tags:
  - AI
created: 2026-06-27
---
## TL;DR
LLMs use the **transformer** (decoder-only) architecture, generating text **autoregressively** one **token** at a time via the **attention** mechanism. Text is tokenized with **Byte Pair Encoding (BPE)**; output tokens cost 3–8× more than input. Inference runs in two phases: **prefill** (drives TTFT) and **decode** (drives TPOT), accelerated by the **KV Cache**. The **context window** is shared by input and output; memory grows quadratically. Generation is shaped by **temperature**, **top-p**, and penalty settings. Key capabilities include **tool calling**, **structured output**, and **reasoning**. Core limitations: **hallucination**, **context rot**, weak exact math, and token-driven cost/latency.
# Large Language Models
## How LLMs work
An LLM is built on the **transformer architecture**. Two ideas explain most behavior you'll observe:
### Encoder vs. Decoder
A transformer has two block types. The *encoder* builds an understanding of input (good for classification, search, entity extraction). The *decoder* generates text one token at a time (good for chat, completion, code). **Most modern LLMs (GPT, Claude, Gemini, Llama) are decoder-only** — they are optimized to continue text, which is why they feel like autocomplete on steroids.
### Attention
A layer that lets the model weigh which earlier words matter when interpreting the current one. This is why context matters so much: the model literally "looks back" at everything in the prompt to decide the next token. The practical consequence — *every token in your prompt influences the output, and longer context costs more compute.*

Models are **autoregressive** — it generates sequentially, token by token, each new token conditioned on all previous ones. That single fact explains latency (long outputs are slow), statefulness (the model has no memory beyond what you resend), and why prompt order and content change results.
## Tokens
A **token** is a chunk of text — a word, a sub-word, or punctuation — that the model actually processes.

Text is split into tokens by a **tokenizer** before the model ever sees it. The dominant method is **Byte Pair Encoding (BPE)**: it starts from characters and merges the most common pairs into single tokens. Common words become one token; rare words, typos, and slang get split into several sub-word tokens (so the model can still handle them).
### Token-counting rules of thumb (English):
- **1 word ≈ 1.3 tokens**, or **1 token ≈ 0.75 words**. A 1,000-word document ≈ 1,300 tokens.
- **Code ≈ 2× more tokens** per word than prose.
- **Non-English text ≈ 2–4× more tokens** per word.
- A single high-resolution **image can cost 1,000+ tokens** once converted for a multimodal model.

APIs bill **per token**, and **output tokens usually cost 3–8× more than input tokens** — the model's *reply* is your main cost driver, not your prompt.

Cutting average output from 500 → 200 tokens per response cuts output cost ~60%.

Because each provider tokenizes differently, the *same text yields different token counts across models* — always count with the target model's tokenizer, and prefer common words (they're more often a single token).
## Text generation flow
Inference runs in two phases:
### Prefill
The model reads your *entire* prompt at once: tokenize → convert tokens to embeddings (vectors that capture meaning) → run them through the network. This is **compute-intensive** and dominates **Time To First Token (TTFT)**. Think "reading the whole question before answering."
### Decode
The model generates the response **one token at a time**, each token attending to all previous ones. This is **memory-intensive** (it must keep prior tokens and their relationships) and determines **Time Per Output Token (TPOT)**.
### KV (Key-Value) Cache
It stores intermediate attention calculations so previously seen tokens aren't recomputed. It trades extra memory for much faster generation and is what makes long-context responses practical.

Internally each step produces **logits** (raw, unnormalized scores per candidate token), which a **softmax** function turns into probabilities. The decoding strategy then picks the next token from that probability distribution.
### Key performance metrics
- **TTFT** — responsiveness; driven by prefill and prompt size.
- **TPOT** — generation speed per token.
- **Throughput** — concurrent requests handled (scaling/cost).
- **VRAM usage** — GPU memory needed.
## Agentic capabilities
Beyond plain text, most production models expose the following capabilities that are leveraged in agentic systems:
### Tool calling
The model can decide to call an external function you define (search, DB query, API call), receive the result, and fold it into its reasoning and final answer. *You* run the code; the model just chooses when and with what arguments. Handle the actual action in your code and grant least privilege — never let the model execute privileged operations directly.
### Structured output
Constrain the response to a defined schema (e.g., JSON). Essential for reliable parsing, classification, and chaining steps. Also useful as a cheap guardrail (e.g., a small model that returns `{"safe": true|false}`).
### Multimodality
Accept/return more than text — images, audio, files. Remember these convert into token-equivalents and consume your context budget.
### Reasoning
Multi-step planning before answering. This is the capability that turns an LLM into an **agent**: it drives the decide → act → observe loop — *which* tool to call, how to interpret results, and *when* to stop and answer.
## Decoding & generation control
After logits are produced, these filters shape token selection (set them per request via the API):
### Temperature
Randomness. `<1` = focused, deterministic, repetitive-safe; `>1` = creative, diverse, more error-prone.

Use low (0–0.3) for extraction/classification/code, higher (0.7–1.0) for brainstorming/creative writing.

>Note: even `temperature=0` is not fully deterministic across API calls.
### Top-k
Only consider the *k* most likely tokens. Caps weirdness but a flat distribution can still admit odd tokens.
### Top-p (nucleus sampling)
Consider the smallest set of tokens whose probabilities sum to *p* (e.g., 0.9).

Adapts to the distribution; generally preferred over **Top-k**.
### Presence penalty
Fixed penalty for any token already used (discourages reusing words at all).
### Frequency penalty
Scaling penalty that grows with how often a token was used (discourages loops). Penalties are applied early, before sampling.

>Control **output length** with token limits (min/max), **stop sequences**, or letting the model emit its end-of-sequence token naturally.
## Context windows
The **context window** is the maximum number of tokens the model can handle **at once, input + output combined**.

Modern models range from `~8K` to `128K`, with some reaching `~1M` tokens.

Practical realities:
- **Memory grows quadratically** and **processing speed drops roughly linearly** as context lengthens — bigger context is slower and pricier.
- **Input and output share the budget.** A `2K`-token prompt leaves less room to reply than a 200-token one. More context = more guidance but less space to answer.
- **Conversation memory is your job.** The model is stateless; in a multi-turn chat you resend the whole history every turn, and tokens accumulate. When the total exceeds the window, chat platforms **silently drop the oldest messages** — early system instructions can vanish mid-conversation.
- **Hidden overhead.** A 500-token system prompt is subtracted from every request's usable budget before the user even types.

Mitigations:
- **server-side compaction** (automatic summarization of older turns)
- **context editing** (keep only the messages relevant to the current step)
- [RAG](retrieval_augmented_generation.md) (retrieve just what's needed instead of stuffing everything in).
## LLMs Limitations
- **Hallucination**: fluent, confident output that is false. Inherent to the architecture; mitigate with RAG, tool calling, and citation grounding.
- **Bias**: reflects patterns in training data; can surface in outputs.
- **Context rot**: degraded recall for content in the middle of long inputs.
- **Weak exact logic/math**: the architecture approximates patterns; it's unreliable for precise arithmetic and float math. Offload to tools.
- **Cost & latency**: both scale with tokens; output tokens dominate cost, long outputs dominate latency.
## Understanding
Considering all of this, LLMs are a powerful tool to leverage when you need reasoning in automated tasks or when you need to handle a lot of information and take action or interpret the meaning of all of it. They can be used to help in your day to day activities but, due to the limitations like Hallucination, Biasing and specially around memory, they can't be assumed to be fully independent leading the decisions and choices in your hand.
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
## References
- [RAG](retrieval_augmented_generation.md)
## Questions
- When is a fine-tuned small model the better economic choice over a large general LLM + good prompting?
- How do different providers' tokenizers diverge in practice, and how much does that shift real cost estimates?
- What are robust, measurable strategies for mitigating context rot in very long (>200K token) contexts?
- How should conversation memory be compacted without losing critical early instructions?
- What is MCP, its uses and when to use it?
## Flashcards
- Q: What architecture do most modern LLMs (GPT, Claude, Gemini, Llama) use, and what makes them "decoder-only"?
- A: They use the transformer architecture in decoder-only mode — optimized to generate text one token at a time by continuing/completing text, like autocomplete.
- Q: What does the attention mechanism do?
- A: It lets the model weigh which earlier tokens matter when generating the current one — every token in the prompt influences output, and longer context costs more compute.
- Q: What does it mean that LLMs are autoregressive?
- A: They generate tokens sequentially, each new token conditioned on all previous ones. This explains latency (long outputs are slow), statelessness (no memory beyond what's resent), and prompt sensitivity.
- Q: What is a token and how are tokens created?
- A: A chunk of text (word, sub-word, or punctuation) the model processes. Tokens are created by a tokenizer using Byte Pair Encoding (BPE): starts from characters and merges the most common pairs.
- Q: What are the token-counting rules of thumb for English?
- A: 1 word ≈ 1.3 tokens; code ≈ 2× more tokens per word; non-English ≈ 2–4× more; a high-res image can cost 1,000+ tokens.
- Q: Why do output tokens cost more than input tokens in LLM APIs?
- A: Output tokens cost 3–8× more because generating each token requires a full decode pass, while input tokens are processed in parallel during prefill.
- Q: What are the two phases of LLM inference?
- A: Prefill (reads the entire prompt at once — compute-intensive, drives TTFT) and Decode (generates one token at a time — memory-intensive, drives TPOT).
- Q: What is the KV Cache and what problem does it solve?
- A: It stores intermediate attention calculations so previously seen tokens aren't recomputed each step, trading memory for much faster generation.
- Q: What are the four key LLM performance metrics?
- A: TTFT (time to first token, driven by prefill), TPOT (time per output token), Throughput (concurrent requests), and VRAM usage (GPU memory).
- Q: What is temperature and when should you use low vs. high values?
- A: Controls randomness. Low (0–0.3) for extraction/classification/code (focused, reliable); high (0.7–1.0) for brainstorming/creative writing (diverse, more error-prone).
- Q: What is the difference between Top-k and Top-p (nucleus) sampling?
- A: Top-k considers only the k most likely tokens; Top-p considers the smallest set of tokens whose probabilities sum to p. Top-p adapts to the distribution and is generally preferred.
- Q: What is the context window, and what key constraint does it impose?
- A: The maximum tokens the model can handle at once (input + output combined). Input and output share the budget, memory grows quadratically, and speed drops linearly with context length.
- Q: Why can early system instructions vanish mid-conversation in multi-turn chat?
- A: The model is stateless — the full history is resent each turn. When total tokens exceed the context window, chat platforms silently drop oldest messages, including early system prompts.
- Q: What are the main LLM limitations to design around?
- A: Hallucination (confident but false output), bias, context rot (poor recall for middle-of-context content), weak exact math/logic, and cost/latency scaling with tokens.
- Q: What is tool calling in LLMs?
- A: The model decides to invoke an external function you define (search, DB, API), receives the result, and folds it into its answer. Your code runs the actual action; the model only chooses when and with what arguments.