---
id: 20260408-text_generation_inference
tags:
  - AI
created: 2026-04-08
status: draft
---
# Text Generation Inference
Process of generating human like text from a **Prompt**. [LLMs](large_language_model.md) use the [training](training_models.md) to formulate responses one word at a time. This sequential generation is what allows LLMs to produce coherent and contextually relevant text.

- **Context Length**: Maximum number of [Tokens](tokenization.md) that the LLM can process at once. Models with different context length capabilities are designed to balance capability with efficiency. The capability is limited by:
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
## Key Performance Metrics
There are four critical metrics that shape implementation decisions:
1. **Time to First Token (TTFT)**: How quickly can you get the first response? Primarily affected by the prefill phase and crucial for user experience.
2. **Time Per Output Token (TPOT)**: How fast can you generate subsequent tokens? Determines the overall generation speed.
3. **Throughput**: How many requests can be handled simultaneously? Affects scalling and cost efficiency.
4. **VRAM Usage**: How much GPU memory do you need?
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
## The context length challenge
Longer context provide more information but come with substantial costs:
- Memory usage: Grows quadratically with context length
- Processing Speed: Decreases linearly with longer contexts
- Resource Allocation: Requires careful balancing of VRAM usage

The key is finding the right balance for your specific use case.
## KV (Key-Value) Cache Optimization
Stores and re-use intermediate calculations:
- Reduces repeated calculations
- Improves generation speed
- Makes Long-Context generation practical

The trade-off is additional memory usage, but benefits outweight the cost.
## Getting context out of numbers
The contextual understanding of an input by the [Transformer Model](transformer_architecture.md) is represented by a **High-Dimensional** vector.

A **High-Dimensional** vector in this interpretation is a vector that generally has three dimensions:
- **Batch size**: The number of sequences processed at a time.
- **Sequence length**: The length of the numerical representation of the sequence.
- **Hidden size**: The vector dimension of each model input.

It's said to be high-dimensional because of the last dimension. The hidden size can be very large (768 is common in smaller models, and in large it can reach 3072 or more).

>Batching is the act of sending multiple sentences through the model, all at once.
## Model Heads: Making sense out of numbers
A head takes the high-dimensional vector of hidden states as input and project them onto a different dimension.

![[nlp_model_inference_workflow.png]]

In this diagram, the model is represented by its embeddings layer and the subsequent layers. The embeddings layer converts each inpu tID in the tokenized input into a vector that represents the associated token. The subsequent layers manipulate those vectors using the attention mechanism to produce the final representation of the sentences.

The model head is an additional component, used to convert the transform predictions to a task-specific output, in this case, sequence classification.

The output for this step are logits.

Those logits are then converted to probabilities (possible outputs based on labels (the word associated with a token ID)).

For example, a model that classifies sentiment on a input has two labels 'NEGATIVE' and 'POSITIVE', let's say that after we converted the head's logits into probabilities we have the following result: `0.0402, 0.9598`.  With this we know that the provided input has 4% chance of being negative and 96% chance of being positive.

Each label receives a probability from the model, the biggest probability is used as the final output.

This are the steps a model goes through:
    - preprocessing with tokenizers
    - passing tokens through the model to generate logits
    - postprocessing, getting the probabilities of each token and producing output.
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
