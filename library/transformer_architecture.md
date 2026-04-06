---
id: 260405-transformer_architecture
tags:
  - AI
created: 2026-04-05
status: draft
---
# Transformer Architecture
Introduced in June of 2017 and originally designed for translations, it's an architecture that define how to train language models on large amounts of raw text in a [Self-Supervised](training_models.md#Self-Supervised%20Learning) fashion that can perform specific language tasks.

A language task can be:
- **Causal Language Modeling**: Predict the next word by reading the previous **n** words.
- **Masked Language Modeling**: The model predicts a masked word in a sentence e.g.: `I've coressed the ___`.

Most tasks follow a similar pattern:
- input data is processed through a model
- output is interpreted for a specific task

Differences lie on how the data is prepared, what architecture variant is used and how output is processed.

The general strategy to achieve better performance and accuracy is to increase model's size and the amount of data the model is pre-trained on.
## Transformer blocks
The architecture determines that a model is composed of two blocks:
- **Encoder (left)**: Receives an input and builds a representation of it. Model is optimized to gather understanding of the input.
- **Decoder (right)**: Uses the generated representation along with other inputs to generate a target sequence. Model is optimized to generate outputs.

Each of these blocks can be used independently:
- **Encoder**: Good for understanding the input. Excels on bidirectional context tasks like text classification, named entity recognition, and question answering.
- **Decoder**: Good for completing sentences. Excels on tasks like text generation.
- **Encoder-Decoder**: Also known as **Sequence-to-Sequence** models, are good for generative tasks that require an input. Excel at translation, summary and question answering.

![[encoder_decoder_model_workflow.png]]
## Attention Layer
Is a layer in the transformer architecture model that drives the model's attention to specific words in the sentence.

In any language, the meaning of the word is affected deeply by the context.

For example, a translation from English to French. Given an input: "You like this course", a translation model will need to also attend to the adjacent word "You" to get the proper translation for the work "like", because in French the verb "like" is conjugated differently depending on the subject.
## Architecture x Checkpoints
- **Architecture**: Skeleton of the model, the definition of each layer and each operation that happens within the model.
- **Checkpoints**: Weights that will be loaded in a given architecture.
- **Model**: Umbrella term, can mean both (architecture or checkpoint).

>An **Architecture** is a succession of mathematical functions to build a **Model** and its **Weights** are the function parameters.
# Understanding
# Trade-offs
# Example
# References
# TL;DR
# Flashcards