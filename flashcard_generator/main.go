package main

import (
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

	f, err := os.OpenFile(
		"test_deck.txt",
		// create files if doesn't exist | open for writing only | truncate file
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString("#separator:Semicolon\n")
	f.WriteString("#columns:Front;Back;GUID;TAGS\n")
	f.WriteString("#deck:alexandria\n")
	f.WriteString("#tags column: 4 \n")
	f.WriteString("#guid column: 3 \n")

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

		fileTags := toStringSlice(metadata["tags"])

		var (
			inFlashcards bool
			question,
			answer string
		)

		qCount := 1

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
					question = content
				} else {
					answer = content
				}

				if question != "" && answer != "" {
					fmt.Fprintf(
						f,
						"%s;%s;%s;%s\n",
						strings.TrimPrefix(question, "Q: "),
						strings.TrimPrefix(answer, "A: "),
						fmt.Sprintf("%s-Q%d", metadata["id"], qCount),
						strings.Join(fileTags, " "),
					)
					qCount += 1
					question = ""
					answer = ""
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
