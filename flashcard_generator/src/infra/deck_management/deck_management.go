package deckmanagement

import (
	"flashcard_generator/src/card"
	"fmt"
	"os"
	"strings"
)

type DeckManagement struct {
	deckFile *os.File
}

func NewDeckManagement(deckFileName string) (DeckManagement, error) {
	deckFile, err := os.OpenFile(
		deckFileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		return DeckManagement{}, err
	}

	return DeckManagement{deckFile}, nil
}

func (dm DeckManagement) WriteHeader() error {
	_, err := dm.deckFile.WriteString("#separator:Semicolon\n")
	if err != nil {
		return err
	}
	_, err = dm.deckFile.WriteString("#columns:Front;Back;GUID;TAGS\n")
	if err != nil {
		return err
	}
	_, err = dm.deckFile.WriteString("#deck:alexandria\n")
	if err != nil {
		return err
	}
	_, err = dm.deckFile.WriteString("#tags column: 4 \n")
	if err != nil {
		return err
	}
	_, err = dm.deckFile.WriteString("#guid column: 3 \n")
	if err != nil {
		return err
	}

	return nil
}

func (dm DeckManagement) AddCard(deckCard card.Card) error {
	fmt.Fprintf(
		dm.deckFile,
		"%s;%s;%s;%s\n",
		deckCard.Question,
		deckCard.Answer,
		deckCard.ID,
		strings.Join(deckCard.Tags, " "),
	)

	return nil
}

func (dm DeckManagement) CloseDeck() error {
	return dm.deckFile.Close()
}
