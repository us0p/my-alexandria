---
id: 260405-transformer_architecture
tags:
  - AI
created: 2026-04-05
status: draft
---
# Transformer Architecture
Introduced in June of 2017, it's an architecture that define how to train language models on large amounts of raw text in a [Self-Supervised](training_models.md#Self-Supervised%20Learning) fashion that can perform specific language tasks.

A language task can be:
- **Causal Language Modeling**: Predict the next word by reading the previous **n** words.
- **Masked Language Modeling**: The model predicts a masked word in a sentence e.g.: `I've coressed the ___`.

The general strategy to achieve better performance and accuracy is to increase model's size and the amount of data the model is pre-trained on.

The architecture determines that a model is composed of two blocks:
- **Encoder (left)**: Receives an input and builds a representation of it. Model is optimized to gather understanding of the input.
- **Decoder (right)**: Uses the generated representation along with other inputs to generate a target sequence. Model is optimized to generate outputs.

Each of these blocks can be used independently:
- **Encoder**: Good for understanding the input. Excels on tasks like text classification, named entity recognition, and question answering.
- **Decoder**: Good for completing sentences. Excels on tasks like text generation.
- **Encoder-Decoder**: Also known as **Sequence-to-Sequence** models, are good for generative tasks that require an input. Excel at translation, summary and question answering.

![[encoder_decoder_model_workflow.png]]
# Understanding
# Trade-offs
# Example
# References
# TL;DR
# Flashcards