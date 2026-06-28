---
id: 20260427-langchain_state
type: practice
status: draft
tags:
  - AI
  - LangChain
created: 2026-04-27
---

# Title
## TL;DR
In summary, short-term memory is a feature that allow us to manage the context our model is going to be used. While it helps keeping the model context under control it add complexity to the application specially if you need to consider message ordering.
Most strategies are around message history and involve removing stale messages or summarizing groups of messages to compress information to fit model context window.
# Short-term Memory (State)
Allows the agent to remember information from previous interactions. It's fundamental for [LLMs](large_language_models.md) since it allows then to learn from feedback, and adapt to user preferences. For more complex tasks this is essential for efficiency and user satisfaction.

>Short-term memory allows you to remember previous interaction within a single **thread or conversation**.

Conversation history is the most common form of short-term memory. Long conversations pose a challenge to [LLMs](large_language_models.md). Even if your model supports the full context length, most LLMs still perform poorly over long context. They get "distracted" by stale or off-topic content, all while suffering from slower response times and higher costs. Because [context windows](text_generation_inference.md#The%20context%20length%20challenge) are limited, many applications can benefit from using technique to remove stale information.
## Usage details
- By storing these in the graph's state, the agent can access the full context for a given conversation while maintaining separation between different threads.
- State is persisted to a database (or memory) using a `checkpointer` so the thread can be resumed at any time.
- Short-term memory updates when the agent is invoked or a step (like a [tool](langchain_tools.md) call) is completed, and the state is read at the start of each step.
- In production you should use a `checkpointer` backed by a database. You can see the [list of `checkpointer` libraries](https://docs.langchain.com/oss/python/langgraph/persistence#checkpointer-libraries) in the docs.

```python
from langgraph.checkpoint.memory import InMemorySaver

agent = create_agent(
	"gpt-5.4",
	checkpointer=InMemorySaver(),
)

agent.invoke(
	{"messages": []},
	{"configurable": {"thread_id": "1"}} # used by the memory saver.
)
```
## Common state management strategies
### Trim messages
Since most LLMs context window limit is denominated in [tokens](tokenization.md). This strategy tries to truncate the history whenever it approaches that limit.

You can specify the number of tokens to keep from the list, as well as the strategy (e.g. keep the last `max_tokens`) to use for handling the boundary.
### Delete messages
Useful when you want to remove **specific messages** or **clear the entire message history**.

When deleting messages, make sure that the resulting message history is valid. Check the limitations of the LLM provider you're using.

For example:
- Some providers expect message history to start with a `user` message.
- Most providers require `assistant` messages with tool calls to be followed by corresponding `tool` result messages.
### Summarize messages
Since the previous methods can result in some information loss, some applications benefit from a approach of summarizing the message history using a chat model.

![[langchain_state_history_summarization_workflow.png]]
## Strategy Selection
- Trim → simplest, keeps recent context (default choice)
- Delete → precise control when removing specific messages
- Summarize → best when long-term context matters but must be compressed
## Understanding
Short-term memory or State, is a functionality that enables more selective context for individual **threads or conversations** within the Agent. By selective, understand that we can control the message history and extra parameters that are shared with a model's single conversation history.
We can apply strategies to remove information that becomes stale or irrelevant keeping the messages and metadata within the model's context length.
## When to Use
- Multi-turn conversations when context matters.
- When model response degrade due to long history.
- When cost/latency from large context becomes significant.
## When NOT to Use
- One shot calls or workflows that don't require many iterations to complete a task.
## Patterns
- State acts as a filtering layer between the conversation history and the model.
- Memory strategies trade completeness (more sparse data) for relevance (less dense data).
- Summarization converts linear history into semantic state. Trimming on the other hand, removes unnecessary information.
## Trade-offs
- Precision x Complexity: Feeding only the necessary to the model can make it more fast and precise but it adds the complexity of the strategies to manage the context.
## Failure Modes
- Removing messages without considering model constraints (message order) or meaning.
- Expecting it to store multi-conversation data or trying to access other conversations history from a different conversation.
## Implementation (Practical)
### Customizing agent memory
Adding additional fields
```python
from langchain.agents import create_agent, AgentState
from langgraph.checkpointer.memory import InMemorySaver

class CustomAgentState(AgentState):
	user_id: str
	preferences: dict

agent = create_agent(
	"gpt-5.4",
	state_schema=CustomAgentState,
	checkpointer=InMemorySave()
)

# Custom state can be passed in invoke
result = agent.invoke(
	{
		"messages": [{"role": "user", "content": "Hello"}],
		"user_id": "user_123",
		"preferences": {"theme": "dark"}
	},
	{"configurable": {"thread_id": "1"}}
)
```
### Trimming messages
To trim message history in an agent, use the `@before_model` [middleware]():
```python
from langchain.messages import RemoveMessage
from langgraph.graph.message import REMOVE_ALL_MESSAGES
from langgraph.checkpointer.memory import InMemorySaver
from langchain.agents import create_agent, AgentState
from langchain.agents.middleware import before_model
from langgraph.runtime import Runtime

@before_model
def trim_messages(state: AgentState, runtime: Runtime) -> dict[str, Any] | None:
	"""Kepp only the last few messages to fit context window."""
	messages = state["messages"]
	
	# logic to select only new messages
	
	return {
		"messages": [
			RemoveMessage(id=REMOVE_ALL_MESSAGES),
			*new_messages
		]
	}

agent = create_agent(
	"model_here",
	middleware=[trim_messages],
	checkpointer=InMemorySaver()
)
```
### Deleting messages
```python
from langchain.messages import RemoveMessage
from langgraph.graph.message import REMOVE_ALL_MESSAGES

# to remove specific messages
def delete_messages(state):
	messages = state["messages"]
	if len(messages) > 2:
		# remove the earliest two messages
		return {"messages": [RemoveMessage(id=m.id) for m in messages[:2]]}

# to remove all messages
def delete_all_messages(state):
	return {"messages": [RemoveMessage(id=REMOVE_ALL_MESSAGES)]}
```
### Summarizing messages
```python
from langchain.agents.middleware import SummarizationMiddleware

agent = create_agent(
	model="gpt-5.4",
	tools=[],
	middleware=[
		SummarizationMiddleware(
			model="gpt-5.4-mini",
			trigger=("tokens", 4000),
			keep=("messages", 20)
		)
	],
	checkpointer=InMemorySaver()
)
```

Check the [`SummarizationMiddleware` docs](https://docs.langchain.com/oss/python/langchain/middleware#summarization) for more configuration options.
## Real-world Usage
- A language training assistant in which you probably don't need a comprehensive memory on all the messages exchanged. Summarization would really help keeping track of what's important without affecting the conversation flow.
## Relationships
### Depends on
- [LLM](large_language_models.md)
- [Agents](llm_agent.md)
- [LangChain Messages](langchain_messages.md)
- [LangChain Tools](langchain_tools.md)
## Iterate on
- Patterns: need to add more depth and possible real world scenarios
- Usage cases: Add more depth to it.
- Trade-offs: Add more depth
## Flashcards
Q: What is short-term memory (State) in an agent system?  
A: It is a mechanism that allows the agent to remember and manage information within a single conversation or thread.
Q: Why is short-term memory important for LLMs?  
A: It enables learning from previous interactions, adapting to user preferences, and improving efficiency and user experience in multi-step tasks.
Q: What is the main challenge with long conversation history?  
A: Long histories can degrade model performance, increase latency and cost, and introduce irrelevant or stale information.
Q: Why must short-term memory be managed carefully?  
A: Because context windows are limited and poorly managed history can confuse the model and reduce response quality.
Q: How is short-term memory typically implemented?  
A: By storing conversation history and related data in the agent’s state, often persisted using a checkpointer.
Q: What is the role of a checkpointer?  
A: It persists the state so that a conversation thread can be resumed later.
Q: When is the agent’s state updated and read?  
A: It is updated after each step or invocation and read at the start of each step.
Q: What is the most common form of short-term memory?  
A: Conversation history (message history).
Q: What is the goal of state management strategies?  
A: To keep relevant information within the model’s context window while removing or compressing less useful data.
Q: What is the trimming strategy?  
A: It removes older messages to keep the total token count within the model’s context limit, usually preserving the most recent messages.
Q: What is the delete strategy?  
A: It removes specific messages or clears the entire history based on precise control needs.
Q: What must be considered when deleting messages?  
A: The resulting message history must remain valid according to the LLM provider’s constraints, such as message ordering.
Q: What is the summarization strategy?  
A: It compresses message history into a shorter form using a model, preserving key information while reducing size.
Q: When is summarization preferred over trimming or deleting?  
A: When long-term context is important but must fit within limited context windows.
Q: What are the main trade-offs between memory strategies?  
A: They balance completeness (more information) against relevance and efficiency (less but more useful information).
Q: How does state act within the agent architecture?  
A: It acts as a filtering layer between the full conversation history and the model.
Q: When should short-term memory be used?  
A: In multi-turn conversations, when context matters, or when long histories impact cost, latency, or performance.
Q: When should short-term memory NOT be used?  
A: In one-shot tasks or workflows that do not require iterative context.
Q: What is a key trade-off when using short-term memory?  
A: Increasing precision and efficiency comes at the cost of added complexity in managing context.
Q: What is a common failure mode in state management?  
A: Removing messages without considering ordering or meaning, which can break model expectations.
Q: Why can’t short-term memory be used across multiple conversations?  
A: Because it is scoped to a single thread and is not designed to share data between different conversations.
Q: How can custom fields be added to the agent’s memory?  
A: By extending the agent’s state schema with additional attributes.
Q: How can trimming be implemented in practice?  
A: By using middleware to filter and replace message history before the model is invoked.
Q: What is the purpose of summarization middleware?  
A: To automatically compress message history when certain thresholds (like token count) are reached.
Q: What is a real-world use case for short-term memory summarization?  
A: A language learning assistant that maintains key context without storing the entire conversation history.