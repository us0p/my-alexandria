---
id: 20260429-retrieval_augmented_generation
type: concept
status: draft
tags:
  - AI
created: 2026-04-29
---
## TL;DR
RAG (Retrieval-Augmented Generation) enhances LLMs by retrieving relevant external information and injecting it into the prompt.

Core idea:
LLM knowledge is static → retrieval makes it dynamic and grounded.
# Retrieval-Augmented Generation - RAG
Retrieval is the process of getting back information based on user request. The common workflow is to **load documents** and then **create semantic representations** of them, known as **knowledge base**, so that we can perform a **semantic search** based on user query to retrieve relevant documents or section of documents.

RAG is the process of adding this extra context to an [LLM](large_language_models.md) so that it can produce more accurate responses.
## Why RAG instead of [Fine-Tuning](training_models.md)
- LLMs have static knowledge (training cutoff)  
- Updating them requires retraining (expensive)  
- RAG allows dynamic knowledge injection without retraining
## Retrieval pipeline
A typical retrieval pipeline looks like this:
![[typical_retrieval_pipeline_workflow.png]]

| Building Blocks         | Responsibility                                                                                                 |
| ----------------------- | -------------------------------------------------------------------------------------------------------------- |
| **Document Loaders**    | Ingest data from external sources into a standard interface.                                                   |
| **Text Splitters**      | Break large documents into smaller chunks that it can fit the model context window.                            |
| [**Embedding Model**]() | Turns pieces of text into numerical vectors so that texts with similar meaning land close in the vector space. |
| **Vector Stores**       | Specialized database for vector(embedding) storage.                                                            |
| **Retriever**           | An interface that return documents given an unstructured query.                                                |
## RAG Architectures
| Architecture    | Description                                                                                                                          |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| **2-Step RAG**  | Retrieval always happens before generation.                                                                                          |
| **Agentic RAG** | A LLM powered [agent](llm_agent.md) decides when and how to retrieve during reasoning. Retrieval is enabled by the use of [Tools](). |
### Hybrid architecture
It combines characteristics of **2-Step RAG** an **Agentic RAG** by introducing **intermediate steps** such as **query preprocessing**, **retrieval validation**, and **post-generation checks**.

Typical components includes:
- **Query enhancement**: Modify the input query to improve retrieval quality. This can involve rewriting unclear queries, generating multiple variations, or expanding queries with additional context.
- **Retrieval validation**: Evaluate if retrieved documents are relevant. If not, the system may refine the query and retrieve again.
- **Answer validation**: Check generated answer for completeness and accuracy with source content. If not, the system can regenerate.

>The architecture usually supports multiple iterations between those components.

This architecture is suitable for:
- Ambiguous applications with underspecified queries.
- System that require validation or quality control steps.
## Understanding
It powers LLM answers by providing extra context by returning relevant information from a **knowledge base** using semantic search.
## When to Use
- When you want to expand the model's knowledge with your own data.
- When you need to gather information from many distinct documents.
## When NOT to Use
- When you don't need the model to have knowledge of your data to perform a task, like generating a `Task` for a todo application.
## Trade-offs
- **Freshness vs Latency**: retrieval improves accuracy but adds overhead
- **Recall vs Precision**: retrieving more documents increases coverage but may add noise
- **Simplicity vs Control**: basic RAG is simple; advanced pipelines add complexity
## Examples
- An application that uses a **knowledge base** and a LLM to answer to company's users about company policies.
- A coding assistant that needs to continuously fetch files from the repository and reason about it, possible doing it many times and adding different files based on conclusions of previous steps.
## References
### Connects with
- [LLM](large_language_models.md)
- [Agents](llm_agent.md)
- [Tools]()
## Contrasts with
- [Fine Tuning](training_models.md)
## Flashcards
Q: What is Retrieval-Augmented Generation (RAG)?  
A: It is a technique that enhances LLM responses by retrieving relevant external information and adding it to the prompt.
Q: What core problem does RAG solve?  
A: It addresses the limitation of static LLM knowledge by enabling dynamic access to external data.
Q: What is the main idea behind retrieval in RAG?  
A: To fetch relevant information from a knowledge base based on a user query using semantic search.
Q: What is a knowledge base in the context of RAG?  
A: A collection of documents transformed into semantic representations (embeddings) for efficient retrieval.
Q: Why are embeddings important in RAG?  
A: They convert text into vectors, allowing semantically similar content to be retrieved based on meaning rather than exact matches.
Q: Why use RAG instead of fine-tuning?  
A: Because RAG allows updating knowledge dynamically without retraining the model, which is expensive and time-consuming.
Q: What are the main steps in a retrieval pipeline?  
A: Loading documents, splitting text, generating embeddings, storing them in a vector store, and retrieving relevant content.
Q: What is the role of document loaders?  
A: To ingest data from external sources into a standardized format.
Q: Why are text splitters used?  
A: To break large documents into smaller chunks that fit within the model’s context window.
Q: What is a vector store?  
A: A specialized database used to store and query embeddings efficiently.
Q: What does a retriever do?  
A: It returns relevant documents based on an unstructured query.
Q: What is 2-Step RAG?  
A: A pipeline where retrieval always occurs before generation.
Q: What is Agentic RAG?  
A: A setup where an agent decides when and how to perform retrieval during reasoning.
Q: What is a hybrid RAG architecture?  
A: A combination of 2-Step and Agentic RAG with additional steps like query enhancement and validation.
Q: What is query enhancement in hybrid RAG?  
A: Modifying or expanding the query to improve retrieval quality.
Q: What is retrieval validation?  
A: Checking whether retrieved documents are relevant and refining the query if needed.
Q: What is answer validation?  
A: Verifying that the generated response is accurate and consistent with retrieved content.
Q: When is hybrid RAG most useful?  
A: In ambiguous scenarios or systems requiring validation and quality control.
Q: How does RAG improve LLM responses?  
A: By grounding them in relevant, external information retrieved at runtime.
Q: When should RAG be used?  
A: When you need to incorporate external or domain-specific knowledge into model responses.
Q: When should RAG NOT be used?  
A: When the task does not require external knowledge, such as simple or self-contained tasks.
Q: What is a key trade-off in RAG regarding freshness and latency?  
A: Retrieval improves accuracy and freshness but increases response time.
Q: What is the trade-off between recall and precision in RAG?  
A: Retrieving more documents increases coverage but may introduce irrelevant information.
Q: What is the trade-off between simplicity and control in RAG systems?  
A: Simple pipelines are easier to implement, while advanced ones offer more control but add complexity.
Q: What is a real-world example of RAG?  
A: A system that answers company policy questions using a knowledge base and an LLM.
Q: How can RAG be used in coding assistants?  
A: By retrieving relevant files from a repository and using them iteratively to improve reasoning and responses.