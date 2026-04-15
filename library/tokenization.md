---
id: 20260412-tokenization
tags: AI
created: 2026-04-12
status: draft
---
# Tokenization
**Token** a word, a segment of a word or a symbol in the [NLP Model's](natural_language_processing.md) vocabulary.

Tokenization is the process of mapping the **Token** to it's numerical representation and adding relevant information to the model (like special characters).
## Vocabulary
A model's vocabulary is the set of word, sub-words or characters that a model knows the meaning for.
## Tokenization steps
1. Split the input into words, sub-words, or symbols (punctuation), the result is called Tokens.
2. Map each token to an integer.
3. Adding additional inputs that may be useful to the model.

>This pre-processing needs to be done in exactly the same way as when the model was pre-trained.

The output of the tokenization step is then converted to **Tensors**. A tensor is basically a `NumPy` array with any number of dimensions.
## Tokenizers
Serve one purpose, translate text into numbers that can be processed by the model and back. Models can only process numbers, so tokenizers need to convert our text inputs to numerical data.

Along side that, all the words that aren't in the vocabulary need to be marked with a custom token.

It's a bad sign if you see that a tokenizer is producing a lot of these tokens which means you're losing information along the way.

Bellow are some strategies to convert tokens into numbers.
### Word-based
- Split raw text into words, punctuation yield different words. 
- Each "word" is assigned one ID starting from 0.
- The ID is used to identify each word. 

Considering that this method treats words like 'dog' and 'dogs' as different, and that we need to have an identifier for each word in the language. It's obvious that the size of the vocabulary becomes very large and related words loses context.

This strategy frequently produces **unknown tokens** making it difficult to preserve context.
### Character-based
Split text into characters rather than words.
- Vocabulary is much smaller.
- There are much fewer **unknown tokens**, since every work can be built from characters.
- It's less meaningful as each character doesn't mean a lot on its own.
- Produces a large amount of tokens to be processed by the model.
### Sub-word tokenization
If follows two principles:
- Frequently used words shouldn't be split into sub-words
- Rare words should be decomposed into **meaningful** sub-words.

For instance, `annoyingly` might be considered a rare word and could be decomposed into `annoying` and `ly`. Which are more likely to appear more frequently as sub-words, while at the same time the meaning of `annoyingly` is kept by the composite meaning of `annoying` and `ly`.

This approach provides a lot of semantic meaning, allowing us to have relatively good coverage with small vocabularies, and close to no **unknown tokens**.

>It's specially useful in agglutinative languages such as Turkish.	
## Understanding
A token is a sub-division of a word that can be mapped to a [NLP model's](natural_language_processing.md) vocabulary ID.

Those tokens are formed by splitting a word into sub-words and then mapping the sub-words to their respective ID and adding any models specific special token.

A **tokenizer** is what convert words into tokens and back, there are some strategies for convert words into tokens like **Word** and **Character** based but the most widely used is **Sub-Word** tokenization that only splits **rare** words into **meaningful** sub-words keeping a short vocabulary that don't produces a lot of **unknown tokens** and keeps the context of the words relationships.
## Trade-offs
- **Vocabulary Size vs Sequence Length**: 
	- Large vocabularies (word based) produces shorter sequences which improves [Attention](transformer_architecture.md#Attention%20Layer) computation but requires more memory and frequently produces **unknown tokens**.
	- Small vocabularies (character based) are capable of handling any word but produces longer sequences which makes [Training](training_models.md) and [Inference](text_generation_inference.md) slower and because of the loss in context, it's harder for the model to learn long dependencies.
- **Granularity vs Semantic Meaning**: The more granular a word is split, the less meaning it have.
- **Language Dependency vs Universality**: Language specific tokenizers have better performance but aren't reusable. Universal tokenizers are reusable but aren't efficient.
- **Training Cost vs Inference Cost**:
	- Large vocabularies are expensive to train but produces faster inferences as the sequence generated are smaller.
	- Small vocabularies are cheaper to train but produces slower inferences as the generated sequences are longer.
## Examples
![[sub_word_tokenization_example.png]]
## References
### Connects with
- [NLP](natural_language_processing.md)
- [LLM](large_language_models.md)
- [Inference](text_generation_inference.md)
- [Model Training](training_models.md)
- [Transformer Architecture](transformer_architecture.md)
## TL;DR
**Token** is a sub-word representation that a model can understand, it's generated by a **Tokenizer** that follow a specific **Tokenization** strategy like **Sub-Wording** to split words in an efficient manner to produce the smallest vocabulary while retaining the biggest meaning for each word.
## Flashcards
- Q: What is a token?
- A: A Token is a sub-word representation of a word which is mapped to a vocabulary ID the model can understand.
- Q: What are the three tokenization steps?
- A: 1. Split each word into sub-words. 2. Map each sub-word to a vocabulary ID. 3. Add special tokens required by the model.
- Q: What is a tokenizer?
- A: A tokenizer is a piece of software that follows a tokenization strategy to convert words into tokens.
- Q: What are the three most common tokenization strategies? And how do they approach tokenization?
- A: 1. Word based: uses complete words, produces a long vocabulary and has a high probability of generating unknown tokens but each token is rich in meaning.. 2. Character based: uses individual characters, has a smaller vocabulary and don't produces unknown tokens but there's no meaning in each token. 3. Sub-Word based: Splits only rare words into sub-words, produces a longer vocabulary but it's capable of keeping a good portion of the meaning of each word and their combination, don't produce a lot of unknown tokens.
- Q: What are the four most important trade-offs of tokenizers?
- A: 1. Vocabulary Size vs Sequence Length. 2. Granularity vs Semantic Meaning. 3. Language Dependency vs Universality. 4. Training Cost vs Inference cost.
