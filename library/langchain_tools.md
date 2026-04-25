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
Very short summary with only essential info.
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
## Understanding
- Explanation in your own words
- Focus on cause and effect
## When to Use
- Situations where this is useful
## When NOT to Use
- Situations where this is overkill or harmful
## Trade-offs
- Limitations
- Costs
- Complexity
## Failure Modes
- Common mistakes
- Misuses
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
## Tool return values
You can choose different return values for your tools:
- `string`: for human-readable results.
- `object`: for structured results the model should parse.
- `Command`: with optional message when you need to write to state.
## Real-world Usage
- Where this was applied
- Link to decisions or systems
## Relationships
### Depends on
- Link to concepts notes that are necessary for this practice

### Enables
- Link to concepts that can be achieved through this practice

### Used in
- Link to decision notes if any
## Questions
- Things still unclear
## Flashcards
- Q:
- A:
