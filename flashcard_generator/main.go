package main

import (
	"flashcard_generator/src/card"
	deckmanagement "flashcard_generator/src/infra/deck_management"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

const LIBRARY_PATH = "../library"

func main() {
	var files []string

	err := getFilesWithCards(&files, LIBRARY_PATH)

	if err != nil {
		log.Fatal(err)
	}

	dm, err := deckmanagement.NewDeckManagement("test_deck.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dm.CloseDeck()

	if err = dm.WriteHeader(); err != nil {
		log.Fatal(err)
	}

	for _, fileName := range files {
		source, err := os.ReadFile(path.Join(LIBRARY_PATH, fileName))
		if err != nil {
			log.Fatal(err)
		}

		p := goldmark.New(
			goldmark.WithExtensions(
				meta.Meta,
			),
		)

		ctx := parser.NewContext()

		node := p.
			Parser().
			Parse(text.NewReader(source), parser.WithContext(ctx))

		metadata := meta.Get(ctx)

		var (
			inFlashcards bool
		)

		qCount := 1
		card := card.Card{
			Tags: toStringSlice(metadata["tags"]),
		}

		ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
			if !entering {
				return ast.WalkContinue, nil
			}

			// 1. Enter/Exit the Flashcards section
			if n.Kind() == ast.KindHeading {
				inFlashcards = string(getNodeText(n, source)) == "Flashcards"
				if !inFlashcards {
					qCount = 1
				}
				return ast.WalkContinue, nil
			}

			// 2. We only care about top-level ListItems inside the Flashcards section
			if inFlashcards && n.Kind() == ast.KindTextBlock {
				content := getNodeText(n, source)

				if strings.HasPrefix(content, "Q") {
					card.Question = strings.TrimPrefix(content, "Q: ")
				} else {
					card.Answer = strings.TrimPrefix(content, "A: ")
				}

				if card.Question != "" && card.Answer != "" {
					card.ID = fmt.Sprintf("%s-Q%d", metadata["id"], qCount)
					if err = dm.AddCard(card); err != nil {
						log.Fatal(err)
					}
					qCount += 1
					card.Question = ""
					card.Answer = ""
				}

				// Skip walking into children since we handled them manually
				return ast.WalkSkipChildren, nil
			}

			return ast.WalkContinue, nil
		})

		//node.Dump(source, 0)
	}
}

func getNodeText(n ast.Node, source []byte) string {
	var buf []byte
	lines := n.Lines()
	for i := 0; i < lines.Len(); i++ {
		segment := lines.At(i)
		buf = append(buf, segment.Value(source)...)
	}
	return string(buf)
}

func toStringSlice(v any) []string {
	switch val := v.(type) {
	case []any:
		out := make([]string, 0, len(val))
		for _, item := range val {
			if s, ok := item.(string); ok {
				out = append(out, s)
			}
		}
		return out
	case []string:
		return val
	case string:
		return []string{val}
	default:
		return nil
	}
}

func getFilesWithCards(files *[]string, dirName string) error {
	entries, err := os.ReadDir(dirName)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			getFilesWithCards(files, entry.Name())
		}
		*files = append(*files, entry.Name())
	}

	return nil
}
