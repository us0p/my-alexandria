---
id: 260405-large_language_models
tags:
  - AI
created: 2026-04-05
status: draft
type: concept
---
# TL;DR
It's a specific area of the NLP study field and it's much more skilled than a single NLP model. It can perform several language tasks with minimal training and can learn by example or be fine-tuned. As any NLP model, it suffers from **bias**, **high training cost** and currently, still has a very short **context window** size. It also suffers from **hallucinations** but lately this is being addressed via **RAG** and **tool calling** strategies.
# Large Language Models (LLM)
It's a subset of [NLP](natural_language_processing.md) with massive training data sets called **parameters**. Most modern LLMs use the [Decoder-Only Architecture](transformer_architecture.md#Transformer%20Blocks) with some of the largest models containing hundreds of billions of parameters.

Different from a NLP model, a LLM has reasoning capabilities and can perform several tasks with minimal training and can learn with examples.

Can also be fine-tuned to perform a wide range of language tasks or to increase its focus.

For example, a LLM can solve math problems and puzzles with several steps as it can plan and act accordingly.
## Agentic capabilities
In addition to text generation, many models support:
- **Tool Calling**: Call external tools and use results in their responses and reasoning.
- **Structured Output**: Model's response is constrained to follow a defined format.
- **Multimodality**: Process and return data other than text.
- **Reasoning**: Perform multi-step reasoning to arrive at a conclusion.

The **Reasoning** capability is what make LLMs suitable for [Agentic applications](llm_agent.md). They drive the agent's decision-making process, determining which tools to call, how to interpret results and when to p[[decoding_strategies]]rovide a final answer.
# Understanding
It's a subset of the NLP study field with massive training data sets. LLMs can perform several language tasks with minimal training. It can also learn by examples or be fine-tuned to perform more specific tasks.
# Trade-offs
Along side all the [limitation of a NLP model](natural_language_processing.md#Trade-offs), a LLM also have the following limitations:
## Context Window
The amount of information the model can process at once. It's limited by the number of [Tokens](tokenization.md) and even with bigger sizes of context window, the model response usually gets a lot less accurate as the size of the context window increases.
## Fine-Tuning
Process of further training a pre-trained language model using additional data. It can be helpful for adapting the model to a specific domain, task, or writing style, but it requires careful consideration of the fine-tuning data and the potential impact on the model's performance and biases.
## Pretraining
Initial process of training language models on a large unlabeled corpus of text. Autoregressive language models are pretrained to predict the next word, given the previous context of text in the document.
## Reinforcement Learning from Human Feedback (RLHF)
Technique used to train a pretrained language model to behave in ways that are consistent with human preferences. Human feedback consists of ranking a set of two or more example texts, and the reinforcement learning process encourages the model to prefer outputs that are similar to the higher-ranked ones.
## Temperature
Is a parameter that controls the randomness of a model's predictions during text generation. Higher temperatures lead to more creative and diverse outputs, allowing for multiple variations in phrasing. Lower temperatures result in more conservative and deterministic outputs that stick to the most probable phrasing and answers. Adjusting the temperature enables users to encourage a language model to explore rare, uncommon, or surprising word choices and sequences, rather than only selecting the most likely predictions.

Users may encounter non-determinism in APIs. Even with temperature set to 0, the results will not be fully deterministic and identical inputs may produce different outputs across API calls.
## Time to first token (TTFT)
performance metric that measures the time it takes for a language model to generate the first token of its output after receiving a prompt. It's an important indicator for interactive applications and real-time systems. Factors that can influence TTFT include model size, hardware capabilities, network conditions, and the complexity of the prompt.
## Tokens
Tokens are the smallest individual units of a language model, and can correspond to words, subwords, characters, or even bytes (in the case of Unicode). When a model is provided with text to evaluate, the text (consisting of a series of characters) is encoded into a series of tokens for the model to process. Larger tokens enable data efficiency during inference and pretraining (and are utilized when possible), while smaller tokens allow a model to handle uncommon or never-before-seen words. The choice of tokenization method can impact the model's performance, vocabulary size, and ability to handle out-of-vocabulary words.
## Chat-Bots
Unlike standalone LLMs that process a single input, chat bots are built to handle ongoing dialogues. They remember previous exchanges, which allows them to generate coherent, context-aware responses—perfect for customer support or multi-step problem-solving. For example, a call to OpenAI API for a `GPT-4o` inference is an AI response while an interaction with their UI interface is a chat bot (it has memory, a single API call doesn't).
### Choosing the Right Model: Chat-bots vs. Non-Chatbots
Your choice depends on your task:
- **Chatbots** are best when you need:
    - **Ongoing conversations:** Where follow-up questions and context retention are crucial.
    - **Complex problem-solving:** Where a dialogue helps refine and clarify answers.
    - **Customer support:** For issues that require multiple interactions and sustained context.
- **Standalone LLMs** are ideal for:
    - **Concise tasks:** Such as sentence completions, short answers, or text summarization.
    - **Quick results without context needs:** When you don't require conversation history.

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
- [Agents](llm_agent.md)
## Questions
- What's the difference between chat models and embedding models.
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
- Q: What are the agentic capabilities an LLM provides?
- A: Tool calling, structured output, multimodaility and reasoning.