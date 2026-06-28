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
Tools can access runtime information through the `ToolRuntime` parameter which is the interface between a tool and the agent's execution environment.

>The parameter `runtime: ToolRuntime` is automatically injected and hidden from the LLM - it won't appear in the tool's schema.
### Core runtime components

| Component                                                                              | Description                                                                                                   | Use case                                            |
| -------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | --------------------------------------------------- |
| [State](https://docs.langchain.com/oss/python/langchain/tools#short-term-memory-state) | Short-term memory - mutable data that exists for the current conversation (messages, counters, custom fields) | Access conversation history, track tool call counts |
| [Context](https://docs.langchain.com/oss/python/langchain/tools#context)               | Immutable configuration passed at invocation time (user IDs, session info)                                    | Personalize responses based on user identity        |
| [Store](https://docs.langchain.com/oss/python/langchain/tools#long-term-memory-store)  | Long-term memory - persistent data that survives across conversations                                         | Save user preferences, maintain knowledge base      |
### Specific use components
| [Stream Writer](https://docs.langchain.com/oss/python/langchain/tools#stream-writer)   | Emit real-time updates during tool execution                                                                                                                                                                                                                                            | Show progress for long-running operations                   |
| -------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------- |
| [Execution Info](https://docs.langchain.com/oss/python/langchain/tools#execution-info) | Identity and retry information for the current execution (thread ID, run ID, attempt number)                                                                                                                                                                                            | Access thread/run IDs, adjust behavior based on retry state |
| [Server Info](https://docs.langchain.com/oss/python/langchain/tools#server-info)       | Server-specific metadata when running on LangGraph Server (assistant ID, graph ID, authenticated user)                                                                                                                                                                                  | Access assistant ID, graph ID, or authenticated user info   |
| **Config**                                                                             | [`RunnableConfig`](https://reference.langchain.com/python/langchain-core/runnables/config/RunnableConfig?_gl=1*1dgzd81*_gcl_au*MTU4OTcyNjY2LjE3NzM4Njc3MjE.*_ga*MTYxODAxODM1OS4xNzczODY3NzIx*_ga_47WX3HKKY2*czE3NzcwNDg5ODkkbzIyJGcxJHQxNzc3MDQ5NzU0JGo2MCRsMCRoMA..) for the execution | Access callbacks, tags, and metadata                        |
| **Tool Call ID**                                                                       | Unique identifier for the current tool invocation                                                                                                                                                                                                                                       | Correlate tool calls for logs and model invocations         |
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
## When to Use
- Need external side effects (API, DB, filesystem)
- Need deterministic operations (math, retrieval)
- Need persistent memory interaction
## When NOT to Use
- Pure reasoning tasks (LLM alone is better)
- When latency matters (tools add overhead)
- When logic can be embedded in prompt
## Patterns
- Bridge interaction between LLMs and external world.
- Provide deterministic results into probabilistic reasoning.
- It composes the execution layer of an Agent.
## Trade-offs
- Adding to many tools can cause the model to become confused and actually degrade performance. Prefer fewer, well-scoped tools
- The same way, adding too much context to a tool might cause the model to become confused as well. Prefer only adding the necessary data and filter based on request.
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
Use [`Command`](#Command) to update the agent's state. Include a [`ToolMessage`](langchain_messages.md) in the update so the model can see the result of the tool call.

You can also use [`Command`](#Command) in other tools to hint the model on what to do next.
```python
from langchain.tools import tool, ToolRuntime
from langchain.messages import ToolMessage
from langchain.agents import AgentState
from langgraph.types import Command
from pydantic import BaseModel


class CustomState(AgentState):
    user_name: str

class CustomContext(BaseModel):
    user_id: str

# this tool performs the change in the state
@tool
def update_user_info(
    runtime: ToolRuntime[CustomContext, CustomState],
) -> Command:
    """Look up and update user info."""
    user_id = runtime.context.user_id
    name = "John Smith" if user_id == "user_123" else "Unknown user"
    return Command(update={
        "user_name": name,
        # update the message history
        "messages": [
            ToolMessage(
                "Successfully looked up user information",
                tool_call_id=runtime.tool_call_id
            )
        ]
    })

# You can force the execution of other tools using the Command return
@tool
def greet(
    runtime: ToolRuntime[CustomContext, CustomState]
) -> str | Command:
    """Use this to greet the user once you found their info."""
    user_name = runtime.state.get("user_name", None)
    if user_name is None:
       return Command(update={
            "messages": [
                ToolMessage(
                    "Please call the 'update_user_info' tool it will get and update the user's name.",
                    tool_call_id=runtime.tool_call_id
                )
            ]
        })
    return f"Hello {user_name}!"
```
>When tools update state variables, consider defining a [reducer](https://docs.langchain.com/oss/python/langgraph/graph-api#reducers) for those fields. Since LLMs can call multiple tools in parallel.
## Real-world Usage
- Use tool for DB lookup instead of embedding all data in prompt.
- Use tool for arithmetic calculation instead of relying on LLM arithmetic.
## Relationships
### Depends on
- [LLM](large_language_models.md)
- [Python type hints]()
- [LangChain Messages](langchain_messages.md)
- [Pydantic]()
- [`docstrings`]()
### Enables
- [Agents](llm_agent.md)
## Iterate on
- Patterns: Add depth
- Relationships: Frameworks and patterns where this enable or is used into.
## Flashcards
Q: What is a LangChain tool?  
A: A LangChain tool is a function with a well-defined signature that an LLM can choose to call to perform actions in the external world.
Q: How does an LLM decide when and how to use a tool?  
A: The LLM uses the tool’s signature, including its name, description, input types, and return type, to determine when to call it, what parameters to pass, and what to expect as output.
Q: Why are type hints required in tool definitions?  
A: Type hints define the tool’s input schema, enabling the model to correctly understand and supply parameters.
Q: What role does the function docstring play in a tool?  
A: The docstring becomes the tool’s description, helping the model understand when and why to use it.
Q: What is the purpose of the ToolRuntime parameter?  
A: ToolRuntime provides access to the execution environment, allowing tools to interact with state, context, memory, and metadata.
Q: What are the core components accessible via ToolRuntime?  
A: State (short-term mutable memory), Context (immutable configuration), and Store (long-term persistent memory).
Q: What is the difference between State and Store in ToolRuntime?  
A: State is short-term and tied to the current conversation, while Store is long-term and persists across conversations.
Q: When should a tool return a string?  
A: When the output is meant to be human-readable text for the model to interpret directly.
Q: When should a tool return an object?  
A: When structured data is needed for the model to reason over specific fields.
Q: When should a tool return a Command?  
A: When the tool needs to update the agent’s state in addition to returning information.
Q: What happens when a tool returns a Command?  
A: The agent’s state is updated, and the updated state becomes available to subsequent steps in the execution.
Q: Why might a ToolMessage be included in a Command return?  
A: To explicitly inform the model that the tool execution succeeded and provide feedback.
Q: When is it appropriate to use tools?  
A: When external side effects, deterministic operations, or persistent memory interactions are required.
Q: When should tools NOT be used?  
A: For pure reasoning tasks, when latency is critical, or when logic can be handled directly in the prompt.
Q: What is a key design principle when defining tools?  
A: Prefer fewer, well-scoped tools to avoid confusing the model and degrading performance.
Q: What is a common trade-off when using many tools?  
A: Too many tools can increase complexity and reduce model performance due to confusion.
Q: What is a common failure mode related to parallel tool execution?  
A: Not using reducers to properly manage state updates from multiple tools running in parallel.
Q: Why should unnecessary metadata be avoided in tools?  
A: Excess metadata increases complexity without adding meaningful value and can confuse the model.
Q: How can tool properties be customized?  
A: By specifying custom names and descriptions in the tool decorator.
Q: How can complex input schemas be defined for tools?  
A: Using Pydantic models or JSON schemas to structure and validate inputs.
Q: What is the main purpose of tools in an agent system?  
A: To bridge the LLM with the external world, enabling actions and providing deterministic data for reasoning.
Q: How do tools complement LLM capabilities?  
A: They provide deterministic execution and external data access, enhancing the probabilistic reasoning of LLMs.
Q: What is a practical use case for tools instead of prompts?  
A: Using a tool for database queries or arithmetic instead of embedding data or relying on the LLM’s internal reasoning.
Q: What is the relationship between tools and agent execution?  
A: Tools compose the execution layer of an agent, enabling it to perform real-world actions beyond text generation.