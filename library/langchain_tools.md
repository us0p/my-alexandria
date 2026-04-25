---
id: 20260424-langchain_tools
type: practice
status: draft
tags:
  - AI
  - LangChain
created: 2026-04-24
---
## TL;DR
LangChain Tools are function that are made available to the LLM which uses the function signature to determine when, and how to use it, allowing it to take actions in the real world.

A tool can have access to the environment metadata by accessing the `runtime: ToolRuntime` object. This object can be safely mutated and it's not shared with the model.
# LangChain Tools
Tools allows [LLMs](large_language_models.md) to take actions in the world by allowing them to interact with resources like databases, execute code, etc.

A tool is nothing more than a function with a well defined signature so that the model can decide whether to call it or not, which parameters supply to the function, and which should it expect as a return value.
## Tool definition
- The function [`docstring`]() becomes the tool's description that helps the model understand when to use it.
- [Type Hints]() are **required** as they define the tool's input schema.
```python
from langchain.tools import tool

@tool
def search_database(query: str, limit: int = 10) -> str:
	"""Search the customer database for records matching the query.
	
	Args:
		query: Search items to look for
		limit: Maximum number of results to return
	"""
	return f"Found {limit} results for '{query}'"
```
## Prebuilt tools
LangChain provides a collection of prebuilt tools and toolkits for common tasks. Check the [documentation](https://docs.langchain.com/oss/python/integrations/tools) for more.
## Access Context
Tools can access runtime information through the `ToolRuntime` parameter which provide access to the core components:

![[langchain_toolruntime_context.png]]

Alongside those it also provides access to:
- **Execution Info**: Identity and retry information for the current execution (thread ID, run ID, attempt number).
- **Server Info**: Server-specific metadata when running on LangGraph Server (assistant ID, graph ID, authenticated user).
- **Config**: [`RunnableConfig`](https://reference.langchain.com/python/langchain-core/runnables/config/RunnableConfig?_gl=1*1dgzd81*_gcl_au*MTU4OTcyNjY2LjE3NzM4Njc3MjE.*_ga*MTYxODAxODM1OS4xNzczODY3NzIx*_ga_47WX3HKKY2*czE3NzcwNDg5ODkkbzIyJGcxJHQxNzc3MDQ5NzU0JGo2MCRsMCRoMA..) for the execution.
- **Tool Call ID**: Unique identifier for the current tool invocation.
### Short-Term memory
A state that exists for the duration of a conversation. It includes the message history and any custom field.

```python
from langchain.tools import tool, ToolRuntime
from langchain.messages import HumanMessage

@tool
def get_last_user_message(runtime: ToolRuntime) -> str:
    """Get the most recent message from the user."""
    messages = runtime.state["messages"]

    # Find the last human message
    for message in reversed(messages):
        if isinstance(message, HumanMessage):
            return message.content

    return "No user messages found"

# Access custom state fields
@tool
def get_user_preference(
    pref_name: str,
    runtime: ToolRuntime
) -> str:
    """Get a user preference value."""
    preferences = runtime.state.get("user_preferences", {})
    return preferences.get(pref_name, "Not set")
```

>The parameter `runtime: ToolRuntime` is automatically injected and hidden from the LLM - it won't appear in the tool's schema.
### Context
Provides immutable configuration data that is passed at invocation time.

```python
from dataclasses import dataclass
from langchain_openai import ChatOpenAI
from langchain.agents import create_agent
from langchain.tools import tool, ToolRuntime


USER_DATABASE = {
    "user123": {
        "name": "Alice Johnson",
        "account_type": "Premium",
        "balance": 5000,
        "email": "alice@example.com"
    }
}

@dataclass
class UserContext:
    user_id: str

@tool
def get_account_info(runtime: ToolRuntime[UserContext]) -> str:
    """Get the current user's account information."""
    user_id = runtime.context.user_id

    if user_id in USER_DATABASE:
        user = USER_DATABASE[user_id]
        return f"Account holder: {user['name']}\nType: {user['account_type']}\nBalance: ${user['balance']}"
    return "User not found"

model = ChatOpenAI(model="gpt-5.4")
agent = create_agent(
    model,
    tools=[get_account_info],
    context_schema=UserContext,
    system_prompt="You are a financial assistant."
)

result = agent.invoke(
    {"messages": [{"role": "user", "content": "What's my current balance?"}]},
    context=UserContext(user_id="user123")
)
```
### Long-Term memory (Store)
`BaseStore` provides persistent storage that survives across conversations. Data remains available across sessions.

The store uses a namespace/key pattern to organize data.

>For production deployments, use a persistent store implementation like `PostgreStore`.

```python
from typing import Any
from langgraph.store.memory import InMemoryStore
from langchain.agents import create_agent
from langchain.tools import tool, ToolRuntime
from langchain_openai import ChatOpenAI

# Access memory
@tool
def get_user_info(user_id: str, runtime: ToolRuntime) -> str:
    """Look up user info."""
    store = runtime.store
    user_info = store.get(("users",), user_id)
    return str(user_info.value) if user_info else "Unknown user"

# Update memory
@tool
def save_user_info(user_id: str, user_info: dict[str, Any], runtime: ToolRuntime) -> str:
    """Save user info."""
    store = runtime.store
    store.put(("users",), user_id, user_info)
    return "Successfully saved user info."

model = ChatOpenAI(model="gpt-5.4")

store = InMemoryStore()
agent = create_agent(
    model,
    tools=[get_user_info, save_user_info],
    store=store
)

# First session: save user info
agent.invoke({
    "messages": [{"role": "user", "content": "Save the following user: userid: abc123, name: Foo, age: 25, email: foo@langchain.dev"}]
})

# Second session: get user info
agent.invoke({
    "messages": [{"role": "user", "content": "Get user info for user with id 'abc123'"}]
})
# Here is the user info for user with ID "abc123":
# - Name: Foo
# - Age: 25
# - Email: foo@langchain.dev
```
### Stream writer
Stream real-time updates from tools during execution.

```python
from langchain.tools import tool, ToolRuntime

@tool
def get_weather(city: str, runtime: ToolRuntime) -> str:
    """Get weather for a given city."""
    writer = runtime.stream_writer

    # Stream custom updates as the tool executes
    writer(f"Looking up data for city: {city}")
    writer(f"Acquired data for city: {city}")

    return f"It's always sunny in {city}!"
```

>If you use `runtime.stream_writter` inside your tool , the tool must be invoked within a LangGraph execution context.
## Tool return values
### String
Use it when the tool should provide plain text for the model, e.g. naturally human-readable text.

```python
from langchain.tools import tool

@tool
def get_weather(city: str) -> str:
	"""Get weather for a city."""
	return f"It is currently sunny in {city}."
```

Behavior:
- Return value is converted to [`ToolMessage`](langchain_messages.md).
- No agent state fields are changed unless the model or another tool does so later.
### Object
Use it when downstream reasoning benefits from explicit fields instead of free-form text.

```python
from langchain.tools import tool

@tool
def get_weather_data(city: str) -> dict:
	"""Get structured weather data for a city."""
	return {
		"city": city,
		"temperature_c": 22,
		"conditions": "sunny"
	}
```

Behavior:
- Returned object is serialized and sent back as tool output.
- The model can read specific fields and reason over them.
- It doesn't update the graph state.
### Command
Use this when the tool is not just returning data, but also mutating agent state.
```python
from langchain.messages import ToolMessage
from langchain.tools import ToolRuntime, tool
from langgraph.types import Command

@tool
def set_language(language: str, runtime: ToolRuntime) -> Command:
	"""Set the preferred response language."""
	return Command(
		update={
			"preferred_language": language,
			"messages": [
				ToolMessage(
					content=f"Language set to {language}.",
					tool_call_id=runtime.tool_call_id,
				)
			]
		}
	)
```

You can return a `Command` with or without including a [`ToolMessage`](langchain_messages.md). If the model needs to see that the tool succeeded, include a [`ToolMessage`](langchain_messages.md) in the update as above.

Behavior:
- The command updates state using `update`.
- Updated state is available to subsequent steps in the same run.
- Use reducers for fields that may be updated by parallel tool calls.
## Understanding
A LangChain tool is just a function that has its signature provided to the model so that it can request its execution based on the information the signature provides like type, name, description and return type.
A tool can use the `runtime: ToolRuntime` parameter to get some environment metadata like, ***State**, **Store**, **Context**, **Stream Writer**, **Execution Info**, **Server Info**, and **Config**. Which can be used to give model more information about user preferences for a particular chat, environment configuration, etc.
## Trade-offs
- Adding to many tools can cause the model to become confused and actually degrade performance.
- The same way, adding too much context to a tool might cause the model to become confused as well.
## Failure Modes
- Not using reducers to specify how state accessed in parallel should be managed.
- Adding too much metadata where it's not needed, increases complexity and doesn't bring real value.
## Implementation (Practical)
### Customize tool properties
```python
@tool("web_search") # Custom name
def search(query: str) -> str:
	"""Search the web for information."""
	return f"Results for: {query}"

print(search.name) # web_search

@tool("calculator", description="Performs arithmetic calculations. Use this for any math problems")
def calc(expressions: str) -> str:
	"""Evaluate mathematical expressions."""
	return str(eval(expression))
```
### Advanced schema definition
You can define complex inputs with [Pydantic]() or [JSON schemas](https://json-schema.org/understanding-json-schema/about).
```python
from pydantic import BaseModel, Field
from typing import Literal

weather_schema = {
	"type": "object",
	"properties": {
		"location": {"type": "string"},
		"units": {"type": "string"},
		"include_forecast": {"type": "boolean"}
	},
	"required": ["location", "units", "include_forecast"]
}

class WeatherInput(BaseModel):
	"""Input for weather queries."""
	location: str = Field(description="City name or coordinates")
	units: Literal["celsius", "fahrenheit"] = Field(
		default="celsius",
		description="Temperature unit preference"
	)
	include_forecast: bool = Field(
		default=False,
		description="Include 5-day forecast"
	)

@tool(args_schema=weather_schema)
# or
@tool(args_schema=WeatherInput)
def get_weather(
	location: str,
	units: str = "celsius",
	include_forecast: bool = False
) -> str:
	"""Get current weather and optional forecast."""
	tempo = 22 if units == "celsius" else 72
	result = f"Current weather in {location}: {temp} degrees {units[0].upper()}"
	if include_forecast:
		result += "\nNext 5 days: Sunny"
	return result
```
### Update state
Use `Command` to update the agent's state. Include a [`ToolMessage`](langchain_messages.md) in the update so the model can see the result of the tool call:
```python
from langchain.agents import AgentState
from langchain.messages import ToolMessage
from langchain.tools import ToolRuntime, tool
from langgraph.types import Command


class CustomState(AgentState):
    user_name: str


@tool
def set_user_name(new_name: str, runtime: ToolRuntime[None, CustomState]) -> Command:
    """Set the user's name in the conversation state."""
    return Command(
        update={
            "user_name": new_name,
            "messages": [
                ToolMessage(
                    content=f"User name set to {new_name}.",
                    tool_call_id=runtime.tool_call_id,
                )
            ],
        }
    )
```
>When tools update state variables, consider defining a [reducer](https://docs.langchain.com/oss/python/langgraph/graph-api#reducers) for those fields. Since LLMs can call multiple tools in parallel.
### Execution Info
Access thread ID, run ID, and retry state from within a tool:
```python
from langchain.tools import tool, ToolRuntime

@tool
def log_execution_context(runtime: ToolRuntime) -> str:
    """Log execution identity information."""
    info = runtime.execution_info
    print(f"Thread: {info.thread_id}, Run: {info.run_id}")
    print(f"Attempt: {info.node_attempt}")
    return "done"
```
### Server Info
**When, and only when** your tool runs on LangGraph Server:
```python
from langchain.tools import tool, ToolRuntime

@tool
def get_assistant_scoped_data(runtime: ToolRuntime) -> str:
    """Fetch data scoped to the current assistant."""
    server = runtime.server_info
    if server is not None:
        print(f"Assistant: {server.assistant_id}, Graph: {server.graph_id}")
        if server.user is not None:
            print(f"User: {server.user.identity}")
    return "done"
```
## Real-world Usage
- A building block of any agentic system.
## Relationships
### Depends on
- [LLM](large_language_models.md)
- [Python type hints]()
- [LangChain Messages](langchain_messages.md)
- [Pydantic]()
### Enables
- [Agents](llm_agent.md)
## Flashcards
- Q: What are LangChain Tools?
- A: Functions made available to an LLM, allowing it to decide when and how to call them to take actions in the real world.
- Q: How does an LLM decide whether to use a tool?
- A: By analyzing the tool’s function signature, including its name, description, input types, and return type.
- Q: What is the purpose of the `runtime: ToolRuntime` object?
- A: To provide access to environment metadata that can be safely mutated and is not shared with the model.
- Q: What makes up a LangChain tool?
- A: A function with a well-defined signature, including a docstring and type hints.
- Q: What role does the docstring play in a tool definition?
- A: It becomes the tool’s description, helping the model understand when to use it.
- Q: Why are type hints required in tool definitions?
- A: They define the tool’s input schema for the model.
- Q: What are prebuilt tools in LangChain?
- A: Ready-to-use tools and toolkits provided by LangChain for common tasks.
- Q: What types of information can be accessed through `ToolRuntime`?
- A: State, store, context, stream writer, execution info, server info, config, and tool call ID.
- Q: What is short-term memory in the context of tools?
- A: A state that exists during a conversation, including message history and custom fields.
- Q: How can a tool access the last user message?
- A: By reading the message history from `runtime.state["messages"]`.
- Q: What is the purpose of the context in `ToolRuntime`?
- A: To provide immutable configuration data passed at invocation time.
- Q: How is user-specific data accessed in tools using context?
- A: Through structured context objects like dataclasses (e.g., `UserContext`).
- Q: What is long-term memory (store) in LangChain tools?
- A: Persistent storage that survives across conversations using a namespace/key pattern.
- Q: What is recommended for production use of long-term memory?
- A: A persistent store implementation like `PostgreStore`.
- Q: How can tools read and write to long-term memory?
- A: Using methods like `store.get()` and `store.put()` via `runtime.store`.
- Q: What is the purpose of the stream writer in tools?
- A: To send real-time updates during tool execution.
- Q: When can `runtime.stream_writer` be used?
- A: Only when the tool is executed within a LangGraph execution context.
- Q: What are the three main types of tool return values?
- A: String, object (dict), and Command.
- Q: When should a tool return a string?
- A: When providing plain, human-readable text output.
- Q: What happens when a tool returns a string?
- A: It is converted into a ToolMessage without modifying agent state.
- Q: When should a tool return an object (dict)?
- A: When structured data is needed for downstream reasoning.
- Q: What is the behavior of object return values?
- A: They are serialized and allow the model to reason over specific fields without updating state.
- Q: When should a tool return a Command?
- A: When it needs to modify the agent’s state.
- Q: What does a Command return value do?
- A: It updates the agent’s state and can include ToolMessages.
- Q: Why might you include a ToolMessage in a Command update?
- A: So the model can see the result of the tool execution.
- Q: What is a key characteristic of LangChain tools?
- A: They expose their function signature to the model for decision-making.
- Q: What kinds of metadata can tools access via ToolRuntime?
- A: Execution info, server info, config, and tool call ID.
- Q: What are some trade-offs of using many tools?
- A: Too many tools or too much context can confuse the model and degrade performance.
- Q: What is a common failure mode related to tool state updates?
- A: Not using reducers to manage state updates when tools run in parallel.
- Q: What is another failure mode when designing tools?
- A: Adding unnecessary metadata that increases complexity without value.
- Q: How can you customize a tool’s name and description?
- A: By passing parameters to the `@tool` decorator.
- Q: How can you define complex input schemas for tools?
- A: Using Pydantic models or JSON Schema.
- Q: What is the benefit of using Pydantic for tool inputs?
- A: It provides structured and validated input definitions.
- Q: How can a tool update agent state in practice?
- A: By returning a Command object with updated fields and optional ToolMessages.
- Q: Why are reducers important when tools update shared state?
- A: Because tools can run in parallel, and reducers define how state conflicts are resolved.
- Q: What kind of execution metadata can be accessed inside a tool?
- A: Thread ID, run ID, and retry attempt number.
- Q: When is server info available in ToolRuntime?
- A: Only when running on LangGraph Server.
- Q: What is the role of LangChain tools in real-world applications?
- A: They are a core building block for agentic systems.
- Q: What do LangChain tools depend on?
- A: LLMs, Python type hints, LangChain messages, and optionally Pydantic.
- Q: What do LangChain tools enable?
- A: The creation of agents that can interact with external systems and take actions.