package flashcardgenerator

import (
	"flashcard_generator/src/domain/card"
	deckmanagement "flashcard_generator/src/domain/deck_management"
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

func Foo() {
	var files []string

	// fs util
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
		// fs util
		source, err := os.ReadFile(path.Join(LIBRARY_PATH, fileName))
		if err != nil {
			log.Fatal(err)
		}

		// external library
		p := goldmark.New(
			goldmark.WithExtensions(
				meta.Meta,
			),
		)

		// external library
		ctx := parser.NewContext()

		// external library
		node := p.
			Parser().
			Parse(text.NewReader(source), parser.WithContext(ctx))

		// eternal library
		metadata := meta.Get(ctx)

		var (
			inFlashcards bool
		)

		qCount := 1
		card := card.Card{
			// card tag creation, should be moved to card entity
			Tags: toStringSlice(metadata["tags"]),
		}

		ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
			// tree walking logic, external library with domain connection
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
				// external library util
				content := getNodeText(n, source)

				// card building logic
				if strings.HasPrefix(content, "Q") {
					card.Question = strings.TrimPrefix(content, "Q: ")
				} else {
					card.Answer = strings.TrimPrefix(content, "A: ")
				}

				// Card storing logic
				if card.Question != "" && card.Answer != "" {
					card.ID = fmt.Sprintf("%s-Q%d", metadata["id"], qCount)
					if err = dm.AddCard(card); err != nil {
						log.Fatal(err)
					}
					qCount += 1
					card.Question = ""
					card.Answer = ""
				}

				// external library util
				// Skip walking into children since we handled them manually
				return ast.WalkSkipChildren, nil
			}

			// external library util
			return ast.WalkContinue, nil
		})
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
