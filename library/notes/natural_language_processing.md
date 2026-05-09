---
id: 260403-natural_language_processing
tags:
  - AI
created: 2026-04-03
status: draft
type: concept
---
## TL;DR
NLP is a study field that tries to make computers understand human language. It's trained on large amount of raw text and are specialized in one of many specific language tasks. A NLP model can suffer from Hallucination and Bias based on training data and usage.
# Natural Language Processing
It's a broad field focused in enabling computers to understand and generate human language. It tries to achieve this by analyzing the meaning of words together and individually.

NLP tasks are executed by a [Transformer Model](transformer_architecture.md) which are language models trained on  a large base of raw text to perform a **specific task**.
## NLP x LLM
NLP, is not the same as LLMs as the focus is to analyze text and generate completion and not to present reasoning capabilities. NLP are usually built with a specific task in mind.

For example, a NLP model wouldn't know how to solve a math problem that takes several steps as it would need reasoning to plan and execute.
## Types of language models
| Task                                   | Suggested Architecture                                                         | Examples      |
| -------------------------------------- | ------------------------------------------------------------------------------ | ------------- |
| Text classification (sentiment, topic) | [Encoder](transformer_architecture.md#Transformer%20Blocks)                    | BERT, RoBERTa |
| Text generation (creative writing)     | [Decoder](transformer_architecture.md#Transformer%20Blocks)                    | GPT, LLaMA    |
| Translation                            | [Encoder-Decoder](transformer_architecture.md#Transformer%20Blocks)            | T5, BART      |
| Summarization                          | [Encoder-Decoder](transformer_architecture.md#Transformer%20Blocks)            | BART, T5      |
| Named entity recognition               | [Encoder](transformer_architecture.md#Transformer%20Blocks)                    | BERT, RoBERTa |
| Question answering (extractive)        | [Encoder](transformer_architecture.md#Transformer%20Blocks)                    | BERT, RoBERTa |
| Question answering (generative)        | [Encoder-Decoder or Decoder](transformer_architecture.md#Transformer%20Blocks) | T5, GPT       |
| Conversational AI                      | [Decoder](transformer_architecture.md#Transformer%20Blocks)<br>                | GPT, LLaMA    |
When considering a language model to use, think about:
1. What kind of understanding does your task need? (Bidirectional or unidirectional)
2. Are you generating new text or analyzing existing text?
3. Do you need to transform one sequence into another?

The answers to these questions will guide you toward the right architecture.
# Understanding
Natural Language Processing is a study field that tries to make computers understand human language in order to perform simple language tasks like sentiment analysis and text completion.
NLP mustn't be mistaken with LLMs as they don't have reasoning capabilities.
# Trade-offs
## Bias
A model needs very large amounts of data in order to be trained, while doing this many researches scrap the internet and feed it all into the model training dataset, this brings all that's useful and all that's not.

Creating not biased models require very careful data cleaning. If you're using a pre-trained model, it can also be biased and fine-tuning it won't make the intrinsic bias disappear.
## Costs
Training a model from scratch require a lot of computational power, while there are already some providers that simplify this infrastructure setup, you still need to pay for those services.
## Hallucinations
Hallucination is when a NLP model generates fluent, confident output that is false or not grounded in reality or source data. The confident generation of this text makes it hard to detect hallucinations, specially if you're not informed about what the model is doing.

It's a natural problem present in all NLP models.
# Examples
A NLP model trained to generate text can generate coherent and relevant text based on prompt.
## References
### Connects with
- [Large Language Models - LLMs](large_language_models.md)
- [Transformer Models](transformer_architecture.md)
- [Tokenizers](tokenization.md)
- [Text Inference](text_generation_inference.md)
- [Model Training](training_models.md)
## Flashcards
- Q: What is a NLP model?
- A: It's a study field that tries to enable computers to understand human language and perform language tasks.
- Q: What's the difference between NLP and LLM?
- A: NLP models aren't trained on the same amount of data as LLMs, because of that LLMs demonstrate a much better reasoning capability and also can perform many language tasks without specific training, where as NLP models are usually trained with a much smaller data and with a specific task in mind.
- Q: What are the 3 limitations that NLP models have?
- A: Bias: Information without filter and treatment. Hallucinations: Fluent and confident generation of information that's not true or doesn't exist. Cost: It's still very expensive to train or host a NLP model.
- Q: What are the 3 questions you should make when you have to decide the language model type you're gonna use?
- A:  1. What kind of understanding does your task need?  2. Are you generating new text or analyzing existing text? 3. Do you need to transform one sequence into another?