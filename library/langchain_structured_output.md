---
id: 20260428-langchain_structured_output
type: practice
status: draft
tags:
  - AI
  - LangChain
created: 2026-04-28
---
## TL;DR
Structured output enforces a predefined schema on LLM responses, enabling reliable parsing, validation, and type-safe integration with applications.

Two strategies:
- `ProviderStrategy`: model-native, most reliable
- `ToolStrategy`: fallback via tool calling, works with any model
# LangChain Structured Output
Allows agents to return data in a specific, predictable format. You can get data in the form of [Pydantic models](), or [`dataclasses`]() that your application can use directly.

LangChain's [`create_agent`]() handles structured output automatically. The generated and validate structured is returned under the **`structured_response`** key as per user specified schema.
## Schema Strategies
You can define the strategy that the agent should use to generate structured output in the `response_format` field. Strategies are:
- `ToolStrategy[StructuredResponseT]`:  Use it when your model doesn't natively support structured output. It allows you to specify how to structure data as well as how to handle model mistakes with error handling.
- `ProviderStrategy[StructuredResponseT]]`: Use it when your model natively supports structured output. **It's the most reliable** as it's enforced by the model itself.
- `type[StructuredResponseT]`: Without any strategy, it selects the best strategy automatically for you based on your model support. It **defaults to `ProviderStrategy`** if supported.
## Error handling
Models can make mistakes when generating structured output via tool calling. For example, when a model generates a structured output that doesn't match the expected schema, the agent provides specific error feedback.
### Error handling strategies
You can customize how errors are handled using the **`ToolStrategy.handle_errors`** parameter.
#### Custom error message
The agent will **always** prompt the model to re-try with a fixed tool message.
```python
ToolStrategy(
	schema=ProductRating,
	handle_errors="please provide a valid rating between 1-5 ..."
)
```
#### Handle specific exceptions only
The agent will only retry if the exception raised if the specified type. All other cases, the exception will be propagated.
```python
ToolStrategy(
	schema=ProductRating,
	handle_errors=ValueError # only retry on ValueError, raise others.
)

# Can also be applied to multiple specific types
ToolStrategy(
	schema=ProductRating,
	handle_errors=(ValueError, TypeError)
)
```
#### Custom error handler
If you want to have a custom error handling logic.
```python
from langchain.agents import create_agent
from langchain.agents.structured_output import ToolStrategy

def custom_error_handler(error: Exception) -> str:
	return f"Error: {str(error)}"

agent = create_agent(
	# ...,
	response_format=ToolStrategy(
		schema=ContactInfo,
		handle_errors=custom_error_handler
	)
)
```
#### No error handling
```python
ToolStrategy(
	schema=ProductRating,
	handle_errors=False # All errors raised
)
```
## Understanding
Structured output constrains LLM responses into a predefined schema, enabling validation and direct integration into typed systems without additional parsing.
## When to Use
- Use it when you need to enforce constraints on model's response. For example a creation agent must return a `Task` object as output, which must have the `priority` field as one of `Literal["low", "medium", "high"]`.
## When NOT to Use
- Don't use it when your output doesn't require any kind of constraints. For example, an user questioning the model for an answer, probably doesn't need a strict format back out of it.
## Patterns
- It's translation layer between the model and external sources.
- Plays the same role as a DTO or POJO, by generating a standard object that can be understood by other layers from the model layer.
## Trade-offs
- Control x Complexity: Adding many structured types to an agent might cause it to become confused. If you have many possible structured outputs it might be a sign that your agent is going too much. Prefer to split it up into small, manageable tasks instead.
## Failure Modes
- Using it when you don't have any constraints on your output.
- Not handling error propagation properly. E.g. using only string error handling providing the same retry even if you have multiple structured options with different constraints.
- Using structured output for highly creative or ambiguous tasks.
- Overly strict schemas can cause frequent generation failures or retries.
## Implementation (Practical)
### Provider Strategy
```python
from pydantic import BaseModel, Field
from langchain.agents import create_agent
from langchain.agents.structured_output import ToolStrategy

class ContactInfo(BaseModel):
	"""Contact information for a person."""
	name: str = Field(description="...")
	email: str = Field(description="...")
	phone: str = Field(description="...")

# Wihout any specific strategy it falls back to ProviderStrategy automatically.
agent = create_agent(
	model="gpt-5.4",
	response_format=ContactInfo # same as response_format = ProviderStrategy(ContactInfo)
)

result = agent.invoke(
	#...
)

print(result["structured_response"])
```
### Tool Strategy
```python
from pydantic import BaseModel, Field
from typing import Literal
from langchain.agents import create_agent
from langchain.agents.structured_output import ToolStrategy

class ProductReview(BaseModel):
	"""Analysis of a product review."""
	rating: int | None = Field(description="...", ge=1, le=5)
	sentiment: Literal["positive", "negative"] = Field(description="...")
	key_points: list[str] = Field(description="...")

agent = create_agent(
	model="gpt-5.4",
	tools=[],
	response_format=ToolStrategy(ProductReview)
)

result = agent.invoke(
	#...
)

print(result["structured_response"])
```
## Real-world Usage
- Using this in a note taking app to return a standard `Note` type with predefined fields and constraints ensuring that data follows a strict interface.
## Relationships
### Depends on
- [LLM](large_language_models.md)
- [Agent](llm_agent.md)
- [LangChain Models](langchain_models.md)
## Iterate on
- Patterns: add depth
## Flashcards
Q: What is structured output in the context of LLMs?  
A: It is a mechanism that enforces a predefined schema on model responses, enabling reliable parsing, validation, and type-safe integration.
Q: What is the main purpose of structured output?  
A: To ensure model responses follow a predictable format that can be directly consumed by applications.
Q: What types can be used to define structured output schemas?  
A: Pydantic models or dataclasses.
Q: How does LangChain return structured output from an agent?  
A: It returns the validated result under the `structured_response` key based on the defined schema.
Q: What are the two main strategies for structured output?  
A: ProviderStrategy and ToolStrategy.
Q: What is ProviderStrategy?  
A: A model-native approach where the model enforces the schema, making it the most reliable strategy.
Q: What is ToolStrategy?  
A: A fallback approach that uses tool calling to enforce structured output, compatible with any model.
Q: What happens if no strategy is explicitly defined?  
A: The system automatically selects the best strategy, defaulting to ProviderStrategy if supported.
Q: When should ProviderStrategy be used?  
A: When the model natively supports structured output and reliability is critical.
Q: When should ToolStrategy be used?  
A: When the model does not support native structured output or when custom error handling is needed.
Q: Why is error handling important in ToolStrategy?  
A: Because models may generate outputs that do not match the expected schema, requiring retries or corrections.
Q: What happens when structured output validation fails?  
A: The agent provides feedback and can prompt the model to retry generating a valid response.
Q: How can error handling be customized in ToolStrategy?  
A: Using the `handle_errors` parameter with options like fixed messages, specific exceptions, custom handlers, or disabling retries.
Q: What does using a fixed error message in error handling do?  
A: It forces the model to retry with a consistent instruction when validation fails.
Q: What is the benefit of handling specific exception types?  
A: It allows retries only for certain errors while propagating others.
Q: What is a custom error handler in ToolStrategy?  
A: A function that defines custom logic for handling and formatting errors.
Q: What happens when error handling is disabled?  
A: All errors are raised without retry attempts.
Q: When should structured output be used?  
A: When responses must follow strict constraints or predefined schemas.
Q: When should structured output NOT be used?  
A: When outputs are free-form and do not require strict formatting or constraints.
Q: What design pattern does structured output resemble?  
A: It acts like a DTO or POJO, providing a standard object interface between the model and application layers.
Q: What is a key trade-off of using structured output?  
A: Increased control and reliability at the cost of added complexity.
Q: What issue can arise from defining too many structured outputs?  
A: It can confuse the model and indicate that the agent should be split into smaller, focused tasks.
Q: What is a common failure mode when using structured output?  
A: Applying it when no constraints are needed, leading to unnecessary complexity.
Q: Why is proper error handling critical in structured output?  
A: Poor error handling can lead to incorrect retries or unhandled exceptions, reducing reliability.
Q: What is a risk of overly strict schemas?  
A: They can cause frequent validation failures and repeated retries.
Q: Why is structured output not ideal for creative tasks?  
A: Because strict schemas limit flexibility and can conflict with ambiguous or open-ended outputs.
Q: What is a practical use case for structured output?  
A: Returning standardized objects, such as a note with predefined fields, ensuring consistent data structure.