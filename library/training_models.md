---
id: 260405-training_models
tags:
  - AI
created: 2026-04-05
status: draft
---
# Training Models
Language models are trained to predict the probability of a word given the surrounding context.
## Pre-training
Act of training a model from scratch. It's done on large amounts of raw data and can take up to several weeks.

![[model_training_workflow.png]]

There are two main approaches for training a language model:
- **Masked Language Modeling (MLM)**: Randomly masks some tokens in the input and trains the model to predict the original tokens based on the surrounding context. The model learns bidirectional context (looking at words both before and after the masked model).
- **Causal Language Modeling (CLM)**: Predicts the next token based on all previous tokens in the sequence.

>Other objectives exist, and fine-tuning typically uses different task-specific training methods.
### Self-Supervised Learning
It's a type of learning where the objective is automatically computed from inputs to the model. Humans aren't needed to label the data. In this phase, the model develops a statistical understanding of the language it has been trained on, but it doesn't know how to perform specifics tasks just yet.
## Transfer Learning
This kind of learning happens when we initialize a model with the **weights** of a pre-trained model. The new model will transfer the knowledge of the first model.
### Fine-Tuning
It's a strategy used to apply **Transfer Learning** in which humans labels a specific data and feed it into a **pre-trained model** so that the model can update some or all of its **weights**  to improve its capabilities in the task-specific data you provided.

Fine-tuning a model require a lot less data than **pre-training** it since the model already have a lot of statistical understanding of the language.
# Understanding
Training a model is done in two steps generally:
1. Pre-training: training a model form the ground
2. Transfer Learning: Set of strategies to improve the learning of the model towards a specific goal.
In the pre-training step the self-supervised learning is a human-less strategy where the model feeds himself.
An example of transfer learning is fine tuning where you have humans label some data and feed it to the model to deeper its understanding and capabilities in a given area.
# Trade-offs
1. **Accuracy x Compute Cost**: Higher accuracy requires larger models which needs more data and a longer training time, this increases compute cost and energy consumption.
2. **Performance x latency**: Bigger models are more performative in the sense they produce more accurate responses but they take more time to produce a response and are harder to deploy on small devices.
3. **Data size x Quality**: Using more data helps the model to generalize better but large datasets are often biased or contain duplicates and noises. Good quality data sets take time and money to be curated.
4. **Bias x Coverage**: Training on real world data improves coverage of language use but also introduces social, cultural and demographic biases.
5. **Speed x Quality**: Large batches can reduce training time but lead to worse generalization. Small batches take longer but provides a better generalization.
6. **Sustainability x Capability**: Training large models consumes massive energy, reducing energy leads to lower performance as we'll have smaller models.
# Example
Leverage a **pre-trained** model on the English language and then **fine-tune** it on rap-composing resulting in a composer model.
# References
## Connects with
- [NLP](natural_language_processing.md)
- [LLM](large_language_models.md)
- [Transformer Architecture](transformer_architecture.md)
# TL;DR
Training a model involves several different strategies from pre-training steps with self-supervised strategies that don't require human intervention to specialized training with transfer learning strategies like fine tuning a model on a specific task. Most important trade-offs are: Accuracy x Compute Cost, Performance x Latency, Data Size x Quality, Bias x Coverage, Speed x Quality, Sustainability x Capability.
# Flashcards
Q: Models are trained to do what in specific?
A: Models are trained to predict the probability of the next word given its surrounding context.
Q: What is the Pre-Training step and its responsibility?
A: The Pre-Training is the step in which you train a model from the ground, the goal is to generate a good generalist model that can be fine-tuned after.
Q: What are the two main approaches for training a model?
A: Masked Language Modeling and Causal Language Modeling.
Q: What is the strategy applied in MLM training?
A: The model's input has random tokens masked and the model is trained to predict the original token based on the surrounding words.
Q: What is the strategy applied in CLM?
A: It masks the last tokens in a sentence and the model must predict the next token based only in the previous tokens and context.
Q: What is self-supervised learning and when it's applied?
A: It's a form of training a model in which the model feeds himself and no human intervention is needed, it allows the model to build a statistical understanding of the language and it's usually applied in the pre-training phase.
Q: What is transfer learning? And when does it happen?
A: Transfer learning is a step in which we apply a strategy to improve the model's understanding of a subject using another already trained source, like using another model weights as base for your model or fine-tuning it with human labeled data.
Q: What is Fine-Tuning? When should it be done?
A: Fine-Tuning is a Transfer Learning strategy in which human labeled data is feed into the model. It's best executed when we want the model to have a mover profound understanding of a specific knowledge or when we want to create task-specific models.
Q: What are the 6 most important trade-offs of training a model?
A:   Accuracy x Compute Cost, Performance x latency, Data size x Quality,  Bias x Coverage, Speed x Quality and Sustainability x Capability.