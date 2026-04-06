package card_test

import (
	"errors"
	"flashcard_generator/src/domain/card"
	"slices"
	"testing"
)

func TestAddTags(t *testing.T) {
	cases := []struct {
		name         string
		tagToAdd     any
		expectedTags []string
	}{
		{
			"non empty tag, strings slice input",
			[]string{"architecture"},
			[]string{"architecture"},
		},
		{
			"empty tag, strings slice input",
			[]string{""},
			[]string{},
		},
		{
			"adding many tags at once, strings slice input",
			[]string{"architecture", "networking"},
			[]string{"architecture", "networking"},
		},
		{
			"non empty tag, any slice input",
			[]any{"architecture"},
			[]string{"architecture"},
		},
		{
			"empty tag, any slice input",
			[]any{""},
			[]string{},
		},
		{
			"adding many tags at once, any slice input",
			[]any{"architecture", "networking"},
			[]string{"architecture", "networking"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			card := card.Card{}

			card.AddTags(c.tagToAdd)
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

func TestAddTagsFailures(t *testing.T) {
	cases := []struct {
		name          string
		tagToAdd      any
		expectedError error
	}{
		{
			"input array with numbers instead of string",
			[]any{123, 456},
			card.ErrNotStringTag,
		},
		{
			"input with numbers instead of string",
			123,
			card.ErrNotStringTag,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			card := card.Card{}

			err := card.AddTags(c.tagToAdd)
			if !errors.Is(err, c.expectedError) {
				t.Errorf("Expected error %s, got %v\n", c.expectedError, err)
			}
		})
	}
}
