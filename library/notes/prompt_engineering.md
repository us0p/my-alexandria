---
id: 20260513-prompt_engineering
type: concept
status: draft
tags:
  - AI
created: 2026-01-30
---
## TL;DR
Prompt engineering is the practice of designing effective inputs for Large Language Models to achieve desired outputs.
# Prompt Engineering
## Terminology
- **Machine Learning (ML)**: A subset of artificial intelligence that allows a machine to automatically learn and improve from large datasets without being explicitly programmed.
- **Deep Learning (DL)**: Subset of ML that uses neural networks with multiple layers to learn from and make decisions based on complex data patterns. The separate layers allow the model to decompose the problem into features that each layer processes individually.
## Tokens
A token is a unit of text that was broken until it became one of word, character sets, or combination of words and punctuation. This is done by a tokenizer. The way the tokens are generated is based on how the LLM was trained. The set of unique tokens that an LLM is trained on is known as vocabulary. Common tokenization methods include: Word tokenization (text split into individual words based on delimiter), Character tokenization (text is split into individual characters), Subword tokenization (text is split into partial words or character sets).

|Token size|Pros|Cons|
|---|---|---|
|Smaller tokens (character or subword tokenization)|- Enables the model to handle a wider range of inputs, such as unknown words, typos, or complex syntax.  <br>- Might allow the vocabulary size to be reduced, requiring fewer memory resources.|- A given text is broken into more tokens, requiring additional computational resources while processing.  <br>- Given a fixed token limit, the maximum size of the model's input and output is smaller.|
|Larger tokens (word tokenization)|- A given text is broken into fewer tokens, requiring fewer computational resources while processing.  <br>- Given the same token limit, the maximum size of the model's input and output is larger.|- Might cause an increased vocabulary size, requiring more memory resources.  <br>- Can limit the model's ability to handle unknown words, typos, or complex syntax.|
After the tokenization is completed, it assigns the token id to each, e.g.:
```plaintext
- `I` (1)
- `heard` (2)
- `a` (3)
- `dog` (4)
- `bark` (5)
- `loudly` (6)
- `at` (7)
- `a` (the "a" token is already assigned an ID of 3)
- `cat` (8)
```

By assigning IDs, text can be represented as a sequence of token IDs. e.g. `[1, 2, 3, 4, 5, 6, 7, 3, 8]`. The sentence "`I heard a cat`" would be represented as `[1, 2, 3, 8]`.

The `TokenID` is used to fetch a vector from a huge table called the **embedding matrix**
```plaintext
1 -> [0.1, -0.8, 1.1, ...]
2 -> [0.7, 0.1, 0.4, ...]
3 -> [-0.3, 0.9, 1.5, ...]
```

So basically the Token IDs are used as a lookup tool to get the embeddings for that token.
The embeddings is where the semantic information lives.

A model can calculate an embedding for the text that contains multiple tokens. It calculates the overall embedding value based on the learned embeddings of the individual tokens. This can be used for semantic document searches.

During output generation, the model predicts a vector value for the next token in the sequence. The model evaluates all potential tokens from these vectors and selects the most probable one to continue the sequence. In practice the model calculate multiple vectors by using various elements of the previous tokens' embeddings.

This is an iterative process, so the model continues this until it predicts the final token. It builds the final output one token at a time.
### How embeddings capture meaning
You can imagine embeddings as coordinates in a giant space (a real embedding usually have hundreds or thousands of dimensions, not just 2).

similar or related words have similar embeddings (or are close together in this space).
### Token limits
Also known as **context window**, it covers both input and output token limits together. Taken together, a model's token limit and tokenization method determine the maximum length of text that can be provided as input or generated as output.

For example, consider a model that has a maximum context window of 100 tokens. The model process the following input text: `I heard a dog bark loudly at a cat`.

By using a word-based tokenization method, the input is 9 tokens. This leaves 91 word tokens available for the output.
By using character-based, the input is 34 tokens (including spaces). This leaves only 66 character tokens available for the output

Generative AI services often use token-based pricing. Pricing might differ between input and output. They also enforce a maximum number of **Token  Per Minute (TPM)**.

>Since each provider has its own tokenization method, the same text can generate different token counts across models.
>It's also work noticing that different models have different context window limitation. Make sure to check your model capacity.

Output tokens almost always cost more than input tokens. The difference is **often 3 to 8 times higher for output**. This means the model’s response is the bigger cost driver, not your prompt.

Shorter, focused prompts that produce concise answers save money at scale. A chat-bot handling 10,000 conversations daily will notice significant cost differences between average reply lengths. Reducing average output from 500 tokens to 200 tokens per response **cuts output costs by 60%**.

>The most common and frequent used the word is, the biggest is the chance of being a single token. So using common words allows you to save tokens.

>The standard tokenization method is Byte Pair Encoding (BPE) or some variation of it. It starts with individual characters and repeatedly merges the most common pairs into single tokens.

Using tokens allow LLMs to handle all words in a vocabulary, even if you use a slang or the phrase has a typo the LLM will be able to understand it by breaking the words into smaller tokens.
### The `0.75` rule for estimating token counts
A rule of thumb: one English word equals roughly 1.3 tokens. Flipped around, one token covers about 0.75 words. A 1k word document becomes approximately 1.3k tokens.

>Different kinds of texts has different token count, for example, plain English words produces around 1.3 tokens on average. Code usually produces a 2x number of tokens compared to the number of words. Non-English languages often reuses 2 to 4 times more tokens per word than English. This is all based on how frequently the words was seen during model training.

