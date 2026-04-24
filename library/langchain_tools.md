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
## Access Context
Tools can access runtime information through the `ToolRuntime` parameter which provide access to the core components:

![[langchain_toolruntime_context.png]]

Alongside those it also provides access to:
- **Execution Info**: Identity and retry information for the current execution (thread ID, run ID, attempt number).
- **Server Info**: Server-specific metadata when running on LangGraph Server (assistant ID, graph ID, authenticated user).
- **Config**: [`RunnableConfig`](https://reference.langchain.com/python/langchain-core/runnables/config/RunnableConfig?_gl=1*1dgzd81*_gcl_au*MTU4OTcyNjY2LjE3NzM4Njc3MjE.*_ga*MTYxODAxODM1OS4xNzczODY3NzIx*_ga_47WX3HKKY2*czE3NzcwNDg5ODkkbzIyJGcxJHQxNzc3MDQ5NzU0JGo2MCRsMCRoMA..) for the execution.
- **Tool Call ID**: Unique identifier for the current tool invocation.
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
