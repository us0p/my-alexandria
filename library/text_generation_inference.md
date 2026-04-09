---
id: 20260408-text_generation_inference
tags:
  - AI
created: 2026-04-08
status: draft
---
# Text Generation Inference
Process of generating human like text from a **Prompt**. [LLMs](large_language_model.md) use the [training](training_models.md) to formulate responses one word at a time. This sequential generation is what allows LLMs to produce coherent and contextually relevant text.

- **Context Length**: Maximum number of [Tokens]() that the LLM can process at once. Models with different context length capabilities are designed to balance capability with efficiency. The capability is limited by:
	1. Architecture Size.
	2. Computational Resources.
	3. Complexity of input and desired output.
- **Prompting**: The models predicts the next word based on the understanding and context of the previous words in your sentence thus making wording or your input crucial.
## Inference 1st step: Pre-fill
Step in which input is processed. The model needs to read all tokens at once which makes it a very compute intensive task.

Think of it like reading all the paragraph before starting to write a response.

It involves three key steps:
1. **Tokenization**: Converting input text into **Tokens**.
2. **Embedding Conversion**: Converting tokens into numerical representation that capture their meaning.
3. **Initial Processing**: Runs the embeddings through the model's neural network to create rich understanding of the context.
## Inference 2nd step: Decode
Where text generation happens. The model generates one token at the time, each new token depending on all the previous tokens.

This process is **auto-regressive**.

Here are some of the steps that happens inside this phase for **each new token**:
- **Attention Computation**: Looking back to all previous tokens to understand context.
- **Probability Calculation**: Determining the likelihood of each possible next token.
- **Token Selection**: Choosing the next token based on calculated probabilities.
- **Continuation Check**: Deciding where to continue or stop generation.

>This phase is **memory intensive** as the model needs to keep the previous generated tokens and their relationships.
## Sampling strategies
Allow us to control the generation process by adjusting the way the model makes its token selection.
## Understanding token selection
**Logits** is the name of the probabilities of all the words in the model's vocabulary.

![[token_selection_workflow.png]]

1. **Raw Logits**: Initial guess about each possible next word.
2. **Temperature Control**: Higher settings (>1) make choices more random and creative, lower settings (<1) make them more focused and deterministic.
3. **Top-p (Nucleus)**: Also known as **Sampling**. Instead of considering all possible words, it only look at the most likely ones that add up to the chosen probability threshold (e.g. top 90%).
4. **Top-k Filtering**: Alternative approach where we only consider the **k** most likely next words.
## Token Penalties
[LLMs](large_language_models.md) tend to repeat themselves, to address this, there are two penalties:
- **Presence Penalty**: Fixed penalty applied to any token that has appeared before, regardless of how often. Helps preventing the model from using the same words.
- **Frequency Penalty**: A scaling penalty that increases based on how often a token has been used. The more a token appears, the less likely it's to be chosen again.

>These penalties are applied early in the token selection process, adjusting the raw probabilities before other sampling strategies are applied.

![[logit_penalty_workflow.png]]
## Controlling generation length
It can be controlled in three ways:
1. **Token limits**: Setting minimum and maximum token counts.
2. **Stop sequences**: Defining specific patterns that signal the end of generation.
3. **End-of-Sequence detection**: Letting the model naturally conclude its response.
## Token Selection Strategies: Beam Search
Instead of looking one token at a time, it explores multiple possible paths at the same time.

How it works:
1. At each step, maintain multiple candidate sequences (typically 5-10).
2. For each candidate compute probabilities for the next token.
3. Keep only the most promising combinations of sequences and next tokens.
4. Continue this process until reaching the desired length or stop condition.
5. Select the sequence with the highest overall probability.

>Length penalties don't impact Beam Search paths. It only influences the choice of sequences in the end towards longer or shorter sequences.

This approach often produces more coherent and grammatically correct text, though requiring more computational power.

![[beam_search_workflow.png]]
## Understanding
- explanation of the concept, using your own words.
- Focus on cause and effect.
Ex:
- This pattern exists because systems are likely to couple business rules and external details...
- The separation allows changing interfaces without having to rewrite central rules...
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
## TL;DR
Very short resume with only the essential information needed.
## Flashcards
- Q: Some question about the notes.
- A: The answer for the question above.
