package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

const LIBRARY_PATH = "../library"

func main() {
	var files []string

	err := getFilesWithCards(&files, LIBRARY_PATH)

	if err != nil {
		log.Fatal(err)
	}

	for _, fileName := range files {
		source, err := os.ReadFile(path.Join(LIBRARY_PATH, fileName))
		if err != nil {
			log.Fatal(err)
		}

		p := goldmark.DefaultParser()

		node := p.Parse(text.NewReader(source))

		ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
			if !entering {
				return ast.WalkContinue, nil
			}

			if heading, ok := n.(*ast.Heading); ok {
				text := extractText(heading, source)
				if text == "Flashcards" {
					fmt.Println(text)
				}
			}

			return ast.WalkContinue, nil
		})

		fmt.Println(fileName)
	}
}

func extractText(n ast.Node, source []byte) string {
	var buf bytes.Buffer

	for c := n.FirstChild(); c != nil; c.NextSibling() {
		if t, ok := c.(*ast.Text); ok {
			buf.Write(t.Segment.Value(source))
		}
	}

	return buf.String()
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
