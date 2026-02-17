package deckmanagement_test

import (
	"flashcard_generator/src/card"
	deckmanagement "flashcard_generator/src/infra/deck_management"
	"fmt"
	"os"
	"path"
	"strings"
	"testing"
)

func mockTargetDeckFile(t *testing.T) string {
	return path.Join(
		t.TempDir(),
		"test_deck.txt",
	)
}

func mockDeckManag(t *testing.T, targetDeck string) deckmanagement.DeckManagement {
	deckManag, err := deckmanagement.NewDeckManagement(targetDeck)
	if err != nil {
		t.Fatal(err)
	}

	return deckManag
}

func TestWriteDeckHeader(t *testing.T) {
	targetDeck := mockTargetDeckFile(t)
	deckManag := mockDeckManag(t, targetDeck)
	defer deckManag.CloseDeck()

	err := deckManag.WriteHeader()
	if err != nil {
		t.Fatal(err)
	}

	deck, err := os.ReadFile(targetDeck)
	if err != nil {
		t.Fatal(err)
	}

	expectedHeader := `#separator:Semicolon
#columns:Front;Back;GUID;TAGS
#deck:alexandria
#tags column: 4 
#guid column: 3 
`

	if string(deck) != expectedHeader {
		t.Errorf(
			"Expected deck header to be:\n%s\n\nGot:\n%s",
			expectedHeader,
			string(deck),
		)
	}
}

func TestAddDeckCard(t *testing.T) {
	targetDeck := mockTargetDeckFile(t)
	deckManag := mockDeckManag(t, targetDeck)
	defer deckManag.CloseDeck()

	_ = deckManag.WriteHeader()

	card := card.Card{
		Question: "Question",
		Answer:   "Answer",
		ID:       "ID",
		Tags:     []string{"Tag"},
	}
	err := deckManag.AddCard(card)
	if err != nil {
		t.Fatal(err)
	}

	deck, err := os.ReadFile(targetDeck)
	if err != nil {
		t.Fatal(err)
	}

	deckLines := strings.Split(string(deck), "\n")

	lastCardAdded := deckLines[len(deckLines)-2]
	expectedLastCardInfo := fmt.Sprintf(
		"%s;%s;%s;%s",
		card.Question,
		card.Answer,
		card.ID,
		strings.Join(card.Tags, " "),
	)

	if lastCardAdded != expectedLastCardInfo {
		t.Errorf(
			"Expected lastCardAdded to be: '%s', got: '%s'",
			expectedLastCardInfo,
			lastCardAdded,
		)
	}

}
