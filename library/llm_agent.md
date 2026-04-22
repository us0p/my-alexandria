---
id: 20260417-llm_agent
tags:
  - AI
type: concept
created: 2026-04-17
status: draft
---
## TL;DR
An agent is an LLM that can take actions in the real world by the use of **Tools** and that can be tuned by managing it's **Memory**.
It adds flexibility and autonomy to your application but increases Cost, Complexity and makes it less reliable.
# Agent
It's a combination of a [LLM](large_language_models.md) with [Tools](), and [Memory]() that can reason about tasks, and iteratively work toward solutions.
## Core components
The core components of an agent are:
- **Model**: The reasoning engine.
- **Tools**: Give the agent ability to interact with the external environment allowing an **LLM** to take actions. 
- **Memory**: Allows the agent to keep more meaningful context besides the conversation history.
## Understanding
An agent is an LLM with capabilities of taking actions in the real world by the use of **tools**. We can also extend the model capabilities by providing more content to it's base data set and also managing it's context window to give more targeted responses.
## Trade-offs
- **Autonomy vs Control**: The more autonomy an agent has, the less control you have over it. **Guardrails** helps but it's impossible to cover every edge-case.
- **Flexibility vs Reliability**: If your agent has too much room for adaptability, this makes it less reliable. Two equal runs can produce distinct outcomes. For applications that require strict consistency this is a problem.
- **Capability vs Complexity**: As your agent grows more capable, your system becomes more complex, and this makes it hard to debug why an agent took a particular decision.
- **Generalization vs Safety**: An agent that can do a lot of tasks is more useful but it's also more error prone and easier to attack. More constrained agents are safer but less useful.
## Examples
An agent that plans and book all reservations for a trip you want to make. It's agentic because it **Plans** (creates the trip itinerary), it **Acts** (buy the tickets), **Observe Results** (monitor reservation steps), and **Iterates** (if anything fails or need further action, it goes back in the previous steps).
## References
### Connects with
- [LLM](large_language_models.md)
- [Tools]()
- [Memory]()
## Questions
- What is a **Guardrail** in a agentic application?
- When should you opt to use an Agentic application rather a normal one?
- What is the difference between an Agent and a RAG pipeline? When should you use one rather the other?
## Flashcards
- Q: What is an Agent?
- A: An Agent is a LLM armed with Tools that allows it to take actions.
- Q: What are the trade-offs of an agentic application?
- A: It increases the flexibility and autonomy of your application but also increases the complexity and reduces the reliability, control and safety of it.
