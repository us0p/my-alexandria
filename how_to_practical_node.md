## High-frequency concepts
Are concepts that are worth keeping reference. 

High-frequency notes should create shortcuits for stuff you need to revisit
often. Usually they refer to stuff that doesn't chage.

High-frequency notes should be optimized for:
- fast-recall
- decision-making
- mental-models

>More important information should have more visual highlight.

### How to identify high-frequency concepts
You can identify then with these questions:
- How often will I need this to think or decide?
- Forgetting about this would hurt your thoght process?
- Should consider about using this or not?
- Can it cause something to fail?

>If you only need something when another happens, it's a sign that it's 
not high-frequency.

## Low-frequency concepts
Low-frequency concepts should be reference notes. They shouldn't be mixed 
with high-frequency notes as this increases the cognitive overload.

Shouldn't be extense. If you're copying the same example of the 
documentation, it's better to store a reference to the documentation as
it's always going to be the most up to date document.

If you're storying something different than the documentation or something 
that's not in the documentation, then it's worth to store it as a 
low-frequency reference note, e.g. a pattern on how to use a specific 
function of a libray, the documentation will explain how to use the 
function, the parameters and side-effects of it, but it might not explain 
how to setup the environment of it in the most efficient manner.

## Pattern-level insights
You should bootstrap patterns early, then refine them as you gain more 
experience.

A pattern is a repeated relationship between components that solves a 
recurring problem.

>Your first version of a pattern will be vague, slightly wrong and 
incomplete. And that's ok, you'll refine this as you gain more experience.

### How to extract patterns early
- Ask "What problem is this solving?"
- The answer is generaly going to be the initial pattern.

For example: 
- "What problem does a LLM tool tries to solve?"
- "LLM can't interact with external systems"
- Pattern: LLM tools bridge between LLM and external world.

### Look for role, not implementation
Patterns comes from roles in a system. A role tells you where and when you
should do something.

For example (LLM tool usage):
- `ToolRuntime`: Access to environment
- `Tool`: Executes action

You can say that: **Tools are the execution layer of an agent**. This is a
pattern.

### Force analogies
Ask, "What does this remind me of?". This will give you a high-quality 
pattern just by mappign to other concepts you already know.
