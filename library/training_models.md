---
id: 260405-training_models
tags:
  - AI
created: 2026-04-05
status: draft
---
# Training Models
![[model_training_workflow.png]]

## Pre-training
Act of training a model from scratch. It's done on large amounts of raw data and can take up to several weeks.
## Self-Supervised Learning
It's a type of learning where the objective is automatically computed from inputs to the model. Humans aren't needed to label the data. In this phase, the model develops a statistical understanding of the language it has been trained on, but it doesn't know how to perform specifics tasks just yet.
## Transfer Learning
This kind of learning happens when we initialize a model with the **weights** of a pre-trained model. The new model will transfer the knowledge of the first model.
### Fine-Tuning
It's a strategy used to apply **Transfer Learning** in which humans labels a specific data and feed it into a **pre-trained model** so that the model can update some or all of its **weights**  to improve its capabilities in the task-specific data you provided.

Fine-tuning a model require a lot less data than **pre-training** it since the model already have a lot of statistical understanding of the language.
# Understanding
# Trade-offs
# Example
Leverage a **pre-trained** model on the English language and then **fine-tune** it on rap-composing resulting in a composer model.
# References
# TL;DR
# Flashcards