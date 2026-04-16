---
id: 20260408-text_generation_inference
tags:
  - AI
created: 2026-04-08
status: draft
type: concept
---
## TL;DR
Text Inference is composed of two phases (**Prefill and Decoding**). The **decoding** phase can the managed by the use of the [Sampling Filters](decoding_strategies.md) and it's directly impacted by the **Context Length**. The [Transformer Model](transformer_architecture.md) is what extracts the meaning of the input and produces the initial **Embeddings** that are converted to [Logits](decoding_strategies.md) that generate the **probabilities** for each token.
# Text Generation Inference
Process of generating human like text from a **Prompt**. [LLMs](large_language_models.md) use the [training](training_models.md) to formulate responses one word at a time. This sequential generation is what allows LLMs to produce coherent and contextually relevant text.

The models predicts the next word based on the understanding and context of the previous words in your sentence thus making the [Prompt]() crucial.
## Inference 1st step: Pre-fill
Step in which input is processed. The model needs to read all tokens at once which makes it a very compute intensive task.

Think of it like reading all the paragraph before starting to write a response.

It involves three key steps:
1. [Tokenization](tokenization.md#Tokenization%20Steps): Converting input text into [Tokens](tokenization.md).
2. **Embedding Conversion**: Converting tokens into numerical representation that capture their meaning.
3. **Initial Processing**: Runs the embeddings through the model's neural network to create [Logits](decoding_strategies.md).
## Inference 2nd step: Decode
Where text generation happens. The model generates one token at the time, each new token depending on all the previous tokens.

This process is **auto-regressive**.

Here are some of the steps that happens inside this phase for **each new token**:
- [Attention Computation](transformer_architecture.md#Attention%20Layer): Looking back to all previous tokens to understand context.
- **Probability Calculation**: Determining the likelihood of each possible next token.
- **Token Selection**: Choosing the next token based on calculated probabilities.
- **Continuation Check**: Deciding where to continue or stop generation.

>This phase is **memory intensive** as the model needs to keep the previous generated tokens and their relationships.
## Sampling strategies
Allow us to control the generation process by adjusting the way the model makes its [Token Selection](decoding_strategies.md).
## Controlling generation length
It can be controlled in three ways:
1. **Token limits**: Setting minimum and maximum token counts.
2. **Stop sequences**: Defining specific patterns that signal the end of generation.
3. **End-of-Sequence detection**: Letting the model naturally conclude its response.
## Key Performance Metrics
There are four critical metrics that shape implementation decisions:
1. **Time to First Token (TTFT)**: How quickly can you get the first response? Primarily affected by the **prefill** phase and crucial for user experience.
2. **Time Per Output Token (TPOT)**: How fast can you generate subsequent tokens? Determines the overall generation speed.
3. **Throughput**: How many requests can be handled simultaneously? Affects scaling and cost efficiency.
4. **VRAM Usage**: How much GPU memory do you need?
## The context length challenge
**Context Length** is the maximum number of [Tokens](tokenization.md) that the LLM can process at once. Models with different context length capabilities are designed to balance capability with efficiency. The capability is limited by:
1. Architecture Size.
2. Computational Resources.
3. Complexity of input and desired output.

Longer context provide more information but come with substantial costs:
- **Memory usage**: Grows quadratically with context length
- **Processing Speed**: Decreases linearly with longer contexts
- **Resource Allocation**: Requires careful balancing of VRAM usage

>The key is finding the right balance for your specific use case.
## KV (Key-Value) Cache Optimization
Stores and re-use intermediate calculations:
- Reduces repeated calculations
- Improves generation speed
- Makes Long-Context generation practical

The trade-off is additional memory usage, but benefits out-weight the cost.
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

In this diagram, the **transformer model** is represented by its embeddings layer and the subsequent layers. The embeddings layer converts each input ID in the tokenized input into a vector that represents the associated token. The subsequent layers manipulate those vectors using the attention mechanism to produce the final representation of the sentences.

The model head is an additional component, used to convert the transformer predictions to a task-specific output, in this case, sequence classification. The output for this step are [Logits](decoding_strategies.md)

Those logits are then converted to probabilities using the **softmax function**.

Each label receives a probability from the model, the biggest probability is used as the final output.

This are the steps a model goes through:
    - pre-processing with [tokenizers](tokenization.md)
    - passing tokens through the model to generate [logits](decoding_strategies.md)
    - post-processing, getting the probabilities of each token and producing output.
## Understanding
The **prefill** phase of text inference is responsible for gathering understanding of the input, for that it uses concepts of [Tokenization](tokenization.md) to convert the input words into **Tokens** and start with some **pre-processing** of the selected tokens to generate the first **Logits**.
The **decoding** phase uses a combination of  [Decoding Strategies](decoding_strategies.md) and **Context management** to get the probabilities and meaning out of the **tokenized input** so that it can pick the **most probable tokens** for the response.
## Trade-offs
The temperature config can impact several aspects of your setup:
- **Diversity vs Reliability**: If your output is diverse it can't be reliable, higher temperature filters produces more diverse responses.
- **Exploration vs Control**: If you want to explore more token options your output might be more fluent but can be more error prone. Top-k and Top-p filters impact directly on this.
- **Latency vs Performance**: The decoding strategy you're going to use will impact directly in the performance of your application. This needs to be fully considered according to the type of application you're going to create.
- **Repetition vs Fluency**: Your response can sound a lot like a robot if you keep repeating the same tokens. If you need more fluent responses (like a chat bot) but this can make your responses more error prone.
## Examples
Consider a model that classifies sentiment on a input has two labels 'NEGATIVE' and 'POSITIVE', let's say that after we converted the head's logits into probabilities we have the following result: `0.0402, 0.9598`.  With this we know that the provided input has 4% chance of being negative and 96% chance of being positive.
## References
### Connects with
- [LLM](large_language_models.md)
- [NLP](natural_language_processing.md)
- [Tokenization](tokenization.md)
- [Transformer Architecture](transformer_architecture.md)
- [Decoding Strategies](decoding_strategies.md)
## Flashcards
- Q: What is Text Inference?
- A: It's the process of generating human like text from a **Prompt**. The model predicts the next word based on the probabilities it learnt during its training.
- Q: What are the two inference steps and their responsibilities?
- A: Prefill: Parse input into tokens and extract meaning. Decode: Select most probable tokens and converting it back to words.
- Q: What are the three key steps of the prefill phase of text inference?
- A: Tokenization, Embedding Conversion, and Initial Processing.
- Q: What are the steps that can happen for each new token in the decoding phase of text inference?
- A: Attention computation, Probability calculation, Token selection, Continuation check.
- Q: What are sampling filters goals? What are some example?
- A: The goal of sampling filters is to control how the model will make its token selection. Some example are: Temperature Scaling, Top-k and Top-p filters and Token Penalties.
- Q: Is it possible to control generation length? How?
- Yes, we can control generation length by using token limits by setting the minimum and maximum token counts or by using a end of sequence token to specific mark a sequence as finished when we need it.
- Q: What are the 4 key performance metrics? What are they important for?
- A: TTFT: how quickly you can get a response. TPOT: how fast can you generate subsequent tokens. Throughput: how many requests can be handled simultaneously. VRAM Usage: how much computing power is needed
- Q: What is context length and why is it important?
- A: It determines the amount of tokens the model can consider at once. It's important because it impacts directly in the performance and the correctness of the output.
- Q: What does impact the desired context length?
- A: Model's architecture size, computational power and relation between input complexity and desired output.
- Q: What are the impacts of increasing the context length?
- A: The memory grows quadratically and processing speed decreases linearly with the size of the context but responses get more accurate.
- Q: What is the KV optimization?
- A: It's a Key-Value Cache that caches intermediate calculations, it takes more memory but significantly increases performance for application.
- Q: Where does the contextual understanding of the input comes from?
- A: It comes from the Transformer model as it creates the token embeddings based on the input. The embeddings layers converts each token into a vector which is manipulated by the attention layer that produces the final representation of the sentences (still a numerical vector) the model heads converts the sentences representations back into Logits and finally converts them back into tokens with their respective probabilities using the softmax function.