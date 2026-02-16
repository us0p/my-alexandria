package card_test

import (
	"flashcard_generator/src/card"
	"slices"
	"testing"
)

func TestAddTag(t *testing.T) {
	cases := []struct {
		name         string
		tagToAdd     []string
		expectedTags []string
	}{
		{
			"non empty tag",
			[]string{"architecture"},
			[]string{"architecture"},
		},
		{
			"empty tag",
			[]string{""},
			[]string{},
		},
		{
			"adding many tags at once",
			[]string{"architecture", "networking"},
			[]string{"architecture", "networking"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			card := card.Card{}

			card.AddTag(c.tagToAdd...)

			if slices.Compare(card.Tags, c.expectedTags) != 0 {
				t.Errorf(
					"Expected tags slice to be equal to %v, got %v",
					c.expectedTags,
					card.Tags,
				)
			}
		})
	}
}
