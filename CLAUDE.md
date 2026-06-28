# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this repo is

A personal knowledge base (Zettelkasten + Generative Notes) with an accompanying Go tool that auto-generates Anki flashcard decks from the notes.

## Commands

All commands run from inside `flashcard_generator/`:

```bash
go run .                        # generate test_deck.txt from ../library
go test ./...                   # run all tests
go test ./src/domain/card/ -run TestName  # run a single test
```

## Note structure

Notes live under `library/` as Markdown files with YAML frontmatter. Three note types exist — use the templates at the repo root:

| Type | Template | Purpose |
|------|----------|---------|
| concept | `atomic_note_template.md` | Definition of a concept |
| practice | `practical_note_template.md` | How/where to apply a concept |
| decision | `decision_note_template.md` | Why one option was chosen over alternatives |

Frontmatter fields: `id`, `type`, `status` (`draft`/`refined`/`validated`/`obsolete`), `tags`, `created`.

Flashcards are extracted from the `## Flashcards` section of any note:
```markdown
## Flashcards
- Q: question text
- A: answer text
```

Card IDs are derived from the note's `id` frontmatter field (`{id}-Q{n}`).

## flashcard_generator architecture

```
main.go                              → entry point
src/domain/card/card.go              → Card entity (Question, Answer, ID, Tags)
src/domain/deck_management/          → writes Anki-compatible semicolon-separated deck file
src/use_cases/flashcard_generator/   → walks library/, parses markdown with goldmark,
                                       extracts flashcard Q/A pairs, calls domain
```

The use case (`flashcard_generator.go`) couples goldmark (external markdown parser) directly to domain logic. The library path is hardcoded as `../library` relative to where the binary runs — always run from inside `flashcard_generator/`.

Output (`test_deck.txt`) is an Anki import file with header fields: `Front;Back;GUID;TAGS`, deck name `alexandria`.

## Note-taking conventions

`how_to_practical_node.md` defines three note categories:
- **High-frequency**: optimize for fast-recall and decision-making
- **Low-frequency**: reference notes, prefer linking to docs over copying them
- **Pattern-level**: extract roles not implementations; bootstrap early and refine
