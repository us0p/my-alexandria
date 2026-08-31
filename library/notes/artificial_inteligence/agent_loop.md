---
id: 20260828-agent_loop
type: concept
status: draft | refined
tags:
  - AI
created: 2026-08-28
---
## TL;DR
Very short resume with only the essential information needed.
# Agent Loop
It's the loop that lets an AI agent to continue working until a task completion.

1. First, the agent gathers fresh data using its tools.
2. Plan/Reason on top of the information and determines what needs to be done.
3. Executes planned action.
4. Evaluates result of previous action.

It keeps executing those 4 steps in a loop and updating its internal state until the agent completes the task.
Each cycle allows the agent to incorporate fresh information (observations) into its reasoning (thought), ensuring that the final answer is well-informed and accurate.

You can think of this cycle as **Thought-Action-Observation**

There are two limitations when creating an agent loop:
- Context window: tool usage, reasoning and file contents can consume a lot of the context window depending on the task.
- Action execution: Calling tools and execution commands that can take a really long time can potentially block the loop and make the agent unresponsive. It's a good idea to give capability for the agent to run those in the background to avoid blocking the agent execution.

1. **Perception**: Involves gathering data through sensors, processing that information, and interpreting it to perform specific tasks. A good perception phase guarantees the agent reason upon facts and not guesses. Processing large volumes de data or interpreting complex information can lead to errors and delay in decision making.
2. **Reason and Plan**: The agent analise the gathered data and determine necessary steps and the order in which they must be executed so that it can achieve the goal. It can also think about what can go wrong and create backup steps.

![[Pasted image 20260828214538.png]]
## Understanding
- explanation of the concept, using your own words.
- Focus on cause and effect.
Ex:
- This pattern exists because systems are likely to couple business rules and external details...
- The separation allows changing interfaces without having to rewrite central rules...
## When to Use
- Situations where this is useful
## When NOT to Use
- Situations where this is overkill or harmful
## Trade-offs
- Limitations
- Costs
- Complexity 
## Examples
## References
### Connects with
Add link to relative notes
### Contrasts with
- Add link to alternatives that tries to solve the same problem
- Always add relation definition like "expands", "contrasts", "depends"
## Questions
- Points that are still not clear.
## Iterate on
- Sections of the document that can be iterated and have it's quality 
improved but need more knowledge to do so.
## Flashcards
- Q: Some question about the notes.
- A: The answer for the question above.
