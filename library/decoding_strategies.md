---
id: 20260413-token_selection
status: draft
tags:
  - AI
created: 2026-04-13
---
# Decoding strategies
![[token_selection_workflow.png]]

1. **Raw Logits**: Is a number that represents a model's raw, unnormalized score for each possible token **before** it gets turned into a probability. A logit only makes sense when compared with other logits and it's not bound between 0 and 1 (it's not a probability). To get a probability out a logit you need to apply **softmax function**.
2. **Temperature Control**: Higher settings (>1) make choices more random and creative, lower settings (<1) make them more focused and deterministic.
3. **Top-p (Nucleus)**: Also known as **Sampling**. Instead of considering all possible words, it only look at the most likely ones that add up to the chosen probability threshold (e.g. top 90%).
4. **Top-k Filtering**: Alternative approach where we only consider the **k** most likely next words.
### Token Penalties
[LLMs](large_language_models.md) tend to repeat themselves, to address this, there are two penalties:
- **Presence Penalty**: Fixed penalty applied to any token that has appeared before, regardless of how often. Helps preventing the model from using the same words.
- **Frequency Penalty**: A scaling penalty that increases based on how often a token has been used. The more a token appears, the less likely it's to be chosen again.

>These penalties are applied early in the token selection process, adjusting the raw probabilities before other sampling strategies are applied.

![[logit_penalty_workflow.png]]
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
Token selection is a group of strategies and filters that we apply to a group of **Token** to select the most probable ones to be present in the response.
## Trade-offs
Each filter in the decoding step applies different trade-off
- **Temperature Scaling**: Too low text becomes more deterministic but dull. Too high, text become more creative but more prone to errors.
- **Top-k**: Prevent's uncommon tokens but can still sneak in weird tokens if they are in the top k tokens.
- **Top-p (sampling)**: Adapts to the distribution as will consider the probability of each token but can still include weird tokens if distribution is flat.
- **Penalties**: Avoid loop and boring responses but if it's too strong can hurt the clarity of the response by avoiding repetitive key terms.
## Examples
Consider the following table of available tokens and their probabilities:

| Token | Probability |
| ----- | ----------- |
| Is    | 0.4         |
| Was   | 0.3         |
| Seems | 0.2         |
| Might | 0.1         |

A greedy decoding strategy would always pick the first token as it has the biggest probability. This kind of strategy is very deterministic as each input has always the same output but it leads to very dull and repetitive responses.
## References
### Connects with
- [LLM](large_language_models.md)
- [NLP](natural_language_processing.md)
- [Text Inference](text_generation_inference.md)
## TL;DR
The model also applies **Penalties** to repetitive tokens to avoid repetition in the response.
A well known strategy that looks at many possible paths instead of doing it sequentially is the **Beam Search** which considers many logits at the same time and pick only the most probable ones.
## Flashcards
- Q: What is a logit?
- A: It's the raw, unnormalized score of a token before passing through the softmax function.
- Q: What is the Temperature Scaling Filter?
- A: It's a filter that determines if the token choices should be more creative (higher settings) or more deterministic (lower settings).
- Q: What is Top-k filtering?
- A: It's a constant value that determines the number of tokens that must be considered. It tries to create more natural responses.
- Q: What is Top-p filtering?
- A: It's a filter that only considers tokens that the probabilities sums up to a certain threshold (usually 90%). It's more flexible that Top-k filtering.
- Q: What are the two token penalties that are usually applied?
- A: Presence Penalty: Fixed penalty that is applied to every token that has already appeared. Frequency Penalty: Scaling penalty that increases as the token keeps being used.
- Q: What is Beam Search and how does it makes it token selections?
- A: It's a decoding strategy that looks at many paths at once instead of doing it linearly. The steps are: 1. Keep a multiple candidate sequence. 2. Calculate the probabilities for the next token of each sequence. 3. Keep only the sequence with the highest overall probabilities. 4. Repeat this until the end and pick the sequence with the highest probability.
- Q: What are the two most important trade-offs to consider in a decoding strategy?
- A: Applying filters to give a more deterministic response can lead the model to more dull and repetitive answers and increasing the randomness of token selection can make the model more error prone.