Token-based billing makes costs transparent and calculable ahead of time. You can estimate expenses before running a task by counting tokens in your planned input and expected output length. But be ware that this is dependent on the tokenizer execution and it's sometimes unintuitive.
### Tokens and response quality
Context window shape output quality in ways that are less obvious. When the model has room to process more tokens, it can maintain coherence over longer responses. Cramming too much into a single prompt leaves fewer tokens for the reply.

This creates a practical trade-off. A detailed 2K token prompt given the model more context but less room to reply. A minimal 200 token prompt provides plenty of output space but less guidance.
### Tokens and conversation memory
In a multi-turn conversation, every previous message stays in the context window. All of these messages accumulate tokens. After enough back-and-forth exchanges, the total exceeds the limit. The usual behavior in chat interfaces is for the platform to silently drop oldest messages causing the model to lose access to instructions or context from early in the conversation.

For long conversations with custom instructions set at the start of the conversation can disappear once the token total grows too large.

Because of this, hidden overhead eats into your limit. All custom prompt that give some kind of instruction consume some tokens invisibly. A chat-bot with a 500-token system prompt starts every conversation with less available context than the started limit suggests.

Also, multimodal inputs add up quickly. Files (including images and audios) get converted into token equivalents before processing. A single high-resolution image can cost over 1K tokens.

>Important to know: More space doesn't guarantee better output. Models can struggle with accuracy over very long inputs. Information buried in the middle of a long context often gets recalled less reliably than content near the start or end.
## Fine-tuning
Fine-tuning is the process of further training a pretrained language model using additional data. This causes the model to start representing and mimicking the patterns and characteristics of the fine-tuning dataset.
Fine-tuning can be useful for adapting a language model to a specific domain, task, or writing style, but it requires careful consideration of the fine-tuning data and the potential impact on the model's performance and biases.
## Latency
Latency, in the context of generative AI and large language models, refers to the time it takes for the model to respond to a given prompt. It is the delay between submitting a prompt and receiving the generated output. Lower latency indicates faster response times, which is crucial for real-time applications, chatbots, and interactive experiences. Factors that can affect latency include model size, hardware capabilities, network conditions, and the complexity of the prompt and the generated response.
## LLM
Large language models (LLMs) are AI language models with many parameters that are capable of performing a variety of surprisingly useful tasks. These models are trained on vast amounts of text data and can generate human-like text, answer questions, summarize information, and more.
## MCP (Model Context Protocol)
Model Context Protocol (MCP) is an open protocol that standardizes how applications provide context to LLMs. Like a USB-C port for AI applications, MCP provides a unified way to connect AI models to different data sources and tools. MCP enables AI systems to maintain consistent context across interactions and access external resources in a standardized manner.
## Pretraining
Pretraining is the initial process of training language models on a large unlabeled corpus of text.
These pretrained models are not inherently good at answering questions or following instructions, and often require deep skill in prompt engineering to elicit desired behaviors. Fine-tuning and RLHF are used to refine these pretrained models, making them more useful for a wide range of tasks.
## RAG (Retrieval augmented generation)
Retrieval augmented generation (RAG) is a technique that combines information retrieval with language model generation to improve the accuracy and relevance of the generated text, and to better ground the model's response in evidence. In RAG, a language model is augmented with an external knowledge base or a set of documents that is passed into the context window. The data is retrieved at run time when a query is sent to the model, although the model itself does not necessarily retrieve the data (but can with tool use and a retrieval function). When generating text, relevant information first must be retrieved from the knowledge base based on the input prompt, and then passed to the model along with the original query. The model uses this information to guide the output it generates. This allows the model to access and utilize information beyond its training data, reducing the reliance on memorization and improving the factual accuracy of the generated text. RAG can be particularly useful for tasks that require up-to-date information, domain-specific knowledge, or explicit citation of sources. However, the effectiveness of RAG depends on the quality and relevance of the external knowledge base and the knowledge that is retrieved at runtime.
## RLHF
Reinforcement Learning from Human Feedback (RLHF) is a technique used to train a pretrained language model to behave in ways that are consistent with human preferences. This can include helping the model follow instructions more effectively or act more like a chatbot. Human feedback consists of ranking a set of two or more example texts, and the reinforcement learning process encourages the model to prefer outputs that are similar to the higher-ranked ones.
## Temperature
Temperature is a parameter that controls the randomness of a model's predictions during text generation. Higher temperatures lead to more creative and diverse outputs, allowing for multiple variations in phrasing and, in the case of fiction, variation in answers as well. Lower temperatures result in more conservative and deterministic outputs that stick to the most probable phrasing and answers. Adjusting the temperature enables users to encourage a language model to explore rare, uncommon, or surprising word choices and sequences, rather than only selecting the most likely predictions.

Users may encounter non-determinism in APIs. Even with temperature set to 0, the results will not be fully deterministic and identical inputs may produce different outputs across API calls.
## TTFT (Time to first token)
Time to First Token (TTFT) is a performance metric that measures the time it takes for a language model to generate the first token of its output after receiving a prompt. It is an important indicator of the model's responsiveness and is particularly relevant for interactive applications, chatbots, and real-time systems where users expect quick initial feedback. A lower TTFT indicates that the model can start generating a response faster, providing a more seamless and engaging user experience. Factors that can influence TTFT include model size, hardware capabilities, network conditions, and the complexity of the prompt.
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
