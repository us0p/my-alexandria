package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/yuin/goldmark"
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

		fmt.Println(fileName)
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
