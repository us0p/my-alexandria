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
The transformer architecture is the definition of the components necessary to build a language model. It uses two blocks known as **Encoder** and **Decoder** to convert input into a representation the model can understand and then uses the decoder to convert it back into words. Those blocks can be used separately and together providing specific capabilities based on the block functionality, for example **Encoder** blocks are useful for task that demands understanding of the input like text classification.
The **attention layer** is what gives the architecture the ability to grasp the context of a sentence based on the most important words in it.
# Trade-offs
- **Compute power**: Training and running models require large amounts of memory which are becoming very expensive with the increasing need for AI setups.
- **Data Hunger**: Can't produce a good result from a poor input without a massive dataset (can't generalize well with low data).
- **Context Window**: Fixed maximum length input.
- **Inference Latency (Auto-regressive Models like GPT)**: [Tokens]() are generated sequentially.
- **Poor Numerical & Logical Precision**: This architecture excels at approximating patterns which is not really good when dealing with exact logic, specially with float numbers.
# Example
BERT is a good example of a transformer model  which is used to analyse the sentiment or topic of a particular text.

It would be able to tell that the sentence "Brazil is a comedy" is actually BAD.
# References
## Connects with
- [NLP](natural_language_processing.md)
- [LLM](large_language_models.md)
- [Training Models](training_models.md)
- [Tokens]()
# TL;DR
It's an architecture that define how to train language models on large amounts of raw text. The architecture determines that a model is composed of two blocks: **Encoder** and  **Decoder** Model is optimized to generate outputs. Each of these blocks can be used independently.
The attention layer is what drives the model's attention to specific words in the sentence and allow him to understand the meaning of a sentence.
# Flashcards
- Q: What is the transformer architecture?
- A: It's an architecture that define how to train language models on large amounts of raw text.
- Q: What are the two language tasks that a language model usually has to perform?
- A: **Causal Language Modeling** and **Masked Language Modeling**.
- Q: What is Causal Language Modeling?
- A: It's a language task that predicts the next word by reading the previous n words.
- Q: What is Masked Language Modeling?
- A: It's a language task that predicts a masked word in a sentence.
- Q: What's the general strategy to increase a transformer model performance?
- A: Increase the model's size by increasing the amount of data the model is trained upon.
- Q: What are the two transformer blocks present in the transformer architecture?
- A: Encoder and Decoder.
- Q: What are the Encoder capabilities and what is this block good at?
- A: This block is responsible for receiving the input and converting it into something the model can understand. It excels at tasks that require input understanding like text classification.
- Q: What are the Decoder capabilities and what's this block good at?
- A: This block is responsible for converting the representation the Encoder generated back into text. It excels at tasks that generate text like question answering.
- Q: What is Attention Layer?
- A: It's a layer in the architecture that allows the model to understand the meaning of a sentence by focusing on specific words in it.
- Q: What's the difference between Architecture, Checkpoints and Model?
- A: Architecture is the definition of each layer that defines the model. Checkpoints are the weights that the architecture functions use to build the Model.
- Q: What are the trade-off of the Transformer Architecture?
- A: Compute Power, Data Hunger, Context Window, Inference Latency and Poor Numerical and Logical Precision.