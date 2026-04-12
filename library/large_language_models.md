---
id: 260405-large_language_models
tags:
  - AI
created: 2026-04-05
status: draft
---
# Large Language Models (LLM)
It's a subset of [NLP](natural_language_processing.md) with massive training data sets called **parameters**. Most modern LLMs use the [Decoder-Only Architecture](transformer_architecture.md#Transformer%20Blocks) with some of the largest models containing hundreds of billions of parameters.

Different from a NLP model, a LLM has reasoning capabilities and can perform several tasks with minimal training and can learn with examples.

Can also be fine-tuned to perform a wide range of language tasks or to increase its focus.

For example, a LLM can solve math problems and puzzles with several steps as it can plan and act accordingly.
# Understanding
It's a subset of the NLP study field with massive training data sets. LLMs can perform several language tasks with minimal training. It can also learn by examples or be fine-tuned to perform more specific tasks.
# Trade-offs
Along side all the [limitation of a NLP model](natural_language_processing.md#Trade-offs), a LLM also have the following limitations:
## Context Window
The amount of information the model can process at once. It's limited by the number of [Tokens](tokenization.md) and even with bigger sizes of context window, the model response usually gets a lot less accurate as the size of the context window increases.

>Hallucinations are lately been addressed via: 
>- [Retrieval-Augmented Generation (RAG)]()
>- [Tools]()
# Examples
There are several LLMs today that can perform a large variety of tasks, from text generation to image, video and audio too, here's a non-extensive list:
- `gpt-4o`
- `claude-sonet`
- `google-gemini`
- `Llama-`
# References
## Connects with
- [Natural Language Processing - NLP](natural_language_processing.md)
- [Transformer Architecture](transformer_architecture.md)
- [Retrieval-Augmented Generation (RAG)]()
- [Tools]()
- [Tokens](tokenization.md)
- [Text Inference](text_generation_inference.md)
- [Model Training](training_models.md)
# TL;DR
It's a specific area of the NLP study field and it's much more skilled than a single NLP model. It can perform several language tasks with minimal training and can learn by example or be fine-tuned. As any NLP model, it suffers from **bias**, **high training cost** and currently, still has a very short **context window** size. It also suffers from **hallucinations** but lately this is being addressed via **RAG** and **tool calling** strategies.
# Flashcards
- Q: What is a LLM?
- A: A Large Language Model is a NLP model with a massive data set called **parameters** having usually billions of parameters each.
- Q: What are the capabilities that a LLM has that a NLP model doesn't?
- A: A LLM can plan and perform a wide variety of language tasks. It's also capable of learning by examples which means they can perform tasks that were not in their dataset.
- Q: What are the limitations of a LLM?
- A: Bias, Cost, Context Window and Hallucinations.
- Q: How can we reduce the hallucination impact on an application?
- A: You can implement a Retrieval-Augmented Generation system and use tool calling capabilities to provide more context to the model so it has a more target content to generate responses on top of it.
- Q: What's the most common architecture for a LLM?
- A: Most LLMs use the Decoder-Only architecture.