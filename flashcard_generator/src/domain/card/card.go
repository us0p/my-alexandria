package card

import (
	"errors"
	"fmt"
)

var ErrNotStringTag = errors.New("Tag is not a string")

type Card struct {
	Question string
	Answer   string
	ID       string
	Tags     []string
}

func (c *Card) AddTags(tags any) error {
	switch val := tags.(type) {
	case []any:
		for _, item := range val {
			if s, ok := item.(string); ok && s != "" {
				c.Tags = append(c.Tags, s)
				continue
			}
			return fmt.Errorf("%w: %v\n", ErrNotStringTag, item)
		}
	case []string:
		for _, tag := range val {
			if tag != "" {
				c.Tags = append(c.Tags, tag)
			}
		}
	case string:
		if val != "" {
			c.Tags = []string{val}
		}
	default:
		return fmt.Errorf("%w: %v\n", ErrNotStringTag, val)
	}

	return nil
}
