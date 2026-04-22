---
id: 20260422-langchain_models
type: practice
status: draft
tags:
  - AI
  - LangChain
created: 2026-04-22
---
# LangChain Models
## TL;DR
The LangChain model definition allow us to configure the behavior of the LLM through a consistent interface across different providers.
We can configure text generation, tool availability, reasoning (when available), add structure to model response and inspect model chain of actions.
## Understanding
LangChain provides a standard interface to interact with many different model providers like OpenAI, Anthropic, HuggingFace, Ollama, etc.

This makes it easy to switch provider and enables the dynamic model selection, allowing you to select the best available model for each specific task.

Each model provider package must be installed separately as follows:
```bash
# pip install "langchain[provider_name]"
pip install "langchain[openai]"
```
## Model Profiles
LangChain chat models expose a dictionary of supported features and capabilities. Refer to the full set of fields in the [API Reference](https://reference.langchain.com/python/langchain-core/language_models/model_profile/ModelProfile?_gl=1*mr4zf3*_gcl_au*MTU4OTcyNjY2LjE3NzM4Njc3MjE.*_ga*MTYxODAxODM1OS4xNzczODY3NzIx*_ga_47WX3HKKY2*czE3NzY4ODk1ODUkbzE4JGcwJHQxNzc2ODg5NTg1JGo2MCRsMCRoMA..)
```python
model.profile
# {
#   "max_input_tokens": 400000,
#   "image_inputs": True,
#   "reasoning_output": True,
#   "tool_calling": True,
#   ...
# }
```

It allows applications to work around model capabilities dynamically:
- **Summarization Middleware** can trigger summarization based on a model's [context window size](text_generation_inference.md#The%20context%20length%20challenge).
- Model inputs can be gated based on supported **modalities** and **maximum input tokens**.
## Multimodal
Certain models can process and return non-textual data. You can pass non-text data with [Content Blocks](). **Content Blocks** can also be present in the [AIMessage]() if the response contains multimodal data.
## Reasoning
**If supported by the underlying model**, you can surface the reasoning process to better understand how the model arrived at its final answer.
```python
response = model.invoke("Why do parrots have colorful feathers?")
reasoning_steps = [b for b in response.content_blocks if b["type"] == "reasoning"]
print(" ".join(step["reasoning"] for step in reasoning_steps))
```

>Depending on the model, you can specify the level of effort it should put into reasoning. Please check the [reference](https://reference.langchain.com/python/integrations/?_gl=1*11qy6pc*_gcl_au*MTU4OTcyNjY2LjE3NzM4Njc3MjE.*_ga*MTYxODAxODM1OS4xNzczODY3NzIx*_ga_47WX3HKKY2*czE3NzY4ODk1ODUkbzE4JGcwJHQxNzc2ODg5NTg1JGo2MCRsMCRoMA..) for your respective chat model.
## Token Usage
Some providers return token usage information as part of the invocation response. When available this information will be included in the [AIMessage]() objects.

You can track aggregate token counts across models in an application using either a callback or context manager.
### Callback handler
```python
from langchain.chat_models import init_chat_model
from langchain_core.callbacks import UsageMetadataCallbackHandler

model_1 = init_chat_model(model="gpt-5.4-mini")
model_2 = init_chat_model(model="claude-haiku-4-5-20251001")

callback = UsageMetadataCallbackHandler()
result_1 = model_1.invoke("Hello", config={"callbacks": [callback]})
result_2 = model_2.invoke("Hello", config={"callbacks": [callback]})
print(callback.usage_metadata)

# {
#     'gpt-5.4-mini': {
#         'input_tokens': 8,
#         'output_tokens': 10,
#         'total_tokens': 18,
#         'input_token_details': {'audio': 0, 'cache_read': 0},
#         'output_token_details': {'audio': 0, 'reasoning': 0}
#     },
#     'claude-haiku-4-5-20251001': {
#         'input_tokens': 8,
#         'output_tokens': 21,
#         'total_tokens': 29,
#         'input_token_details': {'cache_read': 0, 'cache_creation': 0}
#     }
# }
```
### Context Manager
```python
from langchain.chat_models import init_chat_model
from langchain_core.callbacks import get_usage_metadata_callback

model_1 = init_chat_model(model="gpt-5.4-mini")
model_2 = init_chat_model(model="claude-haiku-4-5-20251001")

with get_usage_metadata_callback() as cb:
    model_1.invoke("Hello")
    model_2.invoke("Hello")
    print(cb.usage_metadata)

# same output as above
```

## Configurable models
You can create a **runtime-configurable** if you don't specify `model` and `model_provider` values.
```python
from langchain.chat_models import init_chat_model

configurable_model = init_chat_model(temperature=0)

configurable_model.invoke(
    "what's your name",
    config={"configurable": {"model": "gpt-5-nano"}},  # Run with GPT-5-Nano
)
configurable_model.invoke(
    "what's your name",
    config={"configurable": {"model": "claude-sonnet-4-6"}},  # Run with Claude
)

# Configurable model with default values
first_model = init_chat_model(
        model="gpt-5.4-mini",
        temperature=0,
        configurable_fields=("model", "model_provider", "temperature", "max_tokens"),
        config_prefix="first",  # Useful when you have a chain with multiple models
)

first_model.invoke("what's your name")

# or if you want to use another model
first_model.invoke(
    "what's your name",
    config={
        "configurable": {
            "first_model": "claude-sonnet-4-6",
            "first_temperature": 0.5,
            "first_max_tokens": 100,
        }
    },
)
```
## When to Use
You should use when you need interaction with an LLM model, don't necessary when you're creating an agent.
## When NOT to Use
Shouldn't be used in applications that don't require reasoning capabilities like a simple to-do application.
## Failure Modes
- Not considering vendor limitations like rate-limiting, token consumption.
- Not handling failures properly. LangChain provides an interface for interacting with the model but you still need to handle errors.
## Implementation (Practical)
### Model Initialization
#### Provider agnostic initialization
- Recommended method, avoids vendor lock-in.
- Each provider has its own initialization characteristics. For example, OpenAI requires an `OPENAI_API_KEY` to be set. For each model you plan to use it's a good idea to check the [docs](https://docs.langchain.com/oss/python/integrations/providers/overview).
```python
import os
from langchain.chat_models import init_chat_model

os.environ["OPENAI_API_KEY"] = "<api-key>"

model = init_chat_model("gpt-5.2")
```
#### Provider specific initialization
```python
import os
from langchain_openai import ChatOpenAI

os.environ["OPENAI_API_KEY"] = "<api-key>"

model = ChatOpenAI(model="gpt-5.2")
```
#### Configuring initialization
Configuration parameters is provider specific but usually it'll have:
- `model`
- `api_key`
- `temperature`
- `max_tokens`
- `timeout`: For provider APIs.
- `max_retries`: For provider APIs.
```python
model = init_chat_model(
	"claude-sonnet-4-6",
	temperature=0.7,
	timeout=30,
	max_tokens=1000,
	max_retries=6
)
```
#### Adding Rate Limiting
Some providers applies rate-limiting to their products. To avoid receiving error responses from your provider, you can apply a **local rate limiter** that control the rate at which requests are made.
```python
from langchain_code.rate_limiters import InMemoryRateLimiter

# only limits number of requests, not the size of them.
rate_limiter = InMemoryRateLimiter(
	requests_per_second=0.1, # 1 reuest every 10s
	check_every_n_seconds=0.1, # check every 100ms wheter allowed to make a request
	max_bucket_size=10 # Controls the maximum burst size
)

model = init_chat_model(
	model="gpt-5",
	model_provider="openai",
	rate_limiter=rate_limiter
)
```
### Model invocation
#### Invoke
- Returns a single [AIMessage]().
```python
# call model with a single message or a list of messages
response = model.invoke("Why do parrots have colorful feather?")

# call model with a list of messages
response = model.invoke([
	{"role": "system", "content": "You're a helpful assistant."},
	{"role": "user", "content": "Translate: I love programming."},
	{"role": "assistant", "content": "J'adore la programmation."},
])
```
#### Stream
- Returns output progressively.
- Calling this method returns an [Iterator]() that yield output chunks as they're produced..
- Returns multiple [AIMessageChunk](). Tool calls are built through [ToolCallChunk]().
- Each chunk in a stream is designed to be gathered into a full message via summation.
```python
full = None
for chunk in model.stream("Why parrots have coloful feathers?"):
	full = cunk if full is None else full + chunk	
	print(full.text)

# The
# The sky
# The sky is
# ...

print(full.content_blocks)
# [{"type": "text", "text": "The sky is typically blue..."}]
```
#### Batch
- Batch a collection of independent requests to a model which can be processed in parallel.
```python
responses = model.batch([
	"Why do parrots have colorful feather?",
	"How do airplanes fly?",
	"What is quantum computing?"
])
for response in responses:
	print(response)
```
- By default `batch()` will only return the final output for the **entire batch**. If you want to receive output individually, you can stream results with `batch_as_completed()`. Note that when using `batch_as_completed` result may arrive out of order. Each includes the input index to match back.
```python
for response in model.batch_as_completed([
	"Why do parrots have colorful feather?",
	"How do airplanes fly?",
	"What is quantum computing?"
]):
	print(response)
```
- When processing a large number of inputs with `batch()` or `batch_as_completed()` you may want to control the maximum number of parallel calls. 
```python
model.batch(
	list_of_inputs,
	config={
		"max_concurrency": 5
	}
)
```
#### Invocation Configuration
When invoking a model, you can pass additional configuration through the `config` parameters using a [RunnableConfig](https://reference.langchain.com/python/langchain-core/runnables/config/RunnableConfig?_gl=1*11ua93o*_gcl_au*MTU4OTcyNjY2LjE3NzM4Njc3MjE.*_ga*MTYxODAxODM1OS4xNzczODY3NzIx*_ga_47WX3HKKY2*czE3NzY4NzUxNDIkbzE2JGcxJHQxNzc2ODc1MjQwJGo2MCRsMCRoMA..) dictionary.
### Adding tools to the model
To make [tools]() that you've defined available for the model, you must bind them using `bing_tools()`.
```python
model_with_tools = model.bind_tools([get_weather])

response = model_with_tools.invoke("What's the weather like in Boston?"):
for tool_call in response.tool_calls:
	print(f"Tool: {tool_call["name"]}")
	print(f"Args: {tool_call["args"]}")
```
Some model providers offer **built int tools** that can be enabled via model or invocation parameters (e.g. `ChatOpenAI`, `ChatAnthropic`), check the [provider reference](https://docs.langchain.com/oss/python/integrations/providers/overview) for details.
```python
from langchain.chat_models import init_chat_model

model = init_chat_model("gpt-5.4-mini")

# server side tool call definition
tool = {"type": "web_search"}
model_with_tools = model.bind_tools([tool])

response = model_with_tools.invoke("What was a positive news story from today?")
print(response.content_blocks)
```
### Adding structured output to model's response
Models can be requested to provide their response in a format matching a given schema. LangChain supports the following schema types and methods:
#### [TypedDict]()
```python
from typing_extensions import TypedDict, Annotated

class MovieDict(TypedDict):
    """A movie with details."""
    title: Annotated[str, ..., "The title of the movie"]
    year: Annotated[int, ..., "The year the movie was released"]
    director: Annotated[str, ..., "The director of the movie"]
    rating: Annotated[float, ..., "The movie's rating out of 10"]

model_with_structure = model.with_structured_output(MovieDict)
response = model_with_structure.invoke("Provide details about the movie Inception")
print(response)  
# {'title': 'Inception', 'year': 2010, 'director': 'Christopher Nolan', 'rating': 8.8}
```
#### [JSON Schema](https://json-schema.org/understanding-json-schema/about)
```python
import json

json_schema = {
    "title": "Movie",
    "description": "A movie with details",
    "type": "object",
    "properties": {
        "title": {
            "type": "string",
            "description": "The title of the movie"
        },
        "year": {
            "type": "integer",
            "description": "The year the movie was released"
        },
        "director": {
            "type": "string",
            "description": "The director of the movie"
        },
        "rating": {
            "type": "number",
            "description": "The movie's rating out of 10"
        }
    },
    "required": ["title", "year", "director", "rating"]
}

model_with_structure = model.with_structured_output(
    json_schema,
    method="json_schema",
)
response = model_with_structure.invoke("Provide details about the movie Inception")
print(response)  
# {'title': 'Inception', 'year': 2010, ...}
```
#### [Pydantic]()
```python
from pydantic import BaseModel, Field

class Movie(BaseModel):
	"""A movie with details"""
	title: str = Field(description="The title of the movie")
	year: int = Field(description="The year the movie was release")
	director: str = Field(description="The director of the movie")
	rating: float = Field(description="The movie's rating our of 10")
	
model_with_structure = model.with_structured_output(Movie)
response = model_with_structure.invoke("Provide details about the movie Inception")
print(response)
# Movie(title="Inception", year=2010, director="Christopher Nolan", rating=8.8)
```

>**Pydantic** models provide automatic validation. `TypedDict` and **JSON Schema** require manual validation.
## Real-world Usage
- Where this was applied
- Link to decisions or systems
## Relationships
### Depends on
- [Agents](llm_agent.md)
- [Tools]()
- [Structured Output]()
- [LLMs](large_language_models.md)
- [Messages]()
- [Iterator]()
- [Pydantic]()
- [TypedDict]()
## Flashcards
- Q: What is the main purpose of LangChain model definitions?
- A: To configure the behavior of LLMs through a consistent interface across different providers.
- Q: What aspects of an LLM can be configured using LangChain model definitions?
- A: Text generation, tool availability, reasoning, structured output, and inspection of the model's chain of actions.
- Q: What advantage does LangChain provide when working with multiple model providers?
- A: It offers a standard interface that makes it easy to switch providers and dynamically select the best model for a task.
- Q: How do you install a specific LangChain provider package?
- A: By running a command like `pip install "langchain[provider_name]"`, for example `pip install "langchain[openai]"`.
- Q: What is a model profile in LangChain?
- A: A dictionary that describes a model’s supported features and capabilities, such as max input tokens, multimodal support, reasoning, and tool calling.
- Q: How can applications use model profiles dynamically?
- A: By adapting behavior based on capabilities, such as triggering summarization based on context window size or gating inputs based on supported modalities.
- Q: What is multimodal support in LangChain models?
- A: The ability of certain models to process and return non-text data using content blocks.
- Q: How can you access a model’s reasoning process in LangChain?
- A: By extracting reasoning content blocks from the model’s response when the model supports reasoning.
- Q: Can you control the level of reasoning effort in a model?
- A: Yes, depending on the model, you can specify the level of reasoning effort via configuration.
- Q: How is token usage information handled in LangChain?
- A: Some providers return token usage metadata, which is included in AIMessage objects and can be tracked using callbacks or context managers.
- Q: What are the two main ways to track token usage across models?
- A: Using a callback handler or a context manager.
- Q: What is a runtime-configurable model in LangChain?
- A: A model that allows changing parameters like model name or provider at invocation time instead of initialization.
- Q: How can you configure default and overridable parameters for a model?
- A: By specifying configurable fields and using a config prefix during initialization.
- Q: When should you use LangChain model abstractions?
- A: When you need to interact with an LLM model.
- Q: When should you avoid using LangChain model abstractions?
- A: In applications that don’t require reasoning capabilities, such as simple to-do apps.
- Q: What are common failure modes when using LangChain models?
- A: Ignoring vendor limitations like rate limits and token usage, and not handling errors properly.
- Q: What is the recommended way to initialize a model in a provider-agnostic way?
- A: Using `init_chat_model` with environment variables for API keys.
- Q: What is an example of provider-specific initialization?
- A: Using classes like `ChatOpenAI` from provider-specific packages.
- Q: What are common configuration parameters for model initialization?
- A: Model name, API key, temperature, max tokens, timeout, and max retries.
- Q: How can you handle provider rate limits in LangChain?
- A: By using a local rate limiter such as `InMemoryRateLimiter`.
- Q: What does the `invoke` method do?
- A: It sends a request to the model and returns a single AIMessage response.
- Q: What is the purpose of the `stream` method?
- A: To return output progressively as it is generated, using an iterator of message chunks.
- Q: What is the difference between `invoke` and `stream`?
- A: `invoke` returns a complete response, while `stream` returns partial outputs incrementally.
- Q: What does the `batch` method do?
- A: It processes multiple independent requests in parallel and returns their responses.
- Q: What is `batch_as_completed` used for?
- A: To stream batch results as they complete, possibly out of order.
- Q: How can you control concurrency in batch processing?
- A: By setting the `max_concurrency` parameter in the config.
- Q: What is `RunnableConfig` used for?
- A: To pass additional configuration options when invoking a model.
- Q: How do you make tools available to a model?
- A: By binding them using `bind_tools()`.
- Q: What are built-in tools in some model providers?
- A: Predefined tools like web search that can be enabled via model or invocation parameters.
- Q: What is structured output in LangChain?
- A: A way to enforce model responses to follow a predefined schema.
- Q: What schema types are supported for structured output?
- A: TypedDict, JSON Schema, and Pydantic.
- Q: What is an advantage of using Pydantic for structured output?
- A: It provides automatic validation of the response.
- Q: What is a limitation of TypedDict and JSON Schema compared to Pydantic?
- A: They require manual validation.
- Q: What components does LangChain model usage depend on?
- A: Agents, tools, structured output, LLMs, messages, iterators, Pydantic, and TypedDict.